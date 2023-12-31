package main

import (
	"fmt"
	"strings"
)

type q15lens struct {
	label       string
	focalLength int
}

func compareForQ15(a, b any) bool {
	aV := a.(q15lens)
	bV := b.(q15lens)

	return aV.label == bV.label
}

func q15sol() error {

	input, err := getInputQ15("input15.txt")
	if err != nil {
		return err
	}

	fmt.Println("P1 Answer is: ", solveQ15P1(input))

	fmt.Println("P2 Answer is: ", solveQ15P2(input))

	return nil
}

func solveQ15P2(in []string) int {
	result := 0

	boxes := map[int]*LinkedList{}
	for i := 0; i < 256; i++ {
		boxes[i] = NewLinkedList()
	}

	for _, piece := range in {
		label, boxN, focalL := GetLabelAndBoxNumberAndFocalLength(piece)
		processPieceQ15(label, focalL, boxes[boxN])
	}

	for i := 0; i < 256; i++ {
		r := calculateFocusingPower(i, boxes[i])
		result += r
	}

	return result
}

func processPieceQ15(label string, focalLength int, box *LinkedList) {
	if focalLength < 0 {
		box.Delete(q15lens{label, 0}, compareForQ15)
		return
	}

	box.UpsertAppend(q15lens{label, focalLength}, compareForQ15)
}

func calculateFocusingPower(boxN int, ll *LinkedList) int {
	result := 0

	length := 0
	tmp := ll.First
	for tmp != nil {
		currentValue := tmp.Value.(q15lens)
		length++
		result += (boxN + 1) * length * currentValue.focalLength
		tmp = tmp.next
	}

	return result
}

func GetLabelAndBoxNumberAndFocalLength(in string) (string, int, int) {
	r := 0
	i := 0
	c := rune(0)

	for i, c = range in {
		if c == '-' || c == '=' {
			break
		}
		r = hashChar(c, r)
	}

	focalLength := StringToNumber(in[i+1:])
	if c == '-' {
		focalLength = -1
	}

	return in[:i], r, focalLength

}

func solveQ15P1(in []string) int {
	result := 0
	for _, s := range in {
		r := 0
		for _, c := range s {
			r = hashChar(c, r)
		}
		result += r
	}
	return result
}

func getInputQ15(filePath string) ([]string, error) {
	lines, err := getFileAsListOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	inp := strings.Split(lines[0], ",")

	return inp, nil
}

func hashChar(inp rune, current int) int {
	current += int(inp)
	current *= 17
	current = current % 256
	return current
}
