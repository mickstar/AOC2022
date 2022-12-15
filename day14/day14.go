package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Packet struct {
	A int
	B int
}

const WORLD_WIDTH = 1200

func (p *Packet) String() string {
	return fmt.Sprintf("%d,%d", p.A, p.B)
}

type Chain []Packet

type Cell uint8

const PART_1 = false

const (
	Empty        Cell = '.'
	Rock         Cell = '#'
	SandProducer Cell = '+'
	Sand         Cell = 'o'
)

const (
	Successful Status = iota
	OffEdge
)

type Status int

type World [][]Cell

func (c Chain) String() string {
	s := make([]string, len(c))
	for i, p := range c {
		s[i] = p.String()
	}
	return strings.Join(s, " -> ")
}

func readInput(filename string) ([]Chain, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")

	chains := make([]Chain, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " -> ")
		chain := make([]Packet, len(parts))
		for j, part := range parts {
			n := strings.Split(part, ",")
			chain[j] = Packet{
				A: readInt(n[0]),
				B: readInt(n[1]),
			}
		}
		chains[i] = chain
	}

	return chains, nil
}

func readInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func populateWorld(chains []Chain) World {
	var maxX, maxY int = WORLD_WIDTH, 0

	for _, chain := range chains {
		for _, packet := range chain {
			if packet.A > maxX {
				maxX = packet.A
			}
			if packet.B > maxY {
				maxY = packet.B
			}
		}
	}

	fmt.Println("Generating world with size x:", maxX, " y:", maxY)
	World := make(World, maxY+3)
	for i := range World {
		World[i] = make([]Cell, maxX+1)
	}

	for _, row := range World {
		for i := range World[0] {
			row[i] = Empty
		}
	}

	// now we need to read our chains and populate the rocks
	for _, chain := range chains {
		for i := 0; i < len(chain)-1; i++ {
			packet := chain[i]
			nextPacket := chain[i+1]
			// now we need to draw a line from packet to nextPacket
			// we cannot assume packets are in order so
			// we need to check which is greater and draw the line
			// from the greater to the lesser
			fromA, toA := packet.A, nextPacket.A
			fromB, toB := packet.B, nextPacket.B
			if fromA > toA {
				fromA, toA = toA, fromA
			}
			if fromB > toB {
				fromB, toB = toB, fromB
			}

			for x := fromA; x <= toA; x++ {
				for y := fromB; y <= toB; y++ {
					World[y][x] = Rock
				}
			}
		}
	}

	for i, _ := range World[len(World)-1] {
		World[len(World)-1][i] = Rock
	}

	World[SAND_PRODUCER_Y][SAND_PRODUCER_X] = SandProducer

	return World
}

func (w World) String() string {
	var s strings.Builder
	for _, row := range w {
		for _, cell := range row {
			s.WriteRune(rune(cell))
		}
		s.WriteRune('\n')
	}
	return s.String()
}

const (
	SAND_PRODUCER_X = 500
	SAND_PRODUCER_Y = 0
)

func dropSand(world World, x, y int) Status {
	if world[SAND_PRODUCER_Y][SAND_PRODUCER_X] == Sand {
		return OffEdge
	}
	for y < len(world) && world[y+1][x] == Empty {
		y++
	}

	// we now have the piece of sand one above the empty piece.
	// the unit of sand will now attempt to move diagonally down and left
	if world[y+1][x-1] == Empty {
		return dropSand(world, x-1, y+1)

	} else if world[y+1][x+1] == Empty {
		return dropSand(world, x+1, y+1)
	} else {
		// sand is in the right position.
		world[y][x] = Sand
		return Successful
	}
}

func simulateWorld(world World) {
	// we generate one unit of sand at (500, 1)
	// we need apply gravity to that unit of sand until it reachs the bottom

	var sX, sY int = SAND_PRODUCER_X, SAND_PRODUCER_Y

	i := 0
	for true {
		status := dropSand(world, sX, sY)
		if status == OffEdge {
			break
		}
		i++
	}
	fmt.Println(world)
	fmt.Println("(Part 2) Performed", i, "iterations")
}

func main() {
	chains, err := readInput("day14/day14_input.txt")
	if err != nil {
		panic(err)
	}

	world := populateWorld(chains)
	simulateWorld(world)
}
