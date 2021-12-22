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

	var result uint64
	// part 1
	result = uint64(part1(player1Start, player2Start))
	fmt.Println(result)
	// part 2
	result = part2(player1Start, player2Start)
	fmt.Println(result)
}

func parseInput(lines []string) (player1Start int, player2Start int) {
	player1Start, _ = strconv.Atoi(strings.Split(lines[0], ": ")[1])
	player2Start, _ = strconv.Atoi(strings.Split(lines[1], ": ")[1])
	return
}

const inputPath = "../input.txt"

// TODO Clean up / Refactor
func part1(player1Start int, player2Start int) (result int) {

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

func part2(player1Start int, player2Start int) (result uint64) {

	turn0 := HalfTurn{}
	initialFormation := Formation{
		finished:   false,
		p1Score:    0,
		p1Position: player1Start,
		p2Score:    0,
		p2Position: player2Start,
	}
	turn0.formationUniverses = map[Formation]uint64{}
	turn0.formationUniverses[initialFormation] = 1

	finishingScore := 21

	var countp1, countp2 uint64
	countp1 = 0
	countp2 = 0
	previousHalfTurn := &turn0
	player1Turn := true
	stopPlaying := false
	for !stopPlaying {
		// halfTurns - player1 then player 2 then player 1
		currentHalfTurn := playHalfTurn(previousHalfTurn, finishingScore, player1Turn)
		// nothing to check if no formation universes left
		if len(currentHalfTurn.formationUniverses) == 0 {
			break
		}

		for formation, universes := range currentHalfTurn.formationUniverses {
			if formation.finished {
				if player1Turn {
					countp1 += universes
				} else {
					countp2 += universes
				}
			} else {
				stopPlaying = false
			}
		}

		player1Turn = !player1Turn // toggle whose turn it is
		previousHalfTurn = currentHalfTurn
	}

	if countp1 > countp2 {
		result = countp1
	} else {
		result = countp2
	}
	return
}

type HalfTurn struct {
	previousHalfTurn   *HalfTurn
	formationUniverses map[Formation]uint64
}

type Formation struct {
	finished   bool
	p1Score    int
	p1Position int
	p2Score    int
	p2Position int
}

func playHalfTurn(previousTurn *HalfTurn, winningGoal int, player1Turn bool) (halfTurnCompleted *HalfTurn) {

	currentHalfTurn := HalfTurn{}
	currentHalfTurn.previousHalfTurn = previousTurn
	diceSumCount := generateMapDiceSumCount() // 3 throws produce a sum from 3 to 9 with different frequency
	currentHalfTurn.formationUniverses = map[Formation]uint64{}
	for formation, universes := range previousTurn.formationUniverses {
		if formation.finished { // not available
			continue
		}
		for diceSum, countUniverse := range diceSumCount {
			newFormation := Formation{}
			newFormation.finished = false

			if player1Turn {
				newFormation.p2Position = formation.p2Position
				newFormation.p2Score = formation.p2Score

				newPosition := formation.p1Position + diceSum
				if newPosition > 10 {
					newPosition = newPosition % 10
				}

				newFormation.p1Position = newPosition
				newFormation.p1Score = formation.p1Score + newPosition
				if newFormation.p1Score >= winningGoal {
					newFormation.finished = true
				}
			} else {
				newFormation.p1Position = formation.p1Position
				newFormation.p1Score = formation.p1Score

				newPosition := formation.p2Position + diceSum
				if newPosition > 10 {
					newPosition = newPosition % 10
				}

				newFormation.p2Position = newPosition
				newFormation.p2Score = formation.p2Score + newPosition
				if newFormation.p2Score >= winningGoal {
					newFormation.finished = true
				}
			}

			newUniverses := universes * uint64(countUniverse)
			currentHalfTurn.formationUniverses[newFormation] += newUniverses
		}
	}

	halfTurnCompleted = &currentHalfTurn
	return
}

func generateMapDiceSumCount() (diceSumCount map[int]int) {
	diceSumCount = map[int]int{}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				diceSumCount[i+j+k]++
			}
		}
	}
	return
}
