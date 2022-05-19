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
	e.To = []string{"3224266014@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
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
