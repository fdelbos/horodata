package group

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"time"

	"dev.hyperboloide.com/fred/horodata/services/postgres"
	"github.com/tealeg/xlsx"
)

const query = `
select
   jobs.created, jobs.duration, jobs.comment,
   tasks.name,
   customers.name,
   users.full_name,
   ((jobs.duration * guests.rate) / 360000::float)::numeric(9,2) as cost
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

func exportGenQuery(customer, creator *int64) string {
	const query = `
	select
	   jobs.created, jobs.duration, jobs.comment,
       tasks.name,
       customers.name,
       users.full_name,
	   ((jobs.duration * guests.rate) / 360000::float)::numeric(9,2) as cost
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

func (g *Group) exportMakeQuery(begin, end time.Time, customer, creator *int64) (*sql.Rows, error) {
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
	return rows, err
}

func (g *Group) ExportCSV(w io.Writer, begin, end time.Time, customer, creator *int64) error {
	rows, err := g.exportMakeQuery(begin, end, customer, creator)
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

func (g *Group) ExportXLSX(w io.Writer, begin, end time.Time, customer, creator *int64) error {
	rows, err := g.exportMakeQuery(begin, end, customer, creator)
	if err != nil {
		return err
	}
	defer rows.Close()

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Saisies")
	if err != nil {
		return err
	}

	header := sheet.AddRow()
	header.AddCell().SetString("Date")
	header.AddCell().SetString("Utilisateur")
	header.AddCell().SetString("Dossier")
	header.AddCell().SetString("Tâche")
	header.AddCell().SetString("Durée (en heures)")
	header.AddCell().SetString("Durée (en minutes)")
	header.AddCell().SetString("Coût")
	header.AddCell().SetString("Commentaire")

	for rows.Next() {
		el := &ExportLine{}
		if err := el.Scan(rows.Scan); err != nil {
			return err
		}
		line := sheet.AddRow()
		line.AddCell().SetDate(el.Created)
		line.AddCell().SetString(el.Creator)
		line.AddCell().SetString(el.Customer)
		line.AddCell().SetString(el.Task)
		line.AddCell().SetFloatWithFormat(float64(el.Duration)/3600.0, "0.00")
		line.AddCell().SetInt64(el.Duration / 60)
		line.AddCell().SetFloatWithFormat(el.Cost, "0.00")
		line.AddCell().SetString(el.Comment)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return file.Write(w)
}
