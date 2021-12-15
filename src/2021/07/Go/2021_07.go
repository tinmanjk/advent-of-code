package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	crabHorizontalPosition := parseInput(lines)

	var result int

	// brute force
	result = findResult(crabHorizontalPosition, false, true, true)
	fmt.Println(result)

	result = findResult(crabHorizontalPosition, true, true, true)
	fmt.Println(result)

	// median and min sum triangular distance
	result = findResult(crabHorizontalPosition, false, false, true)
	fmt.Println(result)

	result = findResult(crabHorizontalPosition, true, false, true)
	fmt.Println(result)

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

func findMedian(sliceOfInts []int) (median1 int, median2 int) {

	// not to produce side-effect for initial slice
	copySlice := make([]int, len(sliceOfInts))
	copy(copySlice, sliceOfInts)
	sort.Ints(copySlice)
	middleIndex := len(copySlice) / 2

	// odd number
	if len(copySlice)&1 == 1 {
		median1 = copySlice[middleIndex]
		median2 = median1
		return
	}

	// even number
	median1 = copySlice[middleIndex]
	median2 = copySlice[middleIndex+1]
	return
}

// n²+n / 2 -> where n = |ai - x|
// differentiating the sum and taking into account modulus
// x = sum/N + sign(x-ai)/2N / max sign(x-ai) -> N
// => x = average +-1/2 or two values to check floor(average) + ceiling(average)
func findMinTriangularDistance(sliceOfInts []int) (floorAverage int, ceilingAverage int) {
	average := findAverage(sliceOfInts)
	ceilingAverage = int(math.Ceil(average))
	floorAverage = int(math.Floor(average))
	return
}

func findAverage(sliceOfUints []int) (result float64) {
	var sum int
	for _, v := range sliceOfUints {
		sum += v
	}
	result = float64(sum) / float64(len(sliceOfUints))
	return
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

func findResult(numbers []int, variableRate bool, bruteforce bool, triangularNumbersAware bool) (result int) {

	var lowerBound, upperBound int
	if bruteforce {
		lowerBound, upperBound = findMinAndMax(numbers)
	} else {
		if variableRate {
			//https://en.wikipedia.org/wiki/Triangular_number
			lowerBound, upperBound = findMinTriangularDistance(numbers)
		} else {
			//https://math.stackexchange.com/questions/113270/the-median-minimizes-the-sum-of-absolute-deviations-the-ell-1-norm
			// ∑i=1N|si−x| -> minimize by ∑sign(si-x) = 0 -> which is median
			lowerBound, upperBound = findMedian(numbers)
			// could be unnecessary to assign to both - for consistency
			// alternatively lowerbound, _ , upperBound = lowerBound
		}
	}

	results := make(map[int]int)
	for i := lowerBound; i <= upperBound; i++ {
		for j := 0; j < len(numbers); j++ {
			distance := int(math.Abs(float64(numbers[j] - i)))
			cost := 0
			if variableRate {
				if triangularNumbersAware {
					// triangular numbers n(n+1)/2 -> distsance = n = |ai-x|
					cost = distance * (distance + 1) / 2
				} else {
					additional := 1
					for k := 0; k < distance; k++ {
						cost += additional
						additional++
					}
				}
			} else {
				cost = distance
			}
			results[i] += cost
		}
	}

	// Find minimum in hashmap
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
