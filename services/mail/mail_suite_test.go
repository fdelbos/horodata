package mail_test

import (
	_ "dev.hyperboloide.com/fred/horodata/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestMail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mail Suite")
}
