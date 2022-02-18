package main

import (
	log "github.com/sirupsen/logrus"
)

func setLog(config AppConfig) {
	logger := log.StandardLogger()
	logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	logger.SetLevel(log.DebugLevel)
	/*
		switch app.Config.Log.Level {
		case "error":
			logger.SetLevel(log.ErrorLevel)
		case "warn", "warning":
			logger.SetLevel(log.WarnLevel)
		case "debug":
			logger.SetLevel(log.DebugLevel)
		case "trace":
			logger.SetLevel(log.TraceLevel)
		default:
			logger.SetLevel(log.InfoLevel)
		}
	*/
}

func main() {
	args := parseArgs()
	config := loadConfig(args)
	setLog(config)

	log.Warnln("It works!")
}
