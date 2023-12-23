package main

import (
	"fmt"
	"slices"
	"strings"
)

type Material int

const (
	Seed Material = iota + 1
	Soil
	Fertilizer
	Water
	Light
	Temperature
	Humidity
	Location
)

type Map struct {
	Source      Material
	Destination Material
	Ranges      []Range
}

type Range struct {
	SourceStart      int
	DestinationStart int
	Length           int
}

func (r Range) FindDestination(input int) (int, bool) {
	//	fmt.Println(input, r)
	if input >= r.SourceStart && input < r.SourceStart+r.Length {
		diff := r.DestinationStart - r.SourceStart
		return diff + input, true
	}

	return -1, false
}

type Situation struct {
	Seeds []int
	Maps  map[Material]Map
}

func q5sol() error {
	situation, err := getMap()
	if err != nil {
		return err
	}

	//for _, v := range situation.Maps {
	//	fmt.Println(v)
	//}

	fmt.Println()

	result := solveP2(situation)
	//result := processSeed(14, situation)

	fmt.Printf("Answer is %d", result)

	return nil
}

func solveP1(s Situation) int {
	locations := make([]int, 0)

	for _, seed := range s.Seeds {
		locations = append(locations, processSeed(seed, s))
	}

	// fmt.Println(locations)

	return slices.Min(locations)
}

func solveP2(s Situation) int {
	seeds := make([]int, 0)
	for i := 0; i < len(s.Seeds); i += 2 {
		for j := 0; j < s.Seeds[i+1]; j++ {
			seeds = append(seeds, s.Seeds[i]+j)
		}
	}
	fmt.Println(len(seeds))
	return solveP1(s)
}

func processSeed(seed int, s Situation) int {
	start := 1
	for {
		m, ok := s.Maps[Material(start)]
		if !ok {
			return seed
		}
		for _, r := range m.Ranges {
			if dest, ok := r.FindDestination(seed); ok {
				seed = dest
				break
			}
		}
		start += 1
	}
}

func getMap() (Situation, error) {
	result := Situation{
		Maps: make(map[Material]Map, 0),
	}

	lines, err := getFileAsListOfStrings("input5.txt")
	if err != nil {
		return result, err
	}

	seeds, err := SpaceSeparatedStringOfNumbersToIntSlice(lines[0][strings.Index(lines[0], ":")+1:])
	if err != nil {
		return result, err
	}
	result.Seeds = seeds

	currentMap := Map{0, 0, make([]Range, 0)}
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		if strings.Contains(line, "map") {
			result.Maps[currentMap.Source] = currentMap
			names := strings.Split(line[:strings.Index(line, " ")], "-")
			currentMap = Map{Source: MaterialFromString(names[0]), Destination: MaterialFromString(names[2]), Ranges: make([]Range, 0)}
			continue
		}
		rangeSplit := strings.Split(line, " ")
		currentMap.Ranges = append(currentMap.Ranges, Range{
			DestinationStart: StringToNumber(rangeSplit[0]),
			SourceStart:      StringToNumber(rangeSplit[1]),
			Length:           StringToNumber(rangeSplit[2]),
		})
	}
	result.Maps[currentMap.Source] = currentMap

	return result, nil
}

func MaterialFromString(in string) Material {
	switch in {
	case "seed":
		return Seed
	case "soil":
		return Soil
	case "fertilizer":
		return Fertilizer
	case "water":
		return Water
	case "light":
		return Light
	case "temperature":
		return Temperature
	case "humidity":
		return Humidity
	case "location":
		return Location

	}
	return 0
}
