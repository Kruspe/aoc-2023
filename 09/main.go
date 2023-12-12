package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"strconv"
	"strings"
)

func main() {
	sample := aoc_utils.ReadInput("09/example.txt")
	data := aoc_utils.ReadInput("09/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(sample))
	fmt.Printf("Solution 2: %d\n", Solve2(sample))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", Solve2(data))
}

func createCalculations(line string) map[int][]int {
	split := strings.Split(line, " ")
	var firstLine []int
	for _, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		firstLine = append(firstLine, n)
	}
	calculations := make(map[int][]int)
	calculations[0] = firstLine

	iterator := firstLine
	allEqual := false
	loopCounter := 1
	for !allEqual {
		var r []int
		for i := 0; i < len(iterator)-1; i++ {
			r = append(r, iterator[i+1]-iterator[i])
		}
		calculations[loopCounter] = r
		loopCounter++

		for j := 0; j < len(r)-1; j++ {
			if r[j] != r[j+1] {
				allEqual = false
				iterator = r
				break
			} else {
				allEqual = true
			}
		}
	}
	return calculations
}

func solve1(d []string) int {
	result := 0
	for _, line := range d {
		calculations := createCalculations(line)

		for i := 0; i < len(calculations)-1; i++ {
			prevLine := calculations[len(calculations)-i-2]
			lastNumber := calculations[len(calculations)-i-1][len(calculations[len(calculations)-i-1])-1]

			calculations[len(calculations)-i-2] = append(prevLine, prevLine[len(prevLine)-1]+lastNumber)
		}
		result += calculations[0][len(calculations[0])-1]
	}
	return result
}

func Solve2(d []string) int {
	result := 0
	for _, line := range d {
		calculations := createCalculations(line)

		for i := 0; i < len(calculations)-1; i++ {
			prevLine := calculations[len(calculations)-i-2]
			firstNumber := calculations[len(calculations)-i-1][0]

			calculations[len(calculations)-i-2] = append([]int{prevLine[0] - firstNumber}, prevLine...)
		}
		result += calculations[0][0]
	}
	return result
}
