package mail

import (
	"backend/logger"
	"fmt"
	emailPKG "github.com/jordan-wright/email"
	"net/smtp"
)

// SMTP 实现 email.Driver interface
type SMTP struct{}

// Send 实现 email.Driver interface 的 Send 方法
func (s *SMTP) Send(email Email, config map[string]string) bool {

	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("发送邮件", "发件详情", e)

	err := e.Send(
		fmt.Sprintf("%v:%v", config["smtp_host"], config["smtp_port"]),

		smtp.PlainAuth(
			"",
			config["smtp_username"],
			config["smtp_password"],
			config["smtp_host"],
		),
	)
	if err != nil {
		fmt.Println(err)
		return false
	}

	logger.DebugString("发送邮件", "发件成功", "")
	return true
}
