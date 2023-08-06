package middlewares

import (
	"github.com/Taeho0927/goLogger/logger"
	"runtime"
	"time"
)

const ProjectName = "Taeho's_standard_API"

func SetLogger(handler string, message string) logger.LogContent {
	function, _, line, _ := runtime.Caller(1)
	content := logger.LogContent{
		Project:  ProjectName,
		Handler:  handler,
		Function: runtime.FuncForPC(function).Name(),
		CodeLine: line,
		Message:  message,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
	}
	return content
}
