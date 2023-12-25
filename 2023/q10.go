package main

import (
	"fmt"
	"sort"
)

func q10sol() error {

	input, err := getFileAsListOfStrings("input10.small.txt")
	if err != nil {
		return err
	}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			fmt.Printf("(%d,%d)%s | ", i, j, string(input[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()

	start := findStart(input)

	result, visited := findLoopLength(input, start)

	rp2 := make(map[int][]Coordinate, 0)

	for k := range visited {
		if _, ok := rp2[k.x]; ok {
			rp2[k.x] = append(rp2[k.x], k)
		} else {
			rp2[k.x] = []Coordinate{k}
		}
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if _, ok := visited[Coordinate{i, j}]; ok {
				fmt.Printf("(%d,%d)[%s]  ", i, j, string(input[i][j]))
			} else {
				fmt.Printf("(%d,%d) %s   ", i, j, string(input[i][j]))
			}
		}
		fmt.Println()
	}
	fmt.Println()

	resP2 := 0
	for k, v := range rp2 {
		sort.Slice(v, func(i, j int) bool {
			return v[i].y < v[j].y
		})
		temp := 0
		for i := 1; i < len(v); i += 2 {
			temp += v[i].y - v[i-1].y - 1
		}
		resP2 += temp
		fmt.Println(k, temp, v, len(v))
	}

	fmt.Printf("Answer is P1 %d\n", result/2+1)

	fmt.Printf("Answer is P2 %d\n", resP2)

	return nil
}

func findStart(in []string) Coordinate {
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			if in[i][j] == 'S' {
				return Coordinate{i, j}
			}
		}
	}
	return Coordinate{-1, -1}
}

type stackvalq10 struct {
	depth   int
	cd      Coordinate
	visited map[Coordinate]string
}

func findLoopLength(input []string, begin Coordinate) (int, map[Coordinate]string) {
	st := NewStack()
	start, end := whereToMoveFromS(input, begin)
	st.Push(stackvalq10{0, start, map[Coordinate]string{begin: "S"}})
	for st.Len() > 0 {
		current := st.Pop().(stackvalq10)
		if current.cd == end {
			current.visited[current.cd] = string(input[current.cd.x][current.cd.y])
			return current.depth, current.visited

		}
		possibleMoves := possibleMoves(input, current.cd)
		for _, move := range possibleMoves {
			nmap := copyCoordMap(current.visited)
			if _, ok := current.visited[move]; !ok {
				nmap[current.cd] = string(input[current.cd.x][current.cd.y])
				st.Push(stackvalq10{current.depth + 1, move, nmap})
			}
		}
	}
	return -1, make(map[Coordinate]string)
}

func whereToMoveFromS(input []string, start Coordinate) (Coordinate, Coordinate) {
	paths := make([]Coordinate, 0)
	for d, v := range DIRECTIONS {
		i, j := start.x+v.x, start.y+v.y
		if !isValidIndexQ10(input, Coordinate{i, j}) {
			continue
		}
		switch d {
		case DirectionNorth:
			if input[i][j] == '|' || input[i][j] == '7' || input[i][j] == 'F' {
				paths = append(paths, Coordinate{i, j})
			}
		case DirectionSouth:
			if input[i][j] == '|' || input[i][j] == 'L' || input[i][j] == 'J' {
				paths = append(paths, Coordinate{i, j})
			}
		case DirectionEast:
			if input[i][j] == '-' || input[i][j] == '7' || input[i][j] == 'J' {
				paths = append(paths, Coordinate{i, j})
			}
		case DirectionWest:
			if input[i][j] == '-' || input[i][j] == 'L' || input[i][j] == 'F' {
				paths = append(paths, Coordinate{i, j})
			}
		}

	}

	fmt.Println(paths)

	return paths[0], paths[1]
}

func possibleMoves(input []string, curr Coordinate) []Coordinate {
	result := make([]Coordinate, 0)
	switch input[curr.x][curr.y] {
	case '|':
		// North
		ni, nj := _north.x+curr.x, _north.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '|' || input[ni][nj] == 'F' || input[ni][nj] == '7' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
		// South
		ni, nj = _south.x+curr.x, _south.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '|' || input[ni][nj] == 'L' || input[ni][nj] == 'J' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
	case '-':
		// East
		ni, nj := _east.x+curr.x, _east.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '-' || input[ni][nj] == 'J' || input[ni][nj] == '7' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
		// West
		ni, nj = _west.x+curr.x, _west.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '-' || input[ni][nj] == 'L' || input[ni][nj] == 'F' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
	case 'L':
		// North
		ni, nj := _north.x+curr.x, _north.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '|' || input[ni][nj] == 'F' || input[ni][nj] == '7' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
		// East
		ni, nj = _east.x+curr.x, _east.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '-' || input[ni][nj] == 'J' || input[ni][nj] == '7' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
	case 'J':
		// North
		ni, nj := _north.x+curr.x, _north.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '|' || input[ni][nj] == 'F' || input[ni][nj] == '7' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
		// West
		ni, nj = _west.x+curr.x, _west.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '-' || input[ni][nj] == 'L' || input[ni][nj] == 'F' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
	case 'F':
		// East
		ni, nj := _east.x+curr.x, _east.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '-' || input[ni][nj] == 'J' || input[ni][nj] == '7' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
		// South
		ni, nj = _south.x+curr.x, _south.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '|' || input[ni][nj] == 'L' || input[ni][nj] == 'J' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
	case '7':
		// South
		ni, nj := _south.x+curr.x, _south.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '|' || input[ni][nj] == 'L' || input[ni][nj] == 'J' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
		// West
		ni, nj = _west.x+curr.x, _west.y+curr.y
		if isValidIndexQ10(input, Coordinate{ni, nj}) {
			if input[ni][nj] == '-' || input[ni][nj] == 'L' || input[ni][nj] == 'F' || input[ni][nj] == 'S' {
				result = append(result, Coordinate{ni, nj})
			}
		}
	}
	return result
}

func isValidIndexQ10(input []string, c Coordinate) bool {
	return c.x >= 0 && c.x < len(input) && c.y >= 0 && c.y < len(input[c.x])
}

func copyCoordMap(in map[Coordinate]string) map[Coordinate]string {
	res := make(map[Coordinate]string)

	for k, v := range in {
		res[k] = v
	}

	return res
}
