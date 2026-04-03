package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"advent.of.code/util"
)

type Rng struct {
	start int
	end   int
}

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Q1: %v\n", Q1(inp))
	fmt.Printf("Q2: %v\n", Q2(inp))
}

func Q1(inp []Rng) int {
	result := 0

	for _, r := range inp {
		for i := r.start; i <= r.end; i++ {
			if isIdInvalid(i) {
				result += i
			}
		}
	}

	return result
}

func Q2(inp []Rng) int {
	result := 0

	for _, r := range inp {
		for i := r.start; i <= r.end; i++ {
			if isIdInvalid2(i) {
				// fmt.Println(i)
				result += i
			}
		}
	}

	return result
}

func isIdInvalid(num int) bool {
	numS := fmt.Sprintf("%d", num)
	if len(numS)%2 != 0 {
		return false
	}
	// fmt.Println(numS, numS[:len(numS)/2], numS[(len(numS)/2):])
	return numS[:len(numS)/2] == numS[(len(numS)/2):]
}

func isIdInvalid2(num int) bool {
	numS := fmt.Sprintf("%d", num)
	// fmt.Printf("----%s----\n", numS)
	ok := true
	for chunkSize := 1; chunkSize <= len(numS)/2; chunkSize++ {
		ok = true
		for j := chunkSize; j < len(numS); j += chunkSize {
			if len(numS)%chunkSize != 0 {
				ok = false
				continue
			}
			// fmt.Printf("%s == %s | %d <> %d\n", numS[j-chunkSize:j], numS[j:j+chunkSize], chunkSize, j)
			if numS[j-chunkSize:j] != numS[j:j+chunkSize] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}

	return false
}

func parseInput(filename string) ([]Rng, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := []Rng{}
	ranges := strings.Split(lines[0], ",")
	for _, r := range ranges {
		rs := strings.Split(r, "-")
		result = append(result, Rng{
			start: util.StringToNumber(rs[0]),
			end:   util.StringToNumber(rs[1]),
		})
	}

	return result, nil
}
