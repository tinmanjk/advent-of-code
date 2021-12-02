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

func task02(numbers []int) (result int) {

	for i := 3; i < len(numbers); i++ {
		if numbers[i] > numbers[i-3] {
			result++
		}
	}
	return result
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
		lines = append(lines, strings.TrimRight(sc.Text(), "\n "))
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
