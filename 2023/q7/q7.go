package main

import (
	"fmt"
	"sort"
	"strings"
)

func q7sol() error {
	games, err := getGames("input7.txt")
	if err != nil {
		return err
	}

	sort.Slice(games, func(i, j int) bool {
		return IsLessHand(games[i].hand, games[j].hand)
	})

	result := 0
	for i, g := range games {
		result += g.bid * (i + 1)
	}

	fmt.Println("Answer is ", result)

	return nil
}

var cardValue = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 1,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type Hand struct {
	Cards      []string
	CardGroups map[string]int
}

func (h Hand) getScore() int {
	N := len(h.CardGroups)

	nJ, _ := h.CardGroups["J"]

	if nJ > 0 {
		switch N {
		// Five of a kind
		case 1:
			return 7
		case 2:
			// 6+1 || 5+2
			return 7
		case 3:
			for _, v := range h.CardGroups {
				if v == 3 {
					return 6
				} else if v == 2 {
					if nJ == 2 {
						return 3 + 3
					}
					if nJ == 1 {
						return 3 + 2
					}
				}
			}
		case 4:
			return 2 + 1 + 1

		case 5:
			return 1 + 1
		}
	}

	switch N {
	// Five of a kind
	case 1:
		return 7
	case 2:
		for _, v := range h.CardGroups {
			// Four of a kind
			if v == 1 || v == 4 {
				return 6
			} else { // Full House
				return 5
			}
		}
	case 3:
		for _, v := range h.CardGroups {
			if v == 3 {
				// Three of a kind
				return 4
			} else if v == 2 {
				// Two Pair
				return 3
			}
		}
	case 4:
		// One Pair
		return 2

	case 5:
		// High Card
		return 1
	}
	return 0
}

// true if h1 < h2
func IsLessHand(h1, h2 Hand) bool {
	h1Score := h1.getScore()
	h2Score := h2.getScore()

	if h1Score > h2Score {
		return false
	} else if h2Score > h1Score {
		return true
	}

	for i := 0; i < len(h1.Cards); i++ {
		c1v := cardValue[h1.Cards[i]]
		c2v := cardValue[h2.Cards[i]]
		if c2v > c1v {
			return true
		} else if c1v > c2v {
			return false
		}
	}

	return false
}

type game struct {
	hand Hand
	bid  int
}

func makeHand(in []string) Hand {
	grps := make(map[string]int, 0)
	for _, card := range in {
		if n, ok := grps[card]; ok {
			grps[card] = n + 1
		} else {
			grps[card] = 1
		}
	}
	return Hand{in, grps}
}

func getGames(filename string) ([]game, error) {
	lines, err := getFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([]game, 0)
	for _, line := range lines {
		splits := strings.Split(line, " ")
		bid := StringToNumber(splits[1])
		cards := make([]string, len(splits[0]))
		for i, v := range splits[0] {
			cards[i] = string(v)
		}
		result = append(result, game{makeHand(cards), bid})
	}
	return result, nil
}
