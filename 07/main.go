package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"slices"
	"strconv"
	"strings"
)

func main() {
	sample := aoc_utils.ReadInput("07/example.txt")
	data := aoc_utils.ReadInput("07/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(sample))
	fmt.Printf("Solution 2: %d\n", Solve2(sample))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", Solve2(data))
}

type hand struct {
	cards []string
	bit   int
}

func solve1(d []string) int {
	hands := make(map[int][]hand)
	for _, line := range d {
		split := strings.Split(line, " ")
		bit, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		cards := strings.Split(split[0], "")
		cardCounts := make(map[string]int)
		for _, s := range cards {
			cardCounts[s]++
		}
		found := false
		for _, amount := range cardCounts {
			switch {
			// Five of a kind
			case amount == 5:
				hands[6] = append(hands[6], hand{
					cards: cards,
					bit:   bit,
				})
				found = true
				break
			// Four of a kind
			case amount == 4:
				hands[5] = append(hands[5], hand{
					cards: cards,
					bit:   bit,
				})
				found = true
				break
			// Full house
			case amount == 3 && len(cardCounts) == 2 || amount == 2 && len(cardCounts) == 2:
				hands[4] = append(hands[4], hand{
					cards: cards,
					bit:   bit,
				})
				found = true
				break
			// Three of a kind
			case amount == 3 && len(cardCounts) == 3:
				hands[3] = append(hands[3], hand{
					cards: cards,
					bit:   bit,
				})
				found = true
				break
			// Two pairs
			case amount == 2 && len(cardCounts) == 3:
				hands[2] = append(hands[2], hand{
					cards: cards,
					bit:   bit,
				})
				found = true
				break
			// One pair
			case amount == 2 && len(cardCounts) == 4:
				hands[1] = append(hands[1], hand{
					cards: cards,
					bit:   bit,
				})
				found = true
				break
			}
			if found {
				break
			}
		}
		if !found {
			hands[0] = append(hands[0], hand{
				cards: cards,
				bit:   bit,
			})
		}
	}
	cardRank := map[string]int{
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": 5,
		"7": 6,
		"8": 7,
		"9": 8,
		"T": 9,
		"J": 10,
		"Q": 11,
		"K": 12,
		"A": 13,
	}
	for _, h := range hands {
		slices.SortFunc(h, func(a, b hand) int {
			for i := 0; i < 5; i++ {
				if cardRank[a.cards[i]] == cardRank[b.cards[i]] {
					continue
				}
				if cardRank[a.cards[i]] > cardRank[b.cards[i]] {
					return 1
				}
				return -1
			}
			return 0
		})
	}
	result := 0
	multiplier := 1
	for i := 0; i < 7; i++ {
		if len(hands[i]) == 0 {
			continue
		}
		for _, h := range hands[i] {
			result = result + h.bit*multiplier
			multiplier++
		}
	}
	return result
}

func Solve2(d []string) int {
	hands := make(map[int][]hand)
	for _, line := range d {
		split := strings.Split(line, " ")
		bit, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		cards := strings.Split(split[0], "")
		cardCounts := make(map[string]int)
		for _, s := range cards {
			cardCounts[s]++
		}
		found := false
		jokerAmount := cardCounts["J"]
		for card, amount := range cardCounts {
			if card == "J" {
				// Five of a kind
				if len(cardCounts) == 1 {
					hands[6] = append(hands[6], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				}
				continue
			}
			if jokerAmount > 0 {
				switch {
				// Five of a kind
				case len(cardCounts) == 2:
					hands[6] = append(hands[6], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				// Four of a kind
				case len(cardCounts) == 3 && (amount == 1 || amount+jokerAmount == 4):
					hands[5] = append(hands[5], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				// Full house
				case amount+jokerAmount == 3 && len(cardCounts) == 3 || amount+jokerAmount == 2 && len(cardCounts) == 3:
					hands[4] = append(hands[4], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				// Three of a kind
				case len(cardCounts) == 4 && (amount == 1 || amount == 2):
					hands[3] = append(hands[3], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				// Two pairs
				case amount+jokerAmount == 2 && len(cardCounts) == 4:
					hands[2] = append(hands[2], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				// One pair
				case amount+jokerAmount == 2 && len(cardCounts) == 5:
					hands[1] = append(hands[1], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				}

			} else {
				switch {
				// Five of a kind
				case amount == 5:
					hands[6] = append(hands[6], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				// Four of a kind
				case amount == 4:
					hands[5] = append(hands[5], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				// Full house
				case amount == 3 && len(cardCounts) == 2 || amount == 2 && len(cardCounts) == 2:
					hands[4] = append(hands[4], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				// Three of a kind
				case amount == 3 && len(cardCounts) == 3:
					hands[3] = append(hands[3], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				// Two pairs
				case amount == 2 && len(cardCounts) == 3:
					hands[2] = append(hands[2], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				// One pair
				case amount == 2 && len(cardCounts) == 4:
					hands[1] = append(hands[1], hand{
						cards: cards,
						bit:   bit,
					})
					found = true
					break
				}
			}

			if found {
				break
			}
		}
		if !found {
			hands[0] = append(hands[0], hand{
				cards: cards,
				bit:   bit,
			})
		}
	}
	cardRank := map[string]int{
		"J": 0,
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": 5,
		"7": 6,
		"8": 7,
		"9": 8,
		"T": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}
	for _, h := range hands {
		slices.SortFunc(h, func(a, b hand) int {
			for i := 0; i < 5; i++ {
				if cardRank[a.cards[i]] == cardRank[b.cards[i]] {
					continue
				}
				if cardRank[a.cards[i]] > cardRank[b.cards[i]] {
					return 1
				}
				return -1
			}
			return 0
		})
	}
	result := 0
	multiplier := 1
	for i := 0; i < 7; i++ {
		if len(hands[i]) == 0 {
			continue
		}
		for _, h := range hands[i] {
			result = result + h.bit*multiplier
			multiplier++
		}
	}
	return result
}
