package cache_test

import (
	_ "bitbucket.com/hyperboloide/horo/config"
	"bitbucket.com/hyperboloide/horo/services/cache"
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
