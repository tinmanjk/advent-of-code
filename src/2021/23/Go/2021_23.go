package main

import (
	"aoc/libs/go/ds"
	"aoc/libs/go/inputParse"
	"fmt"
	"log"
	"math"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	board := parseInput(lines)

	// part 1
	result := findResult(board)
	fmt.Println(result)
	// part 2
	// result := findResult(cuboids)
	// fmt.Println("Part 1:", result)
}

func parseInput(lines []string) (board [5][13]rune) {

	// 3 /13 board
	board = [5][13]rune{}
	for i := 1; i < len(lines)-1; i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			board[i-1][j] = rune(line[j])
		}
	}
	return
}

const inputPath = "../input.txt"

// ### AMPHI
type Amphipod struct {
	row   int
	col   int
	Type  string
	moves int
}

func isInCorrectRoom(amph *Amphipod) bool {
	inRoom := (1 <= amph.row && amph.row <= 2)
	if !inRoom {
		return false
	}

	switch amph.Type {
	case "A":
		return amph.col == 3
	case "B":
		return amph.col == 5
	case "C":
		return amph.col == 7
	case "D":
		return amph.col == 9
	default:
		log.Panic("impossible")
	}

	return false
}

// </## AMPHI

// ### STATE
type State struct {
	amphis      []*Amphipod
	board       [5][13]rune // array to be used as key in hashmap
	parentState *State
	finished    bool
	score       int
}

func (state State) checkStateFinished() bool {

	for col := 3; col <= 9; col += 2 {
		room := fmt.Sprintf("%c", ('A' + (col-3)/2))
		for row := 1; row <= 4; row++ {
			if room != string(state.board[row][col]) {
				return false
			}
		}
	}

	return true
}

func (state State) calculateScore() (finalScore int) {
	for _, a := range state.amphis {
		switch a.Type {
		case "A":
			finalScore += a.moves
		case "B":
			finalScore += 10 * a.moves
		case "C":
			finalScore += 100 * a.moves
		case "D":
			finalScore += 1000 * a.moves
		}
	}

	return
}

// for debugging
func (state State) printState() {

	fmt.Println("STATE:")

	for i := 0; i < len(state.board); i++ {
		tempSlice := state.board[i][:]
		fmt.Println(string(tempSlice))
	}
	fmt.Println("/STATE")
	fmt.Println(" ")
}

func buildRootState(matrix [5][13]rune) (rootState State) {
	rootState = State{}
	rootState.amphis = []*Amphipod{}
	rootState.board = matrix
	for j := 3; j <= 9; j += 2 {
		for i := 1; i <= 4; i++ {
			a := Amphipod{i, j, string(matrix[i][j]), 0}
			rootState.amphis = append(rootState.amphis, &a)
		}
	}

	return rootState
}

// </## STATE

func findResult(fixedBoard [5][13]rune) (minScore int) {

	boardScores := map[[5][13]rune]int{}
	boardScores[fixedBoard] = 0

	minScore = math.MaxInt32
	rootState := buildRootState(fixedBoard)
	// queue := ds.Queue{}
	stack := ds.Stack{}
	stack.Push(&rootState)
	// stack.Push(&rootState)
	for !stack.IsEmpty() {
		stateToProcess, err := stack.Pop()
		if err != nil {
			break
		}
		state := stateToProcess.(*State)
		state.score = state.calculateScore()
		state.finished = state.checkStateFinished()
		if state.finished {
			if state.score < minScore {
				minScore = state.score

			}
			continue
		}

		nextStates := generateAllPossibleNextStates(state, minScore)

		// board - state
		// if the board is new it can be overriden
		statesToPush := map[[5][13]rune]*State{}
		for _, st := range nextStates {
			if alreadyIn, ok := statesToPush[st.board]; ok {
				if st.score < alreadyIn.score {
					statesToPush[st.board] = st
				}
			} else {
				statesToPush[st.board] = st
			}
		}

		for _, state := range statesToPush {
			// they all have scores calculated
			if alreadyInScore, ok := boardScores[state.board]; ok {
				if state.score < alreadyInScore {
					boardScores[state.board] = state.score
					stack.Push(state)
				}

				// if not dont bother working on it
			} else {
				boardScores[state.board] = state.score
				stack.Push(state)
			}

		}
	}

	return
}

func generateAllPossibleNextStates(state *State, maxScore int) (nextStates []*State) {
	// GENERATE ALL POSSIBLE MOVES

	// TODO: Smart moving
	for _, a := range state.amphis {
		switch a.row {
		case 0: // hallway
			// only can move into a room that's it
			// and the room should be EMPTY or have its oWN KIND
			destCol := 0
			switch a.Type {
			case "A":
				destCol = 3
			case "B":
				destCol = 5
			case "C":
				destCol = 7
			case "D":
				destCol = 9
			}

			// check if full
			if state.board[1][destCol] != '.' {
				continue
			}

			aType := rune(a.Type[0])
			canMoveIntoRoom := false
			var destRow int

			switch {
			case state.board[2][destCol] == '.' &&
				state.board[3][destCol] == '.' &&
				state.board[4][destCol] == '.':
				// go to bottom
				canMoveIntoRoom = true
				destRow = 4
			case state.board[2][destCol] == '.' &&
				state.board[3][destCol] == '.' &&
				state.board[4][destCol] == aType:
				// go to third spot
				canMoveIntoRoom = true
				destRow = 3
			case state.board[2][destCol] == '.' &&
				state.board[3][destCol] == aType &&
				state.board[4][destCol] == aType:
				// go to third spot
				canMoveIntoRoom = true
				destRow = 2
			case state.board[2][destCol] == aType &&
				state.board[3][destCol] == aType &&
				state.board[4][destCol] == aType:
				// go to third spot
				canMoveIntoRoom = true
				destRow = 1
			}

			if !canMoveIntoRoom {
				continue
			}

			freePath := isPathFree(a, state.board, destRow, destCol)
			if freePath {
				newState := createNewStateForMove(state, a, destRow, destCol)
				newState.score = newState.calculateScore()
				if newState.score < maxScore {
					nextStates = append([]*State{&newState}, nextStates...)
				}
			}

		case 4:
			if state.board[3][a.col] != '.' { // cant move
				continue
			}
			fallthrough
		case 3:
			if state.board[2][a.col] != '.' { // cant move
				continue
			}
			fallthrough
		case 2: // room bottom spot
			if state.board[1][a.col] != '.' { // cant move
				continue
			}
			fallthrough // same left/right hallway mechanics
		case 1: // room top spot -> go into the hallway left or right

			// need to move
			if isInCorrectRoom(a) {
				switch a.row {
				case 4:
					continue // bottom in its room
				case 3:
					// one beneath is the same
					if string(state.board[4][a.col]) == a.Type {
						continue
					}
				case 2:
					if string(state.board[4][a.col]) == a.Type &&
						string(state.board[3][a.col]) == a.Type {
						continue
					}
				case 1:
					if string(state.board[4][a.col]) == a.Type &&
						string(state.board[3][a.col]) == a.Type &&
						string(state.board[2][a.col]) == a.Type {
						continue
					}
				}
			}

			// Go into the HALLWAY
			offSet := 1
			// all left
			for state.board[0][a.col-offSet] == '.' {
				destRow := 0
				destCol := a.col - offSet
				newState := createNewStateForMove(state, a, destRow, destCol)
				// try prepend
				newState.score = newState.calculateScore()
				if newState.score < maxScore {
					nextStates = append([]*State{&newState}, nextStates...)

				}

				offSet++
				// jump over room gate
				switch a.col - offSet {
				case 3, 5, 7, 9:
					offSet++
				}
			}
			// all right
			offSet = 1
			for state.board[0][a.col+offSet] == '.' {
				destRow := 0
				destCol := a.col + offSet
				newState := createNewStateForMove(state, a, destRow, destCol)
				newState.score = newState.calculateScore()
				if newState.score < maxScore {
					nextStates = append([]*State{&newState}, nextStates...)

				}

				offSet++
				// jump over room gate
				switch a.col + offSet {
				case 3, 5, 7, 9:
					offSet++
				}
			}
		}

	}

	return
}

func isPathFree(a *Amphipod, board [5][13]rune, destRow, destCol int) bool {
	if a.col < destCol { // go right
		for col := a.col + 1; col < destCol; col++ { // < because guaranteed free spot in front
			if board[a.row][col] != '.' {
				return false
			}
		}
	} else { // go left
		for col := a.col - 1; col > destCol; col-- {
			if board[a.row][col] != '.' {
				return false
			}
		}
	}

	return true
}

func calculateMoves(a Amphipod, destRow int, destCol int) (moves int) {
	verticalMoves := math.Abs(float64(a.row - destRow))
	horizontalMoves := math.Abs(float64(a.col - destCol))
	return int(verticalMoves + horizontalMoves)
}

func createNewStateForMove(oldState *State, aToMove *Amphipod,
	destRow int, destCol int) (newState State) {

	newAmphis := []*Amphipod{}
	var derefA Amphipod
	for _, amphi := range oldState.amphis {
		if amphi == aToMove {
			derefA = *aToMove // value
			// calculate score here
			derefA.moves += calculateMoves(derefA, destRow, destCol)
			newAmphis = append(newAmphis, &derefA)
			continue
		}
		newAmphi := *amphi
		newAmphis = append(newAmphis, &newAmphi)
	}

	var newBoard [5][13]rune
	for i := range oldState.board {
		for j := 0; j < 13; j++ {
			newBoard[i][j] = oldState.board[i][j]
		}
	}

	derefA.row = destRow
	derefA.col = destCol
	newBoard[destRow][destCol] = rune(derefA.Type[0])
	newBoard[aToMove.row][aToMove.col] = '.'

	newState = State{amphis: newAmphis, board: newBoard}
	newState.parentState = oldState

	return
}
