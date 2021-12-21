package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	player1Start, player2Start := parseInput(lines)
	var result int
	// part 1
	result = findResult(player1Start, player2Start)
	fmt.Println(result)

	// part 2
	// 	result = findResult(imageEnhancement, inputImage, 50)
	// 	fmt.Println(result)
}

func parseInput(lines []string) (player1Start int, player2Start int) {
	player1Start, _ = strconv.Atoi(strings.Split(lines[0], ": ")[1])
	player2Start, _ = strconv.Atoi(strings.Split(lines[1], ": ")[1])
	return
}

const inputPath = "../input.txt"

func findResult(player1Start int, player2Start int) (result int) {

	dice := 0
	player1Score := 0
	player2Score := 0
	direRolledCounter := 0

	for {
		// current Scores
		turnSum := 0
		for i := 0; i < 3; i++ {
			dice++
			direRolledCounter++
			if dice > 100 {
				dice = 1
			}
			turnSum += dice
		}
		turnSum = turnSum % 10
		player1Start += turnSum
		if player1Start > 10 {
			player1Start = player1Start % 10
		}

		player1Score += player1Start

		if player1Score >= 1000 {
			return player2Score * direRolledCounter
		}

		// player2
		turnSum = 0
		for i := 0; i < 3; i++ {
			dice++
			direRolledCounter++
			if dice > 100 {
				dice = 1
			}
			turnSum += dice
		}
		turnSum = turnSum % 10
		player2Start += turnSum
		if player2Start > 10 {
			player2Start = player2Start % 10
		}
		player2Score += player2Start

		if player2Score >= 1000 {
			return player1Score * direRolledCounter
		}

	}
}
