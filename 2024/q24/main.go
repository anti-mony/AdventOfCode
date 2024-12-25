package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"

	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	gates, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Q1(gates))
}

func Q1(gates map[string]Gate) int {
	maxZ := -1
	for k := range gates {
		if strings.HasPrefix(k, "z") {
			maxZ = max(maxZ, util.StringToNumber(k[1:]))
		}
	}

	result := 0
	for bit := range maxZ + 1 {
		start := fmt.Sprintf("z%02d", bit)
		c := calc(gates, gates[start])
		fmt.Print(c)
		result += int(math.Pow(float64(2), float64(bit))) * c
	}

	fmt.Println()

	return result
}

func Q2(gates map[string]Gate) string {

	return ""
}

func calc(gates map[string]Gate, current Gate) int {
	// fmt.Println(current)
	if current.Result == nil {
		return current.Value
	}

	input1 := calc(gates, gates[current.Input1])
	input2 := calc(gates, gates[current.Input2])

	return current.Result(input1, input2)
}

type Gate struct {
	Input1 string
	Input2 string
	Output string
	Value  int
	Result func(a, b int) int
}

func (g Gate) String() string {
	return fmt.Sprintf("[%s] * [%s] -> [%s(%d)]", g.Input1, g.Input2, g.Output, g.Value)
}

func parseInput(filename string) (map[string]Gate, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	gates := map[string]Gate{}
	i := 0
	xb, yb := []int{}, []int{}
	for ; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}
		spls := strings.Split(lines[i], ":")
		value := 0
		if strings.HasSuffix(spls[1], "1") {
			value = 1
		}
		if strings.Contains(spls[0], "x") {
			xb = append(xb, value)
		} else {
			yb = append(yb, value)
		}
		gates[spls[0]] = Gate{Output: spls[0], Value: value}
	}

	fmt.Println(BinaryToNumber(xb), BinaryToNumber(yb))

	fmt.Println("")
	i++

	re := regexp.MustCompile(`(\S+) (OR|XOR|AND) (\S+) -> (\S+)`)
	for ; i < len(lines); i++ {
		f := re.FindStringSubmatch(lines[i])
		gates[f[4]] = Gate{
			Output: f[4],
			Input1: f[1],
			Input2: f[3],
			Result: op[f[2]],
		}
	}
	return gates, nil
}

var op = map[string]func(int, int) int{
	"OR":  OR,
	"AND": AND,
	"XOR": XOR,
}

func AND(a, b int) int {
	if a == 1 && b == 1 {
		return 1
	}
	return 0
}

func OR(a, b int) int {
	if a == 1 || b == 1 {
		return 1
	}
	return 0
}

func XOR(a, b int) int {
	if a != b {
		return 1
	}
	return 0
}

func BinaryToNumber(b []int) int {
	r := 0
	for i := range len(b) {
		r += int(math.Pow(2, float64(i))) * b[i]
	}
	return r
}
