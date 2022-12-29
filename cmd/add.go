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
	Short: "Adds a file to a config file: kaowao config.yaml file.go",
	Long: `Adds a file to a config file: kaowao config.yaml file.go.
Accepts KAOWAO_SALT to hash the checksums and prevent tampering`,
	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		outFile := args[0]
		file := args[1]
		add(file, outFile)
	},
}

func add(targetPath string, configFilePath string) {
	configFile, err := config.ReadConfig(configFilePath)
	if err != nil {
		fmt.Printf("error opening config file %s\n", configFilePath)
		os.Exit(1)
	}

	fileHashes := configFile.Files

	newFileHashes, err := hashutils.HashTarget(targetPath)

	if err != nil {
		fmt.Printf("Error iterating through files: %v\n", err)
		os.Exit(1)
	}
	for _, h := range newFileHashes {
		idx := -1
		for i := range fileHashes {
			if fileHashes[i].Path == h.Path {
				idx = i
				break
			}
		}

		if idx != -1 {
			fileHashes[idx] = config.FileHash{
				Path: h.Path,
				Hash: h.Hash,
			}

		} else {
			fileHashes = append(fileHashes, config.FileHash{
				Path: h.Path,
				Hash: h.Hash,
			})
		}

	}

	// Create a FileHash struct and append it to the slice

	// Marshal the slice of FileHash structs to YAML

	configFile.Files = fileHashes
	config.WriteConfig(configFilePath, *configFile)

}

func init() {
	rootCmd.AddCommand(addCmd)
}
