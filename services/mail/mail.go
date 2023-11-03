package mail

import (
	"fmt"
	"os"

	"github.com/mailgun/mailgun-go"
)

func SendEmail(to, subject, text string) (string, error) {

	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")

	from := fmt.Sprintf("Excited User <mail@%s>", domain)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, text, to)
	_, id, err := mg.Send(m)
	return id, err
}
