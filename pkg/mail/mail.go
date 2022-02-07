package mail

import (
	"fmt"
	"net/smtp"
	"processChecker/pkg/config"
	"strings"
	"time"
)

func Do(c *config.Instance) error {
	s := c.Mail.Smtp
	r := c.Mail.Recipients
	a := fmt.Sprintf("%s:%d", s.Hostname, s.Port)
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Hostname)
	msg := fmt.Sprintf("From: %s\r\n"+
		"Sender: %s\r\n"+
		"To: %s\r\n"+
		"Subject: Process %s exit unexpectedly!\r\n"+
		"\r\n"+
		"Process %s exit unexpectedly!\r\n"+
		"Date: %s\r\n",
		s.Username, s.Username, strings.Join(r, ","), c.Process.Name, c.Process.Name, time.Now())
	fmt.Println(msg)
	err := smtp.SendMail(a, auth, s.Username, r, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
