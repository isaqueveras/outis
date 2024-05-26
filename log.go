package outis

import (
	"log"
	"os"
	"time"
)

type levelLog string

const (
	levelLogInfo  levelLog = "INFO"
	levelLogError levelLog = "EROR"
	levelLogDebug levelLog = "DEBUG"
	levelLogPanic levelLog = "PANIC"
)

type Log struct {
	Level     levelLog  `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Path      string    `json:"path,omitempty"`
}

func setupLogger() Logger {
	return stdLogger{log.New(os.Stderr, "", log.LstdFlags|log.Lmsgprefix)}
}

type stdLogger struct{ log *log.Logger }

func (l stdLogger) Infof(format string, v ...interface{})  { l.log.Printf(format, v...) }
func (l stdLogger) Warnf(format string, v ...interface{})  { l.log.Printf(format, v...) }
func (l stdLogger) Errorf(format string, v ...interface{}) { l.log.Printf(format, v...) }
func (l stdLogger) Debugf(format string, v ...interface{}) { l.log.Printf(format, v...) }
func (l stdLogger) Panicf(format string, v ...interface{}) { l.log.Printf(format, v...) }
