package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"advent.of.code/util"
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

	start := time.Now()
	fmt.Println(Q1(inp, 25))
	fmt.Println("EXT", time.Since(start))
	start = time.Now()
	fmt.Println(Q2(inp, 25))
	fmt.Println("EXT", time.Since(start))
	start = time.Now()
	fmt.Println(Q2(inp, 75))
	fmt.Println("EXT", time.Since(start))
}

func Q1(stones []string, blinks int) int {
	for range blinks {
		newStones := []string{}
		for i := 0; i < len(stones); i++ {
			l := len(stones[i])
			if stones[i] == "0" {
				newStones = append(newStones, "1")
			} else if l%2 == 0 {
				n1 := strconv.Itoa(util.StringToNumber(string(stones[i][:l/2])))
				n2 := strconv.Itoa(util.StringToNumber(string(stones[i][l/2:])))
				newStones = append(newStones, n1, n2)
			} else {
				num := util.StringToNumber(stones[i])
				num *= 2024
				newStones = append(newStones, strconv.Itoa(num))
			}
		}
		stones = newStones
		// fmt.Println(stones)
	}

	return len(stones)
}

func Q2(stones []string, blinks int) int {
	store := map[string]int{}
	for _, stone := range stones {
		if v, found := store[stone]; found {
			store[stone] = v + 1
		} else {
			store[stone] = 1
		}
	}

	for range blinks {
		newStore := map[string]int{}
		for k, v := range store {
			l := len(k)
			if k == "0" {
				newStore["1"] += v
			} else if l%2 == 0 {
				n1 := strconv.Itoa(util.StringToNumber(string(k[:l/2])))
				n2 := strconv.Itoa(util.StringToNumber(string(k[l/2:])))
				newStore[n1] += v
				newStore[n2] += v
			} else {
				newStore[strconv.Itoa(util.StringToNumber(k)*2024)] += v
			}
		}
		store = newStore
	}

	result := 0
	for _, v := range store {
		result += v
	}

	return result
}

func parseInput(filename string) ([]string, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile("[0-9]+")

	return re.FindAllString(lines[0], -1), nil
}
