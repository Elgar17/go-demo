package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "mineemail@qq.com")
	m.SetHeader("To", "test@163.com") //主送(可以发送给多个人)
	// m.SetHeader("Cc", "test1@qq.com")	 //抄送
	// m.SetHeader("Bcc", "test2@qq.com")    // 密送
	m.SetHeader("Subject", "Hello!")
	//  发送格式和内容
	m.SetBody("text/html", "<h1>Hello Go!<h1/>")
	// m.Attach("./test.pdf") // 附件（不是必须）
	// host、 port、邮箱、秘钥（密码）
	d := gomail.NewDialer("smtp.qq.com", 465, "mineemail@qq.com", "thfdggghfdad")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
