package gologger

import (
	"go.uber.org/zap"
	"log"
	"os"
)

type LoggerService struct {
	log *zap.Logger
}

func NewLoggerService() *LoggerService {
	logger, err := zap.NewProduction()
	if nil != err {
		log.Println(err)
		os.Exit(1)
	}

	if len(os.Args) > 1 && os.Args[1] == "debug"{
		logger, err = zap.NewDevelopment()
		if nil != err {
			log.Println(err)
			os.Exit(1)
		}
	}

	defer logger.Sync()

	logger.Debug("Logger Service Initialized")
	return &LoggerService{
		log: logger,
	}
}

func (logger *LoggerService) Info(message string) {
	logger.log.Info(message)
	logger.log.Sync()
}

func (logger *LoggerService) Debug(message string) {
	logger.log.Debug(message)
	logger.log.Sync()
}

func (logger *LoggerService) Error(message string) {
	logger.log.Error(message)
	logger.log.Sync()
}