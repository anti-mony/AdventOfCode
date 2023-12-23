package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	Time        int
	MaxDistance int
}

func q6sol() error {

	inp, err := parseInput()
	if err != nil {
		return err
	}

	result := 1
	for _, in := range inp {
		r := getwaystoWin(in)
		result *= r
	}

	fmt.Printf("Answer is P1 %d \n", result)

	newTimeS := ""
	newDistS := ""

	for i := 0; i < len(inp); i++ {
		newTimeS += strconv.Itoa(inp[i].Time)
		newDistS += strconv.Itoa(inp[i].MaxDistance)
	}

	newTime, _ := strconv.Atoi(newTimeS)
	newDist, _ := strconv.Atoi(newDistS)

	fmt.Printf("Answer P2 : %d", getwaystoWin(Input{newTime, newDist}))

	return nil
}

func getwaystoWin(in Input) int {
	midTime := in.Time / 2
	ways := 0
	for {
		if (midTime)*(in.Time-midTime) > in.MaxDistance {
			ways++
			midTime--
		} else {
			break
		}
	}

	if in.Time%2 == 0 {
		return ways*2 - 1
	}

	return ways * 2
}

func parseInput() ([]Input, error) {

	lines, err := getFileAsListOfStrings("input6.txt")
	if err != nil {
		return nil, err
	}

	times, err := SpaceSeparatedStringOfNumbersToIntSlice(strings.Split(lines[0], ":")[1])
	if err != nil {
		return nil, err
	}
	distances, err := SpaceSeparatedStringOfNumbersToIntSlice(strings.Split(lines[1], ":")[1])
	if err != nil {
		return nil, err
	}

	if len(times) != len(distances) {
		return nil, fmt.Errorf("input parse error")
	}

	result := make([]Input, len(times))

	for i := 0; i < len(times); i++ {
		result[i] = Input{
			times[i],
			distances[i],
		}
	}

	return result, nil

}
