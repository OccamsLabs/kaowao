/*
   Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/OccamsLabs/kaowao/pkg/config"
	"github.com/OccamsLabs/kaowao/pkg/hashutils"
)



// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args:  cobra.ExactArgs(2),

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

		// Read the file contents
		result, err := hashutils.HashForFile(path)

		if err != nil {
			fmt.Printf("error hashing file: %s\n", path)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
