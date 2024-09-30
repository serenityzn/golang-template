package dumblogger

import (
	"github.com/golang-template/pkg/loggers/logiface"
	"github.com/golang-template/pkg/types"
)

type DumbLogger struct {
	logiface.Logiface
	count uint
}

func NewDumbLogger() *DumbLogger {
	return &DumbLogger{}
}

func (d *DumbLogger) Debug(msg string) {
	println("Debug: " + msg)
	d.count++
}

func (d *DumbLogger) Error(msg string) {
	println("Error: " + msg)
	d.count++
}

func (d *DumbLogger) Info(msg string) {
	println("INFO: " + msg)
	d.count++
}

func (d *DumbLogger) SetLogLevel(level string) {
	println(level)
}

func (d *DumbLogger) SetLogOutput(outType types.LogOutput, logFileName string) error {
	return nil
}

func (d *DumbLogger) GetLogsCount() uint {
	return d.count
}
