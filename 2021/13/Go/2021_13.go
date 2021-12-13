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

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int
	mapOfPoints, instructions := parseInput(lines)

	// part 1
	result = findResult(mapOfPoints, instructions, false)
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

const inputPath = "../input.txt"

func findResult(mapOfPoints map[point]point,
	instructions []inst, justFirstInstr bool) (result int) {

	for i := 0; i < len(instructions); i++ {
		if instructions[i].axis == "y" {
			mapOfPoints = foldUp(mapOfPoints, instructions[i].units)

		} else {
			mapOfPoints = foldLeft(mapOfPoints, instructions[i].units)
		}

		if justFirstInstr {
			break
		}
	}
	printReadyPoints(mapOfPoints)
	return len(mapOfPoints)
}

func printReadyPoints(mapOfPoints map[point]point) {
	// find maxX
	// find maxY

	maxX := math.MinInt32
	maxY := math.MinInt32

	for _, p := range mapOfPoints {
		if p.x >= maxX {
			maxX = p.x
		}
		if p.y >= maxY {
			maxY = p.x
		}
	}

	// y-x mi trqbva matrix
	//if point '#' otherwise '.'
	printMatrix := make([][]rune, maxY+1)
	for i := 0; i < len(printMatrix); i++ {
		printMatrix[i] = make([]rune, maxX+1)
		for j := 0; j < len(printMatrix[i]); j++ {
			pointAtCoordinate := point{j, i}
			if _, ok := mapOfPoints[pointAtCoordinate]; ok {
				printMatrix[i][j] = '#'
			} else {
				printMatrix[i][j] = '.'
			}
		}
	}
	lines := make([]string, maxY+1)
	for i := 0; i < len(printMatrix); i++ {
		lines[i] = string(printMatrix[i])
	}

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])

	}
	fmt.Println("pesho")
	// for _, v := range mapOfPoints {
	// 	if condition {

	// 	}
	// }
	//
}

// fold left
// delete.vame
func foldUp(mapOfPoints map[point]point, units int) (foldedMap map[point]point) {
	// if
	// mahame vsichkite na tazi liniq

	// 1.
	// 0,0 i 0,1
	// 0, 13 -> 0,1
	// 0, 14 -> 0,0
	// from y == units, to max
	// units are 0 based ... so 7 = 8 line
	// 2*7 new Y = 2*7-y = 14 - 14 = 0
	// samo y se smenq
	foldedMap = map[point]point{}
	for _, oldPoint := range mapOfPoints {
		if oldPoint.y == units {
			continue // not to be added to new one
			// effectively bye bye
		}
		if oldPoint.y < units {
			// advame bez conversion
			foldedMap[oldPoint] = oldPoint
			continue
		}

		// oldpoint y > units
		oldPoint.y = 2*units - oldPoint.y
		foldedMap[oldPoint] = oldPoint
	}

	return
}

func foldLeft(mapOfPoints map[point]point, units int) (foldedMap map[point]point) {
	foldedMap = map[point]point{}
	for _, oldPoint := range mapOfPoints {
		if oldPoint.x == units {
			continue // not to be added to new one
			// effectively bye bye
		}
		if oldPoint.x < units {
			// advame bez conversion
			foldedMap[oldPoint] = oldPoint
			continue
		}

		// right is affected only
		// oldpoint x > units
		oldPoint.x = 2*units - oldPoint.x
		foldedMap[oldPoint] = oldPoint
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
