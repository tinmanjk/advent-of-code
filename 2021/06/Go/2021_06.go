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
	fishTimeToNew := parseInput(lines)
	var result uint64
	result = findResult(fishTimeToNew, 80)
	fmt.Println(result)
	result = findResult(fishTimeToNew, 256)
	fmt.Println(result)
}

func findResult(fishTimeToNew []int, numberDays int) (result uint64) {

	// use slice as map
	dayBuckets := make([]uint64, 9)

	// Initial state into slice
	for i := 0; i < len(fishTimeToNew); i++ {
		dayBuckets[fishTimeToNew[i]]++
	}

	for i := 0; i < numberDays; i++ {
		newOnes := dayBuckets[0]
		for j := 1; j <= 8; j++ {
			dayBuckets[j-1] = dayBuckets[j]
		}
		dayBuckets[6] += newOnes
		dayBuckets[8] = newOnes
	}

	for _, v := range dayBuckets {
		result += v
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
