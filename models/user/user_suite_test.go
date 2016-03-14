package user_test

import (
	"bitbucket.com/hyperboloide/horo/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var _ = BeforeSuite(func() {
	config.Configure()
})
