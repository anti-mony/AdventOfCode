#!/bin/bash

if [ -z "$1" ]
  then
    echo "question number is a required argument"
    echo "usage: initQuestion <qN>"
    echo example: ./initQuestion q6
    exit 1
fi

mkdir $1
cd $1
touch input.txt input.small.txt
cat >main.go << EOF
package main

import (
	"fmt"
	"log"

	"advent.of.code/util"
)

func main() {
	inp, err := parseInput("input.small.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inp)
}

func parseInput(filename string) ([]string, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
EOF
