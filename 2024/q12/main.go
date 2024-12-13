package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/grid"
	"advent.of.code/list"
	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := util.ReadStringMatrixFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("A1: ", Q1(inp))
	fmt.Println("A2: ", Q2(inp))
}

type region struct {
	name      string
	area      int
	perimeter int
	sides     int
}

func Q1(garden [][]string) int {
	regions := []region{}
	seen := map[grid.Coordinate]bool{}

	for i := range len(garden) {
		for j := range len(garden[i]) {
			c := grid.NewCoordinate(i, j)
			if _, found := seen[c]; !found {
				regions = append(regions, exploreRegion(garden, c, seen))
			}
		}
	}

	result := 0
	for _, region := range regions {
		// fmt.Println(region)
		result += region.area * region.perimeter
	}

	return result
}

var cornerPairs = [][]grid.Direction{
	{grid.DirectionNorth, grid.DirectionEast, grid.DirectionNorthEast},
	{grid.DirectionEast, grid.DirectionSouth, grid.DirectionSouthEast},
	{grid.DirectionSouth, grid.DirectionWest, grid.DirectionSouthWest},
	{grid.DirectionWest, grid.DirectionNorth, grid.DirectionNorthWest},
}

func exploreRegion(garden [][]string, start grid.Coordinate, seen map[grid.Coordinate]bool) region {
	name := garden[start.X][start.Y]
	area := 1
	perimeter := 0
	sidesn := 0
	q := list.NewQueue[grid.Coordinate]()

	seen[start] = true
	q.Push(start)
	for q.Len() > 0 {
		current := q.Pop()
		sides := 4
		sides2 := 0
		for _, d := range grid.DIRECTIONS_STRAIGHT {
			next := current.Add(d)
			_, found := seen[next]
			if inBound(garden, next) && garden[current.X][current.Y] == garden[next.X][next.Y] {
				if !found {
					q.Push(next)
					seen[next] = true
					area++
				}
				sides--
			}
		}

		for _, prs := range cornerPairs {
			d1, d2, d3 := grid.DIRECTIONS[prs[0]], grid.DIRECTIONS[prs[1]], grid.DIRECTIONS[prs[2]]
			n1, n2, n3 := current.Add(d1), current.Add(d2), current.Add(d3)
			v1, v2, v3, cv := "", "", "", garden[current.X][current.Y]
			if inBound(garden, n1) {
				v1 = garden[n1.X][n1.Y]
			}
			if inBound(garden, n2) {
				v2 = garden[n2.X][n2.Y]
			}
			if inBound(garden, n3) {
				v2 = garden[n3.X][n3.Y]
			}
			if cv == v1 && cv == v2 && cv != v3 {
				fmt.Println(name, n1, n2, "yes")
				sides2++
			} else if cv != v1 && cv != v2 {
				fmt.Println(name, n1, n2, "no")
				sides2++
			}
		}

		sidesn += sides2
		perimeter += sides
	}

	return region{area: area, perimeter: perimeter, name: name, sides: sidesn}
}

func inBound(garden [][]string, c grid.Coordinate) bool {
	return c.X >= 0 && c.X < len(garden) && c.Y >= 0 && c.Y < len(garden[c.X])
}

func Q2(garden [][]string) int {
	regions := []region{}
	seen := map[grid.Coordinate]bool{}

	for i := range len(garden) {
		for j := range len(garden[i]) {
			c := grid.NewCoordinate(i, j)
			if _, found := seen[c]; !found {
				regions = append(regions, exploreRegion(garden, c, seen))
			}
		}
	}

	result := 0
	for _, region := range regions {
		fmt.Println(region)
		result += region.area * region.sides
	}

	return result
}
