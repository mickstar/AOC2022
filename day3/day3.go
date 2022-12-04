package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Line struct {
	Line      string
	Sack1     []uint8
	Sack2     []uint8
	EntireBag []uint8
}

/*
 * Read a file into a slice of strings, one string per line.
 * we will make each ascii character to the corresponding value.
 * a-z maps to 1 to 26
 * A-Z maps to 27 to 52
 */
func readLines(filename string) ([]Line, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	input := strings.Split(string(bytes), "\n")
	lines := make([]Line, len(input))
	for i, line := range input {
		n := len(line)
		sack1 := make([]uint8, n/2)
		sack2 := make([]uint8, n/2)
		entireLine := make([]uint8, n)
		for i, c := range line {
			// we will map a-z to 1-26, and A-Z to 27-52
			c := uint8(c)
			if c >= 'a' {
				c -= 'a'
				c += 1
			} else if c >= 'A' {
				c -= 'A'
				c += 27
			}

			if i < n/2 {
				sack1[i] = c
			} else {
				sack2[i-n/2] = c
			}
			entireLine[i] = c
		}

		sort.Slice(sack1, func(i, j int) bool { return sack1[i] < sack1[j] })
		sort.Slice(sack2, func(i, j int) bool { return sack2[i] < sack2[j] })
		sort.Slice(entireLine, func(i, j int) bool { return entireLine[i] < entireLine[j] })

		lines[i] = Line{
			Line:      line,
			Sack1:     sack1,
			Sack2:     sack2,
			EntireBag: entireLine,
		}
	}
	return lines, nil
}

func findMinIndex(v1, v2, v3 uint8) int {
	if v1 < v2 {
		if v1 < v3 {
			return 1
		}
		return 3
	}
	if v2 < v3 {
		return 2
	}
	return 3
}

func calculateScorePart2(sack1 []uint8, sack2 []uint8, sack3 []uint8) int {
	cursor1 := 0
	cursor2 := 0
	cursor3 := 0

	for true {
		if cursor1 >= len(sack1) || cursor2 >= len(sack2) || cursor3 >= len(sack3) {
			break
		}
		if sack1[cursor1] == sack2[cursor2] && sack2[cursor2] == sack3[cursor3] {
			return int(sack1[cursor1])
		}
		minIndex := findMinIndex(sack1[cursor1], sack2[cursor2], sack3[cursor3])
		if minIndex == 1 && cursor1 < len(sack1)-1 {
			cursor1++
		} else if minIndex == 2 && cursor2 < len(sack2)-1 {
			cursor2++
		} else {
			cursor3++
		}
	}

	fmt.Println("sack1", sack1)
	fmt.Println("sack2", sack2)
	fmt.Println("sack3", sack3)
	panic("No common item found")
}

func calculateScore(line Line) int {
	// iterates over sacks and find common items
	score := 0

	sack1Cursor := 0
	sack2Cursor := 0

	n := len(line.Sack1)

	for true {
		if sack1Cursor >= n || sack2Cursor >= n {
			break
		}
		if line.Sack1[sack1Cursor] == line.Sack2[sack2Cursor] {
			score += int(line.Sack1[sack1Cursor])
			sack1Cursor++
			sack2Cursor++
			for true {
				if sack1Cursor > 0 && sack1Cursor < n && line.Sack1[sack1Cursor] == line.Sack1[sack1Cursor-1] {
					sack1Cursor++
				} else {
					break
				}
			}
			continue
		}
		// at this point we need to increment either cursor 1 or cursor 2
		if sack2Cursor == n-1 || line.Sack1[sack1Cursor] < line.Sack2[sack2Cursor] {
			sack1Cursor++
		} else {
			sack2Cursor++
		}
	}
	return score
}

func part2(data []Line) {
	n := len(data)
	total := 0
	for i := 0; i < n; i += 3 {
		score := calculateScorePart2(data[i].EntireBag, data[i+1].EntireBag, data[i+2].EntireBag)
		total += score
	}
	fmt.Println("Part 2 Total is ", total)
}

func part1(data []Line) {
	total := 0
	for _, line := range data {
		total += calculateScore(line)
	}
	fmt.Println("Part 1 Total is ", total)
}

func main() {
	data, err := readLines("day3/day3_input.txt")
	if err != nil {
		panic(err)
	}

	part1(data)
	part2(data)
}
