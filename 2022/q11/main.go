package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"slices"
	"strings"

	"advent.of.code/util"
)

func main() {
	fileName := "input.small.txt"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	monkeys, err := parseInput(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("> Answer P1 : %d \n", solveP1(monkeys, 20, 3))

	fmt.Printf("> Answer P2 : %d \n", solveP2(monkeys, 1000))
}

func solveP2(monkeys []*Monkey, nRounds int) uint64 {
	inspections := make([]uint64, len(monkeys))

	for round := 0; round < nRounds; round++ {
		for mi, m := range monkeys {
			inspections[mi] += uint64(len(m.BigItems))

			for _, item := range m.BigItems {
				worryLevel := *big.NewInt(item.Int64())
				operateWith := *big.NewInt(item.Int64())
				if m.OperateWith != nil {
					operateWith.SetInt64(int64(*m.OperateWith))
				}

				if m.Operation == "+" {
					worryLevel.Add(&worryLevel, &operateWith)
				} else {
					worryLevel.Mul(&worryLevel, &operateWith)
				}

				wMod := big.NewInt(worryLevel.Int64())
				wMod.Mod(wMod, big.NewInt(int64(m.TestDivisibleBy)))
				if wMod.Int64() == 0 {
					monkeys[m.TestPass].BigItems = append(monkeys[m.TestPass].BigItems, worryLevel)
				} else {
					monkeys[m.TestFail].BigItems = append(monkeys[m.TestFail].BigItems, worryLevel)
				}
			}
			monkeys[mi].BigItems = make([]big.Int, 0)

		}
	}

	fmt.Println(inspections)
	slices.Sort(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func solveP1(monkeys []*Monkey, nRounds int, worryLevelRelaxer uint) int {
	inspections := make([]int, len(monkeys))

	for round := 0; round < nRounds; round++ {
		for mi, m := range monkeys {
			inspections[mi] += len(m.Items)

			for _, item := range m.Items {
				worryLevel := operate(item, m.Operation, m.OperateWith)
				relievedWorryLevel := worryLevel / worryLevelRelaxer
				if relievedWorryLevel%m.TestDivisibleBy == uint(0) {
					monkeys[m.TestPass].Items = append(monkeys[m.TestPass].Items, relievedWorryLevel)
				} else {
					monkeys[m.TestFail].Items = append(monkeys[m.TestFail].Items, relievedWorryLevel)
				}
			}
			monkeys[mi].Items = make([]uint, 0)

		}
	}

	slices.Sort(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func operate(existing uint, operation string, operand *int) uint {
	opr := existing
	if operand != nil {
		opr = uint(*operand)
	}

	if operation == "+" {
		return existing + opr
	}

	return existing * opr
}

type Monkey struct {
	Items           []uint
	BigItems        []big.Int
	Operation       string
	OperateWith     *int
	TestDivisibleBy uint
	TestPass        int
	TestFail        int
}

func (m Monkey) String() string {
	return fmt.Sprintf("Starting Items: %v\nOperation: new = old %s %v\nDivisible By %d\nPass: %d \t Fail %d\n", m.Items, m.Operation, m.OperateWith, m.TestDivisibleBy, m.TestPass, m.TestFail)
}

func parseInput(filename string) ([]*Monkey, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}
	result := make([]*Monkey, 0)
	re := regexp.MustCompile(`\d+`)
	for i := 0; i < len(lines); i += 7 {
		itemsStr := re.FindAllString(lines[i+1], -1)
		items := make([]uint, len(itemsStr))
		bigItems := make([]big.Int, len(itemsStr))
		for i, v := range itemsStr {
			items[i] = uint(util.StringToNumber(v))
			bigItems[i] = *big.NewInt(int64(util.StringToNumber(v)))
		}

		opsSplit := strings.Split(strings.Split(lines[i+2], "=")[1], " ")
		var opsInt *int
		if util.IsStringNumber(opsSplit[3]) {
			i := util.StringToNumber(opsSplit[3])
			opsInt = &i
		}

		divisibleBy := util.StringToNumber(re.FindAllString(lines[i+3], -1)[0])
		testPass := util.StringToNumber(re.FindAllString(lines[i+4], -1)[0])
		testFail := util.StringToNumber(re.FindAllString(lines[i+5], -1)[0])

		result = append(result, &Monkey{
			items, bigItems, opsSplit[2], opsInt, uint(divisibleBy), testPass, testFail,
		})

	}

	return result, nil
}
