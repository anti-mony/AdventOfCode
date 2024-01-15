package main

import (
	"fmt"
	"strconv"
)

var _writtenDigits = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func q1sol() error {

	lines, err := getFileAsListOfStrings("input1.txt")
	if err != nil {
		return err
	}

	result := 0

	for _, line := range lines {
		n, err := getNumberFromString(line)
		if err != nil {
			return err
		}
		result += n
	}

	fmt.Printf("Answer is %d \n", result)

	return nil
}

func isNumberWrittenInWords(input string, startIndex int) (bool, int) {
	maxLen := len(input)

	for key, value := range _writtenDigits {
		if len(key)+startIndex-1 < maxLen {
			if input[startIndex:startIndex+len(key)] == key {
				return true, value
			}
		}
	}
	return false, 0
}

func getNumberFromString(in string) (int, error) {

	var firstDecimalDigit, lastDecimalDigit *int
	for i, c := range in {
		if isCharNumber(c) {
			if firstDecimalDigit == nil {
				num, err := strconv.Atoi(string(c))
				if err != nil {
					return 0, err
				}
				firstDecimalDigit = &num
			}
			num, err := strconv.Atoi(string(c))
			if err != nil {
				return 0, err
			}
			lastDecimalDigit = &num
		} else {
			yes, num := isNumberWrittenInWords(in, i)
			if yes {
				if firstDecimalDigit == nil {
					firstDecimalDigit = &num
				}
				lastDecimalDigit = &num
			}
		}
	}

	if firstDecimalDigit != nil && lastDecimalDigit != nil {
		res, err := strconv.Atoi(fmt.Sprintf("%d%d", *firstDecimalDigit, *lastDecimalDigit))
		if err != nil {
			return 0, err
		}
		return res, nil
	}

	return 0, nil
}
