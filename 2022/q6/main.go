package main

import (
	"fmt"
	"log"

	"advent.of.code/util"
)

// This is 4 for P1
const lenStartSeq = 14

func main() {
	signalBuffers, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i, signal := range signalBuffers {
		fmt.Printf("%d) Answer P1: %d \n", i+1, solve(signal))
	}
}

func solve(signal string) int {
	i := 0
	for i < len(signal)-lenStartSeq {
		seen := make(map[byte]bool)
		found := true
		for j := i; j < i+lenStartSeq; j++ {
			if _, ok := seen[signal[j]]; ok {
				found = false
				break
			}
			seen[signal[j]] = true
		}
		if found {
			return i + lenStartSeq
		}
		i++
	}
	return -1
}

func parseInput(fileName string) ([]string, error) {
	lines, err := util.GetFileAsListOfStrings(fileName)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
