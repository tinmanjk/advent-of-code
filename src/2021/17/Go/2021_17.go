package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	var result int
	targetArea := parseInput(lines)
	// part 1
	result = findResult(targetArea, false)
	fmt.Println(result)
	// part 2
	result = findResult(targetArea, true)
	fmt.Println(result)
}

func parseInput(slicesOfLines []string) (targetArea TargetArea) {

	line := slicesOfLines[0]
	firstSplit := strings.Split(line, ":")
	secondSplit := strings.Split(firstSplit[1], ",")
	xSplit := secondSplit[0]
	ySplit := secondSplit[1]
	xSplit = xSplit[3:]
	ySplit = ySplit[3:]
	x1x2 := strings.Split(xSplit, "..")
	y1y2 := strings.Split(ySplit, "..")

	targetArea = TargetArea{}
	targetArea.x1, _ = strconv.Atoi(x1x2[0])
	targetArea.x2, _ = strconv.Atoi(x1x2[1])
	targetArea.y1, _ = strconv.Atoi(y1y2[0])
	targetArea.y2, _ = strconv.Atoi(y1y2[1])

	return
}

type TargetArea struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

type Velocity struct {
	x int
	y int
}

const inputPath = "../input.txt"

func findValidXVelocities(targetArea TargetArea) (validXVelocityStep map[int][]int) {

	validXVelocityStep = map[int][]int{}

	// start at x that doesn't overshoot with 1 step
	// go backwards
	for xVelocity := targetArea.x2; ; xVelocity-- {
		// Loop Break Condition
		maxTravelX := (xVelocity * (xVelocity + 1)) / 2 // n(n+1)/2 formula
		if maxTravelX < targetArea.x1 {
			break
			// decreasing velocity with have even smaller maxTravel
		}

		steps := []int{}
		distanceTraveled := 0
		countStep := 0
		for i := xVelocity; i > 0; i-- {
			distanceTraveled += i
			countStep++
			if targetArea.x1 <= distanceTraveled && distanceTraveled <= targetArea.x2 {
				steps = append(steps, countStep)
			}
			if distanceTraveled > targetArea.x2 {
				break
			}
		}
		if len(steps) != 0 {
			validXVelocityStep[xVelocity] = steps
		}
	}

	return
}

func calcTravelY(initialVelocity int, stepsToMake int) (coorY int, finalVelocity int) {

	finalVelocity = initialVelocity

	for s := 0; s < stepsToMake; s++ {
		coorY += finalVelocity
		finalVelocity--
	}
	return
}

// assuming positive X
func calcTravelX(initialVelocity int, stepsToMake int) (coorX int, finalVelocity int) {

	finalVelocity = initialVelocity

	for s := 0; s < stepsToMake; s++ {
		if finalVelocity == 0 {
			return
		}
		coorX += finalVelocity
		finalVelocity--
	}
	return
}

func findResult(targetArea TargetArea, partTwo bool) (result int) {

	validXVelocitiesSteps := findValidXVelocities(targetArea)

	velocitiesMap := map[Velocity]Velocity{}
	highestY := math.MinInt32
	for x, possibleStepsMadeToTarget := range validXVelocitiesSteps {
		// STEP DETERMINATION
		for _, stepsMadeToTarget := range possibleStepsMadeToTarget {
			// TODO Find Out why targetArea.y1 is good boundary
			for candidateY := targetArea.y1; candidateY <= -targetArea.y1; candidateY++ {

				coorY, yVelocityAtStep := calcTravelY(candidateY, stepsMadeToTarget)
				_, xVelocityAtStep := calcTravelX(x, stepsMadeToTarget)

				if xVelocityAtStep == 0 { // Free Fall Case - x stationary, y increases
					// save going up and then down - cancel each other out
					if yVelocityAtStep > 0 {
						yVelocityAtStep = -yVelocityAtStep - 1
					}
					for {
						coorY += yVelocityAtStep
						if coorY < targetArea.y1 { // "overshoot"
							// return last not-overshooting  yVelocity to try below
							coorY -= yVelocityAtStep
							break
						}

						yVelocityAtStep--
					}
				}

				// we have guaranteed x in
				// find if candidateY's trajectory will be in y bounds
				if targetArea.y1 <= coorY && coorY <= targetArea.y2 {
					velocitiesMap[Velocity{x, candidateY}] = Velocity{x, candidateY}
					if candidateY >= highestY {
						highestY = candidateY
					}
				}
			}
		}

	}

	if partTwo {
		return len(velocitiesMap)
	}

	result = highestY * (highestY + 1) / 2
	return
}
