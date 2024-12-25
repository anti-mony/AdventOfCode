package main

import (
	"fmt"
	"log"
	"os"

	"advent.of.code/util"
)

const (
	ROWS    = 7
	COLS    = 5
	OVERLAP = 5
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

	fmt.Println("A1: ", Q1(getLockAndKeys(inp)))
}

func Q1(locks, keys [][]int) int {
	result := 0
	for _, lock := range locks {
		for _, key := range keys {
			fits := true
			for i := 0; i < COLS; i++ {
				if lock[i]+key[i] > OVERLAP {
					fits = false
					break
				}
			}
			if fits {
				result++
			}
		}
	}
	return result
}

func getLockAndKeys(lorks [][][]string) ([][]int, [][]int) {
	locks := make([][]int, 0)
	keys := make([][]int, 0)

	for _, lork := range lorks {
		if lork[0][0] == "#" {
			// Lock
			lock := make([]int, COLS)
			for j := 0; j < COLS; j++ {
				c := 0
				for i := 1; i < ROWS; i++ {
					if lork[i][j] != "#" {
						break
					}
					c++
				}
				lock[j] = c
			}
			locks = append(locks, lock)
		} else {
			key := make([]int, COLS)
			// Key
			for j := 0; j < COLS; j++ {
				c := 0
				for i := ROWS - 2; i >= 0; i-- {
					if lork[i][j] != "#" {
						break
					}
					c++
				}
				key[j] = c
			}
			keys = append(keys, key)
		}
	}

	return locks, keys
}

func parseInput(filename string) ([][][]string, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := make([][][]string, 0)

	for i := range lines {
		lockOrKey := make([][]string, ROWS)
		if len(lines[i]) == 0 {
			for j := ROWS; j > 0; j-- {
				lockOrKey[ROWS-j] = util.StringToCharSlice(lines[i-j])
			}
			result = append(result, lockOrKey)
		}
	}

	return result, nil
}
