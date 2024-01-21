package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"advent.of.code/list"
	"advent.of.code/util"
)

const RANGE_MAX = uint64(4000)

func main() {

	wkflows, parts, err := getInput("input.small.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("P1 Answer is %d\n", solveP1(parts, wkflows))

	answerP2 := solveP2(wkflows)
	fmt.Printf("P2 Answer is %d\n", answerP2)

	fmt.Printf("Answer matches : %v \n", answerP2 == 167409079868000)
}

func solveP1(parts []Part, wkflows map[string][]Condition) int {
	result := 0
	for _, p := range parts {
		if AcceptPart(p, wkflows) {
			result += p.SumXMAS()
		}
	}

	return result
}

func solveP2(wkflows map[string][]Condition) uint64 {
	stack := list.NewStack()
	stack.Push(stackVar{
		workflow: "in",
		part:     Part{Categories: map[string]int{}},
	})
	result := uint64(0)
	for stack.Len() > 0 {
		current := stack.Pop().(stackVar)
		fmt.Printf("Current Workflow: %s \n", current.workflow)
		for _, c := range wkflows[current.workflow] {
			if c.Operator == ">" {
				if c.Result == "A" {
					result += current.part.MultiplyXMAS()
					continue
				}
				stack.Push(stackVar{workflow: c.Result})
			} else if c.Operator == "<" {
				if c.Result == "A" {
					result += current.part.MultiplyXMAS()
					continue
				}
				stack.Push(stackVar{workflow: c.Result})
			} else if c.Result == "A" {
				result += current.part.MultiplyXMAS()
			} else if c.Result != "R" {
				stack.Push(stackVar{workflow: c.Result})
			}
		}
	}

	return result
}

func AcceptPart(part Part, wkflows map[string][]Condition) bool {
	workflow := wkflows["in"]

	for {
		for idx, condition := range workflow {
			conditionIsTrue := false
			if idx == len(workflow)-1 {
				conditionIsTrue = true
			} else if condition.Operator == ">" {
				if part.Categories[condition.PartCategory] > condition.Comparand {
					conditionIsTrue = true
				}
			} else {
				if part.Categories[condition.PartCategory] < condition.Comparand {
					conditionIsTrue = true
				}
			}
			if conditionIsTrue {
				if condition.Result == "A" {
					return true
				} else if condition.Result == "R" {
					return false
				} else {
					workflow = wkflows[condition.Result]
					break
				}
			}
		}
	}
}

type stackVar struct {
	workflow string
	part     Part
}

type Condition struct {
	Operator     string
	Result       string
	Comparand    int
	PartCategory string
}

type Part struct {
	Categories map[string]int
}

func (m Part) SumXMAS() int {
	res := 0
	for _, v := range m.Categories {
		res += v
	}
	return res
}

func (m Part) MultiplyXMAS() uint64 {
	res := uint64(1)
	for _, c := range []string{"X", "M", "A", "S"} {
		if v, ok := m.Categories[c]; ok {
			res *= uint64(v)
		} else {
			res *= RANGE_MAX
		}
	}
	return res
}

func (m Part) Update(c string, value int) {
	if v, ok := m.Categories[c]; ok {
	} else {
		m.Categories[c] = value
	}

}

func getInput(filename string) (map[string][]Condition, []Part, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, nil, err
	}

	workflows := make(map[string][]Condition)

	i := 0
	for i = 0; i < len(lines); i++ {
		if len(lines[i]) < 1 {
			break
		}
		name, workflow := getWorkflowFromString(lines[i])
		workflows[name] = workflow
	}

	parts := make([]Part, 0)
	for i = i + 1; i < len(lines); i++ {
		parts = append(parts, getPartFromString(lines[i]))
	}

	return workflows, parts, nil
}

func getWorkflowFromString(wk string) (string, []Condition) {
	re := regexp.MustCompile("(.*?){(.*?)}")
	matches := re.FindStringSubmatch(wk)
	wkName := matches[1]

	wkflows := make([]Condition, 0)
	workflowRe := regexp.MustCompile(`(x|m|a|s)(<|>)(\d+):([a-zA-Z]+)`)

	conditions := strings.Split(matches[2], ",")
	for _, wk := range conditions[:len(conditions)-1] {
		wkflowStr := workflowRe.FindStringSubmatch(wk)
		wkflows = append(wkflows, Condition{
			PartCategory: wkflowStr[1],
			Operator:     wkflowStr[2],
			Comparand:    util.StringToNumber(wkflowStr[3]),
			Result:       wkflowStr[4],
		})
	}
	wkflows = append(wkflows, Condition{
		Result: conditions[len(conditions)-1],
	})

	return wkName, wkflows
}

func getPartFromString(pt string) Part {
	re := regexp.MustCompile(`(?P<K>x|m|a|s)=(?P<V>\d+)`)
	p := Part{Categories: make(map[string]int)}
	categories := strings.Split(pt, ",")

	for _, c := range categories {
		match := re.FindStringSubmatch(c)
		p.Categories[match[1]] = util.StringToNumber(match[2])
	}

	return p
}
