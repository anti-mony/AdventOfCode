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
	paths := 0

	var recurse func(start string, path string)

	recurse = func(start string, path string) {
		if strings.HasSuffix(path, "out") {
			// fmt.Println(path)
			paths++
			return
		}

		for _, nxt := range inp[start] {
			recurse(nxt, path+","+nxt)
		}
	}

	recurse("you", "you")

	return paths
}

func Q2(inp InputType) int {
	paths := 0

	var recurse func(start string, path string)

	recurse = func(start string, path string) {
		if strings.HasSuffix(path, "out") {
			if strings.Contains(path, "fft") && strings.Contains(path, "dac") {
				paths++
			}
			return
		}

		for _, nxt := range inp[start] {
			recurse(nxt, path+","+nxt)
		}
	}

	recurse("svr", "svr")

	return paths
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
