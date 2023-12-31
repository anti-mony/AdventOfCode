package main

import (
	"fmt"
	"strings"
)

type Lottery struct {
	Winning map[int]int
	Yours   []int
	Copies  int
}

func q4sol() error {

	cards, err := getLotteries()
	if err != nil {
		return err
	}

	result := playCardGames(cards)

	fmt.Printf("Answer is %d \n", result)
	return nil
}

func playCardGames(cards []Lottery) int {
	result := 0

	for idx, card := range cards {
		found := 0
		for _, yN := range card.Yours {
			if _, ok := card.Winning[yN]; ok {
				found += 1
			}
		}
		for j := 1; j <= found && j+idx < len(cards); j++ {
			cards[j+idx].Copies += card.Copies
		}
	}

	for _, card := range cards {
		result += card.Copies
	}

	return result
}

func findPoints(cards []Lottery) int {
	result := 0

	for _, card := range cards {
		isFirst := true
		score := 0
		for _, yN := range card.Yours {
			if _, ok := card.Winning[yN]; ok {
				if isFirst {
					score += 1
					isFirst = false
				} else {
					score *= 2
				}
			}
		}
		result += score
	}

	return result
}

func getLotteries() ([]Lottery, error) {

	result := make([]Lottery, 0)

	lines, err := getFileAsListOfStrings("input4.txt")
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		split1 := strings.Split(line, ":")
		split2 := strings.Split(split1[1], "|")
		winNumbers, err := SpaceSeparatedStringOfNumbersToIntSlice(split2[0])
		if err != nil {
			return nil, err
		}
		winMap := make(map[int]int, 0)
		for _, winNumber := range winNumbers {
			winMap[winNumber] = winNumber
		}

		yourNumbers, err := SpaceSeparatedStringOfNumbersToIntSlice(split2[1])
		if err != nil {
			return nil, err
		}

		result = append(result, Lottery{winMap, yourNumbers, 1})

	}

	return result, nil

}
