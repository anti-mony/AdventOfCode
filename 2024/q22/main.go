package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"advent.of.code/list"
	"advent.of.code/util"
)

const DAYS = 2001

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(Q1(inp))

	fmt.Println(Q2(inp))
}

func Q1(starts []int) ([]int, int) {
	result := slices.Clone(starts)

	for i := range result {
		for range DAYS {
			result[i] = generateNextSecrectNumber(result[i])
		}
	}

	return result, list.Sum(result)
}

func Q2(starts []int) int {
	BUYERS := len(starts)
	dts_lcl := make([]Value, 0)
	for _, s := range starts {
		dts_lcl = append(dts_lcl, Value{
			actual: s,
			value:  s % 10,
		})
	}

	memo := make([]map[deltas]int, BUYERS)
	for b := range BUYERS {
		memo[b] = map[deltas]int{}
	}

	dts := [][]Value{dts_lcl}
	uniqueDeltas := map[deltas]bool{}

	for d := 1; d < DAYS; d++ {
		dts_lcl := make([]Value, BUYERS)
		for i := range BUYERS {
			secret := generateNextSecrectNumber(dts[d-1][i].actual)
			n := secret % 10
			dts_lcl[i] = Value{
				actual: secret,
				value:  n,
				diff:   n - dts[d-1][i].value,
			}

			if d > 3 {
				dtls := deltas{
					D: dts_lcl[i].diff,
					C: dts[d-1][i].diff,
					B: dts[d-2][i].diff,
					A: dts[d-3][i].diff,
				}
				uniqueDeltas[dtls] = true
				if _, found := memo[i][dtls]; !found {
					memo[i][dtls] = dts_lcl[i].value
				}
			}
		}
		dts = append(dts, dts_lcl)
	}

	bananas := 0
	for unq := range uniqueDeltas {
		tbanana := 0
		for _, buyer := range memo {
			tbanana += buyer[unq]
		}
		if tbanana > bananas {
			bananas = tbanana
			fmt.Println(unq, tbanana)
		}
	}

	return bananas
}

type Value struct {
	actual int
	value  int
	diff   int
}

type deltas struct {
	A int
	B int
	C int
	D int
}

func parseInput(filename string) ([]int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	result := []int{}
	for _, line := range lines {
		result = append(result, util.StringToNumber(line))
	}

	return result, nil
}

func generateNextSecrectNumber(start int) int {
	result := start * 64
	next := mix(start, result)
	next = prune(next)
	result = next / 32
	next = mix(result, next)
	next = prune(next)
	result = next * 2048
	next = mix(next, result)
	next = prune(next)
	return next
}

func mix(secretNumber int, opsRes int) int {
	return secretNumber ^ opsRes
}

func prune(secretNumber int) int {
	return secretNumber % 16777216
}
