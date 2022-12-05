package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Crate struct {
	// ordered top to bottom
	Content *list.List
}

type Move struct {
	From   int
	To     int
	Amount int
}

func readInput(filename string) ([]Crate, []Move, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	data := strings.Split(string(bytes), "\n\n")

	crateLines := strings.Split(data[0], "\n")
	// l is the list of numbers at the bottom of the crates
	l := strings.Split(crateLines[len(crateLines)-1], "   ")
	// last is the last item on this slice
	last := l[len(l)-1]
	nCrates, err := strconv.Atoi(last)
	if err != nil {
		panic(err)
	}

	crates := make([]Crate, nCrates)
	for i := 0; i < nCrates; i++ {
		crates[i] = Crate{Content: list.New()}
	}

	// first we read the crate content
	// sample
	//     [H]         [H]         [V]
	// [R] [T] [T] [R] [G] [W] [F] [W] [L]
	for i := 0; i < len(crateLines)-1; i++ {
		line := crateLines[i]
		for j := 0; j < len(line); j++ {
			if line[j] == '[' {
				j++
				var content = ""
				for line[j] != ']' {
					content += string(line[j])
					j++
				}
				crateIndex := j / 4

				crates[crateIndex].Content.PushBack(content)
			}
		}
	}

	moveLines := strings.Split(data[1], "\n")
	moves := make([]Move, len(moveLines))
	j := 0
	for i, line := range moveLines {
		pieces := strings.Split(line, " ")
		moves[j] = Move{
			Amount: forceReadInt(pieces[1]),
			From:   forceReadInt(pieces[3]),
			To:     forceReadInt(pieces[5]),
		}
		i++
		j++
	}
	return crates, moves, nil
}

func forceReadInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func performMoveOnCrates(crates []Crate, move Move) {
	// get the crate
	fromCrate := crates[move.From-1]
	toCrate := crates[move.To-1]

	for i := 0; i < move.Amount; i++ {
		top := fromCrate.Content.Front()
		if top == nil {
			fmt.Println("cannot move, crate is empty")
			continue
		}
		toCrate.Content.PushFront(top.Value)
		// todo need to check if this is only the first.
		fromCrate.Content.Remove(top)
	}
}
func performMoveOnCratesUsingCrateMover9001(crates []Crate, move Move) {
	// get the crate
	fromCrate := crates[move.From-1]
	toCrate := crates[move.To-1]

	moving := make([]string, move.Amount)
	// if we're moving 2 crates, A,B from A,B,C to X, we want it to result in A,B,X, not B,A,X
	for j := 0; j < move.Amount; j++ {
		curHead := fromCrate.Content.Front()
		if curHead == nil {
			break
		}
		moving[j] = curHead.Value.(string)
		fromCrate.Content.Remove(curHead)
	}
	// now we get an array [A, B], we want to add to front of toCrate in reverse order
	for j := len(moving) - 1; j >= 0; j-- {
		toCrate.Content.PushFront(moving[j])
	}
}

func part1() {
	crates, moves, err := readInput("day5/day5_input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Old Crates")
	for i, crate := range crates {
		fmt.Printf("Crate %d: ", i)
		for e := crate.Content.Front(); e != nil; e = e.Next() {
			fmt.Printf("%s ", e.Value)
		}
		fmt.Println()
	}
	for _, move := range moves {
		performMoveOnCrates(crates, move)
	}

	fmt.Println("New Crates")

	for i, crate := range crates {
		fmt.Printf("Crate %d: ", i)
		for e := crate.Content.Front(); e != nil; e = e.Next() {
			fmt.Printf("%s ", e.Value)
		}
		fmt.Println()
	}
}

func part2() {
	crates, moves, err := readInput("day5/day5_input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Old Crates")

	for i, crate := range crates {
		fmt.Printf("Crate %d: ", i)
		for e := crate.Content.Front(); e != nil; e = e.Next() {
			fmt.Printf("%s ", e.Value)
		}
		fmt.Println()
	}
	for _, move := range moves {
		performMoveOnCratesUsingCrateMover9001(crates, move)
	}

	fmt.Println("New Crates")

	for i, crate := range crates {
		fmt.Printf("Crate %d: ", i)
		for e := crate.Content.Front(); e != nil; e = e.Next() {
			fmt.Printf("%s ", e.Value)
		}
		fmt.Println()
	}

}

func main() {
	part1()
	part2()
}
