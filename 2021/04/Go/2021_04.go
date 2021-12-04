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

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	randomNumbersStrings := strings.Split(lines[0], ",")

	randomNumbers := make([]int, len(randomNumbersStrings))
	for i := 0; i < len(randomNumbersStrings); i++ {
		randomNumbers[i], _ = strconv.Atoi(randomNumbersStrings[i])
	}

	// number of matrices
	numberOfMatrices := ((len(lines) - 1) / 6)
	// numberOfMatrices := 99

	// matrices
	listOfMatrices := make([][][]string, numberOfMatrices)
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
		if len(listOfMatrices[indexMatrix][indexLine]) != 5 {
			for i := 0; i < len(listOfMatrices[indexMatrix][indexLine]); i++ {
				if listOfMatrices[indexMatrix][indexLine][i] == "" {
					listOfMatrices[indexMatrix][indexLine] = RemoveIndex(listOfMatrices[indexMatrix][indexLine], i)
					i--
				}

			}
		}
		indexLine++
	}

	var result int
	// result = task01(randomNumbersStrings, listOfMatrices)
	fmt.Println(result)

	result = task02(randomNumbersStrings, listOfMatrices)

	fmt.Println(result)
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
	// row
	// concatenate
	for i := 0; i < 5; i++ {
		rowConcat := matrix[i][0] + matrix[i][1] + matrix[i][2] + matrix[i][3] + matrix[i][4]
		if rowConcat == "" {
			return true
		}
		colConcat := matrix[0][i] + matrix[1][i] + matrix[2][i] + matrix[3][i] + matrix[4][i]
		if colConcat == "" {
			return true
		}
	}
	return false
}

func calculateScore(randomNumber string, matrix [][]string) (score int) {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
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

// list of matrices

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
func convertFromBinary(binaryString string) (result int) {
	for i := 0; i < len(binaryString); i++ {
		if !(binaryString[i] == '0' || binaryString[i] == '1') {
			log.Panic("Only accept 0 and 1 in string")
		}
		result <<= 1 // make place to the left
		if binaryString[i] == '1' {
			result |= 1
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

func splitLine(line string) (firstNumber int, secondNumber int,
	char rune, password string) {

	// Example line: 5-6 v: hvvgvrm
	lineSplit := strings.Split(line, " ") // should be 3
	numbers := strings.Split(lineSplit[0], "-")
	// TODO: Better Error handling
	firstNumber, err := strconv.Atoi(numbers[0])
	if err != nil {
		log.Panic(err)
	}
	secondNumber, err = strconv.Atoi(numbers[1])
	if err != nil {
		log.Panic(err)
	}

	// use if there are multi-byte unicode chars
	for _, r := range lineSplit[1] {
		char = r
		break
	}

	password = lineSplit[2]
	return
}

func returnSliceOfIntsFromFile(filePath string) (sliceOfLines []int) {
	slicesOfLines := returnSliceOfLinesFromFile(filePath)

	lines := make([]int, 0, len(slicesOfLines))
	for i := 0; i < len(slicesOfLines); i++ {
		trimmed := strings.TrimRight(slicesOfLines[i], "\n ")
		if trimmed == "" {
			continue
		}
		// TODO better Error handling Atoi
		number, err := strconv.Atoi(trimmed)
		if err != nil {
			log.Panic(err)
		}
		lines = append(lines, number)
	}

	return lines
}
