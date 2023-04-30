package utils

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendEmail(to, cc, bcc []string, subject, body, file string) error {
	// QQ 邮箱：
	// SMTP 服务器地址：smtp.11.com (SSL协议端口：465/994 | 非SSL协议端口：25）
	// 163 邮箱：
	// SMTP 服务器地址：smtp.163.com (端口：25）

	alias := viper.GetString("email.alias")
	host := viper.GetString("email.host")
	port := viper.GetInt("email.port")
	username := viper.GetString("email.username")
	password := viper.GetString("email.password") // 授权码

	m := gomail.NewMessage(
	//gomail.SetCharset() // 设置邮件的字符集， 默认UTF-8
	//gomail.SetEncoding() //设置电子邮件编码的邮件设置 默认quoted-printable
	)

	//m.SetHeader("From", username) // 发件人
	m.SetHeader("From", alias+"<"+username+">") // 增加发件人别名
	m.SetHeader("To", to...)                    // 收件人
	m.SetHeader("Cc", cc...)                    // 抄送
	m.SetHeader("Bcc", bcc...)                  // 暗送
	m.SetHeader("Subject", subject)             // 主题
	m.SetBody("text/plain", body)               // 发送纯文本
	//m.SetBody("text/html", message)                // 发送时使用html解析
	if file != "" {
		m.Attach(file) //发送附件
	}
	//m.Attach(file, gomail.Rename(mime.QEncoding.Encode("UTF-8", path.Base(file_name)))) //发送附件并设置附件名，设置编码为UTF-8解决中文乱码

	//创建SMTP客户端，连接到远程的邮件服务器
	d := gomail.NewDialer(host, port, username, password)

	// 关闭SSL协议认证
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 发送邮件
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil

	/*
		使用原始smtp发送邮件
		host := "smtp.qq.com"
		port := "25"
		userName := "*****@qq.com"
		password := "******" // qq邮箱填授权码

		e := &email.Email{
			To:      []string{"******@qq.com", "******1@qq.com"},
			From:    userName,
			Subject: "Email Send Test",
			Text:    []byte("Text Body is, of course, supported!"),
			HTML:    []byte("<h1>This a test email</h1>"),
			Headers: textproto.MIMEHeader{},
		}

		err := e.Send(host+port, smtp.PlainAuth("", userName, password, host))
		if err != nil {
			panic(err)
		}

	*/

}
