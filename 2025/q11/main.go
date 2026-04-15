package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"advent.of.code/util"
)

type InputType map[string][]string

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Q1: ", Q1(inp))
	fmt.Println("Q2: ", Q2(inp))
}

func Q1(inp InputType) int {
	var recurse func(start string, path string) int

	seen := make(map[string]int)

	recurse = func(start string, path string) int {
		if v, found := seen[start]; found {
			return v
		}
		if start == "out" {
			return 1
		}

		res := 0
		for _, nxt := range inp[start] {
			res += recurse(nxt, path+","+nxt)
		}
		seen[start] = res
		return res
	}

	return recurse("you", "you")
}

func Q2(inp InputType) int {
	var recurse func(start string, seenFft bool, seenDac bool) int

	type key struct {
		node    string
		seenFft bool
		seenDac bool
	}

	seen := map[key]int{}

	recurse = func(start string, seenFft bool, seenDac bool) int {
		k := key{start, seenFft, seenDac}
		if v, found := seen[k]; found {
			return v
		}

		if start == "out" {
			if seenFft && seenDac {
				return 1
			}
			return 0
		}

		res := 0
		for _, nxt := range inp[start] {
			res += recurse(
				nxt,
				seenFft || nxt == "fft",
				seenDac || nxt == "dac",
			)
		}
		seen[k] = res
		return res
	}

	return recurse("svr", false, false)
}

func parseInput(filename string) (InputType, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}
	res := make(map[string][]string)
	for _, line := range lines {
		splits := strings.Split(line, ":")
		res[strings.TrimSpace(splits[0])] = strings.Split(strings.TrimSpace(splits[1]), " ")
	}

	return res, nil
}
