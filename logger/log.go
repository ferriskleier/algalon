package logger

import (
	"GoGPT/errorHandler"
	"log"
	"os"
)

func Log(location, user, input string) {
	logFileName := "logger/" + location + ".logger"
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	errorHandler.Handle(err)
	defer func(logFile *os.File) {
		err := logFile.Close()
		errorHandler.Handle(err)
	}(logFile)

	logger := log.New(logFile, "LOG: ", log.Ldate|log.Ltime)
	logger.Printf("[%s] %s", user, input)
}
