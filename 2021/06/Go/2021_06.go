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
	result = findResult(fishAges, 256)
	fmt.Println(result)

	// result = findResult(fishAges, true)
	// fmt.Println(result)
}

// structs
type something struct {
}

// 0 to 6 timer
// at -1 it actually new fish -> + reset of timer
// new fish with timer of 8
// does not start counting down until the next day new...??!
// how many after 80 days
func findResult(fishNext []int, numberDays int) (result int64) {
	// passage of time
	// hashmap
	mapNumbers := make(map[int]int, 0)
	mapNumbers[-1] = 0
	mapNumbers[0] = 0
	mapNumbers[1] = 0
	mapNumbers[2] = 0
	mapNumbers[3] = 0
	mapNumbers[4] = 0
	mapNumbers[5] = 0
	mapNumbers[6] = 0
	mapNumbers[7] = 0
	mapNumbers[8] = 0

	// initial
	for i := 0; i < len(fishNext); i++ {
		mapNumbers[fishNext[i]]++
	}

	for i := 0; i < numberDays; i++ {
		// update map
		for j := 0; j <= 8; j++ {
			mapNumbers[j-1] = mapNumbers[j]
		}
		// vsichkite -1 sa roditeli
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
