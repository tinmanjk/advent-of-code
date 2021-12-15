package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath) // 1 line
	fishTimeToNew := parseInput(lines)
	var result uint64
	var err error
	result, err = findResult(fishTimeToNew, 80)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)

	result, err = findResult(fishTimeToNew, 256)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)

	uint64CausingOverflow := 442
	resultBigInt := findResultBigInt(fishTimeToNew, uint64CausingOverflow)
	fmt.Println(resultBigInt.String())
}

func findResult(fishTimeToNew []int, numberDays int) (result uint64, err error) {

	// use slice as map
	dayBuckets := make([]uint64, 9)

	// Initial state into slice
	for _, v := range fishTimeToNew {
		dayBuckets[v]++
	}

	for i := 0; i < numberDays; i++ {
		newOnes := dayBuckets[0]
		for j := 1; j <= 8; j++ {
			dayBuckets[j-1] = dayBuckets[j]
		}
		dayBuckets[6], err = addUint64(dayBuckets[6], newOnes)
		if err != nil {
			return
		}
		dayBuckets[8] = newOnes
	}

	for _, v := range dayBuckets {
		result, err = addUint64(result, v)
		if err != nil {
			return
		}
	}
	return result, nil
}

func findResultBigInt(fishTimeToNew []int, numberDays int) (result big.Int) {

	// use slice as map
	dayBuckets := make([]big.Int, 9)

	// Initial state into slice
	for _, v := range fishTimeToNew {
		dayBuckets[v].Add(&dayBuckets[v], big.NewInt(1))
	}

	for i := 0; i < numberDays; i++ {
		newOnes := dayBuckets[0]
		for j := 1; j <= 8; j++ {
			dayBuckets[j-1] = dayBuckets[j]
		}
		dayBuckets[6].Add(&dayBuckets[6], &newOnes)
		dayBuckets[8] = newOnes
	}

	for _, v := range dayBuckets {
		result.Add(&result, &v)
	}
	return result
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

func addUint64(nums ...uint64) (result uint64, err error) {
	for _, num := range nums {
		if result > math.MaxUint-num {
			err = errors.New(fmt.Sprintf("uint64 overflow while adding %v together", nums))
			result = 0
			return
		}
		result += num
	}
	return result, nil
}
