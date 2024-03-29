package logger

import (
	"github.com/sirupsen/logrus"
)

type GinBeLoggerConfig struct {
	Service       string
	DebugModeFlag bool
}

type GinBeLogger struct {
	Logger *logrus.Logger
	Config *GinBeLoggerConfig
}

func NewLogger(config *GinBeLoggerConfig) *GinBeLogger {
	logger := logrus.New()
	if config.DebugModeFlag {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.WarnLevel)
	}
	logger.SetFormatter(&logrus.JSONFormatter{})
	return &GinBeLogger{
		Logger: logger,
		Config: config,
	}
}

func (gbl *GinBeLogger) Warn(msg string) {
	gbl.Logger.WithFields(logrus.Fields{
		"service": gbl.Config.Service,
	}).Warn(msg)
}

func (gbl *GinBeLogger) Info(msg string) {
	gbl.Logger.WithFields(logrus.Fields{
		"service": gbl.Config.Service,
	}).Info(msg)
}

func (gbl *GinBeLogger) Error(msg string) {
	gbl.Logger.WithFields(logrus.Fields{
		"service": gbl.Config.Service,
	}).Error(msg)
}
