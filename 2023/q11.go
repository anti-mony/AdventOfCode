package main

import (
	"fmt"
	"strconv"
)

type inputq11 struct {
	grid          [][]string
	rowExtraSpace map[int]int
	colExtraSpace map[int]int
	galaxies      map[string]Coordinate
	gax           []string
}

func q11sol() error {
	input, err := getInputQ11("input11.small.txt")
	if err != nil {
		return err
	}

	// for _, v := range input.grid {
	// 	for _, k := range v {
	// 		fmt.Printf("%5s", k)
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println("Grid size:", len(input.grid), "X", len(input.grid[0]))

	// AdditiveP1 := 1
	AdditiveP2 := 1000000

	result := 0
	for i := 0; i < len(input.galaxies); i++ {
		for j := i + 1; j < len(input.galaxies); j++ {
			res := input.getDistanceBetweenv2(input.gax[i], input.gax[j], AdditiveP2)
			fmt.Printf("(%s,%s)->%d\n", input.gax[i], input.gax[j], res)
			result += res
		}
	}

	fmt.Println("Answer is : ", result)

	return nil
}

func (inp inputq11) getDistanceBetweenv2(start string, end string, ADDITIVE int) int {

	xDist := inp.galaxies[start].x - inp.galaxies[end].x
	if xDist < 0 {
		xDist *= -1
	}

	for k := range inp.rowExtraSpace {
		if inp.galaxies[start].x < k && k < inp.galaxies[end].x || inp.galaxies[start].x > k && k > inp.galaxies[end].x {
			xDist += ADDITIVE - 1
		}
	}

	yDist := inp.galaxies[start].y - inp.galaxies[end].y
	if yDist < 0 {
		yDist *= -1
	}

	for k := range inp.colExtraSpace {
		if inp.galaxies[start].y < k && k < inp.galaxies[end].y || inp.galaxies[start].y > k && k > inp.galaxies[end].y {
			yDist += ADDITIVE - 1
		}
	}

	return xDist + yDist
}

func getInputQ11(filename string) (inputq11, error) {
	lines, err := getFileAsListOfStrings(filename)
	if err != nil {
		return inputq11{}, err
	}

	grid := make([][]string, len(lines))

	galaxies := make(map[string]Coordinate, 0)
	gax := make([]string, 0)

	rowEmptyLines := make(map[int]int)
	colEmptyLines := make(map[int]int)
	numGalaxies := 0

	for i := 0; i < len(lines); i++ {
		empty := true
		grid[i] = make([]string, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			grid[i][j] = string(lines[i][j])
			if lines[i][j] == '#' {
				empty = false
				numGalaxies++
				g := strconv.Itoa(numGalaxies)
				grid[i][j] = g
				galaxies[g] = Coordinate{i, j}
				gax = append(gax, g)
			}
		}
		if empty {
			rowEmptyLines[i] = 2
		}
	}

	for i := 0; i < len(grid); i++ {
		empty := true
		for j := 0; j < len(grid); j++ {
			if grid[j][i] != "." {
				empty = false
			}
		}
		if empty {
			colEmptyLines[i] = 2
		}
	}

	return inputq11{grid, rowEmptyLines, colEmptyLines, galaxies, gax}, nil

}
