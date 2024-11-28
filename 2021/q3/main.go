package main

import (
	"fmt"
	"log"

	"advent.of.code/util"
)

func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	gamma, epsilon := getGammaAndEpsilonRate(input)
	fmt.Println("Answer Q1: ", gamma*epsilon)

	fmt.Println("Answer Q2: ", findRating(input, true)*findRating(input, false))
}

func getGammaAndEpsilonRate(input []string) (gamma, epsilon int) {
	var gammaBin, epsilonBin string

	for i := 0; i < len(input[0]); i++ {
		zeroes, ones := 0, 0
		for j := 0; j < len(input); j++ {
			if input[j][i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		if ones > zeroes {
			gammaBin += "0"
			epsilonBin += "1"
		} else {
			gammaBin += "1"
			epsilonBin += "0"
		}
	}

	return util.ConvertBinaryStringToNumber(gammaBin), util.ConvertBinaryStringToNumber(epsilonBin)
}

func findRating(input []string, most bool) (oxygen int) {
	for i := 0; i < len(input[0]); i++ {
		freq := findFrequent(input, i, most)
		input = filter(input, freq, i)
		if len(input) == 1 {
			return util.ConvertBinaryStringToNumber(input[0])
		}
	}
	return
}

func findFrequent(input []string, index int, most bool) string {
	zeroes, ones := 0, 0
	for j := 0; j < len(input); j++ {
		if input[j][index] == '0' {
			zeroes++
		} else {
			ones++
		}
	}
	if most {
		if zeroes > ones {
			return "0"
		}
		return "1"
	}
	if zeroes > ones {
		return "1"
	}
	return "0"

}

func filter(input []string, bitVal string, index int) []string {
	res := make([]string, 0)
	for _, s := range input {
		if string(s[index]) == bitVal {
			res = append(res, s)
		}
	}
	return res
}

func parseInput(filename string) ([]string, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}
	return lines, nil
}
