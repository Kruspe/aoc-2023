package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"strconv"
	"strings"
)

func main() {
	sample := getMinimumCubes(aoc_utils.ReadInput("02/example.txt"))
	data := getMinimumCubes(aoc_utils.ReadInput("02/data.txt"))

	fmt.Printf("Solution 1: %d\n", solve1(sample))
	fmt.Printf("Solution 2: %d\n", solve2(sample))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

func getMinimumCubes(d []string) map[int]map[string]int {
	amounts := make(map[int]map[string]int)
	for i, line := range d {
		colors := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		split := strings.Split(line, ": ")
		rounds := strings.Split(split[1], "; ")
		for _, round := range rounds {
			roundSplit := strings.Split(round, ", ")
			for _, cube := range roundSplit {
				cubeSplit := strings.Split(cube, " ")
				amount, err := strconv.Atoi(cubeSplit[0])
				if err != nil {
					panic(err)
				}
				if amount > colors[cubeSplit[1]] {
					colors[cubeSplit[1]] = amount
				}
			}
		}
		amounts[i] = colors
	}
	return amounts
}

func solve1(amounts map[int]map[string]int) int {
	config := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	result := 0
	for i, m := range amounts {
		gameNumber := i + 1
		result += gameNumber
		for color, amount := range m {
			if amount > config[color] {
				result -= gameNumber
				break
			}
		}
	}
	return result
}

func solve2(amounts map[int]map[string]int) int {
	result := 0
	for _, m := range amounts {
		result += m["blue"] * m["red"] * m["green"]
	}

	return result
}
