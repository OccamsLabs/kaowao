/*
Copyright © 2022 Andreas Tiefenthaler <contact@occamslabs.com>
*/
package cmd

import (
	"fmt"

	"github.com/OccamsLabs/kaowao/pkg/config"
	"github.com/OccamsLabs/kaowao/pkg/hashutils"
	"github.com/OccamsLabs/kaowao/pkg/result"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans based on a given configuration file: kaowao scan configuration.yaml",
	Long: `Scans based on a given configuration file: kaowao scan configuration.yaml
Accepts KAOWAO_SALT to hash the checksums and prevent tampering`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scan called")
		configFile := args[0]

		scanPath(configFile)
	},
}

func scanPath(configFilePath string) {
	results := []result.ResultInfo{}
	var out result.ResultInfos

	configFile, err := config.ReadConfig(configFilePath)
	if err != nil {
		fmt.Printf("error opening config file %s\n", configFilePath)
		os.Exit(1)
	}

	targets := configFile.Files
	for _, v := range targets {
		var resultHash string
		var err error

		salt := os.Getenv("KAOWAO_SALT")

		if salt != "" {
			resultHash, err = hashutils.SaltedHashForFile(v.Path, salt)
		} else {
			// Read the file contents
			resultHash, err = hashutils.HashForFile(v.Path)
		}

		if err != nil {
			fmt.Printf("error hashing file: %s\n", v.Path)

			results = append(results, result.ResultInfo{
				Path:    v.Path,
				Message: fmt.Sprintf("error hashing: %s", err),
			})
		}

		if resultHash != v.Hash {
			fmt.Printf("Hashes do not match: %s %s %s\n", v.Path, v.Hash, resultHash)

			results = append(results, result.ResultInfo{
				Path:         v.Path,
				Message:      fmt.Sprintf("Hashes do not match: %s %s %s", v.Path, v.Hash, resultHash),
				Hash:         v.Hash,
				ExpectedHash: resultHash,
			})
		}
	}
	out.ScanTime = time.Now().Format(time.RFC3339)
	out.Results = results
	report, _ := result.ToJson(out)
	fmt.Printf("%s\n", report)

	if len(results) != 0 {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
