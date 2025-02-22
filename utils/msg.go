package utils

import (
	"fmt"
	"net/smtp"
	"github.com/nuty/simple-blog/config"

)

func SendEmail(toEmail, content string) error {
	config, _ := config.LoadConfig("config/config.toml")
	to := []string{toEmail}
	subject := "有人评论了你的评论, 内容是：\n"
	message := []byte(subject + "\n" + content)

	auth := smtp.PlainAuth("", config.Email.From, config.Email.From,config.Email.Host)

	err := smtp.SendMail(config.Email.Host+":"+config.Email.Port, auth, config.Email.From, to, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}