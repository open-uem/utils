//go:build darwin

package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func GetJWTKey() (string, error) {
	// First check for environment variable (useful for development)
	if envKey := os.Getenv("OPENUEM_JWT_KEY"); envKey != "" {
		return envKey, nil
	}

	// Open ini file
	configFile := GetConfigFile()
	cfg, err := ini.Load(configFile)
	if err != nil {
		return "", err
	}

	key, err := cfg.Section("JWT").GetKey("Key")
	if err != nil {
		return "", fmt.Errorf("could not read JWT Key from INI")
	}

	return key.String(), nil
}
