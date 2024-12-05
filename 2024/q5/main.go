package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	ordering, manuals, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Asnwer 1: ", Q1(ordering, manuals))

	fmt.Println("Asnwer 2: ", Q2(ordering, manuals))
}

func Q1(orders map[int][]int, manuals [][]int) int {
	result := 0

	for i := 0; i < len(manuals); i++ {
		correct := true
		for j := len(manuals[i]) - 1; j >= 0; j-- {
			v, found := orders[manuals[i][j]]
			if found {
				for k := 0; k < j; k++ {
					for l := 0; l < len(v); l++ {
						if manuals[i][k] == v[l] {
							correct = false
						}
					}
				}
			}
		}
		if correct {
			result += manuals[i][int(len(manuals[i])/2)]
		}
	}

	return result
}

func Q2(orders map[int][]int, manuals [][]int) int {
	result := 0

	incorrects := make([]int, 0)
	for i := 0; i < len(manuals); i++ {
		correct := true
		for j := len(manuals[i]) - 1; j >= 0; j-- {
			v, found := orders[manuals[i][j]]
			if found {
				for k := 0; k < j; k++ {
					for l := 0; l < len(v); l++ {
						if manuals[i][k] == v[l] {
							correct = false
							// fmt.Println(v)
						}
					}
				}
			}
		}
		if !correct {
			incorrects = append(incorrects, i)
		}
	}

	for _, v := range incorrects {
		result += fixOrdering(orders, manuals[v])
	}

	return result
}

func fixOrdering(orders map[int][]int, manual []int) int {
	// just really needed to write the sort function and would have been fine for both 1 & 2
	slices.SortFunc(manual, func(a, b int) int {
		if v, found := orders[a]; found {
			for _, post := range v {
				if post == b {
					return -1
				}
			}
		}
		if v, found := orders[b]; found {
			for _, post := range v {
				if post == b {
					return 1
				}
			}
		}
		return 0
	})

	return manual[int(len(manual)/2)]
}

func parseInput(filename string) (map[int][]int, [][]int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, nil, err
	}

	orders := map[int][]int{}

	i := 0
	for i = 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			break
		}
		spls := strings.Split(line, "|")
		before, after := util.StringToNumber(spls[0]), util.StringToNumber(spls[1])

		_, found := orders[before]
		if found {
			orders[before] = append(orders[before], after)
		} else {
			orders[before] = []int{after}
		}
	}

	manuals := make([][]int, 0)
	for i := i + 1; i < len(lines); i++ {
		line := lines[i]
		nums, err := util.DelimitedStringOfNumbersToIntSlice(line)
		if err != nil {
			return nil, nil, err
		}
		manuals = append(manuals, nums)
	}

	return orders, manuals, nil
}
