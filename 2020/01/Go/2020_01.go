package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"
const finalSum = 2020

func main() {
	numbers := returnSliceOfIntsFromFile(inputPath)
	result := task01(numbers)
	fmt.Println(result)

	result = task02(numbers)
	fmt.Println(result)
}

func returnSliceOfIntsFromFile(filePath string) (sliceOfLines []int) {
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open(filePath)

	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]int, 0)
	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		// TODO better Error handling Atoi
		number, err := strconv.Atoi(strings.TrimRight(sc.Text(), "\n "))
		if err != nil {
			log.Panic(err)
		}

		lines = append(lines, number)
	}

	if err := sc.Err(); err != nil {
		log.Panic(err)
	}

	return lines
}

func task01(numbers []int) (result int) {
	sumPairNumbers := make(map[int]int)

	for _, number := range numbers {
		sumPairNumber := finalSum - number
		// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
		if val, ok := sumPairNumbers[sumPairNumber]; ok {
			result = val * (finalSum - val)
			break
		}
		sumPairNumbers[number] = finalSum - number
	}
	return result
}

type numberIndex struct {
	number int
	index  int
}

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
