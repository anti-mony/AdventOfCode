package main

import (
	"fmt"
)

var (
	sides = [][]int{
		{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1},
	}
)

type Position struct {
	X int
	Y int
}

func q3sol() error {

	schematic, err := makeSchematic()
	if err != nil {
		return err
	}

	//	result := addPartNumbers(schematic)

	//	fmt.Printf("Part1 answer is: %d \n", result)

	result := getGearRatio(schematic)

	fmt.Printf("Part 2 answer is: %d \n", result)

	return nil
}

func makeSchematic() ([][]string, error) {
	schematic := [][]string{}

	lines, err := getFileAsListOfStrings("input3.txt")
	if err != nil {
		return nil, err
	}

	for idx, line := range lines {
		schematic = append(schematic, make([]string, len(line)))
		for i, char := range line {
			schematic[idx][i] = string(char)
		}
	}

	return schematic, nil
}

func addPartNumbers(schematic [][]string) int {
	result := 0

	maxI := len(schematic)
	for i, _ := range schematic {
		for j, _ := range schematic[i] {
			maxJ := len(schematic[i])
			if schematic[i][j] == "." || isStringNumber(schematic[i][j]) {
				continue
			}
			schematic[i][j] = "."
			for _, side := range sides {
				wi, wj := i+side[0], j+side[1]
				if isValidIndex(wi, wj, maxI, maxJ) {
					partNumber := ""
					if isStringNumber(schematic[wi][wj]) {
						for tj := wj; tj >= 0; tj-- {
							if isStringNumber(schematic[wi][tj]) {
								partNumber = schematic[wi][tj] + partNumber
								schematic[wi][tj] = "."
							} else {
								break
							}
						}

						for tj := wj + 1; tj < maxJ; tj++ {
							if isStringNumber(schematic[wi][tj]) {
								partNumber = partNumber + schematic[wi][tj]
								schematic[wi][tj] = "."
							} else {
								break
							}
						}
					}
					result += StringToNumber(partNumber)
				}
			}
		}
	}

	return result
}

func getGearRatio(schematic [][]string) int {
	result := 0

	maxI := len(schematic)

	for i, _ := range schematic {
		for j, _ := range schematic[i] {
			maxJ := len(schematic[i])
			if schematic[i][j] != "*" {
				continue
			}
			positions := make(map[Position]string, 0)
			parts := make([]int, 0)
			for _, side := range sides {
				wi, wj := i+side[0], j+side[1]
				if isValidIndex(wi, wj, maxI, maxJ) {
					if isStringNumber(schematic[wi][wj]) {
						partNumber := ""
						for tj := wj; tj >= 0; tj-- {
							if isStringNumber(schematic[wi][tj]) {
								partNumber = schematic[wi][tj] + partNumber
								positions[Position{X: wi, Y: tj}] = schematic[wi][tj]
								schematic[wi][tj] = "."
							} else {
								break
							}
						}

						for tj := wj + 1; tj < maxJ; tj++ {
							if isStringNumber(schematic[wi][tj]) {
								partNumber = partNumber + schematic[wi][tj]
								positions[Position{X: wi, Y: tj}] = schematic[wi][tj]
								schematic[wi][tj] = "."
							} else {
								break
							}
						}

						parts = append(parts, StringToNumber(partNumber))
					}
				}
			}
			if len(parts) == 2 {
				result += (parts[0] * parts[1])
			}
			for key, value := range positions {
				schematic[key.X][key.Y] = value
			}
		}
	}

	return result
}

func isValidIndex(i, j, maxI, maxJ int) bool {
	return i < maxI && j < maxJ
}
