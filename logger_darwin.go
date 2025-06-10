//go:build darwin

package utils

import (
	"os"
)

type OpenUEMLogger struct {
	LogFile *os.File
}
