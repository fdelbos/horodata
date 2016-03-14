package group_test

import (
	"bitbucket.com/hyperboloide/horo/config"
	"bitbucket.com/hyperboloide/horo/helpers/tests"
	"bitbucket.com/hyperboloide/horo/models/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var Owner *user.User

func TestGroup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Suite")
}

var _ = BeforeSuite(func() {
	config.Configure()
	var err error

	Owner, err = tests.NewUser()
	Ω(err).To(BeNil())
})

var _ = AfterSuite(func() {
	//Ω(tests.CleanupUser(Owner)).To(BeNil())
})
