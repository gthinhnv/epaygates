package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogConfig struct {
	Output   string // console, file
	FilePath string
	Level    string // info, debug, error, etc
}

type Logger struct {
	*logrus.Logger
}

func New(conf *LogConfig) (*Logger, error) {
	var logLevel, err = logrus.ParseLevel(conf.Level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	logInstance := logrus.New()
	logInstance.SetLevel(logLevel)

	if conf.Output == "file" {
		file, err := os.OpenFile(conf.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logInstance.Out = file
		}
	}

	return &Logger{logInstance}, nil
}
