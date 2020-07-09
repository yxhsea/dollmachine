// Package email 封装邮件转发
package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/smtp"
	"strings"
)

type Request struct {
	from     string
	fromName string
	to       []string
	subject  string
	body     io.Reader
}

func NewRequest(subject string, body io.Reader, to ...string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SetFrom(name, addr string) {
	r.fromName = name
	r.from = addr
}

func (r *Request) SendHTML(auth smtp.Auth) error {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("From: %s <%s>\r\n", r.fromName, r.from))
	buf.WriteString("To: " + strings.Join(r.to, ";") + "\r\n")
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", r.subject))
	buf.WriteString("MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n")
	buf.WriteString("\r\n")
	io.Copy(&buf, r.body)
	addr := "smtp.exmail.qq.com:465"
	host, _, _ := net.SplitHostPort(addr)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(r.from); err != nil {
		return err
	}
	if err = c.Rcpt(strings.Join(r.to, ";")); err != nil {
		return err
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	io.Copy(w, &buf)
	if err = w.Close(); err != nil {
		return err
	}
	return c.Quit()
}
