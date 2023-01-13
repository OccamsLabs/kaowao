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

	fileHashes, err := hashutils.HashTarget(directory)

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
