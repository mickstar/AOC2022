package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InstructionCode string

const (
	Noop InstructionCode = "noop"
	addx                 = "addx"
)

type Instruction struct {
	Code InstructionCode
	Arg  int
}

func readFile(filename string) ([]Instruction, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")
	data := make([]Instruction, len(lines))
	for i, line := range lines {
		pieces := strings.Split(line, " ")
		if pieces[0] == "noop" {
			data[i] = Instruction{
				Code: Noop,
			}
		} else if pieces[0] == "addx" {
			data[i] = Instruction{
				Code: addx,
				Arg:  readInt(pieces[1]),
			}
		} else {
			panic("unknown instruction")
		}
	}
	return data, nil
}

func readInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

var x = 1
var cycle = 1

var historicalX = make(map[int]int)

func main() {
	fmt.Println("Hello world")
	data, err := readFile("day10/day10_input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	logCycle(cycle, x)

	for _, instruction := range data {
		//if historicalX[cycle] != 0 {
		//	historicalX[cycle] = x + historicalX[cycle]
		//} else {
		//	historicalX[cycle] = x
		//}

		switch instruction.Code {
		case Noop:
			cycle++
		case addx:
			cycle++
			logCycle(cycle, x)
			cycle++
			x += instruction.Arg
			// when x is 2
			//historicalX[cycle] = instruction.Arg
		}

		logCycle(cycle, x)
	}

	total := 0
	for key := range amounts {
		total += amounts[key]
	}

	fmt.Println("Total:", total)
}

var amounts = make(map[int]int)

func logCycle(cycle int, x int) {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		amounts[cycle] = x * cycle
	}
	printScreen(cycle, x)
}

func printScreen(cycle int, x int) {
	cursor := (cycle - 1) % 40
	if cursor == 0 {
		fmt.Println()
	}
	if cursor == x || cursor == x+1 || cursor == x-1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}
