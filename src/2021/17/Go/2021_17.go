package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	var result int
	targetArea := parseInput(lines)
	// part 1
	result = findResult(targetArea, false, true)
	fmt.Println(result)
	// part 2
	result = findResult(targetArea, true, true)
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

const inputPath = "../input0.txt"

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
		for i := xVelocity; i > 0 && distanceTraveled < targetArea.x2; i-- {
			distanceTraveled += i
			countStep++
			if targetArea.x1 <= distanceTraveled && distanceTraveled <= targetArea.x2 {
				steps = append(steps, countStep)
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

func findResult(targetArea TargetArea, partTwo bool, debug bool) (result int) {

	validXVelocitiesSteps := findValidXVelocities(targetArea)

	if debug {
		validxOnlySorted := debugGetSortedValidXOnly(validXVelocitiesSteps)
		fmt.Println(validxOnlySorted)
		// 6-15 vkluchitelno 20-30 vkluchitelno -> MATCH the ShouldVelocities
	}

	velocitiesMap := map[Velocity]Velocity{}
	highestY := math.MinInt32
	for x, possibleStepsMadeToTarget := range validXVelocitiesSteps {
		// STEP DETERMINATION
		for _, stepsMadeToTarget := range possibleStepsMadeToTarget {
		candidateY:
			// TODO Find Out why targetArea.y1 is good boundary
			for candidateY := targetArea.y1; candidateY <= -targetArea.y1; candidateY++ {

				// we have guaranteed x in
				// find if candidateY's trajectory will be in y bounds
				coorY, _ := calcTravelY(candidateY, stepsMadeToTarget)
				if targetArea.y1 <= coorY && coorY <= targetArea.y2 {
					velocitiesMap[Velocity{x, candidateY}] = Velocity{x, candidateY}
					if candidateY >= highestY {
						highestY = candidateY
						continue
					}
				}

				// 3. TRY STEPPING TOGETHER
				// EQUALIZE X
				_, xVelocityAtStep := calcTravelX(x, stepsMadeToTarget)
				_, yVelocityAtStep := calcTravelY(candidateY, stepsMadeToTarget)

				// Free Fall Case -> x is constant
				if xVelocityAtStep == 0 {
					if yVelocityAtStep > 0 {
						freeFallVelocity := -yVelocityAtStep
						freeFallVelocity--
						yVelocityAtStep = freeFallVelocity
					}
					for {
						coorY += yVelocityAtStep
						yVelocityAtStep--
						if targetArea.y1 <= coorY && coorY <= targetArea.y2 {
							velocitiesMap[Velocity{x, candidateY}] = Velocity{x, candidateY}
							if candidateY >= highestY {
								highestY = candidateY
								break
							}
						}

						if coorY < targetArea.y1 {
							// end for all assumption
							continue candidateY
						}
					}
				}

			}
		}

	}

	if debug {
		fmt.Println(" ")
		shouldVelocities := []Velocity{
			{23, -10}, {25, -9}, {27, -5}, {29, -6}, {22, -6}, {21, -7}, {9, 0}, {27, -7}, {24, -5},
			{25, -7}, {26, -6}, {25, -5}, {6, 8}, {11, -2}, {20, -5}, {29, -10}, {6, 3}, {28, -7},
			{8, 0}, {30, -6}, {29, -8}, {20, -10}, {6, 7}, {6, 4}, {6, 1}, {14, -4}, {21, -6},
			{26, -10}, {7, -1}, {7, 7}, {8, -1}, {21, -9}, {6, 2}, {20, -7}, {30, -10}, {14, -3},
			{20, -8}, {13, -2}, {7, 3}, {28, -8}, {29, -9}, {15, -3}, {22, -5}, {26, -8}, {25, -8},
			{25, -6}, {15, -4}, {9, -2}, {15, -2}, {12, -2}, {28, -9}, {12, -3}, {24, -6}, {23, -7},
			{25, -10}, {7, 8}, {11, -3}, {26, -7}, {7, 1}, {23, -9}, {6, 0}, {22, -10}, {27, -6},
			{8, 1}, {22, -8}, {13, -4}, {7, 6}, {28, -6}, {11, -4}, {12, -4}, {26, -9}, {7, 4},
			{24, -10}, {23, -8}, {30, -8}, {7, 0}, {9, -1}, {10, -1}, {26, -5}, {22, -9}, {6, 5},
			{7, 5}, {23, -6}, {28, -10}, {10, -2}, {11, -1}, {20, -9}, {14, -2}, {29, -7}, {13, -3},
			{23, -5}, {24, -8}, {27, -9}, {30, -7}, {28, -5}, {21, -10}, {7, 9}, {6, 6}, {21, -5},
			{27, -10}, {7, 2}, {30, -9}, {21, -8}, {22, -7}, {24, -9}, {20, -6}, {6, 9}, {29, -5},
			{8, -2}, {27, -8}, {30, -5}, {24, -7},
		}
		fmt.Printf("Should velocities count %v\n", len(shouldVelocities))

		myWrong := []Velocity{}
		for _, mv := range velocitiesMap {
			found := false
			for _, sv := range shouldVelocities {
				if sv == mv {
					found = true
					break
				}
			}
			if !found {
				myWrong = append(myWrong, mv)
			}
		}

		fmt.Printf("My velocities wrong %v\n", len(myWrong))
		fmt.Printf("My velocities count %v\n", len(velocitiesMap))

		fmt.Println(" ")
	}

	if partTwo {
		return len(velocitiesMap)
	}

	result = highestY * (highestY + 1) / 2
	return
}

func debugGetSortedValidXOnly(validXVelocityStep map[int][]int) (validXOnlySorted []int) {

	for k := range validXVelocityStep {
		validXOnlySorted = append(validXOnlySorted, k)
	}

	sort.Ints(validXOnlySorted)
	return
}
