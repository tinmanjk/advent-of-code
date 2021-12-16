package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	var result int
	result = task01(lines)
	fmt.Println(result)

	result = task02(lines)
	fmt.Println(result)
}

func task01(lines []string) (result int) {

	var horizontal int
	var depth int
	for _, line := range lines {

		direction, value := splitLine(line)

		switch direction {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	return horizontal * depth
}

func task02(lines []string) (result int) {

	var horizontal int
	var depth int
	var aim int

	for _, line := range lines {

		direction, value := splitLine(line)

		switch direction {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}

	}

	return horizontal * depth
}

func splitLine(line string) (direction string, value int) {

	lineSplit := strings.Split(line, " ")
	direction = lineSplit[0]
	value, _ = strconv.Atoi(lineSplit[1])
	return
}
