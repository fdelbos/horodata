package mail_test

import (
	"fmt"

	. "dev.hyperboloide.com/fred/horodata/services/mail"
	"github.com/dchest/uniuri"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mail", func() {

	It("should send an email", func() {
		addr := fmt.Sprintf("%s@%s.com", uniuri.NewLen(12), uniuri.NewLen(12))
		// addr := "fred@hyperboloide.com"
		Configure()
		err := NewMessage(
			"test",
			"Test Message",
			"<h1>Test Message</h1>",
			addr)
		Ω(err).To(BeNil())
	})

})
