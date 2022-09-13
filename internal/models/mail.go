package models

import (
	"fmt"
	"net/smtp"
	"strings"
)

// SendMail 使用 SMTP + TLS 协议向用户发送邮件。
func SendMail(dest string, code int) bool {
	auth := smtp.PlainAuth(MailConf.Identity, MailConf.Username, MailConf.Password, MailConf.Host)
	to := []string{dest}
	nickname := "数学学习软件"
	user := MailConf.Username
	subject := "验证码"
	contentType := "Content-Type: text/plain; charset=UTF-8"
	body := fmt.Sprintf("您的数学学习软件注册验证码是%v，请妥善保管", code)
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(MailConf.Host+":465", auth, user, to, msg)
	if err != nil {
		fmt.Println("send mail error: ", err)
		return false
	}
	return true
}
