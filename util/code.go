package util

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"goCommunity/define"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func GetCode() string {
	rand.Seed(time.Now().UnixNano())
	res := ""
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}

func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "goCommunity <godhelptome@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码"
	e.HTML = []byte("<b>" + code + "</b>" + "验证码1分钟内有效，请及时处理。如果不是您在操作请忽略！")
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "godhelptome@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}
