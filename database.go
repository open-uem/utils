package openuem_utils

import (
	"encoding/binary"
	"fmt"

	"github.com/danieljoos/wincred"
	"golang.org/x/sys/windows/registry"
)

func CreatePostgresDatabaseURL() (string, error) {
	var err error
	// Create DATABASE_URL env variable
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\OpenUEM\Server`, registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("could not open registry to search OpenUEM Server entries")
	}
	defer k.Close()

	user, _, err := k.GetStringValue("PostgresUser")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresUser from registry")
	}

	host, _, err := k.GetStringValue("PostgresHost")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresHost from registry")
	}

	port, _, err := k.GetStringValue("PostgresPort")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresPort from registry")
	}

	database, _, err := k.GetStringValue("PostgresDatabase")
	if err != nil {
		return "", fmt.Errorf("could not read PostgresDatabase from registry")
	}

	pass, err := wincred.GetGenericCredential(host + ":" + port)
	if err != nil {
		return "", fmt.Errorf("could not read password from Windows Credential Manager")
	}

	decodedPass := UTF16BytesToString(pass.CredentialBlob, binary.LittleEndian)
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, decodedPass, host, port, database), nil
}
