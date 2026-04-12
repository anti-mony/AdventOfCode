package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"advent.of.code/grid"
	"advent.of.code/util"
)

type InputType []grid.Coordinate

type rectangle struct {
	x1, x2, y1, y2 int
}

func (r rectangle) Area() int {
	return (r.x2 - r.x1 + 1) * (r.y2 - r.y1 + 1)
}

func (r rectangle) Shrink(factor int) rectangle {
	return rectangle{r.x1 + factor, r.x2 - factor, r.y1 + factor, r.y2 - factor}
}

func (r rectangle) Overlaps(other rectangle) bool {
	return max(r.x1, other.x1) <= min(r.x2, other.x2) && max(r.y1, other.y1) <= min(r.y2, other.y2)
}

func newRectange(a, b grid.Coordinate) rectangle {
	x1, x2 := a.X, b.X
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	y1, y2 := a.Y, b.Y
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return rectangle{x1, x2, y1, y2}
}

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
	largest := 0
	for i := 0; i < len(inp); i++ {
		for j := i + 1; j < len(inp); j++ {
			largest = max(largest, (util.Abs(inp[i].X-inp[j].X)+1)*(util.Abs(inp[i].Y-inp[j].Y)+1))
		}
	}

	return largest
}

func Q2(inp InputType) int {
	N := len(inp)
	rects := []rectangle{}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			rects = append(rects, newRectange(inp[i], inp[j]))
		}
	}

	slices.SortFunc(rects, func(a, b rectangle) int {
		return -1 * cmp.Compare(a.Area(), b.Area())
	})

	// 1D rectanges can be edges
	edges := []rectangle{}
	for i, v := range inp {
		edges = append(edges, newRectange(v, inp[(i+1)%N]))
	}

	for _, r := range rects {
		inner := r.Shrink(1)
		if inner.x1 > inner.x2 || inner.y1 > inner.x2 {
			continue
		}

		overlaps := false
		for _, edge := range edges {
			if edge.Overlaps(inner) {
				overlaps = true
				break
			}
		}

		if !overlaps {
			return r.Area()
		}
	}

	return -1
}

func parseInput(filename string) (InputType, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}
	inp := make([]grid.Coordinate, len(lines))
	for i, line := range lines {
		tmp := strings.Split(line, ",")
		inp[i] = grid.NewCoordinate(util.StringToNumber(tmp[1]), util.StringToNumber(tmp[0]))
	}

	return inp, nil
}
