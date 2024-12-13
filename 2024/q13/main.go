package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"

	"advent.of.code/grid"
	"advent.of.code/util"
)

const (
	COST_A = 3
	COST_B = 1
	// PRIZE_ADD = 10000000000000
	PRIZE_ADD = 0000000000000
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

func Q1(games []Game) int {
	result := 0

	for _, g := range games {
		cache := map[grid.Coordinate]int{}
		m := minCostToWinGame(g, grid.NewCoordinate(0, 0), 0, 0, cache)
		if m > 0 {
			result += m
		}
	}

	return result
}

func Q2(games []Game) int {
	result := 0

	for _, g := range games[:1] {
		cache := map[grid.Coordinate]int{
			grid.NewCoordinate(0, 0): 0,
			grid.NewCoordinate(1, 0): COST_A,
			grid.NewCoordinate(0, 1): COST_B,
			grid.NewCoordinate(1, 1): COST_A + COST_B,
		}
		a, b := 1, 1
		i, j := 0, 0

		for i <= g.T.X && j <= g.T.Y {
			fmt.Println(i, j)
			p1, p1Found := cache[grid.NewCoordinate(a-1, b)]
			p2, p2Found := cache[grid.NewCoordinate(a, b-1)]
			cost := math.MaxInt
			if p1Found && p2Found {
				if p1+COST_A < p2+COST_B {
					cost = p1 + COST_A
					i, j = i+g.A.X, j+g.A.Y
					a++
				} else {
					cost = p2 + COST_B
					i, j = i+g.B.X, j+g.B.Y
					b++
				}
			} else if p1Found {
				cost = p1 + COST_A
				i, j = i+g.A.X, j+g.A.Y
				a++
			} else if p2Found {
				cost = p2 + COST_B
				i, j = i+g.B.X, j+g.B.Y
				b++
			}
			cache[grid.NewCoordinate(a, b)] = cost
		}

		// for i <= g.T.X && j <= g.T.Y {
		// 	b, j := 0, 0
		// 	for b*g.B.X <= g.T.X+PRIZE_ADD && b*g.B.Y <= g.T.Y+PRIZE_ADD {
		// 		p1, p1Found := cache[grid.NewCoordinate(a-1, b)]
		// 		p2, p2Found := cache[grid.NewCoordinate(a, b-1)]
		// 		if p1Found && p2Found {
		// 			cache[grid.NewCoordinate(a, b)] = min(p1+COST_A, p2+COST_B)
		// 		} else if p1Found {
		// 			cache[grid.NewCoordinate(a, b)] = p1 + COST_A
		// 		} else if p2Found {
		// 			cache[grid.NewCoordinate(a, b)] = p2 + COST_B
		// 		} else {
		// 			cache[grid.NewCoordinate(a, b)] = a*COST_A + b*COST_B
		// 		}
		// 		b++
		// 	}
		// 	a++
		// }
		fmt.Println(g, cache)
	}

	return result
}

func minCostToWinGame(g Game, current grid.Coordinate, As, Bs int, cache map[grid.Coordinate]int) int {

	if v, found := cache[current]; found {
		return v
	}

	if current.X > g.T.X || current.Y > g.T.Y {
		return -1
	}
	if current.X == g.T.X && current.Y == g.T.Y {
		return As*COST_A + Bs*COST_B
	}
	costA := minCostToWinGame(g, current.Add(g.A), As+1, Bs, cache)
	costB := minCostToWinGame(g, current.Add(g.B), As, Bs+1, cache)

	cost := -1
	if costA < 0 && costB < 0 {
		cost = -1
	} else if costA < 0 {
		cost = costB
	} else if costB < 0 {
		cost = costA
	} else {
		cost = min(costA, costB)
	}

	cache[current] = cost

	return cost
}

type Game struct {
	A grid.Coordinate
	B grid.Coordinate
	T grid.Coordinate
}

func parseInput(filename string) ([]Game, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`[0-9]+`)

	result := make([]Game, 0)
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) != 0 {
			As := re.FindAllString(lines[i], -1)
			Bs := re.FindAllString(lines[i+1], -1)
			Ts := re.FindAllString(lines[i+2], -1)
			result = append(result, Game{
				A: grid.NewCoordinate(util.StringToNumber(As[0]), util.StringToNumber(As[1])),
				B: grid.NewCoordinate(util.StringToNumber(Bs[0]), util.StringToNumber(Bs[1])),
				T: grid.NewCoordinate(util.StringToNumber(Ts[0]), util.StringToNumber(Ts[1])),
			})
			i += 3
		}
	}

	return result, nil
}
