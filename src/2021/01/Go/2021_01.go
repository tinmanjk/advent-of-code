package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
)

const inputPath = "../input.txt"

func main() {
	numbers := inputParse.ReturnSliceOfIntsFromFile(inputPath)
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
