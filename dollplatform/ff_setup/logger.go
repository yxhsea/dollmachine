package ff_setup

import (
	"github.com/sirupsen/logrus"
	"os"
	"github.com/BlueSimle/logcutting/cut_log"
)

func SetupLogger(filePath string) error {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		ForceColors:      true,
		QuoteEmptyFields: true,
		FullTimestamp:    true,
	})

	logrus.AddHook(cut_log.ContextHook{LogPath:filePath})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	return nil
}
