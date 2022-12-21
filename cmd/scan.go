/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/OccamsLabs/kaowao/pkg/config"
	"os"
	"github.com/OccamsLabs/kaowao/pkg/hashutils"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scan called")
		configFile := args[0]

		scanPath(configFile)
	},
}


func scanPath(configFilePath string) {
	configFile, err := config.ReadConfig(configFilePath)
	if err != nil {
		fmt.Printf("error opening config file %s\n", configFilePath)
		os.Exit(1)
	}

	targets := configFile.Files
	for _, v := range targets {
		result, err := hashutils.HashForFile(v.Path)

		if err != nil {
			fmt.Printf("error hashing file: %s\n", v.Path)
		}

		if result != v.Hash {
			fmt.Printf("Hashes do not match: %s %s %s/n", v.Path, v.Hash, result)
		}
	}

}
func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
