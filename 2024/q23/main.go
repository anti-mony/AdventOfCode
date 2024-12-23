package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strings"

	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, _, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(Q1(inp))

	fmt.Println(Q2(inp))
}

type Connection struct {
	C1 string
	C2 string
	C3 string
}

func (c Connection) hash() string {
	x := []string{c.C1, c.C2, c.C3}
	slices.Sort(x)
	return fmt.Sprintf("%v", x)
}

func Q1(adj map[string][]string) int {
	// Fa la la brute force
	result := map[string]bool{}

	for c1 := range adj {
		for c2 := range adj {
			for c3 := range adj {
				if c1 == c2 || c2 == c3 || c3 == c1 {
					continue
				}
				c1a := slices.Index(adj[c2], c1)
				c1b := slices.Index(adj[c3], c1)
				c2a := slices.Index(adj[c1], c2)
				c2b := slices.Index(adj[c3], c2)
				c3a := slices.Index(adj[c1], c3)
				c3b := slices.Index(adj[c2], c3)

				if c1a >= 0 && c1b >= 0 && c2a >= 0 && c2b >= 0 && c3a >= 0 && c3b >= 0 {
					if strings.HasPrefix(c1, "t") || strings.HasPrefix(c2, "t") || strings.HasPrefix(c3, "t") {
						result[Connection{C1: c1, C2: c2, C3: c3}.hash()] = true
					}
				}
			}
		}
	}
	return len(result)
}

func Q2(adj map[string][]string) string {
	comps := slices.Collect(maps.Keys(adj))
	slices.Sort(comps)

	seen := map[string]bool{}

	for _, c := range comps {
		findLargestCluster(adj, c, []string{c}, seen)
	}

	return strings.Join(largestCluster, ",")
}

var largestCluster []string

func findLargestCluster(adj map[string][]string, current string, cluster []string, memo map[string]bool) {
	slices.Sort(cluster)
	if _, found := memo[strings.Join(cluster, ",")]; found {
		return
	}
	if len(cluster) >= len(largestCluster) {
		largestCluster = cluster
	}
	fmt.Println(cluster)

	for _, comp := range adj[current] {
		foundAll := true
		for _, c := range cluster {
			if slices.Index(adj[c], comp) < 0 {
				foundAll = false
				break
			}
		}
		if foundAll {
			findLargestCluster(adj, comp, append(slices.Clone(cluster), comp), memo)
		}
	}

	memo[strings.Join(cluster, ",")] = true

}

func parseInput(filename string) (map[string][]string, []Connection, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, nil, err
	}

	result := map[string][]string{}

	for _, line := range lines {
		c := strings.Split(line, "-")
		result[c[0]] = append(result[c[0]], c[1])
		result[c[1]] = append(result[c[1]], c[0])
	}

	for k := range result {
		slices.Sort(result[k])
	}

	return result, nil, nil
}
