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
	lines := returnSliceOfLinesFromFile(inputPath) // 1 line
	fishAges := parseInput(lines)
	var result int64
	result = findResult(fishAges, 80)
	fmt.Println(result)
	result = findResult(fishAges, 256)
	fmt.Println(result)
}

func findResult(fishNext []int, numberDays int) (result int64) {

	// Initialize map with possible left-days
	mapNumbers := make(map[int]int, 0)
	for i := -1; i <= 8; i++ {
		mapNumbers[i] = 0
	}

	// Initial state into map
	for i := 0; i < len(fishNext); i++ {
		mapNumbers[fishNext[i]]++
	}

	for i := 0; i < numberDays; i++ {
		// update map
		for j := 0; j <= 8; j++ {
			mapNumbers[j-1] = mapNumbers[j]
		}
		mapNumbers[6] += mapNumbers[-1]
		mapNumbers[8] = mapNumbers[-1]
	}
	mapNumbers[-1] = 0

	for _, v := range mapNumbers {
		result += int64(v)
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

func parseInput(slicesOfLines []string) (sliceOfInts []int) {
	line := slicesOfLines[0]
	splitted := strings.Split(line, ",")
	sliceOfInts = make([]int, len(splitted))
	for i := 0; i < len(splitted); i++ {
		sliceOfInts[i], _ = strconv.Atoi(splitted[i])
	}
	return
}
