package main

type Monkey struct {
	id              int
	StartingItems   []int
	Operation       func(int) int
	Test            func(int) bool
	DivisibleBy     int
	MoveToWhenTrue  int
	MoveToWhenFalse int
}

var monkeys []Monkey
var inspectCount map[int]int64

// inputs are hardcoded because i did not want to write an equation parser.
func loadInput() {
	monkeys = make([]Monkey, 8)
	inspectCount = make(map[int]int64)

	monkeys[0] = Monkey{
		id:            0,
		StartingItems: []int{54, 98, 50, 94, 69, 62, 53, 85},
		Operation: func(x int) int {
			return x * 13
		},
		Test: func(x int) bool {
			return x%3 == 0
		},
		DivisibleBy:     3,
		MoveToWhenTrue:  2,
		MoveToWhenFalse: 1,
	}

	monkeys[1] = Monkey{
		id:            1,
		StartingItems: []int{71, 55, 82},
		Operation: func(x int) int {
			return x + 2
		},
		Test: func(x int) bool {
			return x%13 == 0
		},
		DivisibleBy:     13,
		MoveToWhenTrue:  7,
		MoveToWhenFalse: 2,
	}

	monkeys[2] = Monkey{
		id:            2,
		StartingItems: []int{77, 73, 86, 72, 87},
		Operation: func(x int) int {
			return x + 8
		},
		Test: func(x int) bool {
			return x%19 == 0
		},
		DivisibleBy:     19,
		MoveToWhenTrue:  4,
		MoveToWhenFalse: 7,
	}

	monkeys[3] = Monkey{
		id:            3,
		StartingItems: []int{97, 91},
		Operation: func(x int) int {
			return x + 1
		},
		Test: func(x int) bool {
			return x%17 == 0
		},
		DivisibleBy:     17,
		MoveToWhenTrue:  6,
		MoveToWhenFalse: 5,
	}

	monkeys[4] = Monkey{
		id:            4,
		StartingItems: []int{78, 97, 51, 85, 66, 63, 62},
		Operation: func(x int) int {
			// is n % 5 * 17 mod 5 == n * 17 mod 5
			return x * 17
		},
		Test: func(x int) bool {
			return x%5 == 0
		},
		DivisibleBy:     5,
		MoveToWhenTrue:  6,
		MoveToWhenFalse: 3,
	}

	monkeys[5] = Monkey{
		id:            5,
		StartingItems: []int{88},
		Operation: func(x int) int {
			return x + 3
		},
		Test: func(x int) bool {
			return x%7 == 0
		},
		DivisibleBy:     7,
		MoveToWhenTrue:  1,
		MoveToWhenFalse: 0,
	}

	monkeys[6] = Monkey{
		id:            6,
		StartingItems: []int{87, 57, 63, 86, 87, 53},
		Operation: func(x int) int {
			return x * x
		},
		Test: func(x int) bool {
			return x%11 == 0
		},
		DivisibleBy:     11,
		MoveToWhenTrue:  5,
		MoveToWhenFalse: 0,
	}

	monkeys[7] = Monkey{
		id:            7,
		StartingItems: []int{73, 59, 82, 65},
		Operation: func(x int) int {
			return x + 6
		},
		Test: func(x int) bool {
			return x%2 == 0
		},
		DivisibleBy:     2,
		MoveToWhenTrue:  4,
		MoveToWhenFalse: 3,
	}
}

func performRound(divideBy3 bool) {
	for _, monkey := range monkeys {
		for _, item := range monkey.StartingItems {
			inspectCount[monkey.id]++
			newItemScore := monkey.Operation(item)

			if divideBy3 {
				newItemScore = newItemScore / 3
			}

			moveTo := monkey.MoveToWhenTrue
			if !monkey.Test(newItemScore) {
				moveTo = monkey.MoveToWhenFalse
			}

			if !divideBy3 {
				newItemScore = newItemScore % MOD
				if newItemScore == 0 {
					newItemScore = MOD
				}
			}
			monkeys[moveTo].StartingItems = append(monkeys[moveTo].StartingItems, newItemScore)
		}
		// we have processed all the items, so we now clear it from our list
		monkeys[monkey.id].StartingItems = []int{}
	}
}

// MOD multiplication of all divideBys
const MOD = 9_699_690

func part1() {
	loadInput()
	for i := 0; i < 20; i++ {
		performRound(true)
	}

	println("Part 1", calculateAnswer())
}

func part2() {
	loadInput()
	for i := 0; i < 10_000; i++ {
		performRound(false)
	}

	println("Part 2", calculateAnswer())
}

func calculateAnswer() int64 {
	var max1 int64 = 0
	var max2 int64 = 0
	for key := range inspectCount {
		value := inspectCount[key]

		if value > max1 {
			max2 = max1
			max1 = value
		} else if value > max2 {
			max2 = value
		}
	}
	return max1 * max2
}

func main() {
	part1()
	part2()
}
