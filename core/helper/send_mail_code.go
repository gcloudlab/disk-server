package helper

import (
	"crypto/tls"
	"gcloud/core/define"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

// SendMailCode
// 邮箱验证码发送
func SendMailCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "GCloud <gcloud2yesmore@163.com>"
	e.To = []string{mail}
	e.Subject = "欢迎加入 GCloud"
	e.HTML = []byte("您正在注册GCloud云盘账号，请确保为本人操作。验证码将在5分钟后失效，请及时完成注册，验证码：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "gcloud2yesmore@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})

	if err != nil {
		return err
	}
	return nil
}

// 生成随机验证码
func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
