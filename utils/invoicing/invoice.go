package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"dev.hyperboloide.com/fred/horodata/models/billing"
	"dev.hyperboloide.com/fred/horodata/models/errors"
	"dev.hyperboloide.com/fred/horodata/services/mail"
	log "github.com/Sirupsen/logrus"
	"github.com/hyperboloide/qmail/client"
	"github.com/spf13/viper"
)

func ProcessInvoice(stripeId string) error {
	invoice, err := billing.InvoiceByStripeId(stripeId)

	if err == nil && !invoice.Sent {
		log.WithFields(map[string]interface{}{
			"stripe_id": stripeId,
			"id":        invoice.Id,
		}).Info("invoice incomplete; rebuild")
		return buildPdf(invoice)
	} else if err == nil {
		log.WithFields(map[string]interface{}{
			"stripe_id": stripeId,
			"id":        invoice.Id,
		}).Info("invoice already sent")
		return nil
	} else if err != nil && err != errors.NotFound {
		return err
	} else if err := billing.NewInvoice(stripeId); err != nil {
		return err
	} else if invoice, err := billing.InvoiceByStripeId(stripeId); err != nil {
		return err
	} else {
		return buildPdf(invoice)
	}
}

func genMsg(invoice *billing.Invoice) interface{} {

	address, err := invoice.Address()
	if err != nil {
		return err
	}

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
		invoice.FileId(),
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
	return msg
}

func buildPdf(invoice *billing.Invoice) error {

	msg := genMsg(invoice)
	dest := viper.GetString("pdf_host") + viper.GetString("pdf_invoice")
	filePath := path.Join(
		viper.GetString("output"),
		fmt.Sprintf("%s.pdf", invoice.FileId()))

	if jsonStr, err := json.Marshal(msg); err != nil {
		return err
	} else if resp, err := http.Post(dest, "application/json", bytes.NewBuffer(jsonStr)); err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return ErrPDFGEN
	} else if file, err := os.Create(filePath); err != nil {
		return err
	} else {
		defer file.Close()

		if _, err := io.Copy(file, resp.Body); err != nil {
			return err
		}
		resp.Body.Close()

		log.WithFields(map[string]interface{}{
			"file": filePath,
			"id":   invoice.FileId(),
		}).Info("new invoice saved")
	}
	return sendInvoice(invoice)
}

func sendInvoice(invoice *billing.Invoice) error {

	address, err := invoice.Address()
	if err != nil {
		return err
	}

	err = mail.Mailer().Send(client.Mail{
		Dests:    []string{address.Email},
		Subject:  fmt.Sprintf("Facture Horodata du %s", dateToString(invoice.Created)),
		Template: "invoice",
		Data: map[string]interface{}{
			"Date": invoice.Created},
		Files: []string{
			fmt.Sprintf("/invoices/%s.pdf", invoice.FileId())},
	})
	if err != nil {
		return err
	}
	log.WithField("id", invoice.FileId()).Info("invoice sent")
	return invoice.MarkAsSent()
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
