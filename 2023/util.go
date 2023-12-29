package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const MAX_LINE_LENGTH = 65000

/*
i, j --> the directions refer to that
0,0 0,1 0,2
1,0 1,1 1,2
2,0 2,1 2,2
*/

type Direction int

const (
	DirectionNorth Direction = iota + 1
	DirectionEast
	DirectionWest
	DirectionSouth
)

var (
	_north = Coordinate{-1, 0}
	_east  = Coordinate{0, 1}
	_south = Coordinate{1, 0}
	_west  = Coordinate{0, -1}
)

var DIRECTIONS = map[Direction]Coordinate{
	DirectionNorth: _north,
	DirectionEast:  _east,
	DirectionSouth: _south,
	DirectionWest:  _west,
}

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) Add(i Coordinate) Coordinate {
	return Coordinate{c.x + i.x, c.y + i.y}
}

func getFileAsListOfStrings(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	buffer := make([]byte, MAX_LINE_LENGTH)

	scanner := bufio.NewScanner(file)
	scanner.Buffer(buffer, MAX_LINE_LENGTH)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil

}

func isCharNumber(c rune) bool {

	if c >= 48 && c <= 57 {
		return true
	}

	return false
}

func isStringNumber(c string) bool {

	r := []rune(c)

	return isCharNumber(r[0])
}

func StringToNumber(n string) int {
	n = strings.TrimSpace(n)
	res, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}
	return res
}

func SpaceSeparatedStringOfNumbersToIntSlice(in string) ([]int, error) {
	result := make([]int, 0)

	re := regexp.MustCompile("-?[0-9]+")

	for _, number := range re.FindAllString(in, -1) {
		n, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}

	return result, nil
}

func Dedupe[T string | int](inp []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)
	for _, v := range inp {
		if _, ok := seen[v]; !ok {
			result = append(result, v)
			seen[v] = true
		}

	}
	return result
}

func GCD(a, b int) int {
	result := 0
	if a > b {
		result = b
	} else {
		result = a
	}

	for result > 0 {
		if a%result == 0 && b%result == 0 {
			return result
		}
		result--
	}

	return result
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

func print2DArray[T string | int](in [][]T) {
	fmt.Println()
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			fmt.Printf("%2v", in[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func copy2DArray[T string | int](in [][]T) [][]T {
	result := make([][]T, len(in))
	for i := 0; i < len(in); i++ {
		row := make([]T, len(in[i]))
		copy(row, in[i])
		result[i] = row
	}
	return result
}

func areTheseTwoEqual[T string | int](a [][]T, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
