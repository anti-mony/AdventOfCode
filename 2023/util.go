package main

import (
	"bufio"
	"os"
)

const MAX_LINE_LENGTH = 65000

func getFileAsListOfStrings(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	buffer := make([]byte, MAX_LINE_LENGTH)

	scanner := bufio.NewScanner(file)
	scanner.Buffer(buffer, MAX_LINE_LENGTH)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil

}

func isCharNumber(c rune) bool {

	if c >= 48 && c <= 57 {
		return true
	}

	return false
}
