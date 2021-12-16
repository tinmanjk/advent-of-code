package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)

	var result int
	// task 01
	randomNumbersStrings, listOfMatrices := parseInput(lines)
	result = determineResultOfWin(randomNumbersStrings, listOfMatrices, false)

	fmt.Println(result)
	// task 02
	// TODO fix need to re-parse input
	randomNumbersStrings, listOfMatrices = parseInput(lines)
	result = determineResultOfWin(randomNumbersStrings, listOfMatrices, true)
	fmt.Println(result)
}

func parseInput(lines []string) (randomNumbersStrings []string, listOfMatrices [][][]string) {
	randomNumbersStrings = strings.Split(lines[0], ",")

	numberOfMatrices := (len(lines) - 1) / 6

	// matrices
	listOfMatrices = make([][][]string, numberOfMatrices)
	indexMatrix := -1
	indexLine := 0
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			indexMatrix++
			listOfMatrices[indexMatrix] = make([][]string, 5)
			indexLine = 0
			continue
		}
		listOfMatrices[indexMatrix][indexLine] = strings.Split(lines[i], " ")
		// special case for single digit numbers having two spaces - empty space leftover
		if len(listOfMatrices[indexMatrix][indexLine]) != 5 {
			for i := 0; i < len(listOfMatrices[indexMatrix][indexLine]); i++ {
				if listOfMatrices[indexMatrix][indexLine][i] == "" {
					listOfMatrices[indexMatrix][indexLine] = removeIndex(listOfMatrices[indexMatrix][indexLine], i)
					i--
				}

			}
		}
		indexLine++
	}

	return
}

func determineResultOfWin(randomNumbers []string, listOfMatrices [][][]string,
	winLast bool) (result int) {

	numberMatrices := 0
	alreadyWon := make(map[int]int, len(listOfMatrices))
	for i := 0; i < len(randomNumbers); i++ {
		randomNumber := randomNumbers[i]
		// traverse the matrices
		for j := 0; j < len(listOfMatrices); j++ {
			if _, ok := alreadyWon[j]; ok {
				continue
			}
			// inside a matrix by row
			for k := 0; k < 5; k++ {
				// set empty as easiest
				// inside a row by column
				for l := 0; l < 5; l++ {
					if listOfMatrices[j][k][l] == randomNumber {
						listOfMatrices[j][k][l] = ""
						if checkBingo(listOfMatrices[j]) {
							result = calculateScore(randomNumber, listOfMatrices[j])
							if !winLast {
								return
							}
							alreadyWon[j] = j
							numberMatrices++
							if numberMatrices == len(listOfMatrices) {
								return
							}
						}
					}
				}

			}
		}
	}
	return
}

func checkBingo(matrix [][]string) bool {
	for i := 0; i < len(matrix); i++ {
		var rowConcat, colConcat string
		for j := 0; j < len(matrix); j++ {
			rowConcat += matrix[i][j]
			colConcat += matrix[j][i]
		}
		if rowConcat == "" {
			return true
		}
		if colConcat == "" {
			return true
		}
	}
	return false
}

func calculateScore(randomNumber string, matrix [][]string) (score int) {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] != "" {
				converted, _ := strconv.Atoi(matrix[i][j])
				score += converted
			}
		}
	}
	convertedRandomNumber, _ := strconv.Atoi(randomNumber)
	score *= convertedRandomNumber
	return
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
