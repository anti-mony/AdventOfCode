package main

import (
	"fmt"
	"log"

	"advent.of.code/list"
	"advent.of.code/util"
)

func main() {
	fileStructure, err := ParseInput("input.small.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileStructure.Print()
}

type StackVar struct {
	directory *Directory
	depth     int
}

type Directory struct {
	Name        string
	Directories []*Directory
	Files       []*File
	Size        uint64
}

func (d *Directory) Print() {
	stack := list.NewStack[StackVar]()
	stack.Push(StackVar{d, 0})

	for stack.Len() > 0 {
		current := stack.Pop()

		for _, dir := range current.directory.Directories {
			stack.Push(StackVar{dir, current.depth + 1})
		}

		fmt.Printf("%sDIR %s\n", getEmptyStringOfLength(current.depth), current.directory.Name)
		for _, file := range current.directory.Files {
			fmt.Printf("%sFIL %s\n", getEmptyStringOfLength(current.depth+1), file.Name)
		}
	}
}

type File struct {
	Name string
	Size uint64
}

func ParseInput(fileName string) (*Directory, error) {
	_, err := util.GetFileAsListOfStrings(fileName)
	if err != nil {
		return nil, err
	}
	return &Directory{
		Name:  "/",
		Files: []*File{{Name: "File 1", Size: 100}},
		Directories: []*Directory{
			{Name: "A", Files: []*File{{Name: "A1", Size: 100}, {Name: "A2"}}, Directories: []*Directory{{Name: "B"}}},
			{Name: "C", Files: []*File{{Name: "C1", Size: 100}}, Directories: []*Directory{{Name: "D", Files: []*File{{Name: "D1"}}}}},
		},
	}, nil
}

func getEmptyStringOfLength(n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += "  "
	}
	return res
}
