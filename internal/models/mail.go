package models

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

// SendMail 使用 SMTP + TLS 协议向用户发送邮件。
func SendMail(dest string, code int) bool {
	m := gomail.NewMessage()
	m.SetHeader("From", MailConf.Username)
	m.SetHeader("To", dest)
	m.SetHeader("Subject", "验证码")
	m.SetBody("text/html", fmt.Sprintf("您的数学学习软件注册验证码是<br><h1>%v</h1><br>请妥善保管", code))
	d := gomail.NewDialer(MailConf.Host, 465, MailConf.Username, MailConf.Password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
