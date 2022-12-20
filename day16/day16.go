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
}

func (v Valve) String() string {
	if v.name == "" {
		return "nil"
	}

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
	}
}

func readInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

const MAX_MINUTE = 30

var valves map[string]Valve = make(map[string]Valve)

type ValveState struct {
	visited bool
	open    bool
}

func main() {
	valveList, err := readInput("day16/day16_test_input.txt")
	if err != nil {
		panic(err)
	}
	for _, v := range valveList {
		valves[v.name] = v
	}

	valvesOpen := make(map[string]ValveState)
	for _, v := range valveList {
		valvesOpen[v.name] = ValveState{
			visited: false,
			open:    false,
		}
	}

	startingPosition := valveList[0].name
	currentPosition := startingPosition
	// current pressure output for the given minute.
	pressureOutput := 0
	// total pressure output for the given minute
	totalPressureOutput := 0

	fmt.Println(calculateValue(valvesOpen, currentPosition, 1))
	panic("done")
	for minute := 1; minute <= MAX_MINUTE; minute++ {
		totalPressureOutput += pressureOutput
		fmt.Println("=== at ", currentPosition, ", minute ", minute, " we are outputing", pressureOutput, " total pressure output is ", totalPressureOutput, "===")

		// we now need to determine which move is best
		openValve := calculateValueOfOpeningValve(valvesOpen, valves[currentPosition], minute)
		moveValues := make(map[string]int)
		for _, tunnel := range valves[currentPosition].tunnels {
			resetVisited(valvesOpen)
			moveValues[tunnel] = calculateValueOfMovingToValve(valvesOpen, valves[tunnel], minute, 0)
		}
		maxTunnel := ""
		maxValue := 0
		for tunnel, value := range moveValues {
			if value > maxValue {
				maxTunnel = tunnel
				maxValue = value
			}
		}
		fmt.Println("max tunnel is ", maxTunnel, " with value ", maxValue, " open valve value is ", openValve)
		if maxValue > openValve {
			fmt.Println("Moving to ", maxTunnel)
			// we are going to move to the tunnel
			currentPosition = maxTunnel
		} else {
			fmt.Println("Opening valve ", currentPosition)
			// we are going to open the valve
			valvesOpen[currentPosition] = ValveState{
				visited: true,
				open:    true,
			}
			pressureOutput += valves[currentPosition].flowRate
		}
	}
	fmt.Println("total pressure out", totalPressureOutput)
}

func calculateValueOfMovingToValve(valvesOpen map[string]ValveState, v Valve, minute int, candidateMax int) int {
	if v.name == "" {
		panic("empty valve")
	}
	// we are going to move to the tunnel which takes 1 minute
	minute++
	if minute >= MAX_MINUTE {
		return 0
	}

	// now we compare turning on the tap at the new position
	// to moving to other positions.
	valueOpeningTap := calculateValueOfOpeningValve(valvesOpen, v, minute)
	nState := make(map[string]ValveState)
	for k, v := range valvesOpen {
		nState[k] = ValveState{
			visited: v.visited,
			open:    v.open,
		}
	}
	nState[v.name] = ValveState{
		visited: true,
	}
	valueOpeningTap += calculateValueOfOpeningValve(nState, v, minute+1)
	fmt.Println("assessed value of opening tap at ", v.name, " is ", valueOpeningTap)
	old := valvesOpen[v.name]
	valvesOpen[v.name] = ValveState{
		open:    old.open,
		visited: true,
	}

	if valueOpeningTap > candidateMax {
		candidateMax = valueOpeningTap
	}

	moveValues := make(map[string]int)
	for _, tunnel := range v.tunnels {
		cached := getCached(tunnel, minute)
		if cached == -1 || cached > candidateMax {
			moveValues[tunnel] = calculateValueOfMovingToValve(valvesOpen, valves[tunnel], minute, candidateMax)
			setCache(tunnel, minute, moveValues[tunnel])
			if moveValues[tunnel] > candidateMax {
				candidateMax = moveValues[tunnel]
			}
		} else {
			fmt.Println("prevented doing something because of cache")
		}
	}
	maxValue := 0
	maxTunnel := ""
	for tunnel, value := range moveValues {
		if value > maxValue {
			maxValue = value
			maxTunnel = tunnel
		}
	}
	if maxTunnel != "" {
		fmt.Println("assessed value of moving to,", maxTunnel, " from ", v.name, "as", maxTunnel)
	}
	if maxValue > valueOpeningTap {
		fmt.Println("\tvalue of moving to", v.name, "is", maxValue, "and that is from moving immediately after to", maxTunnel)
		return maxValue
	} else {
		fmt.Println("\tvalue of moving to", v.name, "is", valueOpeningTap, "and that is from opening the tap")
		return valueOpeningTap
	}
}

var currentOptimal = 0

func calculateValue(state map[string]ValveState, currentPosition string, minute int) int {
	if minute >= MAX_MINUTE-1 {
		return 0
	}

	valueOfOpeningCurrent := calculateValueOfOpeningValve(state, valves[currentPosition], minute+1)
	nState := setValveAsVisited(state, currentPosition)
	valueOfOpeningCurrent += calculateValue(nState, currentPosition, minute+2)

	setCache(currentPosition, minute, valueOfOpeningCurrent)
	if valueOfOpeningCurrent > currentOptimal {
		currentOptimal = valueOfOpeningCurrent
	}

	maxMoveToValue := 0
	maxTunnel := ""
	for _, tunnel := range valves[currentPosition].tunnels {
		cached := getCached(tunnel, minute)
		if cached == -1 || cached > valueOfOpeningCurrent || cached > currentOptimal {
			valueOfMovingToTunnel := calculateValue(state, tunnel, minute+1)
			fmt.Println("value of moving to", tunnel, "from", currentPosition, "is", valueOfMovingToTunnel, "at minute", minute+1)
			if valueOfMovingToTunnel > maxMoveToValue {
				maxMoveToValue = valueOfMovingToTunnel
				maxTunnel = tunnel
			}
			setCache(tunnel, minute, valueOfMovingToTunnel)
			if valueOfMovingToTunnel > currentOptimal {
				currentOptimal = valueOfMovingToTunnel
			}
		}
	}
	if maxMoveToValue > valueOfOpeningCurrent {
		fmt.Println("assessed moving to", maxTunnel, "as", maxMoveToValue, " is optimal")
		return maxMoveToValue
	} else {
		fmt.Println("assessed opening", currentPosition, "as", valueOfOpeningCurrent, " is optimal")
		return valueOfOpeningCurrent
	}
}

var cache map[string]map[int]int = make(map[string]map[int]int)

func clearCache() {
	cache = make(map[string]map[int]int)
}
func setCache(key string, minute int, value int) {
	if cache[key] == nil {
		cache[key] = make(map[int]int)
	}
	cache[key][minute] = value
}

// given some Valve XX, and minute N,
// returns a guarantee that the value of traveling to XX at N is less than the result.
// returns -1 for no guarentee
func getCached(key string, minute int) int {
	if cache[key] == nil {
		return -1
	}
	closestCalculation := -1
	closestMinute := -1
	for k, v := range cache[key] {
		if minute > k {
			if closestMinute == -1 || k < closestMinute {
				closestMinute = k
				closestCalculation = v
			}
		}
	}

	return closestCalculation
}

func resetVisited(valvesOpen map[string]ValveState) {
	for k, v := range valvesOpen {
		valvesOpen[k] = ValveState{
			visited: false,
			open:    v.open,
		}
	}
}

func setValveAsVisited(valvesOpen map[string]ValveState, valve string) map[string]ValveState {
	nState := make(map[string]ValveState)
	for k, v := range valvesOpen {
		nState[k] = ValveState{
			visited: v.visited,
			open:    v.open,
		}
	}
	nState[valve] = ValveState{
		visited: true,
	}
	return nState
}

func calculateValueOfOpeningValve(valvesOpen map[string]ValveState, valve Valve, minute int) int {
	if valve.name == "" {
		panic("empty valve")
	}
	if valvesOpen[valve.name].open {
		// valve already open - no value in opening it again
		return 0
	}
	return valve.flowRate * (MAX_MINUTE - minute)
}
