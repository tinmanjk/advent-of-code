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
	// lines := returnSliceOfLinesFromFile()
	var result float64
	result = task01(lines)
	fmt.Println(result)

	// result = task02(numbers)
	// fmt.Println(result)
}

func whoWins(lines map[int]rune, mapIndeces *map[int]int) (result rune) {
	count1, count0 := 0, 0
	for _, v := range lines {
		if v == '1' {
			count1++
		} else {
			count0++
		}
	}
	if count1 >= count0 {
		result = '1'
	} else {
		result = '0'
	}

	for k, v := range lines {
		if v != result {
			delete(*mapIndeces, k)
		}
	}
	return
}

func whoWinsCo2(lines map[int]rune, mapIndeces *map[int]int) (result rune) {
	count1, count0 := 0, 0
	for _, v := range lines {
		if v == '1' {
			count1++
		} else {
			count0++
		}
	}
	if count0 <= count1 {
		result = '0'
	} else {
		result = '1'
	}

	for k, v := range lines {
		if v != result {
			delete(*mapIndeces, k)
		}
	}

	return
}

func task01(lines []string) (result float64) {

	// gamma
	// epsilon
	totalLines := len(lines)
	lineLength := len(lines[0])
	// count 1 only
	var res rune
	// hashMap
	mapIndecesOxygen := make(map[int]int)
	for i := 0; i < totalLines; i++ {
		mapIndecesOxygen[i] = i
	}

	for i := 0; i < lineLength; i++ {
		runeLine := make(map[int]rune)
		for j := 0; j < totalLines; j++ {
			if _, ok := mapIndecesOxygen[j]; ok {
				runeLine[j] = rune(lines[j][i])
			}
		}
		res = whoWins(runeLine, &mapIndecesOxygen)
		fmt.Println(res)
	}

	var oxygen string
	// resulta e na mapIndex.a trqbva da e edin
	for k := range mapIndecesOxygen {
		oxygen = lines[k]
	}
	fmt.Println(oxygen)

	mapIndecesCo2 := make(map[int]int)
	for i := 0; i < totalLines; i++ {
		mapIndecesCo2[i] = i
	}

	for i := 0; i < lineLength; i++ {
		runeLine := make(map[int]rune)
		for j := 0; j < totalLines; j++ {
			if _, ok := mapIndecesCo2[j]; ok {
				runeLine[j] = rune(lines[j][i])
			}
		}
		if len(mapIndecesCo2) > 1 {
			res = whoWinsCo2(runeLine, &mapIndecesCo2)
			fmt.Println(res)
		}

	}

	var co2 string
	// resulta e na mapIndex.a trqbva da e edin
	for k := range mapIndecesCo2 {
		co2 = lines[k]
	}
	fmt.Println(co2)

	var gamma, epsilon float64
	gamma = 0
	epsilon = 0
	for i := len(oxygen) - 1; i >= 0; i-- {

		power := float64(len(oxygen) - i - 1)

		if oxygen[i] == '1' {
			gamma += math.Pow(2, power)

		}

		if co2[i] == '1' {
			epsilon += math.Pow(2, power)

		}
	}

	return gamma * epsilon
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
