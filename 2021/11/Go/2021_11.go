package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int

	// part 1
	result = findResult(lines, true)
	fmt.Println(result)

	// // part 2
	result = findResult(lines, false)
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

func findResult(inputData []string, partOne bool) (result int) {
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
