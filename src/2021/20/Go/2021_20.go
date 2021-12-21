package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	var result int
	// part 1
	result = findResult(lines)
	fmt.Println(result)

	// // part 2
	// result = findResult(lines)
	// fmt.Println(result)
}

func parseInput(lines []string) {

	for i := 0; i < len(lines); i++ {
	}
	return
}

const inputPath = "../input0.txt"

func findResult(lines []string) (result int) {

	return
}
