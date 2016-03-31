package cache_test

import (
	_ "dev.hyperboloide.com/fred/horodata/config"
	"dev.hyperboloide.com/fred/horodata/services/cache"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCache(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cache Suite")
}

var _ = BeforeSuite(func() {
	cache.Configure()
})
