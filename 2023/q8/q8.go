package main

import (
	"fmt"
	"regexp"
)

func q8sol() error {

	dmap, err := getDesertMapWithInstructions("input8.txt")
	if err != nil {
		return err
	}

	result := dmap.getStepsWithNStarts()

	fmt.Printf("Answer is %d \n", result)
	return nil
}

type DesertMap struct {
	Instructions string
	DMap         map[string]map[string]string
}

func (d DesertMap) getStepsWithNStarts() int {

	starts := make([]string, 0)
	for k := range d.DMap {
		if string(k[len(k)-1]) == "A" {
			starts = append(starts, k)
		}
	}

	fmt.Println(starts)

	steps := make([]int, len(starts))
	for i, s := range starts {
		steps[i] = d.getSteps(s, true)
	}

	fmt.Println(steps)

	result := 1
	for _, s := range steps {
		result = LCM(result, s)
	}

	return result
}

func allEndInZ(inp []string) bool {
	for _, s := range inp {
		if string(s[len(s)-1]) != "Z" {
			return false
		}
	}
	return true
}

func (d DesertMap) getSteps(startPos string, stopEndsInZ bool) int {
	steps := 0
	nInstructions := len(d.Instructions)

	start := startPos

	for {
		if stopEndsInZ && string(start[len(start)-1]) == "Z" {
			return steps
		} else if start == "ZZZ" {
			return steps
		}
		start = d.DMap[start][string(d.Instructions[steps%nInstructions])]
		steps++
	}

	return steps
}

func getDesertMapWithInstructions(fileName string) (DesertMap, error) {
	lines, err := getFileAsListOfStrings(fileName)
	if err != nil {
		return DesertMap{}, err
	}

	instructions := lines[0]

	dmap := make(map[string]map[string]string)
	re := regexp.MustCompile("[A-Z0-9]+")

	for _, line := range lines[2:] {
		v := re.FindAllString(line, -1)
		m := map[string]string{
			"L": v[1],
			"R": v[2],
		}
		dmap[v[0]] = m
	}

	return DesertMap{
		instructions, dmap,
	}, nil
}
