package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const inputPath = "../input.txt"

func main() {
	numbers := returnSliceOfIntsFromFile(inputPath)
	var result int
	result = task01(numbers)
	fmt.Println(result)

	result = task02(numbers)
	fmt.Println(result)
}

func task01(numbers []int) (result int) {

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			result++
		}
	}
	return result
}

type numberIndex struct {
	number int
	index  int
}

const finalSum int = 2020

// three numbers should sum to 2020 -> multiply
func task02(numbers []int) (result int) {
	twoNumbersSum := make(map[int][]numberIndex)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] < finalSum {
				sum := numbers[i] + numbers[j]
				// override strategy here
				twoNumbersSum[sum] = make([]numberIndex, 2)
				twoNumbersSum[sum][0] = numberIndex{number: numbers[i], index: i}
				twoNumbersSum[sum][1] = numberIndex{number: numbers[j], index: j}
			}
		}
	}

	for index, number := range numbers {
		if twoNumbers, ok := twoNumbersSum[finalSum-number]; ok &&
			index != twoNumbers[0].index && index != twoNumbers[1].index {
			result = number * twoNumbers[0].number * twoNumbers[1].number
			break
		}
	}
	return result
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

func timeMe(toBeTimed func(), repeatTimes int) (duration time.Duration) {
	startTime := time.Now()
	for i := 0; i < repeatTimes; i++ {
		toBeTimed()
	}

	return time.Since(startTime)
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
