package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	crabHorizontalPosition := parseInput(lines)
	var result int

	result = findResult(crabHorizontalPosition, false)
	fmt.Println(result)

	result = findResult(crabHorizontalPosition, true)
	fmt.Println(result)

	// result = findResult(fishTimeToNew, 80)
	// fmt.Println(result)
}

func findAverage(sliceOfUints []int) (result int) {
	var sum int
	for _, v := range sliceOfUints {
		sum += v
	}

	return sum / len(sliceOfUints)
}

func findMinAndMax(sliceOfInts []int) (min int, max int) {
	min = math.MaxInt32
	max = math.MinInt32

	for _, v := range sliceOfInts {
		if min >= v {
			min = v
		}
		if max <= v {
			max = v
		}
	}

	return
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

// bruteforcing
func findResult(numbers []int, variableRate bool) (result int) {
	min, max := findMinAndMax(numbers)
	results := make(map[int]int)
	for i := min; i <= max; i++ {
		for j := 0; j < len(numbers); j++ {
			distance := int(math.Abs(float64(numbers[j] - i)))
			cost := 0
			if variableRate {
				additional := 1
				for k := 0; k < distance; k++ {
					cost += additional
					additional++
				}
			} else {
				cost = distance
			}
			results[i] += cost
		}
	}

	result = math.MaxInt32
	for _, v := range results {
		if result >= v {
			result = v
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
