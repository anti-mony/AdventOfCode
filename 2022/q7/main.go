package main

import (
	"fmt"
	"log"
	"strings"

	"advent.of.code/list"
	"advent.of.code/util"
)

const (
	TOTAL_FILESYSTEM_SIZE = 70000000
	MIN_SIZE_FOR_UPDATE   = 30000000
)

func main() {
	fileStructure, err := ParseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	calculateSizes(fileStructure)
	// fileStructure.Print()

	fmt.Println("Answer P1: ", solveP1(fileStructure, uint64(100000)))
	fmt.Println("Answer P2: ", solveP2(fileStructure))
}

func solveP2(d *Directory) uint64 {
	spaceRequired := MIN_SIZE_FOR_UPDATE - (TOTAL_FILESYSTEM_SIZE - d.Size)

	var recursive func(root *Directory, minSize uint64) uint64
	recursive = func(root *Directory, minSize uint64) uint64 {
		if root == nil {
			return uint64(0)
		}

		for _, dir := range root.SubDirectories {
			res := recursive(dir, minSize)
			if res > 0 {
				return res
			}
		}

		if root.Size > minSize {
			return root.Size
		}

		return 0
	}

	return recursive(d, spaceRequired)
}

func solveP1(d *Directory, maxSize uint64) uint64 {
	res := uint64(0)

	stack := list.NewStack[*Directory]()
	stack.Push(d)

	for stack.Len() > 0 {
		current := stack.Pop()
		// fmt.Printf("Processing %s of size %d \n", current.Name, current.Size)
		for _, dir := range current.SubDirectories {
			stack.Push(dir)
		}

		if current.Size <= maxSize {
			res += current.Size
		}

	}

	return res
}

func calculateSizes(root *Directory) uint64 {
	res := uint64(0)
	if root == nil {
		return res
	}

	for _, d := range root.SubDirectories {
		res += calculateSizes(d)
	}

	for _, file := range root.Files {
		res += file.Size
	}

	root.Size = res

	return res
}

type StackVar struct {
	directory *Directory
	depth     int
}

type Directory struct {
	Name            string
	ParentDirectory *Directory
	SubDirectories  []*Directory
	Files           []*File
	Size            uint64
}

func (d *Directory) Print() {
	stack := list.NewStack[StackVar]()
	stack.Push(StackVar{d, 0})

	for stack.Len() > 0 {
		current := stack.Pop()

		for _, dir := range current.directory.SubDirectories {
			stack.Push(StackVar{dir, current.depth + 1})
		}

		fmt.Printf("%sDIR %s %d\n", getEmptyStringOfLength(current.depth), current.directory.Name, current.directory.Size)
		for _, file := range current.directory.Files {
			fmt.Printf("%sFIL %s %d\n", getEmptyStringOfLength(current.depth+1), file.Name, file.Size)
		}
	}
}

type File struct {
	Name string
	Size uint64
}

func ParseInput(fileName string) (*Directory, error) {
	lines, err := util.GetFileAsListOfStrings(fileName)
	if err != nil {
		return nil, err
	}

	root := &Directory{Name: "/"}
	current := root

	lineN := 1
	for lineN < len(lines) {
		if lines[lineN] == "$ ls" {
			lineN++
			lineStart := lineN
			for lineN < len(lines) && !strings.HasPrefix(lines[lineN], "$") {
				lineN++
			}
			ls(lines[lineStart:lineN], current)
		} else if strings.HasPrefix(lines[lineN], "$ cd") {
			if strings.Contains(lines[lineN], "..") {
				current = current.ParentDirectory
			} else {
				current = findDirectory(lines[lineN][5:], current)
			}
			lineN++
		}
	}

	return root, nil
}

func ls(lines []string, d *Directory) {
	for _, line := range lines {
		if strings.HasPrefix(line, "dir") {
			d.SubDirectories = append(d.SubDirectories, &Directory{Name: line[4:], ParentDirectory: d})
		} else {
			splits := strings.Split(line, " ")
			fileSize := util.StringToNumber(splits[0])
			d.Files = append(d.Files, &File{Name: splits[1], Size: uint64(fileSize)})
		}
	}
}

func findDirectory(name string, d *Directory) *Directory {
	if d == nil {
		return nil
	}

	for _, dir := range d.SubDirectories {
		if dir.Name == name {
			return dir
		}
	}

	return nil
}

func getEmptyStringOfLength(n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += "  "
	}
	return res
}
