package group_test

import (
	"bitbucket.com/hyperboloide/horo/helpers/tests"
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

	It("should test customers", func() {
		Ω(group.CustomerAdd("toto")).To(BeNil())
		customers, err := group.Customers()
		Ω(err).To(BeNil())
		Ω(len(customers)).To(Equal(1))
		c := customers[0]
		Ω(c.Name).To(Equal("toto"))
		c.Name = "titi"
		Ω(c.Update()).To(BeNil())

		customers, err = group.Customers()
		Ω(err).To(BeNil())
		Ω(len(customers)).To(Equal(1))
		c = customers[0]
		Ω(c.Name).To(Equal("titi"))

		id := c.Id
		c.Active = false
		Ω(c.Update()).To(BeNil())
		customers, err = group.Customers()
		Ω(err).To(BeNil())
		Ω(len(customers)).To(Equal(0))

		Ω(group.CustomerAdd("titi")).To(BeNil())
		customers, err = group.Customers()
		Ω(err).To(BeNil())
		Ω(len(customers)).To(Equal(1))
		c = customers[0]
		Ω(c.Name).To(Equal("titi"))
		Ω(c.Id).To(Equal(id))
	})

	It("should test tasks", func() {
		Ω(group.TaskAdd("toto", false)).To(BeNil())
		tasks, err := group.Tasks()
		Ω(err).To(BeNil())
		Ω(len(tasks)).To(Equal(1))
		t := tasks[0]
		Ω(t.Name).To(Equal("toto"))
		Ω(t.CommentMandatory).To(BeFalse())
		t.Name = "tata"
		t.CommentMandatory = true
		Ω(t.Update()).To(BeNil())

		tasks, err = group.Tasks()
		Ω(err).To(BeNil())
		Ω(len(tasks)).To(Equal(1))
		t = tasks[0]
		Ω(t.Name).To(Equal("tata"))
		Ω(t.CommentMandatory).To(BeTrue())

		id := t.Id
		t.Active = false
		Ω(t.Update()).To(BeNil())
		tasks, err = group.Tasks()
		Ω(err).To(BeNil())
		Ω(len(tasks)).To(Equal(0))

		Ω(group.TaskAdd("tata", false)).To(BeNil())
		tasks, err = group.Tasks()
		Ω(err).To(BeNil())
		Ω(len(tasks)).To(Equal(1))
		t = tasks[0]
		Ω(t.Name).To(Equal("tata"))
		Ω(t.CommentMandatory).To(BeFalse())
		Ω(t.Id).To(Equal(id))

		tt, err := group.TaskGet(id)
		Ω(err).To(BeNil())
		Ω(tt.Name).To(Equal("tata"))
		Ω(tt.CommentMandatory).To(BeFalse())
		Ω(tt.Id).To(Equal(id))
	})

	It("should test guests", func() {
		guestUser, err := tests.NewUser()
		Ω(err).To(BeNil())

		Ω(group.GuestAdd(guestUser.Email, "some message", 4242, false, true)).To(BeNil())
		guests, err := group.Guests()
		Ω(err).To(BeNil())
		Ω(len(guests)).To(Equal(1))
		g := guests[0]
		Ω(g.Email).To(Equal(guestUser.Email))
		Ω(g.UserId).To(Equal(guestUser.Id))
		Ω(g.Admin).To(BeFalse())

		g.Admin = true
		Ω(g.Update()).To(BeNil())

		g2, err := group.GuestGetById(g.Id)
		Ω(err).To(BeNil())
		Ω(g2.Admin).To(BeTrue())

		g.Active = false
		Ω(g.Update()).To(BeNil())

		_, err = group.GuestGetById(g.Id)
		Ω(err).ToNot(BeNil())

		Ω(group.GuestAdd(guestUser.Email, "some message", 4242, false, false)).To(BeNil())
		g3, err := group.GuestGetById(g.Id)
		Ω(err).To(BeNil())
		Ω(g3.Id).To(Equal(g.Id))
	})

	It("should test ApiDetail", func() {
		d, err := group.ApiDetail(false)
		Ω(err).To(BeNil())
		Ω(d).ToNot(BeNil())

		d, err = group.ApiDetail(true)
		Ω(err).To(BeNil())
		Ω(d).ToNot(BeNil())
	})
})
