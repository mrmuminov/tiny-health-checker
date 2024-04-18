package utils

import (
	"bufio"
	"fmt"
	"os"
)

// CheckError error function for error handling
func CheckError(e error) bool {
	if e != nil {
		fmt.Println(e)
	}
	return true
}

// readFile function for reading file
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
