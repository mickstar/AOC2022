package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) ([][]int, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\n")
	data := make([][]int, len(lines))
	for i, line := range lines {
		data[i] = make([]int, len(line))
		for j, char := range line {
			data[i][j] = int(char)
		}
	}
	return data, nil
}

func isVisible(x, y int, data [][]int) bool {
	if x < 0 || y < 0 || x >= len(data[0]) || y >= len(data) {
		panic("out of bounds")
	}

	if x == 0 || y == 0 || x == len(data[0])-1 || y == len(data)-1 {
		return true
	}

	value := data[y][x]

	// test To Top
	var isHigher = false
	for j := y - 1; j >= 0; j-- {
		if data[j][x] >= value {
			isHigher = true
			break
		}
	}
	if !isHigher {
		return true
	}
	// need to repeat this 3 times for the other directions

	// test To Top
	isHigher = false
	for j := y + 1; j < len(data); j++ {
		if data[j][x] >= value {
			isHigher = true
			break
		}
	}
	if !isHigher {
		return true
	}

	// test To Top
	isHigher = false
	for i := x - 1; i >= 0; i-- {
		if data[y][i] >= value {
			isHigher = true
			break
		}
	}
	if !isHigher {
		return true
	}
	// need to repeat this 3 times for the other directions

	// test To Top
	isHigher = false
	for i := x + 1; i < len(data); i++ {
		if data[y][i] >= value {
			isHigher = true
			break
		}
	}
	if !isHigher {
		return true
	}

	return false
}

func calculateScenicScore(x, y int, data [][]int) int {
	value := data[y][x]

	if x == 0 || y == 0 || x == len(data[0])-1 || y == len(data)-1 {
		return 0
	}

	scenicScore := 1

	j := y - 1
	for j > 0 {
		if data[j][x] >= value {
			break
		}
		j--
	}
	scenicScore *= y - j

	j = y
	for j < len(data)-1 {
		j++
		if data[j][x] >= value {
			break
		}
	}
	scenicScore *= (j) - y

	i := x
	for i > 0 {
		i--
		if data[y][i] >= value {
			break
		}
	}
	scenicScore *= x - i

	i = x
	for i < len(data[y])-1 {
		i++
		if data[y][i] >= value {
			break
		}
	}
	scenicScore *= i - x

	return scenicScore
}

func part1(data [][]int) {
	count := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {
			if isVisible(x, y, data) {
				count++
			}
		}
	}

	fmt.Println("Part 1", count)
}

func part2(data [][]int) {
	best := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {
			score := calculateScenicScore(x, y, data)
			if score > best {
				best = score
			}
		}
	}

	fmt.Println("Part 2", best)
}

func main() {
	data, err := readFile("day8/day8_input.txt")
	if err != nil {
		panic(err)
	}

	part1(data)
	part2(data)
}
