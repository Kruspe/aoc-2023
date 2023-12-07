package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"slices"
	"strconv"
	"strings"
)

func mappingsFor(m rangeMapping, mappings []map[rangeMapping]rangeMapping) []rangeMapping {
	var r []rangeMapping
	var mapped []int
	for _, mapping := range mappings {
		for source, destination := range mapping {
			offset := destination.from - source.from
			if source.Contains(m.from) && source.Contains(m.to) {
				r = append(r, rangeMapping{
					from: m.from + offset,
					to:   m.to + offset,
				})
				mapped = append(mapped, m.from, m.to)
			} else if source.Contains(m.from) && !source.Contains(m.to) {
				r = append(r, rangeMapping{
					from: m.from + offset,
					to:   source.to + offset,
				})
				mapped = append(mapped, m.from, source.to)
			} else if !source.Contains(m.from) && source.Contains(m.to) {
				r = append(r, rangeMapping{
					from: source.from + offset,
					to:   m.to + offset,
				})
				mapped = append(mapped, source.from, m.to)
			} else if m.Contains(source.from) && m.Contains(source.to) {
				r = append(r, rangeMapping{
					from: destination.from,
					to:   destination.to,
				})
				mapped = append(mapped, source.from, source.to)
			}
		}
	}
	slices.Sort(mapped)
	if len(mapped) == 0 {
		r = append(r, rangeMapping{
			from: m.from,
			to:   m.to,
		})
	} else {
		for i := 0; i < len(mapped); i++ {
			if i == 0 {
				if m.from < mapped[i] {
					r = append(r, rangeMapping{
						from: m.from,
						to:   mapped[i] - 1,
					})
					continue
				}
			}

			if i%2 == 1 {
				if i == len(mapped)-1 {
					if mapped[i]+1 <= m.to {
						r = append(r, rangeMapping{
							from: mapped[i] + 1,
							to:   m.to,
						})
					}
				} else {
					if mapped[i]+1 < mapped[i+1] {
						r = append(r, rangeMapping{
							from: mapped[i],
							to:   mapped[i+1],
						})
					}
				}
			}
		}
	}

	return r
}

type rangeMapping struct {
	from int
	to   int
}

func (r rangeMapping) Contains(key int) bool {
	return r.from <= key && key <= r.to
}

func createMappings(d []string) map[int][]map[rangeMapping]rangeMapping {
	mappings := make(map[int][]map[rangeMapping]rangeMapping)
	separators := aoc_utils.GetSeparators(d)
	for i, separator := range separators {
		var lines []string
		if i+1 != len(separators) {
			lines = d[separator:separators[i+1]]
		} else {
			lines = d[separator:]
		}
		for _, line := range lines[2:] {
			mapping := strings.Split(line, " ")
			destinationStart, err := strconv.Atoi(mapping[0])
			if err != nil {
				panic(err)
			}
			sourceStart, err := strconv.Atoi(mapping[1])
			if err != nil {
				panic(err)
			}
			length, err := strconv.Atoi(mapping[2])
			if err != nil {
				panic(err)
			}

			m := make(map[rangeMapping]rangeMapping)
			m[rangeMapping{
				from: sourceStart,
				to:   sourceStart + length - 1,
			}] = rangeMapping{
				from: destinationStart,
				to:   destinationStart + length - 1,
			}
			mappings[i] = append(mappings[i], m)
		}
	}
	return mappings
}

func main() {
	sample := aoc_utils.ReadInput("05/example.txt")
	data := aoc_utils.ReadInput("05/data.txt")

	fmt.Printf("Solution 1: %d\n", solve1(sample))
	fmt.Printf("Solution 2: %d\n", solve2(sample))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

func solve1(d []string) int {
	ranges := make(map[int][]rangeMapping)
	for _, s := range strings.Split(strings.Split(d[0], "seeds: ")[1], " ") {
		seed, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ranges[0] = append(ranges[0], rangeMapping{
			from: seed,
			to:   seed,
		})
	}

	mappings := createMappings(d)

	for i := 0; i < len(mappings); i++ {
		for _, r := range ranges[i] {
			ranges[i+1] = append(ranges[i+1], mappingsFor(r, mappings[i])...)
		}
	}

	lowestLocation := ranges[len(ranges)-1][0].from
	for _, mapping := range ranges[len(ranges)-1] {
		if mapping.from < lowestLocation {
			lowestLocation = mapping.from
		}
	}
	return lowestLocation
}

func solve2(d []string) int {
	ranges := make(map[int][]rangeMapping)
	seedValues := strings.Split(strings.Split(d[0], "seeds: ")[1], " ")
	for i := 0; i < len(seedValues); i = i + 2 {
		seedStart, err := strconv.Atoi(seedValues[i])
		if err != nil {
			panic(err)
		}
		seedLength, err := strconv.Atoi(seedValues[i+1])
		if err != nil {
			panic(err)
		}
		ranges[0] = append(ranges[0], rangeMapping{
			from: seedStart,
			to:   seedStart + seedLength,
		})
	}

	mappings := createMappings(d)

	for i := 0; i < len(mappings); i++ {
		for _, r := range ranges[i] {
			ranges[i+1] = append(ranges[i+1], mappingsFor(r, mappings[i])...)
		}
	}

	lowestLocation := ranges[len(ranges)-1][0].from
	for _, mapping := range ranges[len(ranges)-1] {
		if mapping.from < lowestLocation {
			lowestLocation = mapping.from
		}
	}
	return lowestLocation
}
