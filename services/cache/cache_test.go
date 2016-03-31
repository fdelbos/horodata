package cache_test

import (
	"dev.hyperboloide.com/fred/horodata/models/errors"
	. "dev.hyperboloide.com/fred/horodata/services/cache"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Cache", func() {

	It("should set some data", func() {
		Ω(SetPackage("test", "kiki", "str", time.Minute)).To(BeNil())
	})

	It("should get some data", func() {
		var str string
		Ω(GetPackage("test", "kiki", &str)).To(BeNil())
		Ω(str).To(Equal("str"))
	})

	It("should update some data", func() {
		Ω(SetPackage("test", "kiki", "str2", time.Minute)).To(BeNil())
		var str string
		Ω(GetPackage("test", "kiki", &str)).To(BeNil())
		Ω(str).To(Equal("str2"))
	})

	It("should delete some data", func() {
		Ω(DelPackage("test", "kiki")).To(BeNil())
		Ω(DelPackage("test", "kiki")).To(BeNil())
		var str string
		Ω(GetPackage("test", "kiki", &str)).To(Equal(errors.NotFound))
	})
})
