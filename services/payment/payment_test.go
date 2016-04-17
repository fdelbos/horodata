package payment_test

import (
	. "dev.hyperboloide.com/fred/horodata/services/payment"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Payment", func() {

	It("should send an event", func() {
		err := NewEvent("evt_181FdWFjT5XSbba34KNBFKc6")
		Ω(err).To(BeNil())
	})
})
