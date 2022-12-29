package hashutils

import (
	"crypto/sha256"
	"fmt"
	"github.com/OccamsLabs/kaowao/pkg/config"
	"io/ioutil"
	"os"
	"path/filepath"
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

func HashTarget(target string) ([]config.FileHash, error) {
	var fileHashes []config.FileHash
	err := filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		var result string
		salt := os.Getenv("KAOWAO_SALT")
		if salt != "" {
			result, err = SaltedHashForFile(path, salt)
		} else {
			// Read the file contents
			result, err = HashForFile(path)
		}

		if err != nil {
			fmt.Printf("error hashing file: %s, %s\n", path, err)
		}

		// Create a FileHash struct and append it to the slice
		fileHashes = append(fileHashes, config.FileHash{
			Path: path,
			Hash: result,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileHashes, nil
}
