package main

import (
	"os"
	"strings"
)

func readFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// isAnyOverlap returns true if any of the strings in the slice is not unique
func isAnyOverlap(slice []string) bool {
	dict := map[string]int{}
	for _, c := range slice {
		dict[c] = dict[c] + 1
		if dict[c] > 1 {
			return true
		}
	}
	return false
}

func part1(data string, uniqueCount int) {

}

func findIndexOfUniqueSubtext(data string, uniqueCount int) int {
	slicedData := strings.Split(data, "")

	for i, _ := range slicedData {
		if i < uniqueCount {
			continue
		}
		// we need to check the uniqueness of each character in the last 4
		// we can either pair match 8 ifs, or use a helper function. i will use a helper function
		lookingAt := slicedData[i-(uniqueCount)+1 : i+1]
		if !isAnyOverlap(lookingAt) {
			println("start of packet is at ", i+1, strings.Join(lookingAt, ""))
			return i + 1
		}
	}
	panic("no unique subtext found")
}

func main() {
	data, err := readFile("day6/day6_input.txt")
	if err != nil {
		panic(err)
	}

	findIndexOfUniqueSubtext(data, 4)
	findIndexOfUniqueSubtext(data, 14)

}
