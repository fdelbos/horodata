package group

import (
	"bitbucket.com/hyperboloide/horo/models/errors"
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"time"
)

type Task struct {
	Id               int64     `json:"id"`
	Created          time.Time `json:"created"`
	Active           bool      `json:"active"`
	GroupId          int64     `json:"group_id"`
	Name             string    `json:"name"`
	CommentMandatory bool      `json:"comment_mandatory"`
}

func (t *Task) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&t.Id,
		&t.Created,
		&t.Active,
		&t.GroupId,
		&t.Name,
		&t.CommentMandatory)
}

func (t *Task) Update() error {
	const query = `
	update tasks
	set active = $2, name = $3, comment_mandatory = $4
	where id = $1;`
	return postgres.Exec(query, t.Id, t.Active, t.Name, t.CommentMandatory)
}

func (g *Group) TaskAdd(name string, cm bool) error {
	task := &Task{}
	const findQuery = `
    select * from tasks where group_id = $1 and name = $2;`

	if err := postgres.QueryRow(task, findQuery, g.Id, name); err == nil {
		task.Active = true
		task.CommentMandatory = cm
		return task.Update()
	} else if err != errors.NotFound {
		return err
	}

	const insertQuery = `
	insert into tasks (group_id, name, comment_mandatory)
	values ($1, $2, $3);`

	return postgres.Exec(insertQuery, g.Id, name, cm)
}

func (g *Group) Tasks() ([]Task, error) {
	const query = `
    select * from tasks
    where group_id = $1 and active = true
	order by name;`

	rows, err := postgres.DB().Query(query, g.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []Task{}
	for rows.Next() {
		i := Task{}
		if err := i.Scan(rows.Scan); err != nil {
			return nil, err
		}
		results = append(results, i)
	}
	return results, rows.Err()
}
