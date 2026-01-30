//go:build darwin

package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type OpenUEMLogger struct {
	LogFile *os.File
}

func (l *OpenUEMLogger) Close() {
	l.LogFile.Close()
}

func NewLogger(logFilename string) *OpenUEMLogger {
	var err error

	logger := OpenUEMLogger{}

	// Get user home directory for logs on macOS
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("[FATAL]: could not get user home directory, reason: %v", err)
	}

	wd := filepath.Join(homeDir, ".openuem", "logs")

	if _, err := os.Stat(wd); os.IsNotExist(err) {
		if err := os.MkdirAll(wd, 0755); err != nil {
			log.Fatalf("[FATAL]: could not create log directory, reason: %v", err)
		}
	}

	logPath := filepath.Join(wd, logFilename)
	logger.LogFile, err = os.Create(logPath)
	if err != nil {
		log.Fatalf("could not create log file: %v", err)
	}

	logPrefix := strings.TrimSuffix(filepath.Base(logFilename), filepath.Ext(logFilename))
	log.SetOutput(logger.LogFile)
	log.SetPrefix(logPrefix + ": ")
	log.SetFlags(log.Ldate | log.Ltime)

	return &logger
}

func NewAuthLogger() *log.Logger {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("[FATAL]: could not get user home directory, reason: %v", err)
	}

	wd := filepath.Join(homeDir, ".openuem", "logs")

	if _, err := os.Stat(wd); os.IsNotExist(err) {
		if err := os.MkdirAll(wd, 0755); err != nil {
			log.Fatalf("[FATAL]: could not create log directory, reason: %v", err)
		}
	}

	logPath := filepath.Join(wd, "auth.log")
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("could not open auth log file: %v", err)
	}

	return log.New(logFile, "auth: ", log.Ldate|log.Ltime)
}
