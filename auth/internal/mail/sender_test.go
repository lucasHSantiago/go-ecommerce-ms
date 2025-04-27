package mail

import (
	"testing"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithMailTrap(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("../../")
	require.NoError(t, err)

	sender := NewMailTrappSender(config.EmailSenderUsername, config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	subject := "a test email"
	content := `
	<h1>Hello World!</h1>
	<p>This is a teste email</p>
	`
	to := []string{"to@example.com"}
	attachFiles := []string{"../../../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
