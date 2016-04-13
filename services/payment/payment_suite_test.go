package payment_test

import (
	"dev.hyperboloide.com/fred/horodata/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPayment(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payment Suite")
}

var _ = BeforeSuite(func() {
	config.Configure()
})
