package group_test

import (
	. "bitbucket.com/hyperboloide/horo/models/group"
	"bitbucket.com/hyperboloide/horo/models/types/listing"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Group", func() {

	var group *Group

	It("should create a group", func() {
		group = &Group{
			OwnerId: Owner.Id,
			Name:    "test",
		}
		Ω(group.Insert()).To(BeNil())
		Ω(group.Url).ToNot(Equal(""))
	})

	It("should get a group", func() {
		g, err := ByUrl(group.Url)
		Ω(err).To(BeNil())
		Ω(g.Url).To(Equal(group.Url))
		Ω(g.OwnerId).To(Equal(Owner.Id))
		group = g

		Ω(group.GetOwner()).To(Equal(Owner))
	})

	It("should get api groups", func() {
		req := &listing.Request{
			Size: 50,
		}
		res, err := ApiByUser(Owner.Id, req)
		Ω(err).To(BeNil())
		By(spew.Sdump(res))
		Ω(res).ToNot(BeNil())
		Ω(res.Size).To(Equal(1))
		Ω(res.Results[0].(*GroupApi).Url).To(Equal(group.Url))
	})

})
