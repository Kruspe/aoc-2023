package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := aoc_utils.ReadInput("01/example.txt")
	input2 := aoc_utils.ReadInput("01/example2.txt")
	data := aoc_utils.ReadInput("01/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(input))
	fmt.Printf("Solution 2: %d\n", solve2(input2))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

func solve1(d []string) int {
	result := 0
	for _, line := range d {
		regex := regexp.MustCompile("[0-9]")
		matches := regex.FindAllString(line, -1)
		n, err := strconv.Atoi(matches[0] + matches[len(matches)-1])
		if err != nil {
			panic(err)
		}
		result += n
	}
	return result
}

func solve2(d []string) int {
	result := 0
	for _, line := range d {
		lowestIndex := 9999999999
		highestIndex := -1
		lowestNumber := ""
		highestNumber := ""
		for _, word := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
			index := strings.Index(line, word)
			if index == -1 {
				continue
			}
			if index < lowestIndex {
				lowestIndex = index
				switch word {
				case "one":
					lowestNumber = "1"
				case "two":
					lowestNumber = "2"
				case "three":
					lowestNumber = "3"
				case "four":
					lowestNumber = "4"
				case "five":
					lowestNumber = "5"
				case "six":
					lowestNumber = "6"
				case "seven":
					lowestNumber = "7"
				case "eight":
					lowestNumber = "8"
				case "nine":
					lowestNumber = "9"
				}
			}
			for {
				nextIndex := strings.Index(line[index+len(word):], word)
				if nextIndex == -1 {
					break
				}
				index += len(word) + nextIndex
			}
			if index > highestIndex {
				highestIndex = index
				switch word {
				case "one":
					highestNumber = "1"
				case "two":
					highestNumber = "2"
				case "three":
					highestNumber = "3"
				case "four":
					highestNumber = "4"
				case "five":
					highestNumber = "5"
				case "six":
					highestNumber = "6"
				case "seven":
					highestNumber = "7"
				case "eight":
					highestNumber = "8"
				case "nine":
					highestNumber = "9"
				}
			}
		}
		regex := regexp.MustCompile("[0-9]")
		matches := regex.FindAllStringIndex(line, -1)
		if len(matches) > 0 {
			firstMatch := matches[0][0]
			if firstMatch < lowestIndex {
				lowestIndex = firstMatch
				lowestNumber = string(line[lowestIndex])
			}
			lastMatch := matches[len(matches)-1][0]
			if lastMatch > highestIndex {
				highestIndex = lastMatch
				highestNumber = string(line[highestIndex])
			}
		}
		number, err := strconv.Atoi(lowestNumber + highestNumber)
		if err != nil {
			panic(err)
		}
		result += number
	}

	return result
}
