package syslog

import (
	"fmt"
	"log/syslog"
	"os"
	"strings"

	"github.com/gogap/logrus"
	"github.com/gogap/logrus/hooks/caller"
)

const (
	// Severity.

	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	LOG_EMERG syslog.Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

// SyslogHook to send logs via syslog.
type SyslogHook struct {
	Writer *syslog.Writer
}

// Creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewSyslogHook("udp", "localhost:514", syslog.LOG_DEBUG, "")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewHook(network, raddr string, priority syslog.Priority, tag string) (*SyslogHook, error) {
	w, err := syslog.Dial(network, raddr, priority, tag)
	return &SyslogHook{w}, err
}

func (hook *SyslogHook) Fire(entry *logrus.Entry) error {
	file, lineNumber := caller.GetCallerIgnoringLogMulti(1)
	if file != "" {
		sep := fmt.Sprintf("%s/src/", os.Getenv("GOPATH"))
		fileName := strings.Split(file, sep)
		if len(fileName) >= 2 {
			file = fileName[1]
		}
	}
	entry.Data["file"] = file
	entry.Data["line"] = lineNumber
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	switch entry.Level {
	case logrus.PanicLevel:
		return hook.Writer.Crit(line)
	case logrus.FatalLevel:
		return hook.Writer.Crit(line)
	case logrus.ErrorLevel:
		return hook.Writer.Err(line)
	case logrus.WarnLevel:
		return hook.Writer.Warning(line)
	case logrus.InfoLevel:
		return hook.Writer.Info(line)
	case logrus.DebugLevel:
		return hook.Writer.Debug(line)
	default:
		return nil
	}
}

func (hook *SyslogHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
