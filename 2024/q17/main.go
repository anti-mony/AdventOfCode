package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"

	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	c, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	c.Execute()
	fmt.Println(c)

	c, err = parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	FindMinA(c)

}

type Instruction struct {
	OppCode     int
	OperandCode int
}

type Computer struct {
	A, B, C int

	Instructions []Instruction

	Output []int
}

func (c *Computer) String() string {
	return fmt.Sprintf("A: %4d | B: %4d | C: %4d\nInstructions(%d): %v\nOutput: %v", c.A, c.B, c.C, len(c.Instructions), c.Instructions, c.Output)
}

func (c *Computer) FlattenInstructions() []int {
	ins := []int{}
	for _, i := range c.Instructions {
		ins = append(ins, i.OppCode, i.OperandCode)
	}
	return ins
}

func (c *Computer) GetComboOperand(index int) int {
	operandCode := c.Instructions[index].OperandCode
	if operandCode >= 0 && operandCode <= 3 {
		return operandCode
	}
	if operandCode == 4 {
		return c.A
	}
	if operandCode == 5 {
		return c.B
	}
	if operandCode == 6 {
		return c.C
	}
	panic("invalid instruction found")
}

func (c *Computer) GetLiteralOperand(index int) int {
	return c.Instructions[index].OperandCode
}

func (c *Computer) GetOperation(index int) int {
	return c.Instructions[index].OppCode
}

func (c *Computer) Execute() {
	instruction := 0

	for instruction != len(c.Instructions) {
		operation := c.GetOperation(instruction)
		switch operation {
		case 0:
			c.A = c.A / int(math.Pow(2, float64(c.GetComboOperand(instruction))))
		case 1:
			c.B = c.B ^ c.GetLiteralOperand(instruction)
		case 2:
			c.B = c.GetComboOperand(instruction) % 8
		case 3:
			if c.A != 0 {
				instruction = c.GetLiteralOperand(instruction)
				continue
			}
		case 4:
			c.B = c.B ^ c.C
		case 5:
			c.Output = append(c.Output, c.GetComboOperand(instruction)%8)
		case 6:
			c.B = c.A / int(math.Pow(2, float64(c.GetComboOperand(instruction))))
		case 7:
			c.C = c.A / int(math.Pow(2, float64(c.GetComboOperand(instruction))))
		}
		instruction++
	}
}

func (c *Computer) Clone() *Computer {
	return &Computer{
		A:            c.A,
		B:            c.B,
		C:            c.C,
		Output:       []int{},
		Instructions: c.Instructions,
	}
}

func FindMinA(c *Computer) int {
	flattenedIns := c.FlattenInstructions()
	a := 0
	for {
		if a == 1000000 {
			return a
		}
		newC := c.Clone()
		newC.A = a
		newC.Execute()
		fmt.Println(newC.Output)
		if slices.Compare(flattenedIns, newC.Output) == 0 {
			fmt.Println("Min A: ", a)
			return a
		}
		a++
	}
}

func parseInput(filename string) (*Computer, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`-?[0-9]+`)

	instructions := []Instruction{}
	ins := re.FindAllString(lines[4], -1)

	for i := 0; i < len(ins)-1; i += 2 {
		instructions = append(instructions, Instruction{
			OppCode:     util.StringToNumber(ins[i]),
			OperandCode: util.StringToNumber(ins[i+1]),
		})
	}

	return &Computer{
		A:            util.StringToNumber(re.FindAllString(lines[0], -1)[0]),
		B:            util.StringToNumber(re.FindAllString(lines[1], -1)[0]),
		C:            util.StringToNumber(re.FindAllString(lines[2], -1)[0]),
		Instructions: instructions,
	}, nil
}
