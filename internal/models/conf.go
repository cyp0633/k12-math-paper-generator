package models

import (
	"log"

	"github.com/go-ini/ini"
)

// MailConf 是用于邮件发送的配置。
var MailConf = struct {
	Identity string
	Username string
	Password string
	Host     string
}{}

// SMSConf 使用 Access Key 和 Access ID 配置阿里云短信 SDK。
var SMSConf = struct {
	AccessKeyID     string
	AccessKeySecret string
}{}

// ServerConf 配置服务器的端口号等信息。
var ServerConf = struct {
	Port string
}{}

// InitConf 使用文件夹内的 conf.ini 文件初始化配置。CLI 无需此操作。
func InitConf() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		log.Printf("配置文件加载失败：%v", err)
	}
	mailSection := cfg.Section("mail")
	{
		MailConf.Identity = mailSection.Key("identity").String()
		MailConf.Username = mailSection.Key("username").String()
		MailConf.Password = mailSection.Key("password").String()
		MailConf.Host = mailSection.Key("host").String()
	}
	smsSection := cfg.Section("sms")
	{
		SMSConf.AccessKeyID = smsSection.Key("access_key_id").String()
		SMSConf.AccessKeySecret = smsSection.Key("access_key_secret").String()
	}
	serverSection := cfg.Section("server")
	{
		ServerConf.Port = serverSection.Key("port").String()
	}
}
