package utils

import (
	"os"

	"gopkg.in/gomail.v2"
)

// SendNotificationEmail sends a plain text email notification.
func SendNotificationEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("FORUM_NOTIFICATION_EMAIL_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Set up the SMTP dialer with TLS
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))

	// Use StartTLS (default behavior on port 587)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// SendHTMLNotificationEmail sends an HTML formatted email notification.
func SendHTMLNotificationEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("FORUM_NOTIFICATION_EMAIL_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	// Set up the SMTP dialer with TLS

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))
	// Use StartTLS (default behavior on port 587)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
