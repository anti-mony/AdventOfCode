package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"advent.of.code/util"
)

type key struct {
	index int
	op    string
}

type InputType map[key][]string

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	// for key, val := range inp {
	// 	fmt.Println(key, val)
	// }

	fmt.Println("Q1: ", Q1(inp))
	fmt.Println("Q2: ", Q2(inp))
}

func Q1(inp InputType) int {
	result := 0
	for key, vals := range inp {
		local := 0
		if key.op == "*" {
			local = 1
		}
		for _, v := range vals {
			vi, _ := strconv.Atoi(v)
			if key.op == "*" {
				local *= vi
			} else {
				local += vi
			}
		}
		result += local
	}
	return result
}

func Q2(inp InputType) int {
	return -1
}

func parseInput(filename string) (InputType, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}
	inp := make(InputType)
	operations := strings.Split(lines[len(lines)-1], " ")
	ops := []string{}
	for _, op := range operations {
		op = strings.TrimSpace(op)
		if op != "" && op != " " {
			ops = append(ops, op)
		}
	}
	for i, op := range ops {
		ops[i] = strings.TrimSpace(op)
		inp[key{i, ops[i]}] = make([]string, len(lines)-1)
	}

	for i, line := range lines[:len(lines)-1] {
		splitsRaw := strings.Split(line, " ")
		var splits []string
		for _, s := range splitsRaw {
			if s != "" && s != " " {
				splits = append(splits, strings.TrimSpace(s))
			}
		}
		for j, split := range splits {
			inp[key{j, ops[j]}][i] = split
		}
	}

	return inp, nil
}
