package logiface

import "github.com/golang-template/pkg/types"

type Logiface interface {
	Debug(msg string)
	Error(msg string)
	Info(msg string)
	SetLogLevel(level string)
	SetLogOutput(outType types.LogOutput, logFileName string) error
	GetLogsCount() uint
}
