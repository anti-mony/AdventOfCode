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

	util.PrintMatrix(inp)

	a1, a2 := Q(inp)

	fmt.Printf("A1:%4d | A2:%4d\n", a1, a2)
}

type region struct {
	name      string
	area      int
	perimeter int
	sides     int
}

func (r region) String() string {
	return fmt.Sprintf("Name: %v | Area: %3d | Perimeter: %3d | Sides: %3d", r.name, r.area, r.perimeter, r.sides)
}

func Q(garden [][]string) (int, int) {
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

	result1 := 0
	result2 := 0
	for _, region := range regions {
		fmt.Println(region)
		result1 += region.area * region.perimeter
		result2 += region.area * region.sides
	}

	return result1, result2
}

var cornerPairs = [][]grid.Direction{
	{grid.DirectionNorth, grid.DirectionEast},
	{grid.DirectionEast, grid.DirectionSouth},
	{grid.DirectionSouth, grid.DirectionWest},
	{grid.DirectionWest, grid.DirectionNorth},
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

		for _, cp := range cornerPairs {
			x, y := grid.DIRECTIONS[cp[0]], grid.DIRECTIONS[cp[1]]
			z := grid.NewCoordinate(x.X+y.X, x.Y+y.Y)
			nx, ny, nz := current.Add(x), current.Add(y), current.Add(z)

			if inBound(garden, nx) && inBound(garden, ny) {
				if garden[current.X][current.Y] != garden[nx.X][nx.Y] && garden[current.X][current.Y] != garden[ny.X][ny.Y] {
					sides2++
				} else if garden[current.X][current.Y] == garden[nx.X][nx.Y] && garden[current.X][current.Y] == garden[ny.X][ny.Y] && inBound(garden, nz) {
					if garden[current.X][current.Y] != garden[nz.X][nz.Y] {
						sides2++
					}
				}
			} else if !inBound(garden, nx) && !inBound(garden, ny) {
				sides2++
			} else if inBound(garden, nx) && garden[current.X][current.Y] != garden[nx.X][nx.Y] {
				sides2++
			} else if inBound(garden, ny) && garden[current.X][current.Y] != garden[ny.X][ny.Y] {
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
