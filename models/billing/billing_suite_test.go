package billing_test

import (
	"bitbucket.com/hyperboloide/ud/config"
	"bitbucket.com/hyperboloide/ud/helpers/tests"
	"bitbucket.com/hyperboloide/ud/models/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
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
