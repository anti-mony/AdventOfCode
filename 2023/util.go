package main

import (
	"bufio"
	"os"
	"regexp"
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
	n = strings.TrimSpace(n)
	res, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}
	return res
}

func SpaceSeparatedStringOfNumbersToIntSlice(in string) ([]int, error) {
	result := make([]int, 0)

	re := regexp.MustCompile("[0-9]+")

	for _, number := range re.FindAllString(in, -1) {
		n, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}

	return result, nil
}

func Dedupe[T string | int](inp []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)
	for _, v := range inp {
		if _, ok := seen[v]; !ok {
			result = append(result, v)
			seen[v] = true
		}

	}
	return result
}

func GCD(a, b int) int {
	result := 0
	if a > b {
		result = b
	} else {
		result = a
	}

	for result > 0 {
		if a%result == 0 && b%result == 0 {
			return result
		}
		result--
	}

	return result
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}
