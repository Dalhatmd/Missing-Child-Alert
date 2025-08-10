package notification

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

type Notification struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}

// EmailConfig holds SMTP configuration loaded from environment variables.
type EmailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

func loadEmailConfig() (EmailConfig, error) {
	cfg := EmailConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USER"),
		Password: os.Getenv("SMTP_PASS"),
		From:     os.Getenv("SMTP_FROM"),
	}
	if cfg.Host == "" || cfg.Port == "" {
		return cfg, errors.New("SMTP_HOST and SMTP_PORT must be set")
	}
	// Use Username as From if From not provided
	if cfg.From == "" {
		cfg.From = cfg.Username
	}
	return cfg, nil
}

// SendEmail sends an HTML email to one or more recipients using SMTP settings from environment variables.
// Required env vars: SMTP_HOST, SMTP_PORT; Optional: SMTP_USER, SMTP_PASS, SMTP_FROM.
func SendEmail(to []string, subject, htmlBody string) error {
	if len(to) == 0 {
		return errors.New("no recipients provided")
	}
	cfg, err := loadEmailConfig()
	if err != nil {
		return err
	}

	// Build MIME message
	headers := map[string]string{
		"From":         cfg.From,
		"To":           strings.Join(to, ", "),
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=\"UTF-8\"",
	}
	var msgBuilder strings.Builder
	for k, v := range headers {
		msgBuilder.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msgBuilder.WriteString("\r\n")
	msgBuilder.WriteString(htmlBody)
	msg := []byte(msgBuilder.String())

	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	// Auth for dev servers
	var auth smtp.Auth
	if cfg.Username != "" {
		auth = smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	}

	return smtp.SendMail(addr, auth, cfg.From, to, msg)
}

// Example convenience helper to format a basic alert email body.
func BuildAlertEmailBody(childName, lastSeenLocation, description, reporterContact string) string {
	return fmt.Sprintf(`
		<h2 style="margin:0 0 12px;color:#1e40af;">Missing Child Alert</h2>
		<p><strong>Child:</strong> %s</p>
		<p><strong>Last seen at:</strong> %s</p>
		<p><strong>Description:</strong> %s</p>
		<p><strong>Reporter contact:</strong> %s</p>
	`,
		childName,
		lastSeenLocation,
		description,
		reporterContact,
	)
}
