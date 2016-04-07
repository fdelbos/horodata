package group

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"

	"dev.hyperboloide.com/fred/horodata/services/postgres"
)

const query = `
select
   jobs.created, jobs.duration, jobs.comment,
   tasks.name,
   customers.name,
   users.full_name,
   ((jobs.duration * guests.rate) / 360000::float) as cost
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
order by jobs.id desc;`

type ExportLine struct {
	Created  time.Time `json:"created"`
	Duration int64     `json:"duration"`
	Comment  string    `json:"comment"`
	Task     string    `json:"task"`
	Customer string    `json:"customer"`
	Creator  string    `json:"creator"`
	Cost     float64   `json:"cost"`
}

func (el *ExportLine) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&el.Created,
		&el.Duration,
		&el.Comment,
		&el.Task,
		&el.Customer,
		&el.Creator,
		&el.Cost)
}

func (el *ExportLine) ToCSVLine() []string {
	return []string{
		el.Created.Format("2006-01-02"),
		el.Creator,
		el.Customer,
		el.Task,
		fmt.Sprintf("%.2f", float64(el.Duration)/3600.0),
		fmt.Sprintf("%d", el.Duration/60),
		fmt.Sprintf("%.2f", el.Cost),
		el.Comment,
	}
}

func (g *Group) ExportCSV(w io.Writer, begin, end time.Time) error {
	rows, err := postgres.DB().Query(query, g.Id, begin, end)
	if err != nil {
		return err
	}
	defer rows.Close()

	cw := csv.NewWriter(w)
	cw.Write([]string{
		"Date",
		"Utilisateur",
		"Dossier",
		"Tâche",
		"Durée (en heures)",
		"Durée (en minutes)",
		"Coût",
		"Commentaire"})
	for rows.Next() {
		el := &ExportLine{}
		if err := el.Scan(rows.Scan); err != nil {
			return err
		} else if err := cw.Write(el.ToCSVLine()); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	cw.Flush()
	return nil
}

func (g *Group) ExportStruct(begin, end time.Time) ([]ExportLine, error) {
	result := make([]ExportLine, 0)

	rows, err := postgres.DB().Query(query, g.Id, begin, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		el := &ExportLine{}
		if err := el.Scan(rows.Scan); err != nil {
			return nil, err
		}
		result = append(result, *el)
	}
	return result, rows.Err()
}
