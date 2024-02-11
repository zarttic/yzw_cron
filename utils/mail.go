package utils

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"yzw_cron/config"
)

func SendEmail(info string) error {
	em := email.NewEmail()
	em.From = config.Conf.MailConfig.From
	em.To = config.Conf.MailConfig.To
	em.Subject = config.Conf.MailConfig.Subject
	em.Text = []byte(info)
	return em.Send(config.Conf.MailConfig.Addr, smtp.PlainAuth(
		config.Conf.MailConfig.Identity, config.Conf.MailConfig.Username, config.Conf.MailConfig.Password, config.Conf.MailConfig.Host))

}
