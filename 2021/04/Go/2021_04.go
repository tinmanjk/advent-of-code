package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	randomNumbersStrings, listOfMatrices := parseInput(lines)
	var result int
	result = task01(randomNumbersStrings, listOfMatrices)
	fmt.Println(result)
	randomNumbersStrings, listOfMatrices = parseInput(lines)
	result = task02(randomNumbersStrings, listOfMatrices)

	fmt.Println(result)
}

func parseInput(lines []string) (randomNumbersStrings []string, listOfMatrices [][][]string) {
	randomNumbersStrings = strings.Split(lines[0], ",")

	randomNumbers := make([]int, len(randomNumbersStrings))
	for i := 0; i < len(randomNumbersStrings); i++ {
		randomNumbers[i], _ = strconv.Atoi(randomNumbersStrings[i])
	}

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

func task01(randomNumbers []string, listOfMatrices [][][]string) (result int) {

	for i := 0; i < len(randomNumbers); i++ {
		randomNumber := randomNumbers[i]
		// traverse the matrices
		for j := 0; j < len(listOfMatrices); j++ {
			// inside a matrix by row
			for k := 0; k < 5; k++ {
				// set empty as easiest
				// inside a row by column
				for l := 0; l < 5; l++ {
					if listOfMatrices[j][k][l] == randomNumber {
						listOfMatrices[j][k][l] = ""
						if checkBingo(listOfMatrices[j]) {
							result = calculateScore(randomNumber, listOfMatrices[j])
							return
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

func task02(randomNumbers []string, listOfMatrices [][][]string) (result int) {

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

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
