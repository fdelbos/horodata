package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"dev.hyperboloide.com/fred/horodata/models/billing"
	"dev.hyperboloide.com/fred/horodata/services/mail"
	"dev.hyperboloide.com/fred/horodata/services/payment"
	"github.com/hyperboloide/qmail/client"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/event"
)

var (
	ErrInvoiceNoId     = errors.New("The invoice has no id")
	ErrInvoiceIdNotStr = errors.New("The invoice id is not a string")
	ErrPDFGEN          = errors.New("pdfgen error")
)

func listenner(buff []byte) error {
	msg := payment.StripeEvent{}
	if err := json.Unmarshal(buff, &msg); err != nil {
		return err
	}

	e, err := event.Get(msg.Id, nil)
	if err != nil {
		return err
	}

	switch e.Type {
	case "invoice.created":
		if idObj, ok := e.Data.Obj["id"]; !ok {
			return ErrInvoiceNoId
		} else if id, ok := idObj.(string); !ok {
			return ErrInvoiceIdNotStr
		} else {
			return DoInvoice(id)
		}
	default:
		return nil
	}
}

func DoInvoice(stripeId string) error {
	if _, err := billing.InvoiceByStripeId(stripeId); err == nil {
		return nil
	} else if err := billing.NewInvoice(stripeId); err != nil {
		return err
	} else if invoice, err := billing.InvoiceByStripeId(stripeId); err != nil {
		return err
	} else if address, err := billing.AddressById(invoice.AddressId); err != nil {
		return err
	} else {

		invoiceId := fmt.Sprintf("HD-%06d", invoice.Id)
		msg := struct {
			Id         string           `json:"id"`
			ClientId   string           `json:"client_id"`
			Created    string           `json:"created"`
			Address    *billing.Address `json:"address"`
			StartDate  string           `json:"start_date"`
			EndDate    string           `json:"end_date"`
			Subtotal   string           `json:"subtotal"`
			Total      string           `json:"total"`
			Tax        string           `json:"tax"`
			TaxPercent string           `json:"tax_percent"`
			Paid       bool             `json:"paid"`
			Lines      []interface{}    `json:"lines"`
		}{
			invoiceId,
			fmt.Sprintf("HD-%09d", invoice.UserId),
			dateToString(invoice.Created),
			address,
			dateToString(invoice.StartDate),
			dateToString(invoice.EndDate),
			numberToDecString(invoice.SubTotal),
			numberToDecString(invoice.Total),
			numberToDecString(invoice.Tax),
			floatToPercent(invoice.TaxPercent),
			invoice.Paid,
			[]interface{}{},
		}
		for _, l := range invoice.Lines {
			msg.Lines = append(msg.Lines, struct {
				Amount      string `json:"amount"`
				UnitPrice   string `json:"unit_price"`
				Quantity    string `json:"quantity"`
				StartDate   string `json:"start_date"`
				EndDate     string `json:"end_date"`
				Description string `json:"description"`
				Title       string `json:"title"`
			}{
				numberToDecString(l.Amount),
				numberToDecString(int64(l.UnitPrice)),
				fmt.Sprintf("%d", l.Quantity),
				dateToString(l.StartDate),
				dateToString(l.EndDate),
				l.Description,
				l.Title,
			})
		}

		if jsonStr, err := json.Marshal(msg); err != nil {
			return err
		} else {
			fmt.Println(string(jsonStr[:]))
		}

		dest := viper.GetString("pdf_host") + viper.GetString("pdf_individual")
		if address.VAT != "" {
			dest = viper.GetString("pdf_host") + viper.GetString("pdf_business")
		}
		path := path.Join(viper.GetString("output"), fmt.Sprintf("%s.pdf", invoiceId))

		if jsonStr, err := json.Marshal(msg); err != nil {
			return err
		} else if resp, err := http.Post(dest, "application/json", bytes.NewBuffer(jsonStr)); err != nil {
			return err
		} else if resp.StatusCode != 200 {
			return ErrPDFGEN
		} else if file, err := os.Create(path); err != nil {
			return err
		} else {
			defer file.Close()

			if _, err := io.Copy(file, resp.Body); err != nil {
				return err
			}
			err := mail.Mailer().Send(client.Mail{
				Dests:    []string{address.Email},
				Subject:  fmt.Sprintf("Facture Horodata du %s", dateToString(invoice.Created)),
				Template: "invoice",
				Data: map[string]interface{}{
					"Date": invoice.Created},
				Files: []string{
					fmt.Sprintf("/invoices/%s.pdf", invoiceId)},
			})
			if err != nil {
				return err
			}
			return invoice.MarkAsSent()
		}
	}
}

func numberToDecString(nb int64) string {
	rem := fmt.Sprintf("%d", nb%100)
	if nb%100 < 10 {
		rem = fmt.Sprintf("0%d", nb%100)
	}
	return fmt.Sprintf("%d,%s", nb/100, rem)
}

func dateToString(d time.Time) string {
	return d.Format("02/01/2006")
}

func floatToPercent(nb float64) string {
	return fmt.Sprint("%.2f", nb)
}
