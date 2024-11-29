#!/bin/bash

if [ -z "$1" ] || [ -z "$2" ]
  then
    echo "question number is a required argument"
    echo "usage: initQuestion YYYY:int DAY:int"
    echo "example: ./initQuestion 2021 6"
    exit 1
fi

basedir=$(dirname "$0")
mkdir "$basedir/$1/q$2"
cd "$basedir/$1/q$2"
touch input.small.txt input.txt
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

session=