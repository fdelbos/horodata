package billing_test

import (
	. "bitbucket.com/hyperboloide/ud/models/billing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Address", func() {

	var addr *Address

	It("should create an address", func() {
		addr = &Address{
			UserId:    Owner.Id,
			FullName:  "Test User",
			Email:     Owner.Email,
			Company:   "test company",
			VAT:       "1234567890",
			Address1:  "rue du test",
			City:      "Test-Ville sur mer",
			Zip:       "12345",
			CountryId: "FR",
		}
		Ω(addr.Insert()).To(BeNil())
	})

	It("should get current address and taxes", func() {
		a, err := CurrentAddress(Owner.Id)
		Ω(err).To(BeNil())
		Ω(a).ToNot(BeNil())
		Ω(a.UserId).To(Equal(Owner.Id))

		ok, err := a.HasTax()
		Ω(err).To(BeNil())
		Ω(ok).To(BeTrue())

		tp, err := a.TaxePercent()
		Ω(err).To(BeNil())
		Ω(tp).To(Equal(20))
	})

	It("should get countries list", func() {
		l, err := CountriesList("en")
		Ω(err).To(BeNil())
		Ω(l).ToNot(BeNil())
		Ω(l[2].Name).To(Equal("Algeria"))

		l, err = CountriesList("fr")
		Ω(err).To(BeNil())
		Ω(l).ToNot(BeNil())
		Ω(l[1].Name).To(Equal("Afrique du Sud"))
		Ω(l[1].Id).To(Equal("ZA"))
	})

})
