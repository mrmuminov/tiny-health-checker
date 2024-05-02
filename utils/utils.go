package utils

import (
	"bufio"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"tiny-healt-checker/structs"
)

// CheckError error function for error handling
func CheckError(e error) bool {
	if e != nil {
		fmt.Println(e)
		return false
	}
	return true
}

// ReadFile function for reading file
func ReadFile(filePath string) (string, error) {
	// open file
	file, err := os.Open(filePath)
	// CheckError error
	CheckError(err)
	// close file
	defer func(file *os.File) {
		err := file.Close()
		CheckError(err)
	}(file)
	// read file
	var lines string
	scanner := bufio.NewScanner(file)
	// read line by line
	for scanner.Scan() {
		lines += scanner.Text() + "\n"
	}
	return lines, scanner.Err()
}

func ParseFlags() string {
	configFile := flag.String("config", "config.yaml", "config file path")
	flag.Parse()
	return *configFile
}

func GetConfigFile(configFile string) string {
	data, err := ReadFile(configFile)
	CheckError(err)
	return data
}

func ParseConfig() structs.Config {
	configFile := ParseFlags()
	data := GetConfigFile(configFile)
	config := structs.Config{}
	err := yaml.Unmarshal([]byte(data), &config)
	CheckError(err)
	return config
}
