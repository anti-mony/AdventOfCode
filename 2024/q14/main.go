package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"regexp"

	"advent.of.code/grid"
	"advent.of.code/util"
)

const (
	HEIGHT_SMALL = 7
	WIDTH_SMALL  = 11
	HEIGHT       = 103
	WIDTH        = 101
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

	// fmt.Println("A1: ", Q1(inp, WIDTH, HEIGHT, 100))
	fmt.Println("A2: ", Q2(inp, WIDTH, HEIGHT))
}

func Q1(robots []*Robot, width, height int, seconds int) int {
	for _, robot := range robots {
		for range seconds {
			robot.Position.X = (((robot.Position.X + robot.Velocity.X) % height) + height) % height
			robot.Position.Y = (((robot.Position.Y + robot.Velocity.Y) % width) + width) % width
		}
	}

	return CountInQuadrants(robots, width, height)
}

func Q2(robots []*Robot, width, height int) int {
	seconds := 0
	for seconds < 10000 {
		for _, robot := range robots {
			robot.Position.X = (((robot.Position.X + robot.Velocity.X) % height) + height) % height
			robot.Position.Y = (((robot.Position.Y + robot.Velocity.Y) % width) + width) % width
		}
		seconds++
		if seconds == 7286 {
			createPng(robots, width, height, seconds)
		}
	}

	return seconds
}

func CountInQuadrants(robots []*Robot, width, height int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		if robot.Position.X < height/2 && robot.Position.Y < width/2 {
			q1++
		} else if robot.Position.X < height/2 && robot.Position.Y > width/2 {
			q2++
		} else if robot.Position.X > height/2 && robot.Position.Y > width/2 {
			q3++
		} else if robot.Position.X > height/2 && robot.Position.Y < width/2 {
			q4++
		}
	}
	fmt.Println(q1, q2, q3, q4)
	return q1 * q2 * q3 * q4
}

func createPng(robots []*Robot, W, H int, index int) {
	g := make([][]string, H)
	for i := 0; i < H; i++ {
		g[i] = make([]string, W)
	}

	img := image.NewRGBA(image.Rect(0, 0, W, H))
	for _, robot := range robots {
		img.Set(robot.Position.Y, robot.Position.X, color.RGBA{0, 0, 0, 255})
	}

	// Save the image to a file
	f, err := os.Create(fmt.Sprintf("images/%d.png", index))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}

}

func PrintRobots(robots []*Robot, W, H int) {
	g := make([][]string, H)
	for i := 0; i < H; i++ {
		g[i] = make([]string, W)
	}

	for _, robot := range robots {
		g[robot.Position.X][robot.Position.Y] = "*"
	}

	util.PrintMatrix(g)
}

type Robot struct {
	Position grid.Coordinate
	Velocity grid.Coordinate
}

func parseInput(filename string) ([]*Robot, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`-?[0-9]+`)

	robots := make([]*Robot, len(lines))
	for i, line := range lines {
		matches := re.FindAllString(line, -1)
		robots[i] = &Robot{
			Position: grid.NewCoordinate(util.StringToNumber(matches[1]), util.StringToNumber(matches[0])),
			Velocity: grid.NewCoordinate(util.StringToNumber(matches[3]), util.StringToNumber(matches[2])),
		}
	}

	return robots, nil
}
