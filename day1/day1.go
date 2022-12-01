package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf = struct {
	Calories []int
	Sum      int
}

func readLines(filename string) ([]Elf, error) {

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	input := string(bytes)
	lines := strings.Split(input, "\n\n")

	elves := make([]Elf, len(lines))

	for i, line := range lines {
		words := strings.Split(line, "\n")
		calories := make([]int, len(words))
		for j, word := range words {
			n, err := strconv.Atoi(word)
			if err != nil {
				return nil, err
			}
			calories[j] = n
		}
		sum := 0
		for j := 0; j < len(calories); j++ {
			sum += calories[j]
		}

		elves[i] = Elf{
			Calories: calories,
			Sum:      sum,
		}
	}

	return elves, nil
}

func part1(elves []Elf) {
	max := 0
	for _, elf := range elves {
		if elf.Sum > max {
			max = elf.Sum
		}
	}
	fmt.Println("Part 1 - Max is ", max)
}

func part2(elves []Elf) {
	sums := make([]int, len(elves))

	for i, elf := range elves {
		sums[i] = elf.Sum
	}

	sort.Ints(sums)
	fmt.Println("Part 2 - Max total", sums[len(sums)-1]+sums[len(sums)-2]+sums[len(sums)-3])

}

func main() {
	data, err := readLines("day1/day1_input.txt")
	if err != nil {
		panic(err)
	}

	part1(data)
	part2(data)
}
