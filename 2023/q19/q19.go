package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"advent.of.code/util"
)

func main() {

	wkflows, parts, err := getInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("P1 Answer is %d\n", solveP1(parts, wkflows))
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
