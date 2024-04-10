package demo

import (
	"log"
	"net/smtp"
)

func InitSmtpAuth() {
	auth := smtp.PlainAuth("", "lishangkun@aliyun.com", "****", "mail.aliyun.com")

	err := smtp.SendMail("mail.aliyun.com:25", auth, "lishangkun@aliyun.com", []string{"873725087@qq.com"}, []byte("This is the email body."))
	if err != nil {
		log.Fatal(err)
	}
}
