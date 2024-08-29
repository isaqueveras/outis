package outis

import (
	"log"
	"os"
)

func setupLogger() ILogger { return logger{log.New(os.Stderr, "", log.LstdFlags)} }

type logger struct{ log *log.Logger }

func (l logger) Infof(format string, v ...interface{})  { l.log.Printf(format, v...) }
func (l logger) Warnf(format string, v ...interface{})  { l.log.Printf(format, v...) }
func (l logger) Errorf(format string, v ...interface{}) { l.log.Printf(format, v...) }
func (l logger) Debugf(format string, v ...interface{}) { l.log.Printf(format, v...) }
func (l logger) Panicf(format string, v ...interface{}) { l.log.Printf(format, v...) }
