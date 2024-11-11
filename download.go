package openuem_utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, filepath string, expectedHash string) error {

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

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("file not found")
	}

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
	if fmt.Sprintf("%x", hash) != expectedHash {
		return fmt.Errorf("checksum doesn't match")
	}

	return nil
}
