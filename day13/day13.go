package main

import (
	"fmt"
	"os"
	"strings"
)

type Packet struct {
	// either has a child packet or a value
	ChildPacket *Packet
	value       int
	hasValue    bool
}

func (p *Packet) String() string {
	if p == nil {
		return "nil"
	}
	if p.hasValue {
		return fmt.Sprintf("[%d]", p.value)
	}
	return fmt.Sprintf("[%s]", p.ChildPacket)
}

func readData(filename string) ([][]*Packet, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	pairs := strings.Split(string(bytes), "\n\n")
	packets := make([][]*Packet, len(pairs))

	for i, pair := range pairs {
		packets[i] = make([]*Packet, 2)
		for j, packet := range strings.Split(pair, "\n") {
			packets[i][j], _ = parsePacket(packet, 0)
		}
	}
	return packets, nil
}

func parsePacket(packet string, start int) ([]*Packet, int) {
	index := start
	for packet[index] != '[' {
		index++
	}

	// now we check the char after the [ to see if it's a number or a [
	// if it's a number, we have a value packet
	// if it's a [, we have a child packet
	if packet[index+1] == '[' {
		// we have a child packet
		innerPacket, endIndex := parsePacket(packet, index+1)
		return &Packet{
			ChildPacket: innerPacket,
		}, endIndex
	} else {
		endIndex := index + 1
		for packet[endIndex] != ']' {
			endIndex++
		}
		values := strings.Split(packet[index+1:endIndex], ",")
	}
}

func main() {

	data, err := readData("day13/day13_input.txt")
	if err != nil {
		panic(err)
	}
	for _, packet := range data {
		fmt.Println(packet)
	}

}
