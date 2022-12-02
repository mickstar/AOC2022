package main

import (
	"fmt"
	"os"
	"strings"
)

type Move uint8

const (
	ROCK     Move = 1
	PAPER         = 2
	SCISSORS      = 3
)

type Result uint8

const (
	PLAYER1_WIN Result = 0
	PLAYER2_WIN        = 1
	DRAW               = 2
)

type Game = struct {
	Player1Move Move
	Player2Move Move
}

// didn't really work out for part 2...
func parseMove(move string) Move {
	switch move {
	case "A", "X":
		return ROCK
	case "B", "Y":
		return PAPER
	case "C", "Z":
		return SCISSORS
	}
	panic("Invalid move")
}

func calculateScore(game Game) (score uint8) {
	score = uint8(game.Player2Move)
	result := play(game.Player1Move, game.Player2Move)
	if result == DRAW {
		// draw
		score += 3
	} else if result == PLAYER1_WIN {
		// player 1 wins
	} else if result == PLAYER2_WIN {
		//player 2 wins
		score += 6
	}
	return
}

// turns out part 2 didn't really match what i expected...
func calculateScorePart2(game Game) (score uint8) {
	// we will map our old game to a new game object using the new ruleset.

	var player2Move Move = ROCK // default
	if game.Player2Move == ROCK {
		// X which means lose
		if game.Player1Move == ROCK {
			player2Move = SCISSORS
		} else if game.Player1Move == PAPER {
			player2Move = ROCK
		} else {
			player2Move = PAPER
		}
	} else if game.Player2Move == PAPER {
		// Y which means draw
		player2Move = game.Player1Move
	} else {
		// Z which means win
		// X which means lose
		if game.Player1Move == ROCK {
			player2Move = PAPER
		} else if game.Player1Move == PAPER {
			player2Move = SCISSORS
		} else {
			player2Move = ROCK
		}
	}

	newGame := Game{
		Player1Move: game.Player1Move,
		Player2Move: player2Move,
	}

	return calculateScore(newGame)
}

func play(move1 Move, move2 Move) Result {
	if move1 == move2 {
		return DRAW
	}
	if move1 == ROCK {
		if move2 == SCISSORS {
			return PLAYER1_WIN
		} else {
			return PLAYER2_WIN
		}
	}
	if move1 == PAPER {
		if move2 == ROCK {
			return PLAYER1_WIN
		} else {
			return PLAYER2_WIN
		}
	}
	if move1 == SCISSORS {
		if move2 == PAPER {
			return PLAYER1_WIN
		} else {
			return PLAYER2_WIN
		}
	}
	panic("Invalid move")
}

func readLines(filename string) ([]Game, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")
	games := make([]Game, len(lines))
	for i, line := range lines {
		moves := strings.Split(line, " ")
		player1Move := parseMove(moves[0])
		player2Move := parseMove(moves[1])

		games[i] = Game{
			Player1Move: player1Move,
			Player2Move: player2Move,
		}
	}

	return games, nil
}

func part1(games []Game) {
	scores := make([]uint8, len(games))
	sum := 0
	for i, game := range games {
		scores[i] = calculateScore(game)
		sum += int(scores[i])
	}
	fmt.Println("Part 1 - Sum is ", sum)

}

func part2(games []Game) {
	scores := make([]uint8, len(games))
	sum := 0
	for i, game := range games {
		scores[i] = calculateScorePart2(game)
		sum += int(scores[i])
	}
	fmt.Println("Part 2 - Sum is ", sum)

}

func main() {
	data, err := readLines("day2/day2_input.txt")
	if err != nil {
		panic(err)
	}

	part1(data)
	part2(data)
}
