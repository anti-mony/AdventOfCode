package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func isStringNumber(c string) bool {

	r := []rune(c)

	return isCharNumber(r[0])
}

func StringToNumber(n string) int {
	res, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}
	return res
}

func SpaceSeparatedStringOfNumbersToIntSlice(in string) ([]int, error) {
	result := make([]int, 0)

	in = strings.ReplaceAll(in, "  ", " ")
	in = strings.TrimSpace(in)
	numbers := strings.Split(in, " ")

	for _, number := range numbers {
		n, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}

	return result, nil
}
