package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	From int
	To   int
}

type ElfLine struct {
	Range1 Range
	Range2 Range
}

func readInput(filename string) ([]ElfLine, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\n")
	elfLines := make([]ElfLine, len(lines))

	for i, line := range lines {
		ranges := strings.Split(line, ",")
		pair1 := strings.Split(ranges[0], "-")
		pair2 := strings.Split(ranges[1], "-")

		elfLines[i] = ElfLine{
			Range1: Range{
				From: forceReadInt(pair1[0]),
				To:   forceReadInt(pair1[1]),
			},
			Range2: Range{
				From: forceReadInt(pair2[0]),
				To:   forceReadInt(pair2[1]),
			},
		}
	}
	return elfLines, nil
}

func forceReadInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func contains(r1, r2 Range) bool {
	if r1.From >= r2.From && r1.To <= r2.To {
		return true
	}
	if r2.From >= r1.From && r2.To <= r1.To {
		return true
	}
	return false
}

func anyOverlap(r1, r2 Range) bool {
	if (r1.To >= r2.From && r1.From <= r2.To) || (r2.To >= r1.From && r2.From <= r1.To) {
		return true
	}
	return false
}

func part1(data []ElfLine) {
	count := 0
	for _, line := range data {
		if contains(line.Range1, line.Range2) {
			count++
		}
	}
	fmt.Println("Part 1 Count", count)
}

func part2(data []ElfLine) {
	count := 0
	for _, line := range data {
		if anyOverlap(line.Range1, line.Range2) {
			count++
		}
	}
	fmt.Println("Part 2 Count", count)
}

func main() {
	data, err := readInput("day4/day4_input.txt")

	if err != nil {
		panic(err)
	}

	part1(data)
	part2(data)
}
