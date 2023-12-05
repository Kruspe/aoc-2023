package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"slices"
	"strconv"
	"strings"
)

func main() {
	sample := aoc_utils.ReadInput("04/example.txt")
	data := aoc_utils.ReadInput("04/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(sample))
	fmt.Printf("Solution 2: %d\n", solve2(sample))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

func solve1(d []string) int {
	result := 0
	for _, line := range d {
		cleanedLine := strings.Replace(line, "  ", " ", -1)
		worth := 0
		numbers := strings.Split(strings.Split(cleanedLine, ": ")[1], " | ")
		winningNumbers := strings.Split(numbers[0], " ")
		myNumbers := strings.Split(numbers[1], " ")

		for _, number := range myNumbers {
			for _, winningNumber := range winningNumbers {
				if number == winningNumber {
					if worth == 0 {
						worth = 1
					} else {
						worth = worth * 2
					}
				}
			}
		}
		result += worth
	}
	return result
}

func solve2(d []string) int {
	cardResults := make(map[int]int)
	for _, line := range d {
		cleanedLine := strings.Replace(strings.Replace(line, "   ", " ", -1), "  ", " ", -1)
		split := strings.Split(cleanedLine, ": ")
		cardNumber, err := strconv.Atoi(strings.Split(split[0], "Card ")[1])
		if err != nil {
			panic(err)
		}

		numbers := strings.Split(split[1], " | ")
		winningNumbers := strings.Split(numbers[0], " ")
		myNumbers := strings.Split(numbers[1], " ")

		counter := 0
		for _, number := range myNumbers {
			for _, winningNumber := range winningNumbers {
				if number == winningNumber {
					counter += 1
				}
			}
		}

		cardResults[cardNumber] = counter
	}

	cards := make(map[int]int)
	for i := 1; i < len(d)+1; i++ {
		cards[i] = 1
	}

	var keys []int
	for key := range cards {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	for _, key := range keys {
		for i := 0; i < cards[key]; i++ {
			for j := 0; j < cardResults[key]; j++ {
				cards[key+j+1] += 1
			}
		}
	}

	result := 0
	for _, amount := range cards {
		result += amount
	}

	return result
}
