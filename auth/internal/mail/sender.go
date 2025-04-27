package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "sandbox.smtp.mailtrap.io"
	smtpServerAddress = "sandbox.smtp.mailtrap.io:2525"
)

type EmailSender interface {
	SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error
}

type MailTrapSender struct {
	name              string
	fromEmailUsername string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewMailTrappSender(fromEmailUsername string, name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &MailTrapSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailUsername: fromEmailUsername,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *MailTrapSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.fromEmailUsername, sender.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
