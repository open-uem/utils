//go:build linux

package utils

import (
	"gopkg.in/ini.v1"
)

func CreatePostgresDatabaseURL() (string, error) {
	var err error

	// Open ini file
	cfg, err := ini.Load("/etc/openuem-server/openuem.ini")
	if err != nil {
		return "", err
	}

	key, err := cfg.Section("DB").GetKey("PostgresUrl")
	if err != nil {
		return "", err
	}

	return key.String(), nil
}
