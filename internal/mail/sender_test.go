package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teamkweku/code-odessey/config"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := config.LoadConfig("../../.env")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="teamkuuku@gmail.com">Tech School</a></p>
	`
	to := []string{"teamkuuku@gmail.com"}
	attachFiles := []string{"../../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
