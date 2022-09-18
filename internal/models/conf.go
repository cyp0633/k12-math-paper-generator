package models

import (
	"log"

	"github.com/go-ini/ini"
)

// MailConf 是用于邮件发送的配置。
var MailConf = struct {
	Identity string // 发件人身份，默认留空即可
	Username string // 发件人邮箱
	Password string // 发件人邮箱密码
	Host     string // SMTP 服务器地址
}{}

// SMSConf 使用 Access Key 和 Access ID 配置阿里云短信 SDK。
var SMSConf = struct {
	AccessKeyID     string // 阿里云短信服务的 Access Key ID
	AccessKeySecret string // 阿里云短信服务的 Access Key Secret
}{}

// ServerConf 配置服务器的端口号等信息。
var ServerConf = struct {
	Port string // 服务器端口号
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
