package mail_test

import (
	. "bitbucket.com/hyperboloide/horo/services/mail"
	"fmt"
	"github.com/dchest/uniuri"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mail", func() {

	It("should send an email", func() {
		addr := fmt.Sprintf("%s@%s.com", uniuri.NewLen(12), uniuri.NewLen(12))
		err := NewMessage(
			"test",
			"Test Message",
			"<h1>Test Message</h1>",
			addr)
		Ω(err).To(BeNil())
	})

})
