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
	parsedInput := parseInput(lines)

	// part 1
	result = findResult(parsedInput, 100, false)
	fmt.Println(result)

	// part 2
	result = findResult(parsedInput, 100, true)
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

const inputPath = "../input.txt"
const paddedValue = -1

const noFlashMarker = 0
const toFlashMarker = 1
const finishedFlashingMarker = 2

const flashThreshold = 9

func findResult(inputData [][]int, numberSteps int, partTwo bool) (result int) {

	paddedData := addPaddings(inputData, paddedValue)
	inputDataRowCount := len(paddedData)
	// assuming equal number of columns
	inputDataColumnCount := len(paddedData[0])
	flashState := makeFlashStateMatrix(inputDataRowCount, inputDataColumnCount)

	waitingForSimultaneousFlash := partTwo
	for step := 1; (step <= numberSteps) || waitingForSimultaneousFlash; step++ {
		// 1. Step mark
		energizeAndMark := func(i int, j int) (breaks bool, breakReturns bool, returnValue interface{}) {
			energizeAndMark(paddedData, flashState, i, j)
			return
		}
		iteratePaddedMatrix(paddedData, energizeAndMark)

		// 2. NeighbourMarks iteratively
		flashAndMarkNeighbors := func(i int, j int) (breaks bool, breakReturns bool, returnValue interface{}) {
			if (flashState)[i][j] != toFlashMarker {
				return // effectively continue
			}
			energizeAndMarkNeighbors(i, j, paddedData, flashState)
			flashState[i][j] = finishedFlashingMarker
			return
		}

		for !allFinishedFlashing(flashState) {
			iteratePaddedMatrix(paddedData, flashAndMarkNeighbors)
		}

		// 3. Set Flashed Ones to 0
		finalSweep := func(i int, j int) (breaks bool, breakReturns bool, returnVal interface{}) {
			if flashState[i][j] == finishedFlashingMarker {
				paddedData[i][j] = 0
				flashState[i][j] = noFlashMarker
				result++
			}
			return
		}
		iteratePaddedMatrix(paddedData, finalSweep)

		if !partTwo {
			continue
		}

		foundNonFlashIterator := func(i, j int) (breaks bool, breakReturns bool, returnValue interface{}) {
			if paddedData[i][j] > 0 {
				return true, true, true // breakreturn
			}
			return // continue
		}
		iterationBreakReturn, foundNonFlash := iteratePaddedMatrix(paddedData, foundNonFlashIterator)
		if iterationBreakReturn && foundNonFlash.(bool) {
			continue
		}
		return step
	}
	return
}

func makeFlashStateMatrix(totalRows int, totalCols int) (flashMatrix [][]int) {

	flashMatrix = make([][]int, totalRows)
	for i := 0; i < totalRows; i++ {
		flashMatrix[i] = make([]int, totalCols)
		for j := 0; j < totalCols; j++ {
			flashMatrix[i][j] = noFlashMarker
		}
	}

	return
}

// return defaults = continue
type iteratorSignature = func(i int, j int) (breaks bool, breakReturns bool, returnValue interface{})

func iteratePaddedMatrix(paddedMatrix [][]int, iterator iteratorSignature) (breakReturns bool, retValue interface{}) {

	rowNumber := len(paddedMatrix)
	columnNumber := len(paddedMatrix[0])

outerLoop:
	for j := 1; j < rowNumber-1; j++ {
		for k := 1; k < columnNumber-1; k++ {
			breaks, breakReturns, retVal := iterator(j, k)
			// if blank return continue
			if breakReturns {
				return breakReturns, retVal
			}
			if breaks {
				break outerLoop
			}
		}
	}
	return // false, nil
}

func energizeAndMark(inputData [][]int, flashMatrix [][]int, i int, j int) {
	if flashMatrix[i][j] == noFlashMarker {
		inputData[i][j]++
		if inputData[i][j] > flashThreshold {
			flashMatrix[i][j] = toFlashMarker
		}
	}
}

func energizeAndMarkNeighbors(i int, j int, inputData [][]int, flashMatrix [][]int) {
	// neighbor offsets
	for k := -1; k <= 1; k++ {
		for m := -1; m <= 1; m++ {
			if k == 0 && m == 0 {
				continue
			}
			energizeAndMark(inputData, flashMatrix, i+k, j+m)
		}
	}
}

func allFinishedFlashing(flashMatrix [][]int) (result bool) {

	result = true
	breakReturns, anyToFlash := iteratePaddedMatrix(flashMatrix,
		func(i, j int) (breaks bool, breakReturns bool, atLeastOneToFlash interface{}) {
			if flashMatrix[i][j] == toFlashMarker {
				return true, true, true
			}
			return //false, false, nil
		})
	if breakReturns {
		result = !anyToFlash.(bool)
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
