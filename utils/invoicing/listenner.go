package main

import (
	"encoding/json"
	"errors"

	"dev.hyperboloide.com/fred/horodata/services/payment"
	"github.com/stripe/stripe-go/event"
	"github.com/stripe/stripe-go/invoice"
)

var (
	ErrInvoiceNoId     = errors.New("The invoice has no id")
	ErrInvoiceIdNotStr = errors.New("The invoice id is not a string")
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
			return GetInvoice(id)
		}
	default:
		return nil
	}
}

func GetInvoice(id string) error {
	i, err := invoice.Get(id, nil)
	if err != nil {
		return err
	}

	//c, err := i.Customer.ID

	return nil
}
