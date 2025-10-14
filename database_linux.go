//go:build linux

package utils

import (
	"fmt"
	"log"
	"net/url"
	"text/template"

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
	username := url.PathEscape(user.String())

	host, err := cfg.Section("DB").GetKey("PostgresHost")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresHost from INI")
	}
	hostname := url.PathEscape(host.String())

	port, err := cfg.Section("DB").GetKey("PostgresPort")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresPort from INI")
	}
	dbPort := url.PathEscape(port.String())

	database, err := cfg.Section("DB").GetKey("PostgresDatabase")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresDatabase from INI")
	}
	databaseName := url.PathEscape(database.String())

	pass, err := cfg.Section("DB").GetKey("PostgresPassword")
	if err != nil {
		return "", fmt.Errorf("could not read password from Windows Credential Manager")
	}
	password := url.PathEscape(pass.String())
	log.Println(pass.String(), password, url.QueryEscape(pass.String()), template.URLQueryEscaper(pass.String()))

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, hostname, dbPort, databaseName), nil

}
