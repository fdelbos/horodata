package user

import (
	"time"

	"dev.hyperboloide.com/fred/horodata/services/mail"
	"dev.hyperboloide.com/fred/horodata/services/postgres"
	"github.com/dchest/uniuri"
	"github.com/hyperboloide/qmail/client"
)

type PasswordRequest struct {
	Id      int64     `json:"id"`
	Created time.Time `json:"created"`
	UserId  int64     `json:"user_id"`
	Active  bool      `json:"active"`
	Url     string    `json:"url"`
}

func (pr *PasswordRequest) Scan(scanFn func(dest ...interface{}) error) error {
	return scanFn(
		&pr.Id,
		&pr.Created,
		&pr.UserId,
		&pr.Active,
		&pr.Url)
}

func (pr *PasswordRequest) GetUser() (*User, error) {
	return ById(pr.UserId)
}

func (pr *PasswordRequest) IsValid() bool {
	if pr.Created.Add(time.Hour * 2).Before(time.Now()) {
		return false
	}
	return pr.Active
}

func (pr *PasswordRequest) Invalidate() error {
	const query = `UPDATE password_requests SET active = false WHERE id = $1;`
	return postgres.Exec(query, pr.Id)
}

func (u *User) NewPasswordRequest() error {
	pr := &PasswordRequest{
		Created: time.Now(),
		UserId:  u.Id,
		Active:  true,
		Url:     uniuri.NewLen(40),
	}
	const query = `
    INSERT INTO password_requests (
    	created,
    	user_id,
    	active,
    	url)
	VALUES ($1, $2, $3, $4);`
	if err := postgres.Exec(query, pr.Created, pr.UserId, pr.Active, pr.Url); err != nil {
		return err
	}

	return mail.Mailer().Send(client.Mail{
		Dests:    []string{u.Email},
		Subject:  "Réinitialisation du mot de passe sur Horodata",
		Template: "reset_password",
		Data: map[string]interface{}{
			"link": pr.Url,
		},
	})
}

func GetPasswordRequest(url string) (*PasswordRequest, error) {
	pr := &PasswordRequest{}
	query := `
	SELECT id, created, user_id, active, url
	FROM password_requests
	WHERE url = $1;`
	return pr, postgres.QueryRow(pr, query, url)
}
