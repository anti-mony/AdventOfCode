package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"advent.of.code/util"
)

type display struct {
	signals []string
	output  []string
}

var NUMS = map[int]int{
	0: 6,
	1: 2,
	2: 5,
	3: 5,
	4: 4,
	5: 5,
	6: 6,
	7: 3,
	8: 7,
	9: 6,
}

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer Q1", solveQ1(inp))

	fmt.Println("Answer Q2", solveQ2(inp))
}

func solveQ1(displays []display) int {
	result := 0

	for _, d := range displays {
		for _, o := range d.output {
			l := len(o)
			if l == 2 || l == 4 || l == 3 || l == 7 {
				result++
			}
		}
	}

	return result
}

func solveQ2(displays []display) int {
	result := 0

	for _, d := range displays {
		for _, o := range d.signals {
			fmt.Println(o, len(o))
		}
		fmt.Println("X-------------------------X")
	}

	return result
}

func parseInput(filename string) ([]display, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([]display, 0)
	for _, line := range lines {
		splits := strings.Split(line, "|")
		signals := make([]string, 0)
		output := make([]string, 0)
		for _, s := range strings.Split(splits[0], " ") {
			sr := []rune(s)
			slices.Sort(sr)
			s = strings.TrimSpace(string(sr))
			if len(s) > 0 {
				signals = append(signals, s)
			}
		}
		for _, s := range strings.Split(splits[1], " ") {
			sr := []rune(s)
			slices.Sort(sr)
			s = strings.TrimSpace(string(sr))
			if len(s) > 0 {
				output = append(output, s)
			}
		}
		result = append(result, display{signals: signals, output: output})
	}

	return result, nil
}
