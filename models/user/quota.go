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

	cacheQuotas    = "models.users.quotas"
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

type Quotas struct {
	Created *time.Time `json:"-"`
	Plan    string     `json:"plan"`
	Limits  Limits     `json:"limits"`
}

func (q *Quotas) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&q.Created,
		&q.Plan)
}

func (u User) Quotas() (*Quotas, error) {
	quotas := &Quotas{}

	if err := cache.GetPackage(cacheQuotas, fmt.Sprintf("%d", u.Id), quotas); err == nil {
		return quotas, nil
	}

	const query = `
	select created, plan
	from quotas
	where user_id = $1;`

	if err := postgres.QueryRow(quotas, query, u.Id); err != nil {
		return nil, err
	}
	quotas.Limits = PlansLimits[quotas.Plan]
	return quotas, quotas.saveInCache(u.Id)
}

func (u User) QuotaChangePlan(plan string) error {
	const query = `
	update quotas
	set plan = $2
	where user_id = $1;`

	if err := postgres.Exec(query, u.Id, plan); err != nil {
		return err
	}
	cache.DelPackage(cacheQuotas, fmt.Sprintf("%d", u.Id))
	return u.UsageJobsReset()
}

func (q *Quotas) saveInCache(userId int64) error {
	id := fmt.Sprintf("%d", userId)
	return cache.SetPackage(cacheQuotas, id, q, time.Hour*4)
}

func (u User) UsageGroups() (int, error) {
	const query = `
	select count(id) from groups_active where owner_id = $1;`
	var count int
	return count, postgres.DB().QueryRow(query, u.Id).Scan(&count)
}

func (u User) QuotaCanAddGroup() (bool, error) {
	quotas, err := u.Quotas()
	if err != nil {
		return false, err
	}
	usage, err := u.UsageGroups()
	return usage < quotas.Limits.Groups, err
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
	quotas, err := u.Quotas()
	if err != nil {
		return false, err
	}
	usage, err := u.UsageGuests()
	return usage < quotas.Limits.Guests, err
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
	quotas, err := u.Quotas()
	if err != nil {
		return false, err
	}
	usage, err := u.UsageJobs()
	return usage < quotas.Limits.Jobs, err
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
