package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int

	var parsedInput = parseInput(lines)
	paddedInput := addPaddings(parsedInput, -1000)
	// part 1
	result = findResult(paddedInput, false)
	fmt.Println(result)

	// // part 2
	parsedInput = parseInput(lines)
	paddedInput = addPaddings(parsedInput, -1000)
	result = findResult(paddedInput, true)
	fmt.Println(result)
}

func parseInput(slicesOfLines []string) (solutionsData [][]int) {
	lengthTotalLines := len(slicesOfLines)
	lenghSingleLine := len(slicesOfLines[0]) // should be the same for all
	solutionsData = make([][]int, lengthTotalLines)
	for i := 0; i < lengthTotalLines; i++ {
		solutionsData[i] = make([]int, lenghSingleLine)
		for j := 0; j < lenghSingleLine; j++ {
			solutionsData[i][j] = int(slicesOfLines[i][j] - '0')
		}
	}

	return
}

func addPaddings(solutionsData [][]int, paddedValue int) (paddedSolutionsData [][]int) {

	paddedSolutionsData = make([][]int, len(solutionsData))
	for i := 0; i < len(solutionsData); i++ {
		paddedSolutionsData[i] = append([]int{paddedValue}, solutionsData[i]...)
		paddedSolutionsData[i] = append(paddedSolutionsData[i], paddedValue)
	}

	lenghtPaddedSingleLine := len(solutionsData[0]) + 2 // should be the same for all
	paddedBeginRow := make([]int, lenghtPaddedSingleLine)
	for i := 0; i < len(paddedBeginRow); i++ {
		paddedBeginRow[i] = paddedValue
	}
	paddedEndRow := make([]int, lenghtPaddedSingleLine)
	copy(paddedEndRow, paddedBeginRow)

	paddedMatrix := make([][]int, 0)
	paddedMatrix = append(paddedMatrix, paddedBeginRow)

	paddedSolutionsData = append(paddedMatrix, paddedSolutionsData...)
	paddedSolutionsData = append(paddedSolutionsData, paddedEndRow)

	return
}

// energy levels of octopuses
// flashes of light
// single step
// 1. ++ energy level
// 2. if > 9 flashes
// 3. all adjacent ++ -> diagonally adjacent
// 3a -> back to 2 if adjacent rises to >9 - flashes - ad infinitum
// flash -> set to 0...and remaining there I guess

func produceFlashInitialMatrix(totalRows int, totalCols int) (flashMatrix [][]int) {

	alreadyFlashedInitial := 1
	flashMatrix = make([][]int, totalRows)
	for i := 0; i < totalRows; i++ {
		flashMatrix[i] = make([]int, totalCols)
		for j := 0; j < totalCols; j++ {
			flashMatrix[i][j] = alreadyFlashedInitial
		}
	}

	return
}

func pulse(inputData [][]int, alreadyFlashed [][]int) (numbeFlashes int) {

	for j := 1; j < len(alreadyFlashed)-1; j++ {
		for k := 1; k < len(alreadyFlashed[j])-1; k++ {
			if alreadyFlashed[j][k] == 0 {
				// all positions
				if inputData[j-1][k-1] != 10 {
					inputData[j-1][k-1]++
					if inputData[j-1][k-1] > 9 {
						numbeFlashes++
						alreadyFlashed[j-1][k-1] = 0
					}
				}
				if inputData[j-1][k] != 10 {
					inputData[j-1][k]++
					if inputData[j-1][k] > 9 {
						numbeFlashes++
						alreadyFlashed[j-1][k] = 0
					}
				}
				if inputData[j-1][k+1] != 10 {
					inputData[j-1][k+1]++
					if inputData[j-1][k+1] > 9 {
						numbeFlashes++
						alreadyFlashed[j-1][k+1] = 0
					}
				}
				if inputData[j][k-1] != 10 {
					inputData[j][k-1]++
					if inputData[j][k-1] > 9 {
						numbeFlashes++
						alreadyFlashed[j][k-1] = 0
					}
				}
				if inputData[j][k+1] != 10 {
					inputData[j][k+1]++
					if inputData[j][k+1] > 9 {
						numbeFlashes++
						alreadyFlashed[j][k+1] = 0
					}
				}
				if inputData[j+1][k-1] != 10 {
					inputData[j+1][k-1]++
					if inputData[j+1][k-1] > 9 {
						numbeFlashes++
						alreadyFlashed[j+1][k-1] = 0
					}
				}
				if inputData[j+1][k] != 10 {
					inputData[j+1][k]++
					if inputData[j+1][k] > 9 {
						numbeFlashes++
						alreadyFlashed[j+1][k] = 0
					}
				}
				if inputData[j+1][k+1] != 10 {
					inputData[j+1][k+1]++
					if inputData[j+1][k+1] > 9 {
						numbeFlashes++
						alreadyFlashed[j+1][k+1] = 0
					}
				}

				alreadyFlashed[j][k] = -1
			}
		}
	}
	return
}

func IsFinished(alreadyFlashed [][]int) bool {
	for j := 1; j < len(alreadyFlashed)-1; j++ {
		for k := 1; k < len(alreadyFlashed[j])-1; k++ {
			if alreadyFlashed[j][k] == 0 {
				return false
			}
		}
	}
	return true
}

const inputPath = "../input.txt"

func findResult(inputData [][]int, partTwo bool) (result int) {

	// alreadyFlashed
	// full of 1s
	alreadyFlashed := produceFlashInitialMatrix(len(inputData), len(inputData))
	waitingForAllFlash := partTwo
	// total flashes after 100 steps
	for i := 0; (i < 100) || waitingForAllFlash; i++ {
		if partTwo {
			sim := 0
			for j := 1; j < len(inputData)-1; j++ {
				for k := 1; k < len(inputData[j])-1; k++ {
					sim += inputData[j][k]
				}
			}
			if sim == 0 {
				return i
			}
		}

		// 1. increase everybody and first Flash 2.
		for j := 1; j < len(inputData)-1; j++ {
			for k := 1; k < len(inputData[j])-1; k++ {
				inputData[j][k]++
				if inputData[j][k] > 9 {
					result++
					// FLASH record
					// assume that it stays there!!!!!
					alreadyFlashed[j][k] = 0 // good becauase that's
					// inputData[j][k] = 0
				}
			}
		}

		for !IsFinished(alreadyFlashed) {
			result += pulse(inputData, alreadyFlashed)
		}
		// 3. all adjacent ++
		for j := 1; j < len(inputData)-1; j++ {
			for k := 1; k < len(inputData[j])-1; k++ {
				if alreadyFlashed[j][k] == -1 {
					inputData[j][k] = 0
				}
			}
		}
		alreadyFlashed = produceFlashInitialMatrix(len(inputData), len(inputData))
	}
	return
}

func returnSliceOfLinesFromFile(filePath string) (sliceOfLines []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	rawBytes, err := io.ReadAll(file)
	if err != nil {
		log.Panic(err)
	}

	lines := strings.Split(string(rawBytes), "\n")

	return lines
}
