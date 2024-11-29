package util

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const MAX_LINE_LENGTH = 65000

func GetFileAsListOfStrings(filePath string) ([]string, error) {
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

func IsCharNumber(c rune) bool {
	if c >= 48 && c <= 57 {
		return true
	}
	return false
}

func IsStringNumber(c string) bool {
	r := []rune(c)
	return IsCharNumber(r[0])
}

func StringToNumber[T string | rune](n T) int {
	sn := string(n)
	sn = strings.TrimSpace(sn)
	res, err := strconv.Atoi(sn)
	if err != nil {
		return 0
	}
	return res
}

func DelimitedStringOfNumbersToIntSlice(in string) ([]int, error) {
	return SpaceSeparatedStringOfNumbersToIntSlice(in)
}

func SpaceSeparatedStringOfNumbersToIntSlice(in string) ([]int, error) {
	result := make([]int, 0)

	re := regexp.MustCompile("-?[0-9]+")

	for _, number := range re.FindAllString(in, -1) {
		n, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil
}

func CopyStringToIntMap(in map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range in {
		result[k] = v
	}
	return result
}

func ConvertBinaryStringToNumber(input string) int {
	result := float64(0)
	N := len(input) - 1

	for i, c := range input {
		result += math.Pow(2, float64(N-i)) * float64(StringToNumber(c))
	}

	return int(result)
}

func Abs[T int | float64](v T) T {
	if v < 0 {
		return -1 * v
	}
	return v
}
