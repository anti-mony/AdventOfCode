package main

import (
	"fmt"
	"log"

	"advent.of.code/list"
	"advent.of.code/util"
)

func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer P1: ", solveP1(input))

	fmt.Println("Answer P2: ", solveP2(input, 3))
}

func solveP2(sacks []RuckSack, groupSize int) int {
	result := 0
	for i := 0; i < len(sacks); i += groupSize {
		commonElem := findCommonElementInRucksacks(sacks[i : i+groupSize])
		result += getItemPriority(commonElem)

	}
	return result
}

func solveP1(sacks []RuckSack) int {
	result := 0

	for _, sack := range sacks {
		result += getItemPriority(findCommonElementInCompartments(sack))
	}

	return result
}

func findCommonElementInRucksacks(rss []RuckSack) rune {
	intersection := make([]rune, 0)
	intersection = append(intersection, rss[0].JoinCompartments()...)

	for _, rs := range rss[1:] {
		intersection = list.Intersection(intersection, rs.JoinCompartments())
	}

	return intersection[0]
}

func findCommonElementInCompartments(rs RuckSack) rune {
	intersection := list.Intersection(rs.Compartments[0].Items, rs.Compartments[1].Items)
	return intersection[0]
}

type Compartment struct {
	Items []rune
}

func (c Compartment) String() string {
	result := ""
	for _, ch := range c.Items {
		result += string(ch) + " "
	}
	result += " | "
	return result
}

type RuckSack struct {
	Compartments []Compartment
}

func (r RuckSack) JoinCompartments() []rune {
	result := make([]rune, 0)
	for _, c := range r.Compartments {
		result = append(result, c.Items...)
	}
	return result
}

func (r RuckSack) String() string {
	res := "\n"
	for _, compt := range r.Compartments {
		res += fmt.Sprintf("%v ", compt)
	}
	res += "\n"
	return res
}

func getItemPriority(item rune) int {
	priority := int(item)
	if priority > 96 {
		return priority - 96
	}
	return priority + 26 - 64
}

func parseInput(fileName string) ([]RuckSack, error) {
	lines, err := util.GetFileAsListOfStrings(fileName)
	if err != nil {
		return nil, err
	}

	result := make([]RuckSack, len(lines))

	for i, line := range lines {
		result[i] = makeRuckSackFromString(line, 2)
	}

	return result, nil
}

func makeRuckSackFromString(line string, compartments int) RuckSack {
	rucksack := RuckSack{
		Compartments: make([]Compartment, compartments),
	}

	numberOfItemsInEachCompartment := len(line) / compartments

	for charIdx, char := range line {
		rucksack.Compartments[charIdx/numberOfItemsInEachCompartment].Items = append(
			rucksack.Compartments[charIdx/numberOfItemsInEachCompartment].Items, char)
	}

	return rucksack
}
