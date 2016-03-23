package group

import (
	"bitbucket.com/hyperboloide/horo/services/postgres"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"time"
)

func exportGenQuery(customer, creator *int64) string {
	const query = `
	select
	   jobs.created, jobs.duration, jobs.comment,
       tasks.name,
       customers.name,
       users.full_name
	from
        jobs
        join tasks on jobs.task_id = tasks.id
        join customers on jobs.customer_id = customers.id
        join guests on jobs.creator_id = guests.id
        join users on guests.user_id = users.id
	where
			jobs.group_id = $1
		and jobs.created > $2
		and jobs.created < $3
		%s
	order by jobs.id desc`

	if customer != nil && creator != nil {
		return fmt.Sprintf(query, "and customer_id = $6 and creator_id = $7")
	} else if customer != nil {
		return fmt.Sprintf(query, "and customer_id = $6")
	} else if creator != nil {
		return fmt.Sprintf(query, "and creator_id = $6")
	}
	return fmt.Sprintf(query, "")
}

type ExportLine struct {
	Created  time.Time
	Duration int64
	Comment  string
	Task     string
	Customer string
	Creator  string
}

func (el *ExportLine) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&el.Created,
		&el.Duration,
		&el.Comment,
		&el.Task,
		&el.Customer,
		&el.Creator)
}

func (el *ExportLine) ToLine() []string {
	return []string{
		el.Created.Format("2006-01-02"),
		el.Creator,
		el.Customer,
		el.Task,
		fmt.Sprintf("%.2f", float64(el.Duration)/3600.0),
		fmt.Sprintf("%d", el.Duration/60),
		el.Comment,
	}
}

func (g *Group) Export(w io.Writer, begin, end time.Time, customer, creator *int64) error {
	query := exportGenQuery(customer, creator)

	var rows *sql.Rows
	var err error
	if customer != nil && creator != nil {
		rows, err = postgres.DB().Query(query, g.Id, begin, end, *customer, *creator)
	} else if customer != nil {
		rows, err = postgres.DB().Query(query, g.Id, begin, end, *customer)
	} else if creator != nil {
		rows, err = postgres.DB().Query(query, g.Id, begin, end, *creator)
	} else {
		rows, err = postgres.DB().Query(query, g.Id, begin, end)
	}

	if err != nil {
		return err
	}
	defer rows.Close()

	cw := csv.NewWriter(w)
	cw.UseCRLF = true
	cw.Write([]string{
		"Date",
		"Utilisateur",
		"Dossier",
		"Tâche",
		"Durée (en heures)",
		"Durée (en minutes)",
		"Commentaire"})
	for rows.Next() {
		el := &ExportLine{}
		if err := el.Scan(rows.Scan); err != nil {
			return err
		} else if err := cw.Write(el.ToLine()); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	cw.Flush()
	return nil
}
