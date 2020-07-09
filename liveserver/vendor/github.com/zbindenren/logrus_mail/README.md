# Mail Hook for Logrus [![GoDoc](http://godoc.org/github.com/zbinderen/logrus_mail?status.svg)](http://godoc.org/github.com/zbindenren/logrus_mail) [![Go Report Card](https://goreportcard.com/badge/github.com/zbindenren/logrus_mail)](https://goreportcard.com/report/github.com/zbindenren/logrus_mail)

In some deployments, you'll want to report errors by email. If you add this hook, an email will send for the following levels:

* Error
* Fatal
* Panic

The subject is of the form `APPLICATION_NAME - LEVEL` and the body contains the timestamp and the message.

## Installation

Install the package with go:

```go
go get github.com/zbindenren/logrus_mail
```

## Usage

For `APPLICATION_NAME`, substitute a short string that will identify your application or service in the logs.

```go
import (
  "log/syslog"
  "github.com/Sirupsen/logrus"
  "github.com/zbindenren/logrus_mail"
)

func main() {
  log       := logrus.New()
  // if you do not need authentication for your smtp host
  hook, err := logrus_mail.NewMailHook("APPLICATION_NAME", "HOST", PORT, "FROM", "TO")

  if err == nil {
    log.Hooks.Add(hook)
  }
}
```

Example with authentication:
```go
  // if you need authentication for your smtp host
  hook, err := logrus_mail.NewMailAuthHook("APPLICATION_NAME", "HOST", PORT, "FROM", "TO", "USERNAME", "PASSWORD")
```

If you want to send mails with gmail:
```go
 hook, err := logrus_mail.NewMailAuthHook("testapp", "smtp.gmail.com", 587, "user.name@gmail.com", "user.name@gmail.com", "user.name", "password")
```

If you get the following error:
```
Failed to fire hook: 534 5.7.14 <https://accounts.google.com/ContinueSignIn?sarp=1&scc=1&plt=AKgnsbt7D
5.7.14 N0zOlIl3SKJo5pXolT2E87etIZvTL03NXXQI4vST_GvPFo5p5OPvn6XqnQNgJsneZytvRa
5.7.14 nWhV3Qy6cynKd7_s0KRGlGNKI25t15FH9v5ztyFZ80dnM-qXDqRvMr8_pGaYrHKfI9rRB2
5.7.14 3VJLfZAiBBBe0L3IKG6sy8QEFRLylxNLiTvighE2qAfcSnxAxz5kDs1fB0szYnlryuBN0B
5.7.14 43SATbFjeep4iMzAWvVJ9hoGwrqI> Please log in via your web browser and
5.7.14 then try again.
5.7.14  Learn more at
5.7.14  https://support.google.com/mail/answer/78754 t8sm37578500wjy.41 - gsmtp
```

Check the Setting **Allow less secure apps** in the gmail account settings.
