package main

import (
	log "github.com/sirupsen/logrus"
	"logcutting/cut_log"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		ForceColors:      true,
		QuoteEmptyFields: true,
		FullTimestamp:    true,
	})
	log.AddHook(cut_log.ContextHook{LogPath:"/var/cut_log/log_cutting/"})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.Level(5))
}

func main() {
	log.Debug("A group of walrus emerges from the ocean")
	log.Info("A group of walrus emerges from the ocean")
	log.Warn("A group of walrus emerges from the ocean")
	log.Error("A group of walrus emerges from the ocean")
	log.Fatal("A group of walrus emerges from the ocean")
}
