package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"advent.of.code/grid"
	"advent.of.code/util"
)

const (
	COST_A    = 3
	COST_B    = 1
	PRIZE_ADD = 10000000000000
	// PRIZE_ADD = 0000000000000
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

	fmt.Println(Q(inp, 0))
	fmt.Println(Q(inp, PRIZE_ADD))
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

func Q(games []Game, additive int) int {
	result := 0

	for _, game := range games {
		game.T.X += additive
		game.T.Y += additive
		nA := game.GetAPresses()
		nB := game.GetBPresses(nA)

		if nA*game.A.X+nB*game.B.X == game.T.X && nA*game.A.Y+nB*game.B.Y == game.T.Y {
			result += nA*3 + nB
		}
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

func (g Game) GetAPresses() int {
	return (g.T.X*g.B.Y - g.T.Y*g.B.X) / (g.A.X*g.B.Y - g.A.Y*g.B.X)
}

func (g Game) GetBPresses(nA int) int {
	return (g.T.Y - nA*g.A.Y) / g.B.Y
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
