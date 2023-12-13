package main

import (
	"fmt"
	aoc_utils "github.com/kruspe/aoc-utils"
	"regexp"
	"slices"
	"strconv"
	"unicode"
)

func main() {
	sample := newSolver(aoc_utils.PadStringArray(aoc_utils.ReadInput("03/example.txt"), "."))
	data := newSolver(aoc_utils.PadStringArray(aoc_utils.ReadInput("03/data.txt"), "."))

	fmt.Printf("Solution 1: %d\n", sample.solve1())
	fmt.Printf("Solution 2: %d\n", sample.solve2())

	fmt.Printf("Solution 1: %d\n", data.solve1())
	fmt.Printf("Solution 2: %d\n", data.solve2())
}

type solver struct {
	d       []string
	numbers map[int][][]int
	symbols []coordinates
}

func newSolver(d []string) solver {
	numbers := make(map[int][][]int)
	regex := regexp.MustCompile("[0-9]+")
	symbols := make([]coordinates, 0)
	for i, line := range d {
		n := regex.FindAllStringIndex(line, -1)
		for _, number := range n {
			numbers[i] = append(numbers[i], []int{number[0], number[len(number)-1]})
		}
		for j := 0; j < len(line); j++ {
			if string(line[j]) != "." && !unicode.IsNumber([]rune(string(line[j]))[0]) {
				symbols = append(
					symbols, coordinates{
						x: j,
						y: i,
					},
				)
			}
		}
	}
	return solver{
		d:       d,
		numbers: numbers,
		symbols: symbols,
	}
}

type coordinates struct {
	x int
	y int
}

func (s solver) solve1() int {
	myNumbers := make(map[int][][]int)
	for i, n := range s.numbers {
		clone := slices.Clone(n)
		myNumbers[i] = clone
	}
	result := 0
	for _, symbol := range s.symbols {
		if unicode.IsNumber([]rune(string(s.d[symbol.y-1][symbol.x-1]))[0]) {
			number, i := s.getNumber(symbol.x-1, symbol.y-1, myNumbers)

			if i != -1 {
				myNumbers[symbol.y-1] = append(myNumbers[symbol.y-1][:i], myNumbers[symbol.y-1][i+1:]...)
				result += number
			}
		}
		if unicode.IsNumber([]rune(string(s.d[symbol.y-1][symbol.x]))[0]) {
			number, i := s.getNumber(symbol.x, symbol.y-1, myNumbers)
			if i != -1 {
				myNumbers[symbol.y-1] = append(myNumbers[symbol.y-1][:i], myNumbers[symbol.y-1][i+1:]...)
				result += number
			}
		}
		if unicode.IsNumber([]rune(string(s.d[symbol.y-1][symbol.x+1]))[0]) {
			number, i := s.getNumber(symbol.x+1, symbol.y-1, myNumbers)
			if i != -1 {
				myNumbers[symbol.y-1] = append(myNumbers[symbol.y-1][:i], myNumbers[symbol.y-1][i+1:]...)
				result += number
			}
		}
		if unicode.IsNumber([]rune(string(s.d[symbol.y][symbol.x-1]))[0]) {
			number, i := s.getNumber(symbol.x-1, symbol.y, myNumbers)
			if i != -1 {
				myNumbers[symbol.y] = append(myNumbers[symbol.y][:i], myNumbers[symbol.y][i+1:]...)
				result += number
			}
		}
		if unicode.IsNumber([]rune(string(s.d[symbol.y][symbol.x+1]))[0]) {
			number, i := s.getNumber(symbol.x+1, symbol.y, myNumbers)
			if i != -1 {
				myNumbers[symbol.y] = append(myNumbers[symbol.y][:i], myNumbers[symbol.y][i+1:]...)
				result += number
			}
		}
		if unicode.IsNumber([]rune(string(s.d[symbol.y+1][symbol.x-1]))[0]) {
			number, i := s.getNumber(symbol.x-1, symbol.y+1, myNumbers)
			if i != -1 {
				myNumbers[symbol.y+1] = append(myNumbers[symbol.y+1][:i], myNumbers[symbol.y+1][i+1:]...)
				result += number
			}
		}
		if unicode.IsNumber([]rune(string(s.d[symbol.y+1][symbol.x]))[0]) {
			number, i := s.getNumber(symbol.x, symbol.y+1, myNumbers)
			if i != -1 {
				myNumbers[symbol.y+1] = append(myNumbers[symbol.y+1][:i], myNumbers[symbol.y+1][i+1:]...)
				result += number
			}
		}
		if unicode.IsNumber([]rune(string(s.d[symbol.y+1][symbol.x+1]))[0]) {
			number, i := s.getNumber(symbol.x+1, symbol.y+1, myNumbers)
			if i != -1 {
				myNumbers[symbol.y+1] = append(myNumbers[symbol.y+1][:i], myNumbers[symbol.y+1][i+1:]...)
				result += number
			}
		}
	}
	return result
}

func (s solver) getNumber(
	x, y int,
	numbers map[int][][]int,
) (int, int) {

	for i, n := range numbers[y] {
		if n[0] <= x && x < n[1] {
			x, err := strconv.Atoi(s.d[y][n[0]:n[1]])
			if err != nil {
				panic(err)
			}
			return x, i
		}
	}
	return 0, -1
}

func (s solver) solve2() int {
	result := 0
	for _, symbol := range s.symbols {
		if string(s.d[symbol.y][symbol.x]) == "*" {
			r := make(map[int]int)
			if unicode.IsNumber([]rune(string(s.d[symbol.y-1][symbol.x-1]))[0]) {
				number, i := s.getNumber(symbol.x-1, symbol.y-1, s.numbers)
				if i != -1 {
					r[number] = number
				}
			}
			if unicode.IsNumber([]rune(string(s.d[symbol.y-1][symbol.x]))[0]) {
				number, i := s.getNumber(symbol.x, symbol.y-1, s.numbers)
				if i != -1 {
					r[number] = number
				}
			}
			if unicode.IsNumber([]rune(string(s.d[symbol.y-1][symbol.x+1]))[0]) {
				number, i := s.getNumber(symbol.x+1, symbol.y-1, s.numbers)
				if i != -1 {
					r[number] = number
				}
			}
			if unicode.IsNumber([]rune(string(s.d[symbol.y][symbol.x-1]))[0]) {
				number, i := s.getNumber(symbol.x-1, symbol.y, s.numbers)
				if i != -1 {
					r[number] = number
				}
			}
			if unicode.IsNumber([]rune(string(s.d[symbol.y][symbol.x+1]))[0]) {
				number, i := s.getNumber(symbol.x+1, symbol.y, s.numbers)
				if i != -1 {
					r[number] = number
				}
			}
			if unicode.IsNumber([]rune(string(s.d[symbol.y+1][symbol.x-1]))[0]) {
				number, i := s.getNumber(symbol.x-1, symbol.y+1, s.numbers)
				if i != -1 {
					r[number] = number
				}
			}
			if unicode.IsNumber([]rune(string(s.d[symbol.y+1][symbol.x]))[0]) {
				number, i := s.getNumber(symbol.x, symbol.y+1, s.numbers)
				if i != -1 {
					r[number] = number
				}
			}
			if unicode.IsNumber([]rune(string(s.d[symbol.y+1][symbol.x+1]))[0]) {
				number, i := s.getNumber(symbol.x+1, symbol.y+1, s.numbers)
				if i != -1 {
					r[number] = number
				}
			}
			if len(r) == 2 {
				ratio := 1
				for key := range r {
					ratio = ratio * key
				}
				result += ratio
			}
		}

	}
	return result
}
