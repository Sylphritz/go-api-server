package email

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/sylphritz/go-api-server/pkg/config"
)

func getEmailAuth() smtp.Auth {
	appConfig := config.GetConfig()

	return smtp.PlainAuth("", appConfig.SmtpUsername, appConfig.SmtpPassword, appConfig.SmtpHost)
}

func SendTestEmail() {
	appConfig := config.GetConfig()

	log.Println("Sending test email...")
	log.Println("via: " + appConfig.SmtpHost + ":" + fmt.Sprintf("%v", appConfig.SmtpPort))

	err := smtp.SendMail(
		appConfig.SmtpHost+":"+fmt.Sprintf("%v", appConfig.SmtpPort),
		getEmailAuth(),
		"dear123000@gmail.com",
		[]string{"sylphritz@gmail.com"},
		[]byte("From: dear123000@gmail.com\r\n"+
			"To: sylphritz@gmail.com\r\n"+
			"Subject: Did you get the email?\r\n"+
			"\r\n"+
			"Thanks for your help!\r\n",
		),
	)

	if err != nil {
		log.Fatalf("Error sending test email: %v", err)
	}

	log.Println("Successfully sending test email")
}
