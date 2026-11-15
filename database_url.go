package utils

import (
	"fmt"
	"net/url"
)

func postgresDatabaseURL(username, password, hostname, port, database string) string {
	return fmt.Sprintf("postgres://%s@%s:%s/%s",
		url.UserPassword(username, password).String(),
		url.PathEscape(hostname),
		url.PathEscape(port),
		url.PathEscape(database))
}
