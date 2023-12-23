package main

import "fmt"

func q9sol() error {

	input, err := getInput("input9.txt")
	if err != nil {
		return err
	}

	result := 0
	for _, arr := range input {
		result += getNextValue(arr, false)
	}

	fmt.Println("Answer is :", result)
	return nil

}

func getNextValue(inp []int, doEnd bool) int {
	storage := make([][]int, 1)
	storage[0] = inp

	i := 0

	for !areAllZeros(storage[i]) {
		below := make([]int, len(storage[i])-1)
		for j := 1; j < len(storage[i]); j++ {
			below[j-1] = storage[i][j] - storage[i][j-1]
		}
		storage = append(storage, below)
		i++
	}

	res := 0

	if doEnd {
		for i := len(storage) - 1; i >= 0; i-- {
			res += storage[i][len(storage[i])-1]
		}
	} else {
		for i := len(storage) - 2; i >= 0; i-- {
			res = storage[i][0] - res
		}
	}
	return res
}

func areAllZeros(in []int) bool {

	for _, v := range in {
		if v != 0 {
			return false
		}
	}

	return true
}

func getInput(fileName string) ([][]int, error) {

	lines, err := getFileAsListOfStrings(fileName)
	if err != nil {
		return nil, err
	}

	result := make([][]int, len(lines))

	for idx, line := range lines {
		ints, err := SpaceSeparatedStringOfNumbersToIntSlice(line)
		if err != nil {
			return nil, err
		}
		result[idx] = ints
	}

	return result, nil
}
