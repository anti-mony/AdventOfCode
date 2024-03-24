package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"advent.of.code/util"
)

func main() {
	fileName := "input.small.txt"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	input, err := parseInput(fileName)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range input {
		fmt.Println(p)
	}

	// fmt.Println("> Answer P1: ", solveP1(input))

	// fmt.Println("> Answer P2: ", solveP2(input))
}

type Value struct {
	List     []Value
	Singular int
}

func (v Value) String() string {
	if v.List == nil {
		return fmt.Sprintf("%d", v.Singular)
	}
	return fmt.Sprintf("%v", v.List)
}

type Pair struct {
	Left, Right Value
}

func (p Pair) String() string {
	return fmt.Sprintf("First : %v\nSecond: %v\n", p.Left, p.Right)
}

func parseInput(filename string) ([]Pair, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([]Pair, 0)
	for i := 0; i < len(lines); i += 3 {
		result = append(result, Pair{
			Left:  getValueFromString(lines[i]),
			Right: getValueFromString(lines[i+1]),
		})
	}

	return result, nil
}

func getValueFromString(line string) Value {
	line = line[1 : len(line)-1]
	v := Value{
		List: make([]Value, 0),
	}

	for _, sp := range Split(line) {
		if strings.HasPrefix(sp, "[") {
			v.List = append(v.List, getValueFromString(sp))
		} else {
			v.List = append(v.List, Value{Singular: util.StringToNumber(sp)})
		}
	}

	return v
}

func Split(line string) []string {
	result := []string{}
	i := 0
	for i < len(line) {
		if string(line[i]) == "[" {
			st := i
			ct := 1
			i++
			for ct != 0 && i < len(line) {
				if string(line[i]) == "[" {
					ct++
				} else if string(line[i]) == "]" {
					ct--
				}
				i++
			}
			result = append(result, line[st:i])
		} else if string(line[i]) != "," {
			st := i
			i++
			for i < len(line) && string(line[i]) != "," {
				i++
			}
			result = append(result, line[st:i])
		}
		i++
	}

	return result
}
