package tools

import (
	"bufio"
	"log"
	"os"
)

// FileExists checks if a file exists and is not a directory
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FolderExists checks if a folder exists
func FileScaner(inputFile string) *bufio.Scanner {
	var (
		finput  *os.File
		scanner *bufio.Scanner
		err     error
	)

	if FileExists(inputFile) {
		finput, err = os.Open(inputFile)
		if err != nil {
			log.Fatalf("Could read input file '%s': %s\n", inputFile, err)
		}
		scanner = bufio.NewScanner(finput)
		return scanner
	} else {
		log.Fatalf("输入的文件有问题")
	}
	return nil
}


