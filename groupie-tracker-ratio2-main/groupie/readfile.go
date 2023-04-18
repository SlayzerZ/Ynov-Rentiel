package groupie

import (
	"bufio"
	"fmt"
	"os"
)

// Reads the file at a certain index
func Read(file string, index int) string {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Print("err")
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for i := 0; i < index; i++ {
		if fileScanner.Scan() {
		} else {
			return ""
		}
	}
	readFile.Close()
	return fileScanner.Text()
}

// Count the number of words in the file
func Count(file string) int {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Print("err")
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	count := 0
	for fileScanner.Scan() {
		count++
	}
	readFile.Close()
	return count
}
