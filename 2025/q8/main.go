package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"advent.of.code/list"
	"advent.of.code/util"
)

type C struct {
	x, y, z int
}

type DistancePair struct {
	a int
	b int
	d int
}

type InputType []C

type circuitJunctionCounts struct {
	cI int
	nJ int
}

const nCircuits = 1000

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

func distance(a, b C) int {
	return (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y) + (a.z-b.z)*(a.z-b.z)
}

func Q1(inp InputType) int {
	h := list.NewHeap[DistancePair](func(a, b DistancePair) bool {
		return a.d < b.d
	})

	circuitCount := 0

	circuits := make([]int, len(inp))
	for i := range len(inp) {
		circuits[i] = -1
	}

	for i := 0; i < len(inp); i++ {
		for j := i + 1; j < len(inp); j++ {
			h.Push(DistancePair{i, j, distance(inp[i], inp[j])})
		}
	}

	for range nCircuits {
		if h.IsEmpty() {
			break
		}
		least := h.Pop()
		if circuits[least.a] == -1 && circuits[least.b] == -1 {
			circuits[least.a], circuits[least.b] = circuitCount, circuitCount
			circuitCount++
		} else if circuits[least.a] == -1 && circuits[least.b] != -1 {
			circuits[least.a] = circuits[least.b]
		} else if circuits[least.a] != -1 && circuits[least.b] == -1 {
			circuits[least.b] = circuits[least.a]
		} else {
			lower := circuits[least.a]
			higher := circuits[least.b]
			if circuits[least.a] > circuits[least.b] {
				lower = circuits[least.b]
				higher = circuits[least.a]
			}
			for i := range len(circuits) {
				if circuits[i] == higher {
					circuits[i] = lower
				}
			}
		}
	}

	circuitCompCounter := make(map[int]int)
	for _, v := range circuits {
		if v == -1 {
			continue
		}
		circuitCompCounter[v] += 1
	}

	sortable := make([]circuitJunctionCounts, 0, 10)
	for k, v := range circuitCompCounter {
		sortable = append(sortable, circuitJunctionCounts{k, v})
	}

	slices.SortFunc(sortable, func(a, b circuitJunctionCounts) int {
		return -1 * cmp.Compare(a.nJ, b.nJ)
	})

	sLen := len(sortable)
	if sLen <= 1 {
		return sortable[0].nJ
	} else if sLen <= 2 {
		return sortable[0].nJ * sortable[1].nJ
	}

	return sortable[0].nJ * sortable[1].nJ * sortable[2].nJ
}

func Q2(inp InputType) int {
	h := list.NewHeap[DistancePair](func(a, b DistancePair) bool {
		return a.d < b.d
	})

	circuitCount := 0

	circuits := make([]int, len(inp))
	for i := range len(inp) {
		circuits[i] = -1
	}

	for i := 0; i < len(inp); i++ {
		for j := i + 1; j < len(inp); j++ {
			h.Push(DistancePair{i, j, distance(inp[i], inp[j])})
		}
	}

	for {
		if h.IsEmpty() {
			break
		}
		least := h.Pop()
		if circuits[least.a] == -1 && circuits[least.b] == -1 {
			circuits[least.a], circuits[least.b] = circuitCount, circuitCount
			circuitCount++
		} else if circuits[least.a] == -1 && circuits[least.b] != -1 {
			circuits[least.a] = circuits[least.b]
		} else if circuits[least.a] != -1 && circuits[least.b] == -1 {
			circuits[least.b] = circuits[least.a]
		} else {
			lower := circuits[least.a]
			higher := circuits[least.b]
			if circuits[least.a] > circuits[least.b] {
				lower = circuits[least.b]
				higher = circuits[least.a]
			}
			for i := range len(circuits) {
				if circuits[i] == higher {
					circuits[i] = lower
				}
			}
		}
		connected := true
		for i := 1; i < len(circuits); i++ {
			if circuits[i] != circuits[i-1] {
				connected = false
				break
			}
		}
		if connected {
			return inp[least.a].x * inp[least.b].x
		}
	}

	circuitCompCounter := make(map[int]int)
	for _, v := range circuits {
		if v == -1 {
			continue
		}
		circuitCompCounter[v] += 1
	}

	sortable := make([]circuitJunctionCounts, 0, 10)
	for k, v := range circuitCompCounter {
		sortable = append(sortable, circuitJunctionCounts{k, v})
	}

	slices.SortFunc(sortable, func(a, b circuitJunctionCounts) int {
		return -1 * cmp.Compare(a.nJ, b.nJ)
	})

	sLen := len(sortable)
	if sLen <= 1 {
		return sortable[0].nJ
	} else if sLen <= 2 {
		return sortable[0].nJ * sortable[1].nJ
	}

	return sortable[0].nJ * sortable[1].nJ * sortable[2].nJ
}

func parseInput(filename string) (InputType, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}
	coords := make([]C, len(lines))
	for i, line := range lines {
		splits := strings.Split(line, ",")
		x, _ := strconv.Atoi(splits[0])
		y, _ := strconv.Atoi(splits[1])
		z, _ := strconv.Atoi(splits[2])
		coords[i] = C{x, y, z}
	}

	return coords, nil
}
