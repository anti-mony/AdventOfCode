package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/grid"
	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer 1: ", Q1(inp))

	fmt.Println("Answer 2: ", Q2(inp))
}

func Q1(inp [][]string) int {
	xmas := 0

	for i := range len(inp) {
		for j := range len(inp[i]) {
			if inp[i][j] == "X" {
				for _, d := range grid.DIRECTIONS {
					ni, nj := i, j
					found := true
					for _, c := range []string{"M", "A", "S"} {
						ni += d.X
						nj += d.Y
						if ni >= 0 && ni < len(inp) && nj >= 0 && nj < len(inp[i]) {
							if inp[ni][nj] != c {
								found = false
								break
							}
						} else {
							found = false
						}
					}
					if found {
						xmas++
					}
				}
			}
		}
	}

	return xmas
}

func Q2(inp [][]string) int {
	xmas := 0

	for i := range len(inp) {
		for j := range len(inp[i]) {
			if inp[i][j] == "A" {
				tli, tlj := i-1, j-1
				tri, trj := i-1, j+1
				bri, brj := i+1, j+1
				bli, blj := i+1, j-1

				if !inBound(tli, tlj, len(inp), len(inp[i])) ||
					!inBound(tri, trj, len(inp), len(inp[i])) ||
					!inBound(bli, blj, len(inp), len(inp[i])) ||
					!inBound(bri, brj, len(inp), len(inp[i])) {
					continue
				}

				if inp[tli][tlj] == "M" && inp[tri][trj] == "M" && inp[bri][brj] == "S" && inp[bli][blj] == "S" {
					xmas++
				} else if inp[tli][tlj] == "S" && inp[tri][trj] == "M" && inp[bri][brj] == "M" && inp[bli][blj] == "S" {
					xmas++
				} else if inp[tli][tlj] == "S" && inp[tri][trj] == "S" && inp[bri][brj] == "M" && inp[bli][blj] == "M" {
					xmas++
				} else if inp[tli][tlj] == "M" && inp[tri][trj] == "S" && inp[bri][brj] == "S" && inp[bli][blj] == "M" {
					xmas++
				}

			}
		}
	}

	return xmas
}

func inBound(i, j, lenI, lenJ int) bool {
	return i >= 0 && i < lenI && j >= 0 && j < lenJ
}

func parseInput(filename string) ([][]string, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}
	inp := make([][]string, len(lines))
	for i, line := range lines {
		inp[i] = util.StringToCharSlice(line)
	}

	return inp, nil
}
