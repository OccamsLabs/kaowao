package hashutils

import (
	"fmt"
	"crypto/sha256"
	"io/ioutil"
)


func HashForFile(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}


	// Compute the SHA256 hash of the file contents
	hash := sha256.Sum256(contents)

	newHash := fmt.Sprintf("%x",hash)
	return newHash, nil
}
