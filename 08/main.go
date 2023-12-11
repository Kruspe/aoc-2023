package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"strings"
)

func main() {
	sample := aoc_utils.ReadInput("08/example.txt")
	sample2 := aoc_utils.ReadInput("08/example2.txt")
	data := aoc_utils.ReadInput("08/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(sample))
	fmt.Printf("Solution 2: %d\n", Solve2(sample2))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", Solve2(data))
}

func solve1(d []string) int {
	directions := strings.Split(d[0], "")
	nodes := make(map[string][]string)
	for _, n := range d[2:] {
		split := strings.Split(n, " = ")
		replace := strings.Replace(strings.Replace(split[1], "(", "", -1), ")", "", -1)
		nodes[split[0]] = strings.Split(replace, ", ")
	}

	node := "AAA"
	counter, stepper := 0, 0
	for {
		if node == "ZZZ" {
			break
		}
		if stepper == len(directions) {
			stepper = 0
		}

		if directions[stepper] == "L" {
			node = nodes[node][0]
		} else {
			node = nodes[node][1]
		}
		counter++
		stepper++
	}
	return counter
}

func Solve2(d []string) int {
	directions := strings.Split(d[0], "")
	nodes := make(map[string][]string)
	for _, n := range d[2:] {
		split := strings.Split(n, " = ")
		replace := strings.Replace(strings.Replace(split[1], "(", "", -1), ")", "", -1)
		nodes[split[0]] = strings.Split(replace, ", ")
	}

	var nodesToNavigate []string
	for key := range nodes {
		if string(key[2]) == "A" {
			nodesToNavigate = append(nodesToNavigate, key)
		}
	}

	var node string
	var steps []int
	for _, n := range nodesToNavigate {
		counter, stepper := 0, 0
		node = n
		for {
			if string(node[2]) == "Z" {
				steps = append(steps, counter)
				break
			}
			if stepper == len(directions) {
				stepper = 0
			}

			if directions[stepper] == "L" {
				node = nodes[node][0]
			} else {
				node = nodes[node][1]
			}
			counter++
			stepper++
		}
	}

	lcm, err := aoc_utils.LCM(steps)
	if err != nil {
		panic(err)
	}
	return lcm
}
