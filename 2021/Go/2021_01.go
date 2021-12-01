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
