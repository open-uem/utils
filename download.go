package openuem_utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, filepath string, expectedHash []byte) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// Get hash
	hash, err := GetSHA256Sum(filepath)
	if err != nil {
		return err
	}

	// Check hash
	if string(hash) != string(expectedHash) {
		return fmt.Errorf("checksum doesn't match")
	}

	return nil
}
