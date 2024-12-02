package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"advent.of.code/util"
)

const MAX_DELTA = 3
const MIN_DELTA = 1

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer 1:", Q1(inp))

	fmt.Println("Answer 2:", Q2(inp))

}

func Q1(reports [][]int) int {
	safeReports := 0

	for i := range len(reports) {
		ascending := true
		safe := true
		for j := 1; j < len(reports[i]); j++ {
			if j == 1 && reports[i][j-1]-reports[i][j] > 0 {
				ascending = false
			}
			delta := reports[i][j-1] - reports[i][j]
			if ascending {
				if delta > 0 || (util.Abs(delta) < MIN_DELTA || util.Abs(delta) > MAX_DELTA) {
					safe = false
					break
				}
			} else {
				if delta < 0 || (util.Abs(delta) < MIN_DELTA || util.Abs(delta) > MAX_DELTA) {
					safe = false
					break
				}
			}
		}
		if safe {
			safeReports++
		}
	}

	return safeReports
}

func Q2(reports [][]int) int {
	safeReports := 0
	unsafeReports := make([]int, 0)
	for i := range len(reports) {
		ascending := true
		safe := true
		for k := 0; k < len(reports[i]); k++ {

		}
		for j := 1; j < len(reports[i]); j++ {
			if j == 1 && reports[i][j-1]-reports[i][j] > 0 {
				ascending = false
			}
			delta := reports[i][j-1] - reports[i][j]
			if ascending {
				if delta > 0 || (util.Abs(delta) < MIN_DELTA || util.Abs(delta) > MAX_DELTA) {
					safe = false
					unsafeReports = append(unsafeReports, i)
					break
				}
			} else {
				if delta < 0 || (util.Abs(delta) < MIN_DELTA || util.Abs(delta) > MAX_DELTA) {
					safe = false
					unsafeReports = append(unsafeReports, i)
					break
				}
			}
		}
		if safe {
			safeReports++
		}
	}

	for _, v := range unsafeReports {
		for k := 0; k < len(reports[v]); k++ {
			safe := true
			tmpReport := append(slices.Clone(reports[v][:k]), reports[v][k+1:]...)
			ascending := true
			for j := 1; j < len(tmpReport); j++ {
				if j == 1 && tmpReport[j-1]-tmpReport[j] > 0 {
					ascending = false
				}
				delta := tmpReport[j-1] - tmpReport[j]
				if ascending {
					if delta > 0 || (util.Abs(delta) < MIN_DELTA || util.Abs(delta) > MAX_DELTA) {
						safe = false
						break
					}
				} else {
					if delta < 0 || (util.Abs(delta) < MIN_DELTA || util.Abs(delta) > MAX_DELTA) {
						safe = false
						break
					}
				}
			}
			if safe {
				safeReports++
				break
			}
		}
	}

	return safeReports
}

func parseInput(filename string) ([][]int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([][]int, len(lines))

	for i, line := range lines {
		result[i], err = util.DelimitedStringOfNumbersToIntSlice(line)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
