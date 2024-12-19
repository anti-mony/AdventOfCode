package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	availablePatterns, patternsToDisplay, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(availablePatterns, patternsToDisplay)

	fmt.Println("Q1: ", Q1(availablePatterns, patternsToDisplay))
	fmt.Println("Q2: ", Q2(availablePatterns, patternsToDisplay))
}

func Q1(availablePatterns []string, patternsToDisplay []string) int {
	result := 0

	for _, patternToDisplay := range patternsToDisplay {
		seen := map[string]bool{}
		if canDisplayPattern(availablePatterns, patternToDisplay, seen) {
			result++
		}
	}
	return result
}

func Q2(availablePatterns []string, patternsToDisplay []string) int {
	result := 0

	for _, patternToDisplay := range patternsToDisplay {
		seen := map[string]int{}
		result += numberOfWaysDisplayPattern(availablePatterns, patternToDisplay, seen)
	}

	return result
}

func canDisplayPattern(availablePatterns []string, patternToDisplay string, seen map[string]bool) bool {
	if result, found := seen[patternToDisplay]; found {
		return result
	}

	if len(patternToDisplay) == 0 {
		return true
	}

	for _, availablePattern := range availablePatterns {
		if strings.HasPrefix(patternToDisplay, availablePattern) {
			if canDisplayPattern(availablePatterns, patternToDisplay[len(availablePattern):], seen) {
				seen[patternToDisplay] = true
				return true
			}
		}
	}

	seen[patternToDisplay] = false
	return false
}

func numberOfWaysDisplayPattern(availablePatterns []string, patternToDisplay string, seen map[string]int) int {
	if result, found := seen[patternToDisplay]; found {
		return result
	}

	if len(patternToDisplay) == 0 {
		return 1
	}

	nWays := 0

	for _, availablePattern := range availablePatterns {
		if strings.HasPrefix(patternToDisplay, availablePattern) {
			nWays += numberOfWaysDisplayPattern(availablePatterns, patternToDisplay[len(availablePattern):], seen)
		}
	}

	seen[patternToDisplay] = nWays
	return nWays
}

func parseInput(filename string) ([]string, []string, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, nil, err
	}

	availablePatterns := make([]string, 0)
	for _, p := range strings.Split(lines[0], ",") {
		availablePatterns = append(availablePatterns, strings.TrimSpace(p))
	}

	patternsToDisplay := make([]string, 0)
	patternsToDisplay = append(patternsToDisplay, lines[2:]...)

	return availablePatterns, patternsToDisplay, nil
}
