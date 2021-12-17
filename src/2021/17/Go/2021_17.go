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
			// decreasing velocity will produce even smaller maxTravel
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

// to implement start coord parameter
func calcPossibleStartingYvelocities(inclLowBound int, inclUpperBound int, steps int) (possibleYs []int) {

	for targetY := inclLowBound; targetY <= inclUpperBound; targetY++ {
		numberStartY := 1            // y
		numbeDecreases := 0          // -1
		for s := 1; s < steps; s++ { // no decrease during first step
			numberStartY++
			numbeDecreases++
		}
		sumDecreases := -((numbeDecreases) * (numbeDecreases + 1) / 2)
		// numberStartY * startY + sumDecreases = targetY
		if (targetY-sumDecreases)%numberStartY == 0 { // no remainder - dealing with integers
			startY := (targetY - sumDecreases) / numberStartY
			possibleYs = append(possibleYs, startY)
		}
	}

	return
}

func findResult(targetArea TargetArea, partTwo bool) (result int) {

	validXVelocitiesSteps := findValidXVelocities(targetArea)

	velocitiesMap := map[Velocity]Velocity{}
	highestY := math.MinInt32
	for x, possibleStepsMadeToTarget := range validXVelocitiesSteps {
		for _, stepsMadeToTarget := range possibleStepsMadeToTarget {
			if x != stepsMadeToTarget { // 1 X to fixed number of Ys based on size of y area
				possibleYs := calcPossibleStartingYvelocities(targetArea.y1, targetArea.y2, stepsMadeToTarget)
				for _, y := range possibleYs {
					velocitiesMap[Velocity{x, y}] = Velocity{x, y}
					if y >= highestY {
						highestY = y
					}
				}
			} else { // Free Fall Case - x stationary, y increases -> Many possible Ys for 1 X
				// still need to determine boundaries for free fall case??
				for candidateY := targetArea.y1; candidateY <= -targetArea.y1; candidateY++ {
					coorY, yVelocityAtStep := calcTravelY(candidateY, stepsMadeToTarget)

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

					if targetArea.y1 <= coorY && coorY <= targetArea.y2 {
						velocitiesMap[Velocity{x, candidateY}] = Velocity{x, candidateY}
						if candidateY >= highestY {
							highestY = candidateY
						}
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
