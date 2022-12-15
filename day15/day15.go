package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Sensor struct {
	Position      Coordinate
	ClosestBeaken Coordinate
}

func readData(filename string) ([]Sensor, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\n")
	sensors := make([]Sensor, len(lines))

	//regex := regexp.MustCompile(`Sensor at x=(\d+), y=(\d+): closest beacon is at x=(\d+), y=(\d+)`)
	regex := regexp.MustCompile(`x=(?P<sx>-?\d+), y=(?P<sy>-?\d+).*x=(?P<dx>-?\d+), y=(?P<dy>-?\d+)`)
	for i, line := range lines {
		res := regex.FindAllStringSubmatch(line, -1)
		sx, sy, bx, by := res[0][1], res[0][2], res[0][3], res[0][4]
		fmt.Println(sx, sy, bx, by)
		sensors[i] = Sensor{
			Position:      Coordinate{X: readInt(sx), Y: readInt(sy)},
			ClosestBeaken: Coordinate{X: readInt(bx), Y: readInt(by)},
		}
	}
	return sensors, nil
}

func readInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func distance(p1, p2 Coordinate) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

type Cell uint8

const (
	SensorCell   Cell = 'S'
	BeaconCell   Cell = 'B'
	GreyZoneCell Cell = '#' // Greyzone cells are cells that cannot have beacons.
)

type World map[Coordinate]Cell

func simulateWorld(world *World, sensor Sensor) {
	mWorld := *world
	mWorld[sensor.Position] = SensorCell
	mWorld[sensor.ClosestBeaken] = BeaconCell
	d := distance(sensor.Position, sensor.ClosestBeaken)
	// now we will assign every value s.x -d to s.x + d with a dis(x,z) <= d
	for x := sensor.Position.X - d; x <= sensor.Position.X+d; x++ {
		for y := sensor.Position.Y - d; y <= sensor.Position.Y+d; y++ {
			_, exists := mWorld[Coordinate{X: x, Y: y}]
			if exists {
				continue
			}
			if distance(sensor.Position, Coordinate{X: x, Y: y}) <= d {
				mWorld[Coordinate{X: x, Y: y}] = GreyZoneCell
			}
		}
	}
	// done.
}

func getSize(world World) (minX int, maxX int, minY, maxY int) {
	minX, minY, maxX, maxY = 0, 0, 0, 0
	for k := range world {
		if k.X < minX {
			minX = k.X
		}
		if k.Y < minY {
			minY = k.Y
		}
		if k.X > maxX {
			maxX = k.X
		}
		if k.Y > maxY {
			maxY = k.Y
		}
	}
	return
}

func (w World) String() string {
	// bit tricky to print the world as we need to know the dimmensions.
	minX, maxX, minY, maxY := getSize(w)
	s := ""

	for y := minY; y <= maxY; y++ {
		s += fmt.Sprintf("%d - ", y)
		for x := minX; x <= maxX; x++ {
			cell, ok := w[Coordinate{X: x, Y: y}]
			if !ok {
				cell = '.'
			}
			s += string(cell)
		}
		s += "\n"
	}
	return s
}

func main() {
	data, err := readData("day15/day15_input.txt")
	if err != nil {
		panic(err)
	}

	//Y_COORD := 10
	Y_COORD := 2_000_000

	world := make(World)
	for _, sensor := range data {
		world[sensor.Position] = SensorCell
		world[sensor.ClosestBeaken] = BeaconCell
	}

	minX, maxX := -5_000_000, 5_000_000

	greyCount := 0
	for x := minX; x <= maxX; x++ {
		for _, sensor := range data {
			d := distance(sensor.Position, sensor.ClosestBeaken)
			if (distance(sensor.Position, Coordinate{X: x, Y: Y_COORD}) <= d) {
				_, ok := world[Coordinate{X: x, Y: Y_COORD}]
				if !ok {
					greyCount++
					break
				}
			}

		}
	}
	fmt.Println(greyCount)

}
