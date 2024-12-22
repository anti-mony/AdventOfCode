package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	codes, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("A1 ", Q1(codes))
}

type Stroke struct {
	c      grid.Coordinate
	nSteps int
}

type Path struct {
	c    grid.Coordinate
	step string
}

type CoordinatePair struct {
	Start grid.Coordinate
	End   grid.Coordinate
}

type Store struct {
	Start grid.Coordinate
	End   grid.Coordinate
	Paths [][]string
}

var NUMPAD = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"#", "0", "A"},
}

var NUMPAD_P = map[string]grid.Coordinate{
	"0": grid.NewCoordinate(3, 1),
	"A": grid.NewCoordinate(3, 2),
	"1": grid.NewCoordinate(2, 0),
	"2": grid.NewCoordinate(2, 1),
	"3": grid.NewCoordinate(2, 2),
	"4": grid.NewCoordinate(1, 0),
	"5": grid.NewCoordinate(1, 1),
	"6": grid.NewCoordinate(1, 2),
	"7": grid.NewCoordinate(0, 0),
	"8": grid.NewCoordinate(0, 1),
	"9": grid.NewCoordinate(0, 2),
}

var KEYPAD = [][]string{
	{"#", "^", "A"},
	{"<", "v", ">"},
}

var KEYPAD_P = map[string]grid.Coordinate{
	"^": grid.NewCoordinate(0, 1),
	"A": grid.NewCoordinate(0, 2),
	"<": grid.NewCoordinate(1, 0),
	"v": grid.NewCoordinate(1, 1),
	">": grid.NewCoordinate(1, 2),
}

var numpadShortestPaths = map[CoordinatePair][]string{}
var keypadShortestPaths = map[CoordinatePair][]string{}
var result = []string{}

func Q1(codes []string) int {

	FindAllPairShortestPaths(NUMPAD_P, NUMPAD, numpadShortestPaths)
	FindAllPairShortestPaths(KEYPAD_P, KEYPAD, keypadShortestPaths)

	result := 0
	for _, code := range codes {
		numpadPaths := list.Dedupe(NumpadCodeToSequence(code))
		keypadPaths := []string{}
		secondKeypadPaths := []string{}
		for _, p := range numpadPaths {
			r := KeypadCodeToSequence(p)
			keypadPaths = append(keypadPaths, r...)
		}
		keypadPaths = list.Dedupe(keypadPaths)
		for _, p := range keypadPaths {
			r := KeypadCodeToSequence(p)
			secondKeypadPaths = append(secondKeypadPaths, r...)
		}

		minl := math.MaxInt
		for _, r := range secondKeypadPaths {
			minl = min(minl, len(r))
		}
		n := util.StringToNumber(code[:len(code)-1])
		fmt.Println(code, minl, n)
		result += n * minl
	}

	return result
}

func NumpadCodeToSequence(code string) []string {
	result := numpadShortestPaths[CoordinatePair{Start: NUMPAD_P["A"], End: NUMPAD_P[string(code[0])]}]

	for i := 1; i < len(code); i++ {
		start := NUMPAD_P[string(code[i-1])]
		end := NUMPAD_P[string(code[i])]
		new := []string{}
		for _, ep := range result {
			for _, p := range numpadShortestPaths[CoordinatePair{Start: start, End: end}] {
				new = append(new, ep+p)
			}
		}
		result = new
	}

	return result
}

func KeypadCodeToSequence(code string) []string {
	result := keypadShortestPaths[CoordinatePair{Start: KEYPAD_P["A"], End: KEYPAD_P[string(code[0])]}]

	for i := 1; i < len(code); i++ {
		start := KEYPAD_P[string(code[i-1])]
		end := KEYPAD_P[string(code[i])]
		newPaths := []string{}
		for _, ep := range result {
			for _, p := range keypadShortestPaths[CoordinatePair{Start: start, End: end}] {
				newPaths = append(newPaths, ep+p)
			}
		}
		result = newPaths
	}

	return result
}

func FindAllPairShortestPaths(options map[string]grid.Coordinate, pad [][]string, resStore map[CoordinatePair][]string) {
	for _, v1 := range options {
		for _, v2 := range options {
			FindShortestPaths(pad, v1, v2)
			resStore[CoordinatePair{
				Start: v1,
				End:   v2,
			}] = slices.Clone(result)
			result = []string{}
		}
	}
}

func FindShortestPaths(pad [][]string, start, end grid.Coordinate) {
	prevs := map[grid.Coordinate][]Path{}
	seen := map[grid.Coordinate]int{}

	hp := list.NewHeap[Stroke](func(a, b Stroke) bool {
		return a.nSteps <= b.nSteps
	})
	hp.Push(Stroke{c: start})

	for hp.Len() > 0 {
		current := hp.Pop()
		if _, found := seen[current.c]; found {
			continue
		}
		seen[current.c] = current.nSteps

		for dir := range grid.DIRECTIONS_STRAIGHT {
			next := current.c.MoveTowards(dir)
			_, found := seen[next]
			if inBound(pad, next) && !found && pad[next.X][next.Y] != "#" {

				hp.Push(Stroke{
					c:      next,
					nSteps: current.nSteps + 1,
				})
				prevs[next] = append(prevs[next], Path{
					c:    current.c,
					step: grid.DirectionToRLUD(dir),
				})
			}

		}
	}

	findAllPaths(prevs, []string{}, end, start)
}

func findAllPaths(adjacenyList map[grid.Coordinate][]Path, currentPath []string, current, end grid.Coordinate) {
	if current == end {
		slices.Reverse(currentPath)
		result = append(result, strings.Join(append(currentPath, "A"), ""))
		return
	}

	for _, adj := range adjacenyList[current] {
		findAllPaths(adjacenyList, append(currentPath, adj.step), adj.c, end)
	}
}

func inBound(pad [][]string, c grid.Coordinate) bool {
	return c.X >= 0 && c.X < len(pad) && c.Y >= 0 && c.Y < len(pad[c.X])
}
