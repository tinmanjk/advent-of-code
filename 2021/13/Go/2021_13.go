package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int
	mapOfPoints, instructions := parseInput(lines)

	// part 1
	result = findResult(mapOfPoints, instructions, true)
	fmt.Println(result)

	// // part 2
	// result = findResult(&graph, true)
	// fmt.Println(result)
}

type point struct {
	x int
	y int
}

type inst struct {
	axis  string
	units int
}

// list of points

func parseInput(slicesOfLines []string) (mapOfPoints map[point]point,
	instructions []inst) {

	// points
	mapOfPoints = map[point]point{}

	i := 0
	for ; i < len(slicesOfLines); i++ {
		line := slicesOfLines[i]
		if line == "" {
			i++
			break
		}
		splitted := strings.Split(line, ",")
		x, _ := strconv.Atoi(splitted[0])
		y, _ := strconv.Atoi(splitted[1])
		p := point{x, y}
		mapOfPoints[p] = p
	}
	// we have i

	//instructions
	for j := i; j < len(slicesOfLines); j++ {
		line := slicesOfLines[j]
		if line == "" {
			break
		}
		splitted := strings.Split(line, " ")
		instruction := strings.Split(splitted[2], "=")

		axis := instruction[0]
		units, _ := strconv.Atoi(instruction[1])
		instr := inst{axis, units}
		instructions = append(instructions, instr)
	}

	return
}

const inputPath = "../input0.txt"

func findResult(mapOfPoints map[point]point,
	instructions []inst, justFirstInstr bool) (result int) {

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
