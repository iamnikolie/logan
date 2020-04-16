package logrus

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetLevel() {
	logLevel := log.WarnLevel
	if level := os.Getenv("LOGAN_LOG_LEVEL"); level != "" {
		switch level {
		case "info":
			logLevel = log.InfoLevel
			break
		case "debug":
			logLevel = log.DebugLevel
			break
		case "error":
			logLevel = log.ErrorLevel
			break
		case "fatal":
			logLevel = log.FatalLevel
			break
		case "trace":
			logLevel = log.TraceLevel
			break
		case "warn":
			logLevel = log.WarnLevel
			break
		}
	}
	log.SetLevel(logLevel)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}
