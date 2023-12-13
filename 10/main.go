package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"slices"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func main() {
	sample := aoc_utils.ReadInput("10/example.txt")
	sample2 := aoc_utils.ReadInput("10/example2.txt")
	data := aoc_utils.ReadInput("10/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(sample))
	fmt.Printf("Solution 2: %d\n", Solve2(sample2))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", Solve2(data))
}

func findStartPoint(pipeMap [][]string) *coordinate {
	for y, line := range pipeMap {
		for x, p := range line {
			if p == "S" {
				return &coordinate{
					x: x,
					y: y,
				}
			}
		}
	}
	return nil
}

func generateLoop(pipeMap [][]string, startLocation coordinate) map[int]coordinate {
	steps := make(map[int]coordinate)
	steps[0] = startLocation
	c := getFirstStep(pipeMap, steps[0])
	steps[1] = *c

	for {
		lastStep := steps[len(steps)-1]
		stepBeforeLastStep := steps[len(steps)-2]

		var nextStep coordinate
		if lastStep.x-stepBeforeLastStep.x == 1 {
			switch pipeMap[lastStep.y][lastStep.x] {
			case "-":
				nextStep = coordinate{
					x: lastStep.x + 1,
					y: lastStep.y,
				}
			case "J":
				nextStep = coordinate{
					x: lastStep.x,
					y: lastStep.y - 1,
				}
			case "7":
				nextStep = coordinate{
					x: lastStep.x,
					y: lastStep.y + 1,
				}
			}
		} else if lastStep.x-stepBeforeLastStep.x == -1 {
			switch pipeMap[lastStep.y][lastStep.x] {
			case "-":
				nextStep = coordinate{
					x: lastStep.x - 1,
					y: lastStep.y,
				}
			case "L":
				nextStep = coordinate{
					x: lastStep.x,
					y: lastStep.y - 1,
				}
			case "F":
				nextStep = coordinate{
					x: lastStep.x,
					y: lastStep.y + 1,
				}
			}
		} else if lastStep.y-stepBeforeLastStep.y == 1 {
			switch pipeMap[lastStep.y][lastStep.x] {
			case "|":
				nextStep = coordinate{
					x: lastStep.x,
					y: lastStep.y + 1,
				}
			case "L":
				nextStep = coordinate{
					x: lastStep.x + 1,
					y: lastStep.y,
				}
			case "J":
				nextStep = coordinate{
					x: lastStep.x - 1,
					y: lastStep.y,
				}
			}
		} else if lastStep.y-stepBeforeLastStep.y == -1 {
			switch pipeMap[lastStep.y][lastStep.x] {
			case "|":
				nextStep = coordinate{
					x: lastStep.x,
					y: lastStep.y - 1,
				}
			case "F":
				nextStep = coordinate{
					x: lastStep.x + 1,
					y: lastStep.y,
				}
			case "7":
				nextStep = coordinate{
					x: lastStep.x - 1,
					y: lastStep.y,
				}
			}
		}
		if nextStep == startLocation {
			break
		} else {
			steps[len(steps)] = nextStep
		}
	}
	return steps
}

func getFirstStep(pipeMap [][]string, startLocation coordinate) *coordinate {
	up := pipeMap[startLocation.y-1][startLocation.x]
	right := pipeMap[startLocation.y][startLocation.x+1]
	down := pipeMap[startLocation.y+1][startLocation.x]
	left := pipeMap[startLocation.y][startLocation.x-1]
	if up == "|" || up == "7" || up == "F" {
		return &coordinate{
			x: startLocation.x,
			y: startLocation.y - 1,
		}
	}
	if right == "-" || right == "J" || right == "7" {
		return &coordinate{
			x: startLocation.x + 1,
			y: startLocation.y,
		}
	}
	if down == "|" || down == "L" || down == "J" {
		return &coordinate{
			x: startLocation.x,
			y: startLocation.y + 1,
		}
	}
	if left == "-" || left == "L" || left == "F" {
		return &coordinate{
			x: startLocation.x - 1,
			y: startLocation.y,
		}
	}
	return nil
}

func generatePipeMap(d []string) [][]string {
	var pipeMap [][]string
	for _, line := range d {
		pipeMap = append(pipeMap, strings.Split(line, ""))
	}
	return aoc_utils.Pad2dStringArray(pipeMap, ".")
}

func solve1(d []string) int {
	pipeMap := generatePipeMap(d)

	loop := generateLoop(pipeMap, *findStartPoint(pipeMap))
	return len(loop) / 2
}

func Solve2(d []string) int {
	pipeMap := generatePipeMap(d)
	loop := generateLoop(pipeMap, *findStartPoint(pipeMap))

	newMap := make([][]string, 0)
	emptyRow := strings.Split(strings.Repeat("#", len(pipeMap[0])*2), "")
	for _, line := range pipeMap {
		newMap = append(newMap, strings.Split(strings.Join(line, "#")+"#", ""), slices.Clone(emptyRow))
	}

	for _, l := range loop {
		switch pipeMap[l.y][l.x] {
		case "|":
			newMap[l.y*2+1][l.x*2] = "X"
			newMap[l.y*2-1][l.x*2] = "X"
		case "-":
			newMap[l.y*2][l.x*2-1] = "X"
			newMap[l.y*2][l.x*2+1] = "X"
		case "L":
			newMap[l.y*2-1][l.x*2] = "X"
			newMap[l.y*2][l.x*2+1] = "X"
		case "J":
			newMap[l.y*2-1][l.x*2] = "X"
			newMap[l.y*2][l.x*2-1] = "X"
		case "7":
			newMap[l.y*2][l.x*2-1] = "X"
			newMap[l.y*2+1][l.x*2] = "X"
		case "F":
			newMap[l.y*2][l.x*2+1] = "X"
			newMap[l.y*2+1][l.x*2] = "X"
		}
		newMap[l.y*2][l.x*2] = "X"
	}
	aoc_utils.FloodFill(newMap, []string{"X"}, " ", 0, 0)

	result := 0
	for _, line := range newMap {
		result += strings.Count(strings.Join(line, ""), ".")
		result += strings.Count(strings.Join(line, ""), "|")
		result += strings.Count(strings.Join(line, ""), "-")
		result += strings.Count(strings.Join(line, ""), "L")
		result += strings.Count(strings.Join(line, ""), "J")
		result += strings.Count(strings.Join(line, ""), "7")
		result += strings.Count(strings.Join(line, ""), "F")
	}

	return result
}
