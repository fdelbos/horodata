package mail_test

import (
	_ "bitbucket.com/hyperboloide/horo/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestMail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mail Suite")
}
