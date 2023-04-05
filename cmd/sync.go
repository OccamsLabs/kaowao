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

// syncCmd represents the init command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync a config file: kaowao config.yaml",
	Long: `Sync a config file: kaowao config.yaml.
Accepts KAOWAO_SALT to hash the checksums and prevent tampering`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sync called")
		configFile := args[0]

		sync(configFile)
	},
}

func sync(configFilePath string) {
	configFile, err := config.ReadConfig(configFilePath)
	if err != nil {
		fmt.Printf("error opening config file %s\n", configFilePath)
		os.Exit(1)
	}

	fileHashes := configFile.Files

	var newFileHashes []config.FileHash

	for _, value := range fileHashes {
		targetPath := value.Path
		if targetPath != "" {
			newFileHash, err := hashutils.HashTarget(targetPath)
			if err != nil {
				fmt.Printf("Error iterating through files: %v, removing \n", err)
				// remove files that errored out
				// append a filehash with empty hash value
				// sort out in the next step
				deleteHash := config.FileHash{Path: targetPath, Hash: ""}
				newFileHashes = append(newFileHashes, deleteHash)
			} else {
				newFileHashes = append(newFileHashes, newFileHash...)
			}
		}
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
			if h.Hash == "" {
				fileHashes = append(fileHashes[:idx], fileHashes[idx+1:]...)
			} else if fileHashes[idx].Hash != h.Hash {
				fmt.Printf("updating: %s\n", h.Path)
				fileHashes[idx] = config.FileHash{
					Path: h.Path,
					Hash: h.Hash,
				}
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
	rootCmd.AddCommand(syncCmd)
}
