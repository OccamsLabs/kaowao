/*
   Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
  "crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
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

type FileHash struct {
	Path string `yaml:"path"`
	Hash string `yaml:"hash"`
}



func scan(directory string, outFile string) {

  var fileHashes []FileHash

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
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}


		// Compute the SHA256 hash of the file contents
		hash := sha256.Sum256(contents)

    newHash := fmt.Sprintf("%x",hash)

		// Create a FileHash struct and append it to the slice
		fileHashes = append(fileHashes, FileHash{
			Path: path,
			Hash: newHash,
		})

		return nil
	})

	if err != nil {
		fmt.Printf("Error iterating through files: %v\n", err)
		os.Exit(1)
	}

	// Marshal the slice of FileHash structs to YAML
	yamlBytes, err := yaml.Marshal(fileHashes)
	if err != nil {
		fmt.Printf("Error marshalling to YAML: %v\n", err)
		os.Exit(1)
	}

	// Write the YAML to the output file
	err = ioutil.WriteFile(outFile, yamlBytes, 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		os.Exit(1)
	}
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
