package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
)

type shape [][]int
type region struct {
	rows, cols int
	qShapes    []int
}

type Input struct {
	shapes  []shape
	regions []region
}

type InputType Input

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	// for _, s := range inp.shapes[:1] {
	// 	util.PrintMatrix(s)
	// }

	// for _, r := range inp.regions {
	// 	fmt.Println(r)
	// }

	fmt.Println("Q1: ", Q1(inp))
	fmt.Println("Q2: ", Q2(inp))
}

func Q1(inp InputType) int {
	res := 0
	allso := map[int][][]grid.Coordinate{}
	for i, v := range inp.shapes {
		allso[i] = getShapeOrientations(v)
	}

	for _, r := range inp.regions {
		totalshapearea := 0
		for j, v := range r.qShapes {
			totalshapearea += len(allso[j][0]) * v
		}
		if totalshapearea <= r.cols*r.rows {
			res += 1
		}
	}

	return res

	for i, r := range inp.regions {
		fmt.Printf("\nRegion %d: %dx%d, pieces: %v\n", i, r.rows, r.cols, r.qShapes)
		if r.DoesFit(allso) {
			fmt.Println("YES ", i)
			res += 1
		} else {
			fmt.Println("NO  ", i)
		}
	}
	return res
}

func hash(in []grid.Coordinate) string {
	res := ""
	clone := slices.Clone(in)
	slices.SortStableFunc(clone, func(a, b grid.Coordinate) int {
		x := cmp.Compare(a.X, b.X)
		if x == 0 {
			return cmp.Compare(a.Y, b.Y)
		}
		return x
	})
	for _, c := range clone {
		res += fmt.Sprintf("%d%d", c.X, c.Y)
	}

	return res
}

func getShapeOrientations(s shape) [][]grid.Coordinate {
	shapeOptions := [][]grid.Coordinate{}
	for _, degree := range []int{0, 90, 180, 270} {
		r := util.Rotate(s, degree)
		shapeOptions = append(shapeOptions,
			grid.GetFilled(r, func(in int) bool { return in == 1 }),
			grid.GetFilled(util.FlipVertically(r), func(in int) bool { return in == 1 }),
			grid.GetFilled(util.FlipHorizontally(r), func(in int) bool { return in == 1 }),
		)
	}
	unique := [][]grid.Coordinate{}
	seen := list.NewSet[string]()

	for _, so := range shapeOptions {
		hsh := hash(so)
		if seen.Contains(hsh) {
			continue
		} else {
			unique = append(unique, so)
			seen.Add(hsh)
		}
	}
	return unique
}

func (r region) DoesFit(shapes map[int][][]grid.Coordinate) bool {
	surface := util.MakeMatrix[bool](r.rows, r.cols)
	pieces := make([]int, 0)
	for shapeIdx, qty := range r.qShapes {
		for k := 0; k < qty; k++ {
			pieces = append(pieces, shapeIdx)
		}
	}

	var recurse func(decided int, pieces []int) bool

	recurse = func(decided int, pieces []int) bool {
		if len(pieces) == 0 {
			return true
		}

		// Find the first cell we haven't decided about yet
		// "decided" is just an index into the flattened grid
		rows, cols := len(surface), len(surface[0])
		for decided < rows*cols && surface[decided/cols][decided%cols] {
			decided++
		}
		if decided == rows*cols {
			return false // no room left but pieces remain
		}

		row, col := decided/cols, decided%cols

		// Option 1: skip this cell (leave empty)
		surface[row][col] = true
		if recurse(decided+1, pieces) {
			surface[row][col] = false
			return true
		}
		surface[row][col] = false

		// Option 2: place a piece covering this cell
		tried := map[int]bool{}
		for i, shapeId := range pieces {
			if tried[shapeId] {
				continue
			}
			tried[shapeId] = true

			for _, orient := range shapes[shapeId] {
				for _, anchor := range orient {
					dr, dc := row-anchor.X, col-anchor.Y
					if canPlace(surface, orient, dr, dc) {
						place(surface, orient, dr, dc, true)
						remaining := append(pieces[:i:i], pieces[i+1:]...)
						if recurse(decided+1, remaining) {
							place(surface, orient, dr, dc, false)
							return true
						}
						place(surface, orient, dr, dc, false)
					}
				}
			}
		}

		return false
	}

	return recurse(0, pieces)
}

func findFirstEmpty(surface [][]bool) (int, int, bool) {
	R, C := len(surface), len(surface[0])
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if !surface[i][j] {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func canPlace(surface [][]bool, shape []grid.Coordinate, r, c int) bool {
	rows, cols := len(surface), len(surface[0])
	for _, cell := range shape {
		r, c := cell.X+r, cell.Y+c
		if r < 0 || r >= rows || c < 0 || c >= cols || surface[r][c] {
			return false
		}
	}
	return true
}

func place(surface [][]bool, shape []grid.Coordinate, r, c int, val bool) bool {
	for _, cell := range shape {
		r, c := cell.X+r, cell.Y+c
		surface[r][c] = val
	}
	return true
}

func Q2(inp InputType) int {
	return -1
}

func parseInput(filename string) (InputType, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return InputType{}, err
	}

	lineIdx := 0
	shapes := make([]shape, 0)
	regions := make([]region, 0)

	for lineIdx = 0; lineIdx < len(lines); lineIdx += 5 {
		if strings.Contains(lines[lineIdx], "x") {
			break
		}

		s := make([][]int, 3)
		for i := range 3 {
			s[i] = make([]int, 3)
		}

		si := 0
		for _, l := range lines[lineIdx+1 : lineIdx+4] {
			for li, v := range l {
				s[si][li] = 0
				if v == '#' {
					s[si][li] = 1
				}
			}
			si++
		}

		shapes = append(shapes, s)
	}

	for l := lineIdx; l < len(lines); l++ {
		splits := strings.Split(lines[l], ":")
		qS, err := util.SpaceSeparatedStringOfNumbersToIntSlice(splits[1])
		if err != nil {
			return InputType{}, err
		}
		splits2 := strings.Split(splits[0], "x")
		r := region{
			rows:    util.StringToNumber(splits2[1]),
			cols:    util.StringToNumber(splits2[0]),
			qShapes: qS,
		}

		regions = append(regions, r)
	}

	return InputType{
		shapes:  shapes,
		regions: regions,
	}, nil
}
