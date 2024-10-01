package logiface

import (
	"github.com/golang-template/pkg/types"
	"net/http"
)

type Logiface interface {
	Debug(msg string)
	Error(msg string)
	Info(msg string)
	SetLogLevel(level types.LogLevel)
	SetLogOutput(outType types.LogOutput, logFileName types.LogName) error
	GetLogsCount() uint
	ServeHTTP(next http.Handler) http.Handler
}
