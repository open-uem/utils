//go:build linux

package openuem_utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type OpenUEMLogger struct {
	LogFile *os.File
}

func NewLogger(logFilename string) *OpenUEMLogger {
	var err error

	logger := OpenUEMLogger{}

	// Get executable path to store logs
	wd := "/var/logs/openuem-server"

	logPath := filepath.Join(wd, logFilename)
	logger.LogFile, err = os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("could not create log file: %v", err)
	}

	logPrefix := strings.TrimSuffix(filepath.Base(logFilename), filepath.Ext(logFilename))
	log.SetOutput(logger.LogFile)
	log.SetPrefix(logPrefix + ": ")
	log.SetFlags(log.Ldate | log.Ltime)

	return &logger
}

func (l *OpenUEMLogger) Close() {
	l.LogFile.Close()
}
