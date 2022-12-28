/*
Copyright © 2022 Andreas Tiefenthaler <contact@occamslabs.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/OccamsLabs/kaowao/pkg/config"
	"github.com/OccamsLabs/kaowao/pkg/hashutils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes directory to a config file: kaowao directory config.yaml",
	Long: `Initializes directory to a config file: kaowao directory config.yaml
Accepts KAOWAO_SALT to hash the checksums and prevent tampering`,
	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		directory := args[0]
		outFile := args[1]
		scan(directory, outFile)
	},
}

func scan(directory string, outFile string) {
	var fileHashes []config.FileHash
	var configFile config.ConfigFile
	configFile.Version = 1

	// Open the output file for writing
	f, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening output file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// Iterate through all files in the directory
	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		var result string
		salt := os.Getenv("KAOWAO_SALT")
		if salt != "" {
			result, err = hashutils.SaltedHashForFile(path, salt)
		} else {
			// Read the file contents
			result, err = hashutils.HashForFile(path)
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
		fmt.Printf("Error iterating through files: %v\n", err)
		os.Exit(1)
	}

	// Marshal the slice of FileHash structs to YAML

	configFile.Files = fileHashes
	config.WriteConfig(outFile, configFile)

}

func init() {
	rootCmd.AddCommand(initCmd)
}
