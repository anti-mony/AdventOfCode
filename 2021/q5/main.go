package main

import (
	"fmt"
	"log"
	"regexp"

	"advent.of.code/grid"
	"advent.of.code/util"
)

func main() {
	inp, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer P1: ", solveQ1(inp))
	fmt.Println("Answer P2: ", solveQ2(inp))
}

func solveQ1(lines []Line) int {
	result := 0
	points := make(map[grid.Coordinate]int)
	for _, line := range lines {
		if line.Start.X == line.End.X {
			// Horizontal Line
			lS, lE := line.Start.Y, line.End.Y
			if lS > lE {
				lS, lE = lE, lS
			}
			for i := lS; i <= lE; i++ {
				c := grid.NewCoordinate(line.Start.X, i)
				if v, found := points[c]; found {
					points[c] = v + 1
				} else {
					points[c] = 1
				}
			}
		} else if line.Start.Y == line.End.Y {
			// Vertical Line
			lS, lE := line.Start.X, line.End.X
			if lS > lE {
				lS, lE = lE, lS
			}
			for i := lS; i <= lE; i++ {
				c := grid.NewCoordinate(i, line.Start.Y)
				if v, found := points[c]; found {
					points[c] = v + 1
				} else {
					points[c] = 1
				}
			}
		}
	}

	for _, v := range points {
		if v > 1 {
			result++
		}
	}

	return result
}

func solveQ2(lines []Line) int {
	result := 0
	points := make(map[grid.Coordinate]int)
	for _, line := range lines {
		dx, dy := 0, 0
		xD, yD := line.Start.X-line.End.X, line.Start.Y-line.End.Y
		if xD > 0 {
			dx = -1
		} else if xD < 0 {
			dx = 1
		}
		if yD > 0 {
			dy = -1
		} else if yD < 0 {
			dy = 1
		}
		sx, sy := line.Start.X, line.Start.Y
		for {
			c := grid.NewCoordinate(sx, sy)
			if v, found := points[c]; found {
				points[c] = v + 1
			} else {
				points[c] = 1
			}
			if sx == line.End.X && sy == line.End.Y {
				break
			}
			sx += dx
			sy += dy
		}
	}

	for _, v := range points {
		if v > 1 {
			result++
		}
	}

	return result
}

type Line struct {
	Start grid.Coordinate
	End   grid.Coordinate
}

func parseInput(filename string) ([]Line, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	res := make([]Line, len(lines))
	re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	for idx, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		res[idx] = Line{
			Start: grid.NewCoordinate(util.StringToNumber(matches[0][1]),
				util.StringToNumber(matches[0][2])),
			End: grid.NewCoordinate(util.StringToNumber(matches[0][3]),
				util.StringToNumber(matches[0][4])),
		}
	}

	return res, nil
}
