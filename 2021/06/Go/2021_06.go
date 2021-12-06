package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputPath = "../input0.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)

	lineSegmentSlice := parseInput(lines)
	var result int
	result = findResult(lineSegmentSlice, false)
	fmt.Println(result)

	result = findResult(lineSegmentSlice, true)
	fmt.Println(result)
}

// structs
type something struct {
}

func parseInput(slicesOfLines []string) (sliceOfLineSegments []something) {

	return
}

func findResult(sliceOfLineSegments []something, includeDiagonal bool) (result int) {

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
