package email

import (
	"log"
	"net/smtp"

	"github.com/Minhvn98/ecommerce-fashion/models"
)

// sender data
var from string = "khanh1tb98@gmail.com"
var password string = "mneohqyglldnmuvc"

// smtp - Simple Mail Transfer Protocol
var host string = "smtp.gmail.com"
var port string = "587"
var address string = host + ":" + port

func SendMail(receiver string, mail models.Mail) {
	subject := "Subject: " + mail.Subject + "\n"
	to := []string{receiver}
	message := []byte(subject + mail.Body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		log.Println("err:", err)
	}
}
