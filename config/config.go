package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Command      string   `json:"command" yaml:"command"`
	Files        []string `json:"files" yaml:"files"`
	IgnoredFiles []string `json:"ignoredFiles" yaml:"ignoredFiles"`
}

var extensionErr error = errors.New("configuration file must be in json or yaml format")
var commandMissingErr error = errors.New("command is a string and must be provided in the configuration file")
var filesMissingError error = errors.New("files is an array of string and must be provided in the configuration file")

// getConfigFileExtension returns the extension of the config file if this format is handled by the program.
func getConfigFileExtension(fp string) (ext string, err error) {
	ext = filepath.Ext(fp)
	if ext != ".json" && ext != ".yaml" && ext != ".yml" {
		return "", extensionErr
	}
	return ext, nil
}

// ParseConfigFile take the filepath as a parameter.
// If this file exists and his extension is supported
// the file is parsed and the Config is returned.
func ParseConfigFile(fp string) (conf Config, err error) {
	ext, err := getConfigFileExtension(fp)
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.Open(fp)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()

	fileContent, err := ioutil.ReadAll(configFile)
	if err != nil {
		return Config{}, err
	}

	if ext == ".json" {
		if err = json.Unmarshal(fileContent, &conf); err != nil {
			return Config{}, err
		}
	} else if ext == ".yaml" || ext == ".yml" {
		if err = yaml.Unmarshal(fileContent, &conf); err != nil {
			return Config{}, err
		}
	}

	if conf.Command == "" {
		return Config{}, commandMissingErr
	}
	if len(conf.Files) <= 0 {
		return Config{}, filesMissingError
	}

	return conf, nil
}
