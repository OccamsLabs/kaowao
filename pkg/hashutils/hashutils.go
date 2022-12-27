package hashutils

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func HashForFile(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Compute the SHA256 hash of the file contents
	hash := sha256.Sum256(contents)

	newHash := fmt.Sprintf("%x", hash)
	return newHash, nil
}

func SaltedHashForFile(path string, salt string) (string, error) {
	hashForFile, err := HashForFile(path)
	if err != nil {
		return "", err
	}
	saltedHash := fmt.Sprintf("%s_%s", hashForFile, salt)

	newSaltedHash := sha256.Sum256([]byte(saltedHash))

	newHash := fmt.Sprintf("%x", newSaltedHash)
	return newHash, nil
}
