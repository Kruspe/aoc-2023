package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"math"
	"regexp"
	"slices"
	"strings"
)

func main() {
	sample := aoc_utils.ReadInput("11/example.txt")
	data := aoc_utils.ReadInput("11/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(sample))
	fmt.Printf("Solution 2: %d\n", Solve2(sample))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", Solve2(data))
}

type coordinate struct {
	x int
	y int
}

func createMap(d []string, expansion int) [][]string {
	tempMap := make([][]string, len(d))
	starMap := make([][]string, 0)
	for i := 0; i < len(d[0]); i++ {
		hasGalaxy := false
		for j, l := range d {
			if l[i] == '#' {
				hasGalaxy = true
			}
			tempMap[j] = append(tempMap[j], string(l[i]))
		}
		if hasGalaxy {
			continue
		}
		for j := range d {
			tempMap[j] = append(tempMap[j], strings.Repeat(".", expansion))
		}
	}
	for _, line := range tempMap {
		starMap = append(starMap, line)
		if slices.Contains(line, "#") {
			continue
		}
		for i := 0; i < expansion; i++ {
			starMap = append(starMap, line)
		}
	}
	return starMap
}

func getGalaxyCoordinates(starMap [][]string) []coordinate {
	var galaxyCoordinates []coordinate
	regex := regexp.MustCompile("#+")
	for y, line := range starMap {
		if slices.Contains(line, "#") {
			for _, x := range regex.FindAllStringIndex(strings.Join(line, ""), -1) {
				galaxyCoordinates = append(galaxyCoordinates, coordinate{x[0], y})
			}
		}
	}
	return galaxyCoordinates
}

func solve1(d []string) int {
	starMap := createMap(d, 1)
	galaxyCoordinates := getGalaxyCoordinates(starMap)

	pairs := aoc_utils.GetPairs(galaxyCoordinates)

	result := 0
	for _, pair := range pairs {
		result += int(math.Abs(float64(pair[0].y - pair[1].y)))
		result += int(math.Abs(float64(pair[0].x - pair[1].x)))
	}

	return result
}

func Solve2(d []string) int {
	starMap := createMap(d, 999999)
	galaxyCoordinates := getGalaxyCoordinates(starMap)

	pairs := aoc_utils.GetPairs(galaxyCoordinates)

	result := 0
	for _, pair := range pairs {
		result += int(math.Abs(float64(pair[0].y - pair[1].y)))
		result += int(math.Abs(float64(pair[0].x - pair[1].x)))
	}

	return result
}
