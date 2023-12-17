package main

import (
	"fmt"
	"slices"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

type Round struct {
	Red   int
	Green int
	Blue  int
}

func (r *Round) IsPossible() bool {
	return r.Red <= MAX_RED && r.Green <= MAX_GREEN && r.Blue <= MAX_BLUE
}

type Game struct {
	ID     int
	Rounds []Round
}

func q2sol() error {

	games, err := getGamesFromFile()
	if err != nil {
		return err
	}

	resultP1 := 0
	resultP2 := 0

	for _, game := range games {
		isValid := true
		reds := make([]int, len(game.Rounds))
		greens := make([]int, len(game.Rounds))
		blues := make([]int, len(game.Rounds))
		for i, round := range game.Rounds {
			if !round.IsPossible() {
				isValid = false
			}
			reds[i] = round.Red
			greens[i] = round.Green
			blues[i] = round.Blue
		}

		resultP2 += slices.Max(reds) * slices.Max(blues) * slices.Max(greens)

		if isValid {
			resultP1 += game.ID
		}

	}

	fmt.Printf("Part 1: Answer is: %d | Part 2: Answer is: %d \n", resultP1, resultP2)

	return nil
}

func getGamesFromFile() ([]Game, error) {

	lines, err := getFileAsListOfStrings("input2.txt")
	if err != nil {
		return nil, err
	}

	games := make([]Game, 0)

	for _, line := range lines {
		games = append(games, makeGameFromString(line))
	}

	return games, nil

}

func makeGameFromString(input string) Game {
	gameSplit := strings.Split(input, ":")

	gameIDByte := strings.Split(strings.TrimSpace(gameSplit[0]), " ")[1]
	gameID := StringToNumber(string(gameIDByte))

	roundsString := strings.Split(gameSplit[1], ";")

	rounds := make([]Round, 0)

	for _, roundStr := range roundsString {
		round := Round{}
		cubes := strings.Split(roundStr, ",")
		for _, cube := range cubes {
			color := strings.Split(strings.TrimSpace(cube), " ")
			switch string(color[1]) {
			case "red":
				round.Red = StringToNumber(string(color[0]))
			case "blue":
				round.Blue = StringToNumber(string(color[0]))
			case "green":
				round.Green = StringToNumber(string(color[0]))
			}

		}
		rounds = append(rounds, round)
	}

	return Game{ID: gameID, Rounds: rounds}

}
