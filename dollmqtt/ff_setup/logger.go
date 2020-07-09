package ff_setup

import (
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/BlueSimle/logcutting/cut_log"
)

func SetupLogger(filePath string) error {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		ForceColors:      true,
		QuoteEmptyFields: true,
		FullTimestamp:    true,
	})
	log.AddHook(cut_log.ContextHook{LogPath:filePath})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	return nil
}
