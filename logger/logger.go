package logger

import (
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Service   string
	DebugMode bool
}

type ginBeLogger struct {
	Logger *logrus.Logger
	Config *LoggerConfig
}

func NewLogger(config *LoggerConfig) *ginBeLogger {
	logger := logrus.New()
	if config.DebugMode {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.ErrorLevel)
	}
	logger.SetFormatter(&logrus.JSONFormatter{})
	return &ginBeLogger{
		Logger: logger,
		Config: config,
	}
}

func (gbl *ginBeLogger) Warn(msg string) {
	gbl.Logger.WithFields(logrus.Fields{
		"service": gbl.Config.Service,
	}).Warn(msg)
}

func (gbl *ginBeLogger) Info(msg string) {
	gbl.Logger.WithFields(logrus.Fields{
		"service": gbl.Config.Service,
	}).Info(msg)
}
