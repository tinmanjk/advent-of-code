package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputPath = "../input0.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int

	result = task01(lines)
	fmt.Println(result)

	result = task02(lines)
	fmt.Println(result)
}

func task01(lines []string) (result int) {
	return
}

func task02(lines []string) (result int) {
	return
}

func convertFromBinary(binaryString string) (result int) {
	for i := 0; i < len(binaryString); i++ {
		if !(binaryString[i] == '0' || binaryString[i] == '1') {
			log.Panic("Only accept 0 and 1 in string")
		}
		result <<= 1 // make place to the left
		if binaryString[i] == '1' {
			result |= 1
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
