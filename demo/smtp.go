package demo

import (
	"bytes"
	"log"
	"net/smtp"
)

func InitSmtp() {
	client, err := smtp.Dial("smtp.mxhichina.com:465")
	if err != nil {
		log.Fatal(err)
	}
	client.Mail("873725087@11.com")
	client.Rcpt("lishangkun@aliyun.com")
	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("This is the email body.")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
