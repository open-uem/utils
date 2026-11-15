//go:build linux

package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func CreatePostgresDatabaseURL() (string, error) {
	var err error

	// Open ini file
	cfg, err := ini.Load("/etc/openuem-server/openuem.ini")
	if err != nil {
		return "", err
	}

	user, err := cfg.Section("DB").GetKey("PostgresUser")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresUser from INI")
	}

	host, err := cfg.Section("DB").GetKey("PostgresHost")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresHost from INI")
	}

	port, err := cfg.Section("DB").GetKey("PostgresPort")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresPort from INI")
	}

	database, err := cfg.Section("DB").GetKey("PostgresDatabase")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresDatabase from INI")
	}

	pass, err := cfg.Section("DB").GetKey("PostgresPassword")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresPassword from INI")
	}

	return postgresDatabaseURL(user.String(), pass.String(), host.String(), port.String(), database.String()), nil
}
