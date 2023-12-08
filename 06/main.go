package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sample := aoc_utils.ReadInput("06/example.txt")
	data := aoc_utils.ReadInput("06/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(sample))
	fmt.Printf("Solution 2: %d\n", solve2(sample))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

func solve1(d []string) int {
	regex := regexp.MustCompile("[0-9]+")
	var durations []int
	var records []int
	for _, s := range regex.FindAllString(d[0], -1) {
		d, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		durations = append(durations, d)
	}
	for _, s := range regex.FindAllString(d[1], -1) {
		r, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		records = append(records, r)
	}

	errorMargin := calculateErrorMargin(durations, records)

	return errorMargin
}

func calculateErrorMargin(durations []int, records []int) int {
	errorMargin := 1
	for i := 0; i < len(durations); i++ {
		solutionCounter := 0
		for j := 1; j < durations[i]; j++ {
			distance := (durations[i] - j) * j
			if distance > records[i] {
				solutionCounter++
			}
		}
		errorMargin = errorMargin * solutionCounter

	}
	return errorMargin
}

func solve2(d []string) int {
	regex := regexp.MustCompile("[0-9\\s]+")
	duration, err := strconv.Atoi(strings.Replace(regex.FindString(d[0]), " ", "", -1))
	if err != nil {
		panic(err)
	}
	record, err := strconv.Atoi(strings.Replace(regex.FindString(d[1]), " ", "", -1))
	if err != nil {
		panic(err)
	}

	return calculateErrorMargin([]int{duration}, []int{record})
}
