package models

import (
	"fmt"
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

// SendSms 用于向客户端发送验证码短信，返回是否成功。
func SendSms(phone string, code int) bool {
	config := sdk.NewConfig()
	credentials := credentials.NewAccessKeyCredential(SMSConf.AccessKeyID, SMSConf.AccessKeySecret)
	client, err := dysmsapi.NewClientWithOptions("cn-hangzhou", config, credentials)
	if err != nil {
		log.Printf("创建短信客户端失败：%v", err)
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.TemplateParam = "{\"code\":\"" + fmt.Sprint(code) + "\"}"
	request.SignName = "阿里云短信测试"
	request.TemplateCode = "SMS_154950909"
	response, err := client.SendSms(request)
	if err != nil {
		log.Printf("发送短信失败：%v", err)
		return false
	}
	log.Printf("response is %#v\n", response)
	return true
}
