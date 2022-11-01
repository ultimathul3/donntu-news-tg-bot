package logger

import (
	"errors"
	"log"
	"os"
)

type Logger struct {
	outFile *os.File
	logger  *log.Logger
}

var (
	Log *Logger
)

func New(filename string) (*Logger, error) {
	logger := &Logger{}
	var err error

	logger.outFile, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	logger.logger = log.New(logger.outFile, "", log.LstdFlags)
	if logger.logger == nil {
		return nil, errors.New("failed to create logger object")
	}

	return logger, nil
}

func (l *Logger) Info(v ...any) {
	l.logger.Println(v...)
}
