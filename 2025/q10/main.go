package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"advent.of.code/list"
	"advent.of.code/util"
)

type Machine struct {
	indicatorlights []int
	buttons         [][]int
	joltages        []int
}

func (m Machine) String() string {
	return fmt.Sprintf("%v | %v | %v", m.indicatorlights, m.buttons, m.joltages)
}

func (m Machine) GetTarget() int {
	res := 0

	for i := len(m.indicatorlights) - 1; i >= 0; i-- {
		res += int(math.Pow(2, float64(len(m.indicatorlights)-i-1))) * m.indicatorlights[i]
	}

	return res
}

func (m Machine) GetButton(index int) int {
	res := 0

	for i := len(m.buttons[index]) - 1; i >= 0; i-- {
		res += int(math.Pow(2, float64(len(m.buttons[index])-i-1))) * m.buttons[index][i]
	}

	return res
}

type InputType []Machine

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
	res := 0

	for _, m := range inp {
		r := m.GetFewestPresses()
		res += r
	}
	return res
}

func (m Machine) GetFewestPresses() int {
	type state struct {
		lights  int
		presses int
	}

	visited := map[int]bool{0: true}
	q := list.NewQueue[state]()
	q.Push(state{0, 0})

	for q.Len() > 0 {
		curr := q.Pop()
		if curr.lights == m.GetTarget() {
			return curr.presses
		}

		for i := range m.buttons {
			nxt := curr.lights ^ m.GetButton(i)
			if !visited[nxt] {
				q.Push(state{nxt, curr.presses + 1})
				visited[nxt] = true
			}
		}
	}

	return -1
}

func Q2(inp InputType) int {
	return -1
}

func parseInput(filename string) (InputType, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	res := make([]Machine, len(lines))

	for i, line := range lines {
		res[i] = parseLine(line)
	}

	return res, nil
}

func parseLine(in string) Machine {
	lightsEnd := strings.Index(in, "]")
	var lights []int
	var buttons [][]int

	for i := 1; i < lightsEnd; i++ {
		if in[i] == '.' {
			lights = append(lights, 0)
		} else {
			lights = append(lights, 1)
		}
	}
	in = in[lightsEnd+1:]
	start := strings.Index(in, "(")
	end := strings.Index(in, ")")
	for start != -1 && end != -1 {
		ns, err := util.DelimitedStringOfNumbersToIntSlice(in[start:end])
		if err != nil {
			panic(err)
		}
		netButtons := make([]int, len(lights))
		for _, v := range ns {
			netButtons[v] = 1
		}
		buttons = append(buttons, netButtons)
		in = in[end+1:]
		start = strings.Index(in, "(")
		end = strings.Index(in, ")")
	}

	joltages, err := util.DelimitedStringOfNumbersToIntSlice(in)
	if err != nil {
		panic(err)
	}

	return Machine{
		indicatorlights: lights,
		buttons:         buttons,
		joltages:        joltages,
	}
}

/*
1,0,0,0,1
0,1,1,0,1
0,0,0,1,0
*/
