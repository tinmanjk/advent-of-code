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

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	result := task01(lines)
	fmt.Println(result)
	result = task02(lines)
	fmt.Println(result)
}

func returnSliceOfLinesFromFile(filePath string) (sliceOfLines []string) {
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open(filePath)

	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Panic(err)
	}

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

func task01(lines []string) (countValid int) {

	for _, line := range lines {

		lowerBound, upperBound, char, password := splitLine(line)

		count := 0
		for _, r := range password {
			if r == char {
				count++
			}
		}

		if count >= lowerBound && count <= upperBound {
			countValid++
		}
	}

	return countValid
}

func task02(lines []string) (countValid int) {

	for _, line := range lines {

		firstIndex, secondIndex, char, password := splitLine(line)
		// 0-basing the indeces
		firstIndex--
		secondIndex--

		count := 0
		for i, r := range password {
			if i != firstIndex && i != secondIndex {
				continue
			}
			if r == char {
				count++
			}
		}

		if count == 1 {
			countValid++
		}
	}

	return countValid
}
