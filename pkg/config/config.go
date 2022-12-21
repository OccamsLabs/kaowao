package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
	"os"
)

type ConfigFile struct {
	Version int `yaml:"version"`
	RepoUrl string `yaml:"repo_url"`
	Files []FileHash `yaml:"files"`
}

type FileHash struct {
	Path string `yaml:"path"`
	Hash string `yaml:"hash"`
}

func ReadConfig(filename string) (*ConfigFile, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &ConfigFile{}
	err = yaml.Unmarshal(buf, c)

	if err != nil {
        return nil, fmt.Errorf("in file %q: %w", filename, err)
    }

    return c, err
}


func WriteConfig(filename string, config ConfigFile) {
	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		fmt.Printf("Error marshalling to YAML: %v\n", err)
		os.Exit(1)
	}

	// Write the YAML to the output file
	err = ioutil.WriteFile(filename, yamlBytes, 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		os.Exit(1)
	}
}
