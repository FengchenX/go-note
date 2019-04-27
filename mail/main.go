package main

import (
	"fmt"
	"gopkg.in/gomail.v1"
	"strconv"
)

func main() {

}

//info.From:发件地址或账号
//to:  收件地址
//body:邮件内容
//port:邮件服务器端口
////info.Host:邮件服务器地址
//info.Password:密码
func GoMail(to, body string) error {
	info := MailInfoGet()

	if info.Host == "" || info.Port == "" || info.From == "" || info.Password == "" || info.Header == "" || info.Footer == "" {
		fmt.Println("Pleaase Finish Mail Config!")
		return nil
	}
	port, _ := strconv.Atoi(info.Port)

	msg := gomail.NewMessage()
	msg.SetHeader("From", info.From)
	msg.SetHeader("To", to)
	//  msg.SetAddressHeader("Cc", "dan@example.com", "Dan")
	msg.SetHeader("Subject", info.Header)
	msg.SetBody("text/html", body)

	//
	mailer := gomail.NewMailer(info.Host, info.From, info.Password, port)
	if err := mailer.Send(msg); err != nil {
		fmt.Println("Has err:", err)
		return err
	}
	return nil
}

type MailInfo struct {
	Port     string
	From     string
	Header   string
	Password string
	Host     string
	Footer   string
}

func MailInfoGet() MailInfo {
	return MailInfo{
		Host:     "",
		Port:     "",
		From:     "",
		Header:   "",
		Password: "",
	}
}
