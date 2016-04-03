package user

import (
	"fmt"
	"time"

	"dev.hyperboloide.com/fred/horodata/services/cache"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
)

const (
	QuotaFree   = "free"
	QuotaSmall  = "small"
	QuotaMedium = "medium"
	QuotaLarge  = "large"

	cacheQuota     = "models.users.quota"
	cacheJobsUsage = "models.users.jobs_usage"
)

type Limits struct {
	Jobs   int `json:"jobs"`
	Guests int `json:"guests"`
	Groups int `json:"groups"`
}

var PlansLimits = map[string]Limits{
	QuotaFree: {
		Jobs:   15,
		Guests: 2,
		Groups: 1,
	},
	QuotaSmall: {
		Jobs:   500,
		Guests: 10,
		Groups: 2,
	},
	QuotaMedium: {
		Jobs:   1500,
		Guests: 30,
		Groups: 5,
	},
	QuotaLarge: {
		Jobs:   5000,
		Guests: 100,
		Groups: 15,
	},
}

type Quota struct {
	Created *time.Time `json:"-"`
	Plan    string     `json:"plan"`
	Limits  Limits     `json:"limits"`
}

func (q *Quota) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&q.Created,
		&q.Plan)
}

func (u User) Quota() (*Quota, error) {
	quota := &Quota{}

	if err := cache.GetPackage(cacheQuota, fmt.Sprintf("%d", u.Id), quota); err == nil {
		return quota, nil
	}

	const query = `
	select created, plan
	from quotas
	where user_id = $1;`

	if err := postgres.QueryRow(quota, query, u.Id); err != nil {
		return nil, err
	}
	quota.Limits = PlansLimits[quota.Plan]
	return quota, quota.saveInCache(u.Id)
}

func (q *Quota) saveInCache(userId int64) error {
	id := fmt.Sprintf("%d", userId)
	return cache.SetPackage(cacheQuota, id, q, time.Hour*4)
}

func (u User) UsageGroups() (int, error) {
	const query = `
	select count(id) from groups_active where owner_id = $1;`
	var count int
	return count, postgres.DB().QueryRow(query, u.Id).Scan(&count)
}

func (u User) QuotaCanAddGroup() (bool, error) {
	quota, err := u.Quota()
	if err != nil {
		return false, err
	}
	usage, err := u.UsageGroups()
	return usage < quota.Limits.Groups, err
}

func (u User) UsageGuests() (int, error) {
	const query = `
	select count(*)
	from (
		select distinct email
		from guests
		where active = true and group_id in (
			select id from groups_active where owner_id = $1
		)
	) as temp;`
	var count int
	return count, postgres.DB().QueryRow(query, u.Id).Scan(&count)
}

func (u User) QuotaCanAddGuest() (bool, error) {
	quota, err := u.Quota()
	if err != nil {
		return false, err
	}
	usage, err := u.UsageGuests()
	return usage < quota.Limits.Guests, err
}

func (u User) usageJobsKey() string {
	return fmt.Sprintf("%d.%d",
		u.Id,
		time.Now().YearDay())
}

func (u User) UsageJobs() (int, error) {
	var count int
	if err := cache.GetPackage(cacheJobsUsage, u.usageJobsKey(), &count); err == nil {
		return count, nil
	}
	const query = `
	select count(id) from jobs
	where created > current_date and group_id in (
		select id from groups_active where owner_id = $1
	);`

	if err := postgres.DB().QueryRow(query, u.Id).Scan(&count); err != nil {
		return count, err
	}
	return count, cache.SetPackage(cacheJobsUsage, u.usageJobsKey(), count, time.Hour)
}

func (u User) QuotaCanAddJob() (bool, error) {
	quota, err := u.Quota()
	if err != nil {
		return false, err
	}
	usage, err := u.UsageJobs()
	return usage < quota.Limits.Jobs, err
}

func (u User) UsageJobsIncr() error {
	return cache.IncrPackage(cacheJobsUsage, u.usageJobsKey())
}

func (u User) UsageJobsDecr() error {
	// for side effect in case not in cache
	if _, err := u.UsageJobs(); err != nil {
		return err
	}
	return cache.DecrPackage(cacheJobsUsage, u.usageJobsKey())
}

func (u User) UsageJobsReset() error {
	return cache.DelPackage(cacheJobsUsage, u.usageJobsKey())
}
