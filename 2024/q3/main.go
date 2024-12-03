package main

import (
	"fmt"
	"os"
	"regexp"

	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	res, err := Q1(filename)
	fmt.Println("Answer 1:", res, err)

	res, err = Q2(filename)
	fmt.Println("Answer 2:", res, err)
}

func Q1(filename string) (int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return -1, err
	}

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	res := 0

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, r := range matches {
			res += util.StringToNumber(r[1]) * util.StringToNumber(r[2])
		}
	}

	return res, nil
}

func Q2(filename string) (int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return -1, err
	}

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	re2 := regexp.MustCompile(`[0-9]+`)
	res := 0

	multiply := true
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			if match == "do()" {
				multiply = true
			} else if match == "don't()" {
				multiply = false
			} else {
				if multiply {
					submatches := re2.FindAllString(match, -1)
					res += util.StringToNumber(submatches[0]) * util.StringToNumber(submatches[1])
				}
			}
		}
	}

	return res, nil
}
