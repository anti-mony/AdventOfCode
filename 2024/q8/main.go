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
	fmt.Println(Q1(inp))
	fmt.Println(Q2(inp))
}

func Q1(antennas [][]string) int {
	antinodes := map[grid.Coordinate]bool{}

	ants := map[string][]grid.Coordinate{}

	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas[i]); j++ {
			if antennas[i][j] != "." {
				ants[antennas[i][j]] = append(ants[antennas[i][j]], grid.NewCoordinate(i, j))
			}
		}
	}

	for _, v := range ants {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				dx, dy := v[j].X-v[i].X, v[j].Y-v[i].Y
				pa1X, pa1Y := (v[j].X + dx), (v[j].Y + dy)
				if inBound(antennas, pa1X, pa1Y) {
					antinodes[grid.NewCoordinate(pa1X, pa1Y)] = true
				}
				pa2X, pa2Y := (v[i].X - dx), (v[i].Y - dy)
				if inBound(antennas, pa2X, pa2Y) {
					antinodes[grid.NewCoordinate(pa2X, pa2Y)] = true
				}
			}
		}
	}

	return len(antinodes)
}

func Q2(antennas [][]string) int {
	antinodes := map[grid.Coordinate]bool{}

	ants := map[string][]grid.Coordinate{}

	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas[i]); j++ {
			if antennas[i][j] != "." {
				ants[antennas[i][j]] = append(ants[antennas[i][j]], grid.NewCoordinate(i, j))
			}
		}
	}

	for _, v := range ants {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				antinodes[v[i]] = true
				antinodes[v[j]] = true
				dx, dy := v[j].X-v[i].X, v[j].Y-v[i].Y
				pa1X, pa1Y := (v[j].X + dx), (v[j].Y + dy)
				for inBound(antennas, pa1X, pa1Y) {
					antinodes[grid.NewCoordinate(pa1X, pa1Y)] = true
					pa1X, pa1Y = pa1X+dx, pa1Y+dy
				}
				pa2X, pa2Y := (v[i].X - dx), (v[i].Y - dy)
				for inBound(antennas, pa2X, pa2Y) {
					antinodes[grid.NewCoordinate(pa2X, pa2Y)] = true
					pa2X, pa2Y = pa2X-dx, pa2Y-dy
				}
			}
		}
	}

	return len(antinodes)
}

func inBound(antennas [][]string, i, j int) bool {
	return i >= 0 && i < len(antennas) && j >= 0 && j < len(antennas[i])
}

func parseInput(filename string) ([][]string, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}
	lab := make([][]string, len(lines))
	for i, line := range lines {
		lab[i] = util.StringToCharSlice(line)
	}

	return lab, nil
}
