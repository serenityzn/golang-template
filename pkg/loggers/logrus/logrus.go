package logrus

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-template/pkg/loggers/logiface"
	"github.com/golang-template/pkg/types"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"time"
)

const logName = "logrus"

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
func (l *LogrusLogger) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Build the full URL format: http://ip:port/uri_path.
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		fullURL := url.URL{
			Scheme: scheme,
			Host:   r.Host,
			Path:   r.URL.Path,
		}

		// Wrap the ResponseWriter to capture the status code.
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		// Log request information before calling the next handler.
		l.log.WithFields(logrus.Fields{
			"full_url":   fullURL.String(),
			"method":     r.Method,
			"remote":     r.RemoteAddr,
			"user-agent": r.UserAgent(),
		}).Info("request started")

		// Call the next handler.
		next.ServeHTTP(ww, r)

		// Log request information after the request is completed, including the status code.
		l.log.WithFields(logrus.Fields{
			"full_url":     fullURL.String(),
			"method":       r.Method,
			"status":       ww.Status(),
			"path":         r.URL.Path,
			"remote":       r.RemoteAddr,
			"user-agent":   r.UserAgent(),
			"duration":     time.Since(start),
			"responseSize": ww.BytesWritten(),
		}).Info("request completed")
	})
}
