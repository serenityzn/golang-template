package logrus

import (
	"fmt"
	"github.com/golang-template/pkg/loggers/logiface"
	"github.com/golang-template/pkg/types"
	"github.com/sirupsen/logrus"
	"os"
)

const logName = "logrus"

var logger = LogrusLogger{}

type LogrusLogger struct {
	log   *logrus.Logger
	count uint
	logiface.Logiface
}

func NewLogrusLogger() *LogrusLogger {
	fmt.Printf("LOGRUS INITIALISING...\n\n")
	mlog := logrus.New()
	mlog.Out = os.Stdout
	mlog.SetLevel(logrus.DebugLevel)
	return &LogrusLogger{
		log:   mlog,
		count: 0,
	}
}

func (l *LogrusLogger) Debug(msg string) {
	l.log.WithFields(logrus.Fields{
		"logger": logName,
	}).Debug(msg)
	l.count++
}

func (l *LogrusLogger) Error(msg string) {
	l.log.WithFields(logrus.Fields{
		"logger": logName,
	}).Error(msg)
}

func (l *LogrusLogger) Info(msg string) {
	l.log.WithFields(logrus.Fields{
		"logger": logName,
	}).Info(msg)
}

func (l *LogrusLogger) SetLogLevel(level string) {
	if level == "error" {
		l.log.SetLevel(logrus.ErrorLevel)
	} else if level == "debug" {
		l.log.SetLevel(logrus.DebugLevel)
	} else {
		l.log.SetLevel(logrus.InfoLevel)
	}
	l.count++
}

func (l *LogrusLogger) SetLogOutput(outType types.LogOutput, logFileName string) error {
	switch outType {
	case types.Stdout:
		l.log.Out = os.Stdout
	case types.Fileout:
		logFile, err := getLogFile(logFileName)
		if err != nil {
			return err
		}
		l.log.SetOutput(logFile)
	}

	return nil
}

func (l *LogrusLogger) GetLogsCount() uint {
	return l.count
}

func getLogFile(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		file, err := os.Create(name)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return file, nil
}
