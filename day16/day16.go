package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	name     string
	flowRate int
	tunnels  []string
	isOpen   bool
}

func (v Valve) String() string {
	return fmt.Sprintf("Valve %s has flow rate=%d; tunnels lead to valves %s", v.name, v.flowRate, strings.Join(v.tunnels, ", "))
}

// sample input Valve ZN has flow rate=0; tunnels lead to valves SD, ZV
func readInput(filename string) ([]Valve, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")
	valves := make([]Valve, len(lines))
	for i, line := range lines {
		valves[i] = parseLine(line)
	}
	return valves, nil
}

func parseLine(line string) Valve {
	exp, err := regexp.Compile(`Valve (\w+) has flow rate=(\d+); tunnel[s]* lead[s]* to valve[s]* (.*)`)
	if err != nil {
		panic(err)
	}
	m := exp.FindAllStringSubmatch(line, -1)
	valve := m[0][1]
	flow := readInt(m[0][2])
	valvesString := m[0][3]

	valves := strings.Split(valvesString, ", ")
	return Valve{
		name:     valve,
		flowRate: flow,
		tunnels:  valves,
		isOpen:   false,
	}
}

func readInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

const MAX_MINUTE = 60

func main() {
	valveList, err := readInput("day16/day16_input.txt")
	if err != nil {
		panic(err)
	}
	valves := make(map[string]Valve)

	for _, v := range valveList {
		valves[v.name] = v
	}

	valvesOpen := make(map[string]bool)

	startingPosition := valveList[0].name
	currentPosition := startingPosition
	pressureOutput := 0
	for minute := 1; minute <= MAX_MINUTE; minute++ {
		// we now need to determine which move is best
		openValve := calculateValueOfOpeningValve(valvesOpen, valves[currentPosition], minute)
		moveValues := make(map[string]int)
		for _, tunnel := range valves[currentPosition].tunnels {
			moveValues[tunnel] = calculateValueOfMovingToValve(valvesOpen, valves[currentPosition], minute)
		}
		maxTunnel := ""
		maxValue := 0
		for tunnel, value := range moveValues {
			if value > maxValue {
				maxTunnel = tunnel
				maxValue = value
			}
		}
		if maxValue > openValve {
			// we are going to move to the tunnel
			currentPosition = maxTunnel
			minute++
		} else {
		}
	}
}

func calculateValueOfMovingToValve(v Valve, minute int) int {
	return 0
}

func calculateValueOfOpeningValve(valve Valve, minute int) int {
	return valve.flowRate * (MAX_MINUTE - minute)
}
