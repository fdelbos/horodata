package main

import (
	"encoding/json"
	"errors"

	"dev.hyperboloide.com/fred/horodata/services/payment"
	log "github.com/Sirupsen/logrus"
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
			log.WithField("id", id).Info("processing invoice")
			return ProcessInvoice(id)
		}
	default:
		log.WithField("type", e.Type).Info("dropping unsupported event")
		return nil
	}
}
