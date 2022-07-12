package gomail

import (
	"crypto/rand"
	"encoding/hex"

	"gopkg.in/gomail.v2"
)

type Gomail struct {
	Email    string
	Password string
}

func NewGomail(Email, Password string) *Gomail {
	return &Gomail{Email: Email, Password: Password}
}

func (mail *Gomail) SendMail(to, cc, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mail.Email)
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", to, cc)
	m.SetHeader("Subject", "Hello "+cc)
	m.SetBody("text/html", message)

	d := gomail.NewDialer("smtp.gmail.com", 587, mail.Email, mail.Password)

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}

func (mail *Gomail) GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
