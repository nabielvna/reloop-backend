package services

import (
	"log"
	"os"
	"sync"
)

type LoggerService struct {
	logger *log.Logger
}

var (
	loggerInstance *LoggerService
	loggerOnce     sync.Once
)

func GetLogger() *LoggerService {
	loggerOnce.Do(func() {
		logger := log.New(os.Stdout, "[RELOOP] ", log.LstdFlags|log.Lshortfile)
		loggerInstance = &LoggerService{logger: logger}
	})
	return loggerInstance
}

func (ls *LoggerService) Info(message string) {
	ls.logger.Printf("INFO: %s", message)
}

func (ls *LoggerService) Error(message string) {
	ls.logger.Printf("ERROR: %s", message)
}

func (ls *LoggerService) Debug(message string) {
	ls.logger.Printf("DEBUG: %s", message)
}
