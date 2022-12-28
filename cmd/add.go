/*
Copyright © 2022 Andreas Tiefenthaler <contact@occamslabs.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/OccamsLabs/kaowao/pkg/config"
	"github.com/OccamsLabs/kaowao/pkg/hashutils"
	"github.com/spf13/cobra"
)

// addCmd represents the init command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a file to a config file: kaowao file.go config.yaml",
	Long: `Adds a file to a config file: kaowao file.go config.yaml.
Accepts KAOWAO_SALT to hash the checksums and prevent tampering`,
	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		file := args[0]
		outFile := args[1]
		add(file, outFile)
	},
}

func add(targetFile string, configFilePath string) {

	configFile, err := config.ReadConfig(configFilePath)
	if err != nil {
		fmt.Printf("error opening config file %s\n", configFilePath)
		os.Exit(1)
	}

	fileHashes := configFile.Files

	var result string
	salt := os.Getenv("KAOWAO_SALT")
	if salt != "" {
		result, err = hashutils.SaltedHashForFile(targetFile, salt)
	} else {
		// Read the file contents
		result, err = hashutils.HashForFile(targetFile)
	}

	if err != nil {
		fmt.Printf("error hashing file: %s, %s\n", targetFile, err)
	}

	idx := -1
	for i := range fileHashes {
		if fileHashes[i].Path == targetFile {
			idx = i
		}
	}

	if idx != -1 {
		fileHashes[idx] = config.FileHash{
			Path: targetFile,
			Hash: result,
		}

	} else {
		fileHashes = append(fileHashes, config.FileHash{
			Path: targetFile,
			Hash: result,
		})

	}

	// Create a FileHash struct and append it to the slice

	// Marshal the slice of FileHash structs to YAML

	configFile.Files = fileHashes
	config.WriteConfig(configFilePath, *configFile)

}

func init() {
	rootCmd.AddCommand(addCmd)
}
