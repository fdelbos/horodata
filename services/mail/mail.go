package mail

import (
	"bytes"
	"fmt"
	"text/template"

	"dev.hyperboloide.com/fred/horodata/services/mail/messages"
	"github.com/russross/blackfriday"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

var (
	sender       string
	smtpHost     string
	smtpPort     int
	smtpUser     string
	smtpPassword string
)

func Configure() {
	sender = viper.GetString("mail_sender")
	smtpHost = viper.GetString("mail_smtp_host")
	smtpPort = viper.GetInt("mail_smtp_port")
	smtpUser = viper.GetString("mail_smtp_user")
	smtpPassword = viper.GetString("mail_smtp_password")
}

func NewMessage(subject, text, html string, to string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", text)
	m.AddAlternative("text/html", html)
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)
	return d.DialAndSend(m)
}

type File struct {
	Name   string `json:"name"`
	Bucket string `json:"bucket"`
}

type Mail struct {
	Dest     string      `json:"dest"`
	Subject  string      `json:"subject"`
	Template string      `json:"template"`
	Data     interface{} `json:"data"`
}

func (m Mail) Send() error {
	file := fmt.Sprintf("%s.md", m.Template)
	textBuff := &bytes.Buffer{}

	if b, err := messages.Asset(file); err != nil {
		return err
	} else if tmpl, err := template.New(file).Parse(string(b[:])); err != nil {
		return err
	} else if err := tmpl.Execute(textBuff, m.Data); err != nil {
		return err
	}

	html := blackfriday.MarkdownCommon(textBuff.Bytes())

	return NewMessage(
		m.Subject,
		string(textBuff.Bytes()[:]),
		string(html[:]),
		m.Dest)
	return nil

}
