package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputPath = "../input.txt"
const blockValue = 9

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	inputData := parseInput(lines)
	var result int

	// part 1
	result = findResult(inputData)
	fmt.Println(result)

	// // part 2
	// result = findResult(inputData)
	// fmt.Println(result)
}

func parseInput(slicesOfLines []string) (solutionsData []int) {

	solutionsData = make([]int, len(slicesOfLines))
	for i := 0; i < len(slicesOfLines); i++ {
		// line := slicesOfLines[i]
		// // 10 string
		// // 4 string
		// inputL := 0
		// solutionsData[i] = inputL

	}

	return
}

func findResult(inputData []int) (result int) {
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
