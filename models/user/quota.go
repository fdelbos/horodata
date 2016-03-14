package user

import (
	"bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/services/cache"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"database/sql"
	"fmt"
	"time"
)

const (
	QuotaFree   = "free"
	QuotaSmall  = "small"
	QuotaMedium = "medium"
	QuotaCustom = "custom"

	cacheQuota = "models.users.quota"
	cacheUsage = "models.users.usage"
)

var PlansLimits = map[string]Limits{
	QuotaFree: {
		Instances: 100,
		Forms:     10,
		Roles:     5,
		Files:     100 << 20, // 100 MO
	},
	QuotaSmall: {
		Instances: 2000,
		Forms:     50,
		Roles:     25,
		Files:     25 << 30, // 25 GO
	},
	QuotaMedium: {
		Instances: 10000,
		Forms:     100,
		Roles:     50,
		Files:     100 << 30, // 100 GO
	},
}

type Limits struct {
	Instances int64 `json:"instances"`
	Forms     int64 `json:"forms"`
	Roles     int64 `json:"roles"`
	Files     int64 `json:"files"`
}

func (l *Limits) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&l.Instances,
		&l.Forms,
		&l.Roles,
		&l.Files,
	)
}

type Bonus struct {
	Id          int64      `json:"id"`
	Created     *time.Time `json:"created"`
	Description string     `json:"description"`
	Limits
}

type Quota struct {
	Created *time.Time `json:"-"`
	Plan    string     `json:"plan"`
	Limits
}

func (q *Quota) saveInCache(userId int64) error {
	id := fmt.Sprintf("%d", userId)
	return cache.SetPackage(cacheQuota, id, q, time.Hour*2)
}

type Usage struct {
	Limits
}

func (u *Usage) saveInCache(userId int64) error {
	id := fmt.Sprintf("%d", userId)
	return cache.SetPackage(cacheUsage, id, u, time.Hour*2)
}

func (u *User) Bonus() ([]Bonus, error) {
	bonus := make([]Bonus, 0)
	query := `
    SELECT id, created, description, instances, forms, roles, files
    FROM quotas_bonus WHERE user_id = $1`

	rows, err := postgres.DB().Query(query, u.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Bonus
		if err := rows.Scan(
			&b.Id,
			&b.Created,
			&b.Description,
			&b.Instances,
			&b.Forms,
			&b.Roles,
			&b.Files); err != nil {
			return nil, err
		}
		bonus = append(bonus, b)
	}
	return bonus, rows.Err()
}

func (u *User) GetQuota() (*Quota, error) {
	quota := &Quota{}
	if err := cache.GetPackage(cacheQuota, fmt.Sprintf("%d", u.Id), quota); err == nil {
		return quota, nil
	}

	query := `SELECT created, plan FROM quotas WHERE user_id = $1`
	err := postgres.DB().QueryRow(query, u.Id).Scan(&quota.Created, &quota.Plan)
	if err == sql.ErrNoRows {
		return nil, errors.NotFound
	} else if err != nil {
		return nil, err
	}

	if quota.Plan != QuotaCustom {
		quota.Limits = PlansLimits[quota.Plan]
	} else {
		var l Limits
		query := `SELECT instances, forms roles, files FROM quotas_custom WHERE user_id = $1`
		if err := postgres.QueryRow(&l, query, u.Id); err != nil {
			return nil, err
		}
		quota.Limits = l
	}
	bonus, err := u.Bonus()
	if err != nil {
		return nil, err
	}

	for _, b := range bonus {
		quota.Instances += b.Instances
		quota.Forms += b.Forms
		quota.Roles += b.Roles
		quota.Files += b.Files
	}

	return quota, quota.saveInCache(u.Id)
}

func (u *User) GetUsage() (*Usage, error) {
	usage := &Usage{}
	if err := cache.GetPackage(cacheUsage, fmt.Sprintf("%d", u.Id), usage); err == nil {
		return usage, nil
	}
	query := `SELECT instances, forms, roles, files FROM usages WHERE user_id = $1`
	if err := postgres.QueryRow(usage, query, u.Id); err != nil {
		return nil, err
	}
	return usage, usage.saveInCache(u.Id)
}

func (u *User) AddBonus(desc string, l *Limits) error {
	query := `
    INSERT INTO quotas_bonus
    (user_id, description, instances, forms, roles, files)
    VALUES ($1, $2, $3, $4, $5, $6);`
	if stmt, err := postgres.DB().Prepare(query); err != nil {
		return err
	} else if _, err := stmt.Exec(
		u.Id,
		desc,
		l.Instances,
		l.Forms,
		l.Roles,
		l.Files); err != nil {
		return err
	}
	return cache.DelPackage(cacheQuota, fmt.Sprintf("%d", u.Id))
}

func (u *User) CanAddUsage(l *Limits) (bool, error) {
	if quota, err := u.GetQuota(); err != nil {
		return false, err
	} else if usage, err := u.GetUsage(); err != nil {
		return false, err
	} else {
		switch {
		case l.Instances+usage.Instances > quota.Instances:
			return false, nil
		case l.Forms+usage.Forms > quota.Forms:
			return false, nil
		case l.Roles+usage.Roles > quota.Roles:
			return false, nil
		case l.Files+usage.Files > quota.Files:
			return false, nil
		default:
			return true, nil
		}
	}
}

func (u *User) AddUsage(l *Limits) error {
	query := `
    UPDATE usages
    SET
        instances = instances + $2,
        forms = forms + $3,
        roles = roles + $4,
        files = files + $5
    WHERE user_id = $1`

	if err := postgres.Exec(query, u.Id, l.Instances, l.Forms, l.Roles, l.Files); err != nil {
		return err
	}
	return cache.DelPackage(cacheUsage, fmt.Sprintf("%d", u.Id))
}
