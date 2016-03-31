package billing_test

import (
	"testing"

	"dev.hyperboloide.com/fred/horodata/config"
	"dev.hyperboloide.com/fred/horodata/helpers/tests"
	"dev.hyperboloide.com/fred/horodata/models/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var Owner *user.User

func TestBilling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Billing Suite")
}

var _ = BeforeSuite(func() {
	config.Configure()
	var err error

	Owner, err = tests.NewUser()
	Ω(err).To(BeNil())
})

var _ = AfterSuite(func() {
	Ω(tests.CleanupUser(Owner)).To(BeNil())
})
