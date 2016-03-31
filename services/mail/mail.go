package mail

import (
	"dev.hyperboloide.com/fred/horodata/services/mail/messages"
	"bytes"
	"fmt"
	mailgun "github.com/mailgun/mailgun-go"
	"github.com/russross/blackfriday"
	"github.com/spf13/viper"
	"text/template"
)

func newMailer() mailgun.Mailgun {
	return mailgun.NewMailgun(
		viper.GetString("mail_domain"),
		viper.GetString("mail_key"),
		"")
}

func NewMessage(subject, text, html string, to ...string) error {
	mailer := newMailer()
	msg := mailer.NewMessage(
		viper.GetString("mail_sender"),
		subject,
		text,
		to...)
	if html != "" {
		msg.SetHtml(html)
	}
	_, _, err := mailer.Send(msg)
	return err
}

type File struct {
	Name   string `json:"name"`
	Bucket string `json:"bucket"`
}

type Mail struct {
	Dests    []string    `json:"dest"`
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
		m.Dests...)

}
