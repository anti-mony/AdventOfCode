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
	fmt.Println("Q2: ", Q2(filename))
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

func Q2(filename string) int {
	lines, _ := util.GetFileAsListOfStrings(filename)
	opsLine := util.StringToCharSlice(lines[len(lines)-1])
	prev := 0
	i := 0
	result := 0
	for i = 1; i < len(opsLine); i++ {
		if opsLine[i] != "*" && opsLine[i] != "+" {
			continue
		}

		vals := parseColumn(lines, prev, i-1)
		localRes, _ := strconv.Atoi(vals[0])
		for _, v := range vals[1:] {
			vv, _ := strconv.Atoi(v)
			if opsLine[prev] == "+" {
				localRes += vv
			} else {
				localRes *= vv
			}
		}
		// fmt.Println("-->", vals, ">>", localRes)
		result += localRes
		prev = i
	}

	vals := parseColumn(lines, prev, len(opsLine))
	localRes, _ := strconv.Atoi(vals[0])
	for _, v := range vals[1:] {
		vv, _ := strconv.Atoi(v)
		if opsLine[prev] == "+" {
			localRes += vv
		} else {
			localRes *= vv
		}
	}
	// fmt.Println("-->", vals, ">>", localRes)

	result += localRes

	return result
}

func parseColumn(lines []string, start, end int) []string {
	tvals := [][]string{}
	maxLen := 0
	for _, line := range lines[:len(lines)-1] {
		w := line[start:end]
		tvals = append(tvals, util.StringToCharSlice(w))
		maxLen = max(maxLen, len(w))
	}

	vals := []string{}
	for i := range maxLen {
		num := ""
		for _, v := range tvals {
			num += v[i]
		}
		if num != "" {
			vals = append(vals, strings.TrimSpace(num))
		}
	}

	return vals
}

func getNums(vals []string) []string {
	result := []string{}
	maxI := 0
	for _, v := range vals {
		maxI = max(maxI, len(v))
	}

	return result
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
