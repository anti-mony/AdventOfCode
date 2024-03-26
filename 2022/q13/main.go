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

	fmt.Println("> Answer P1: ", solveP1(input))

	// fmt.Println("> Answer P2: ", solveP2(input))
}

func solveP1(pairs []Pair) int {
	result := 0
	for i, pair := range pairs {
		if compare(pair.Left, pair.Right, 1) < 0 {
			result += i + 1
			// fmt.Println("# In ORDER #", i+1, "# # #")
		}
	}

	return result
}

func compare(left, right Value, depth int) int {
	// fmt.Println(makeNLenString(depth), left, right)
	if left.List == nil {
		// Left is Singular
		if right.List == nil {
			return left.Singular - right.Singular
		} else {
			return compare(newListValue(left.Singular), right, depth+1)
		}
	} else {
		// Left is a list
		if right.List == nil {
			// Right is Singular
			return compare(left, newListValue(right.Singular), depth+1)
		}
	}

	i := 0
	for i < len(left.List) && i < len(right.List) {
		r := compare(left.List[i], right.List[i], depth+1)
		if r != 0 {
			return r
		}
		i++
	}

	return len(left.List) - len(right.List)
}

func makeNLenString(n int) string {
	r := ""
	for i := 0; i < n; i++ {
		r += ">"
	}
	return r
}

type Value struct {
	List     []Value
	Singular int
}

func newValue(singular *int, list []Value) Value {
	if singular != nil {
		return Value{Singular: *singular}
	}
	return Value{List: list}
}

func newListValue(v int) Value {
	return Value{
		List: []Value{{Singular: v}},
	}
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
