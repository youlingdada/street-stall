package utils

import (
	"crypto/tls"
	// "fmt"

	"github.com/youlingdada/street-stall/config"
	"gopkg.in/gomail.v2"
)

const (
	REG_EMAIL          = "REG_EMMAIL_"         // 用户注册时邮箱验证码前缀
	UPDATE_PWD_EMAIL   = "UPDATE_PWD_EMAIL_"   // 用户修改密码邮箱验证码前缀
	UPDATE_EMAIL_EMAIL = "UPDATE_INFO_EMAIL_"  // 用户更新信息邮箱验证码前缀
	UPDATE_PHONE_EMAIL = "UPDATE_PHONE_EMAIL_" // 用户更新电话号码邮箱验证码前缀
)

var Dialer *gomail.Dialer

type Email struct {
	From    string // 邮件发起者
	To      string // 邮件接受者
	Subject string // 邮件主题
	Context string // 邮件内容
}

func init() {
	app := config.GlobalConfig
	host, port := app.Email.Host, app.Email.Port
	username, password := app.Email.Username, app.Email.Password
	d := gomail.NewDialer(host, int(port), username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	Dialer = d
}

func SendEmail(email Email) error {
	app := config.GlobalConfig
	m := gomail.NewMessage()
	m.SetHeader("From", app.Email.Username)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Context)
	return Dialer.DialAndSend(m)
}
