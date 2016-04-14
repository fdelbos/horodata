package billing

import (
	"fmt"
	"time"

	"dev.hyperboloide.com/fred/horodata/services/postgres"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/invoice"
)

type InvoiceLine struct {
	Id          int64
	StripeId    string
	InvoiceId   int64
	Amount      int64
	UnitPrice   uint64
	Quantity    int64
	StartDate   time.Time
	EndDate     time.Time
	Description string
	Title       string
}

func (i *InvoiceLine) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&i.Id,
		&i.StripeId,
		&i.InvoiceId,
		&i.Amount,
		&i.UnitPrice,
		&i.Quantity,
		&i.StartDate,
		&i.EndDate,
		&i.Description,
		&i.Title)
}

type Invoice struct {
	Id         int64
	StripeId   string
	Created    time.Time
	UserId     int64
	AddressId  int64
	StartDate  time.Time
	EndDate    time.Time
	SubTotal   int64
	Total      int64
	Tax        int64
	TaxPercent float64
	Paid       bool
	Sent       bool
	ChargeId   *string
	Lines      []InvoiceLine
}

func (i *Invoice) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&i.Id,
		&i.StripeId,
		&i.Created,
		&i.UserId,
		&i.AddressId,
		&i.StartDate,
		&i.EndDate,
		&i.SubTotal,
		&i.Total,
		&i.Tax,
		&i.TaxPercent,
		&i.Paid,
		&i.Sent,
		&i.ChargeId)
}

func (i *Invoice) MarkAsSent() error {
	const query = `
	update invoices
	set sent = true
	where id = $1`
	return postgres.Exec(query, i.Id)
}

func (i *Invoice) FileId() string {
	return fmt.Sprintf("HD-%06d", i.Id)
}

func (i *Invoice) Address() (*Address, error) {
	return AddressById(i.AddressId)
}

func InvoiceByStripeId(id string) (*Invoice, error) {
	invoice := &Invoice{}
	const queryInvoice = `
    select * from invoices
    where stripe_id = $1;`
	if err := postgres.QueryRow(invoice, queryInvoice, id); err != nil {
		return nil, err
	}

	const queryLines = `
    select * from invoice_lines
    where invoice_id = $1
    order by id;`

	rows, err := postgres.DB().Query(queryLines, invoice.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		i := InvoiceLine{}
		if err := i.Scan(rows.Scan); err != nil {
			return nil, err
		}
		invoice.Lines = append(invoice.Lines, i)
	}
	return invoice, rows.Err()
}

func InvoiceById(id int64) (*Invoice, error) {
	invoice := &Invoice{}
	const queryInvoice = `
    select * from invoices
    where id = $1;`
	if err := postgres.QueryRow(invoice, queryInvoice, id); err != nil {
		return nil, err
	}

	const queryLines = `
    select * from invoice_lines
    where invoice_id = $1
    order by id;`

	rows, err := postgres.DB().Query(queryLines, invoice.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		i := InvoiceLine{}
		if err := i.Scan(rows.Scan); err != nil {
			return nil, err
		}
		invoice.Lines = append(invoice.Lines, i)
	}
	return invoice, rows.Err()
}

func NewInvoice(stripeId string) error {
	if i, err := invoice.Get(stripeId, nil); err != nil {
		return err
	} else if customer, err := CustomerByStripeId(i.Customer.ID); err != nil {
		return err
	} else if address, err := customer.Address(); err != nil {
		return err
	} else {

		newInvoice := &Invoice{
			StripeId:   i.ID,
			UserId:     customer.UserId,
			AddressId:  address.Id,
			StartDate:  time.Unix(i.Start, 0),
			EndDate:    time.Unix(i.End, 0),
			SubTotal:   i.Subtotal,
			Total:      i.Total,
			Tax:        i.Tax,
			TaxPercent: i.TaxPercent,
			Paid:       i.Paid,
		}
		if i.Charge != nil {
			newInvoice.ChargeId = &i.Charge.ID
		}

		newLines := []InvoiceLine{}
		params := &stripe.InvoiceLineListParams{ID: i.ID}
		params.Filters.AddFilter("limit", "", "200")
		iter := invoice.ListLines(params)
		for iter.Next() {
			l := iter.InvoiceLine()
			newLines = append(newLines, InvoiceLine{
				StripeId:    l.ID,
				Amount:      l.Amount,
				UnitPrice:   l.Plan.Amount,
				Quantity:    l.Quantity,
				StartDate:   time.Unix(l.Period.Start, 0),
				EndDate:     time.Unix(l.Period.End, 0),
				Description: i.Desc,
				Title:       l.Plan.Statement,
			})
		}

		tx, err := postgres.DB().Begin()
		if err != nil {
			return err
		}
		defer tx.Rollback()

		stmtInvoice, err := tx.Prepare(`
            insert into invoices
                (stripe_id, user_id, address_id, start_date, end_date, subtotal, total, tax, tax_percent, paid, charge_id)
            values
                ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
            returning id;`)
		if err != nil {
			return err
		}
		defer stmtInvoice.Close()

		stmtLine, err := tx.Prepare(`
            insert into invoice_lines
                (stripe_id, invoice_id, amount, unit_price, quantity, start_date, end_date, description, title)
            values
                ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
		if err != nil {
			return err
		}
		defer stmtLine.Close()

		err = stmtInvoice.QueryRow(
			newInvoice.StripeId,
			newInvoice.UserId,
			newInvoice.AddressId,
			newInvoice.StartDate,
			newInvoice.EndDate,
			newInvoice.SubTotal,
			newInvoice.Total,
			newInvoice.Tax,
			newInvoice.TaxPercent,
			newInvoice.Paid,
			newInvoice.ChargeId).Scan(&newInvoice.Id)
		if err != nil {
			return err
		}

		for _, line := range newLines {
			_, err := stmtLine.Exec(
				line.StripeId,
				newInvoice.Id,
				line.Amount,
				line.UnitPrice,
				line.Quantity,
				line.StartDate,
				line.EndDate,
				line.Description,
				line.Title)
			if err != nil {
				return err
			}
		}
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}

type InvoicePreview struct {
	Id        int64     `json:"id"`
	Created   time.Time `json:"created"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Total     int64     `json:"total"`
	Paid      bool      `json:"paid"`
	Sent      bool      `json:"sent"`
}

func (i *InvoicePreview) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&i.Id,
		&i.Created,
		&i.StartDate,
		&i.EndDate,
		&i.Total,
		&i.Paid,
		&i.Sent)
}

func InvoicePreviewsByUserId(id int64) ([]InvoicePreview, error) {
	const query = `
    select
        id, created, start_date, end_date, total, paid, sent
    from invoices
    where user_id = $1
    order by id desc;`

	rows, err := postgres.DB().Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	previews := []InvoicePreview{}
	for rows.Next() {
		i := InvoicePreview{}
		if err := i.Scan(rows.Scan); err != nil {
			return nil, err
		}
		previews = append(previews, i)
	}
	return previews, rows.Err()
}
