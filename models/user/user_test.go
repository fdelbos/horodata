package user_test

import (
	"bitbucket.com/hyperboloide/horo/models/errors"
	. "bitbucket.com/hyperboloide/horo/models/user"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"github.com/dchest/uniuri"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {

	user := &User{
		Login: uniuri.NewLen(12),
		Email: uniuri.NewLen(12) + "@test.com",
	}

	It("should create a user", func() {
		Ω(user.Insert()).To(BeNil())
	})

	It("should find user", func() {
		u, err := ByLogin(user.Login)
		Ω(err).To(BeNil())
		Ω(u.Id).To(Equal(user.Id))

		u, err = ByEmail(user.Email)
		Ω(err).To(BeNil())
		Ω(u.Login).To(Equal(user.Login))

		u, err = ById(u.Id)
		Ω(err).To(BeNil())
		Ω(u.Id).To(Equal(user.Id))

		// again for cache
		u, err = ById(u.Id)
		Ω(err).To(BeNil())
		Ω(u.Login).To(Equal(user.Login))
	})

	It("should update password", func() {
		err := user.UpdatePassword("new password")
		Ω(err).To(BeNil())
		user, err = ByLogin(user.Login)
		ok, err := user.CheckPassword("new password")
		Ω(ok).To(BeTrue())
		Ω(err).To(BeNil())
		ok, err = user.CheckPassword("wrong")
		Ω(ok).To(BeFalse())
		Ω(err).To(BeNil())
	})

	It("should find user", func() {
		u, err := ByLogin(user.Login)
		Ω(err).To(BeNil())
		Ω(u.Login).To(Equal(user.Login))

		u, err = ByEmail(user.Email)
		Ω(err).To(BeNil())
		Ω(u.Login).To(Equal(user.Login))

		u, err = ById(u.Id)
		Ω(err).To(BeNil())
		Ω(u.Login).To(Equal(user.Login))
	})

	It("should work with quotas", func() {
		u, err := ByLogin(user.Login)
		Ω(err).To(BeNil())

		quota, err := u.GetQuota()
		Ω(err).To(BeNil())
		Ω(quota).ToNot(BeNil())
		Ω(quota.Forms).To(Equal(PlansLimits["free"].Forms))

		u.AddBonus("test", &Limits{Forms: 2})
		quota, err = u.GetQuota()
		Ω(err).To(BeNil())
		Ω(quota).ToNot(BeNil())
		Ω(quota.Forms).To(Equal(PlansLimits["free"].Forms + 2))
	})

	It("should work with usage", func() {
		u, err := ByLogin(user.Login)
		Ω(err).To(BeNil())

		usage, err := u.GetUsage()
		Ω(err).To(BeNil())
		Ω(usage).ToNot(BeNil())
		Ω(usage.Forms).To(Equal(int64(0)))

		ok, err := u.CanAddUsage(&Limits{Forms: 1})
		Ω(err).To(BeNil())
		Ω(ok).To(BeTrue())

		err = u.AddUsage(&Limits{Forms: 1000})
		Ω(err).To(BeNil())

		usage, err = u.GetUsage()
		Ω(err).To(BeNil())
		Ω(usage).ToNot(BeNil())
		Ω(usage.Forms).To(Equal(int64(1000)))

		ok, err = u.CanAddUsage(&Limits{Forms: 1})
		Ω(err).To(BeNil())
		Ω(ok).To(BeFalse())
	})

	It("should gen reset password request", func() {
		u, err := ByLogin(user.Login)
		Ω(err).To(BeNil())
		Ω(u.NewPasswordRequest()).To(BeNil())
	})

	It("should work with a session", func() {
		user, err := ByLogin(user.Login)
		Ω(err).To(BeNil())

		id, err := NewSession(user, "localhost")
		Ω(err).To(BeNil())

		By(id)

		session, err := GetSession(id)
		Ω(err).To(BeNil())
		Ω(session).ToNot(BeNil())
		Ω(session.UserId).To(Equal(user.Id))
		Ω(session.Active).To(BeTrue())

		u, err := session.GetUser()
		Ω(session.UserId).To(Equal(u.Id))

		err = session.Close()
		Ω(err).To(BeNil())

		session, err = GetSession(id)
		Ω(err).To(BeNil())
		Ω(session.Active).To(BeFalse())
	})

	It("should update profile", func() {
		user, err := ByLogin(user.Login)
		Ω(err).To(BeNil())
		user.FullName = "Test User"
		user.Organization = "test llc"
		user.Website = "test.com"
		user.About = "somthing about me..."
		Ω(user.UpdateProfile()).To(BeNil())

		user2, err := ByLogin(user.Login)
		Ω(err).To(BeNil())
		Ω(user2.FullName).To(Equal("Test User"))
		Ω(user2.Organization).To(Equal("test llc"))
		Ω(user2.Website).To(Equal("test.com"))
		Ω(user2.About).To(Equal("somthing about me..."))
	})

	It("should clean", func() {
		postgres.DB().Exec("DELETE FROM users WHERE login = $1", user.Login)
		_, err := ByLogin(user.Login)
		Ω(err).To(Equal(errors.NotFound))
	})

})
