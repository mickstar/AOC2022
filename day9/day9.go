package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	Direction Direction
	Amount    int
}

type Direction uint8

const (
	L Direction = 'L'
	R           = 'R'
	U           = 'U'
	D           = 'D'
)

func readDirection(c uint8) Direction {
	switch c {
	case 'L':
		return L
	case 'R':
		return R
	case 'U':
		return U
	case 'D':
		return D
	default:
		panic("invalid direction")
	}
}

func readFile(filename string) ([]Move, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\n")
	moves := make([]Move, len(lines))
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		moves[i] = Move{
			Direction: readDirection(parts[0][0]),
			Amount:    amount,
		}
	}
	return moves, nil
}

// wtf Golang...
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func performMove(direction Direction, knots int) {
	originalPosition := positions[0]

	switch direction {
	case D:
		positions[0] = Position{originalPosition.X, originalPosition.Y - 1}
	case U:
		positions[0] = Position{originalPosition.X, originalPosition.Y + 1}
	case L:
		positions[0] = Position{originalPosition.X - 1, originalPosition.Y}
	case R:
		positions[0] = Position{originalPosition.X + 1, originalPosition.Y}
	}

	// now we need to perform the knot move for each knot
	for i := 1; i <= knots; i++ {
		current := positions[i]
		leading := positions[i-1]
		if isNextTo(current, leading) {
			// do nothing
			break
		}

		var xDelta, yDelta int = 0, 0

		if current.X < leading.X {
			xDelta = 1
		} else if current.X > leading.X {
			xDelta = -1
		}
		// check if leading Y is higher than current Y, if so we need to move up
		if current.Y < leading.Y {
			yDelta = 1
		} else if current.Y > leading.Y {
			yDelta = -1
		}
		positions[i] = Position{current.X + xDelta, current.Y + yDelta}
	}

	return
}

func isNextTo(a, b Position) bool {
	dX := Abs(a.X - b.X)
	dY := Abs(a.Y - b.Y)

	return dX <= 1 && dY <= 1
}

// can use a complex number but this is easier to read
type Position struct {
	X, Y int
}

var positions = make([]Position, 10)
var loggedPositions map[string]int = make(map[string]int)

func logPosition(position Position) {
	key := fmt.Sprintf("%d_%d", position.X, position.Y)
	loggedPositions[key] = 1

	for key, _ := range loggedPositions {
		if loggedPositions[key] == 0 {
			continue
		}
	}
}

func countUniquePositions() int {
	count := 0
	for key, _ := range loggedPositions {
		if loggedPositions[key] > 0 {
			count++
		}
	}
	return count
}

func part1(moves []Move) {

	for i := 0; i < 2; i++ {
		positions[i] = Position{0, 0}
	}

	logPosition(positions[0])
	for _, move := range moves {
		for i := move.Amount; i > 0; i-- {
			performMove(move.Direction, 1)
			logPosition(positions[1])
		}
	}

	fmt.Println("Part 1", countUniquePositions())
}

// debug method to bring the board, supply some X,Y as (40,40) for test input 2.
func printGrid(X, Y int, knots int) {
	grid := make([][]string, X)
	for i := range grid {
		grid[i] = make([]string, Y)
	}
	for i := 0; i <= knots; i++ {
		// add X/2 and Y/2 to center the grid and deal with negative values.
		grid[positions[i].X+X/2][positions[i].Y+Y/2] = fmt.Sprintf("%d", i)
	}

	for i := Y - 1; i >= 0; i-- {
		for j := 0; j < X; j++ {
			if grid[j][i] == "" {
				grid[j][i] = "."
			}
			fmt.Print(grid[j][i])
		}
		fmt.Println("")
	}
}

func part2(moves []Move) {
	loggedPositions = make(map[string]int)

	for i := 0; i < 10; i++ {
		positions[i] = Position{0, 0}
	}
	logPosition(positions[9])
	for _, move := range moves {
		for i := move.Amount; i > 0; i-- {
			performMove(move.Direction, 9)
			logPosition(positions[9])
		}
	}

	fmt.Println("Part 2", countUniquePositions())
}

func main() {
	moves, err := readFile("day9/day9_input.txt")
	if err != nil {
		panic(err)
	}

	part1(moves)
	part2(moves)

}
