package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"advent.of.code/util"
)

func main() {
	filename := "input.small.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	inp, err := parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}
	handleDefrag1(inp)

	inp, err = parseInput(filename)
	if err != nil {
		log.Fatal(err)
	}

	handleDefrag2(inp)
}

func handleDefrag1(inp []int) {
	defragged := defrag(inp)

	slices.SortFunc(defragged, func(a, b block) int {
		z := a.Index - b.Index
		if z != 0 {
			return z
		}
		return b.ID - a.ID
	})

	res := ""
	for _, v := range defragged {
		for i := 0; i < v.Size; i++ {
			res += fmt.Sprintf("%d", v.ID)
		}
	}

	fmt.Println(checksum(defragged))
}

func handleDefrag2(inp []int) {
	defragged := defrag2(inp)

	slices.SortFunc(defragged, func(a, b block) int {
		z := a.Index - b.Index
		if z != 0 {
			return z
		}
		return b.ID - a.ID
	})

	idx := 0
	for _, d := range defragged {
		idx += d.Index
		// fmt.Println(d, idx)
	}

	res := ""
	for _, v := range defragged {
		for i := 0; i < v.Size; i++ {
			if v.ID < 0 {
				res += "."
			} else {
				res += fmt.Sprintf("%d", v.ID)
			}
		}
	}

	// fmt.Println(res)
	fmt.Println(checksum(defragged))
}

type block struct {
	Index int
	ID    int
	Size  int
}

func (b block) String() string {
	return fmt.Sprintf("IDX: %3d -> ID:%3d SIZE:%3d", b.Index, b.ID, b.Size)
}

func defrag(disk []int) []block {
	blocks := make([]block, 0)

	first := 1
	last := len(disk) - 1
	if len(disk)%2 == 0 {
		last = len(disk) - 2
	}

	for first < last {
		emptySpace := disk[first]
		spaceToMove := disk[last]
		// fmt.Printf("first: %3d, last:%3d, emptySpace:%d, spaceToMove:%d\n", first, last, emptySpace, spaceToMove)
		if emptySpace > spaceToMove {
			b := block{
				Index: first,
				ID:    last / 2,
				Size:  spaceToMove,
			}
			// fmt.Println(b)
			blocks = append(blocks, b)
			disk[first] -= spaceToMove
			disk[last] = 0
			last -= 2
		} else {
			b := block{
				Index: first,
				ID:    last / 2,
				Size:  emptySpace,
			}
			// fmt.Println(b)
			blocks = append(blocks, b)
			disk[first] = 0
			disk[last] -= emptySpace
			first += 2
		}
	}

	for i := 0; i < len(disk); i += 2 {
		if disk[i] == 0 {
			continue
		}
		blocks = append(blocks, block{
			Index: i,
			ID:    i / 2,
			Size:  disk[i],
		})
	}

	return blocks
}

func defrag2(disk []int) []block {
	orig := slices.Clone(disk)
	blocks := make([]block, 0)

	lasto := len(disk) - 1
	if len(disk)%2 == 0 {
		lasto = len(disk) - 2
	}

	for first := 1; first < len(disk); first += 2 {
		if disk[first] == 0 {
			continue
		}
		last := lasto
		for first < last {
			if disk[last] == 0 {
				last -= 2
				continue
			}
			emptySpace := disk[first]
			spaceToMove := disk[last]
			// fmt.Printf("first: %3d, last:%3d, emptySpace:%d, spaceToMove:%d\n", first, last, emptySpace, spaceToMove)
			if emptySpace >= spaceToMove {
				b := block{
					Index: first,
					ID:    last / 2,
					Size:  spaceToMove,
				}
				// fmt.Println(b)
				blocks = append(blocks, b)
				disk[first] -= spaceToMove
				disk[last] = 0
			}
			last -= 2
		}
	}

	for i := 0; i < len(disk); i++ {
		if i%2 == 0 {
			if disk[i] == 0 && orig[i] != 0 {
				blocks = append(blocks, block{
					Index: i,
					ID:    -1,
					Size:  orig[i],
				})
			} else {
				blocks = append(blocks, block{
					Index: i,
					ID:    i / 2,
					Size:  disk[i],
				})
			}

		} else if i%2 == 1 {
			blocks = append(blocks, block{
				Index: i,
				ID:    -1,
				Size:  disk[i],
			})
		}

	}

	return blocks
}

func checksum(blocks []block) int {
	res := 0
	idx := 0
	for _, v := range blocks {
		for i := 0; i < v.Size; i++ {
			id := v.ID
			if id < 0 {
				id = 0
			}
			res += idx * id
			idx++
		}
	}
	return res
}

func parseInput(filename string) ([]int, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	return util.StringOfNumbersToIntSlice(lines[0])
}
