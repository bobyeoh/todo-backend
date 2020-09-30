package utils

import (
	"crypto/tls"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// SendMail godoc
func SendMail(to string, title string, content string) error {
	server := os.Getenv("MAIL_SERVER")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	// get server params
	message := gomail.NewMessage()
	message.SetHeader("From", username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", title)
	message.SetBody("text/html", content)
	dialer := gomail.NewDialer(server, port, username, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return dialer.DialAndSend(message)
}
