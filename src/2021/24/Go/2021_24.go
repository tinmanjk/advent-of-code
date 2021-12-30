package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	fifth, sixth, sixteenth := parseInput(lines)

	var result uint64
	// part 1
	result = findResult(fifth, sixth, sixteenth, "max")
	fmt.Println(result)
	// part 2
	result = findResult(fifth, sixth, sixteenth, "min")
	fmt.Println(result)
}

func parseInput(lines []string) (fifth []int,
	sixth []int, sixteeneth []int) {

	lines = append([]string{" "}, lines...)
	fifth = make([]int, 15) // fifth row one based
	sixth = make([]int, 15)
	sixteeneth = make([]int, 15)
	for i := 0; i < 14; i++ {
		fifthLine := lines[18*i+5]
		fifthLineSplited := strings.Split(fifthLine, " ")
		fifthLineConverted, _ := strconv.Atoi(fifthLineSplited[2])
		fifth[i+1] = fifthLineConverted // 1 based

		sixthLine := lines[18*i+6]
		sixthLineSplited := strings.Split(sixthLine, " ")
		sixthLineConverted, _ := strconv.Atoi(sixthLineSplited[2])
		sixth[i+1] = sixthLineConverted // 1 based

		sixteenethLine := lines[18*i+16]
		sixteenethLineSplited := strings.Split(sixteenethLine, " ")
		sixteenethLineConverted, _ := strconv.Atoi(sixteenethLineSplited[2])
		sixteeneth[i+1] = sixteenethLineConverted // 1 based
	}
	return
}

const inputPath = "../input.txt"

func findResult(_1or26 []int, add []int, booster []int, valueToFind string) (result uint64) {

	depthZ := make([]int64, 15)        // 1 based
	currentSolution := make([]int, 15) // 1 based

	switch valueToFind {
	case "max":
		for i := 9; i >= 1; i-- {
			foundSolution := dfs(1, i, _1or26, add, booster, depthZ, currentSolution, valueToFind)

			if foundSolution {
				return convertSliceIntDigitsToUint64(currentSolution)
			}
		}
	case "min":
		for i := 1; i <= 9; i++ {
			foundSolution := dfs(1, i, _1or26, add, booster, depthZ, currentSolution, valueToFind)

			if foundSolution {
				return convertSliceIntDigitsToUint64(currentSolution)
			}
		}
	}

	return
}

func convertSliceIntDigitsToUint64(digits []int) (result uint64) {

	multiplier := 1
	for i := len(digits) - 1; i >= 0; i-- {
		result += uint64(digits[i] * multiplier)
		multiplier *= 10
	}
	return
}

func dfs(currentDepth int, currentDigit int,
	_1or26 []int, add []int, booster []int, depthZ []int64,
	currentSolution []int, valueToFind string) (foundSolution bool) {

	currentSolution[currentDepth] = currentDigit

	newZ := int64(0)
	z := depthZ[currentDepth-1]
	z_Mod26 := z % 26
	z_Div26 := z / 26
	z_Mul26 := z * 26
	digitAddDiff := int64(currentDigit - add[currentDepth])
	digitBooster := int64(currentDigit + booster[currentDepth])

	// in this case it will have a similar z as the input
	// however this case will not allow for division which is needed
	// to balance the many multiplication given the MONAD
	if _1or26[currentDepth] == 26 && z_Mod26 != digitAddDiff {
		return false
	}

	switch {
	case _1or26[currentDepth] == 1:
		newZ = z_Mul26 + digitBooster
	case _1or26[currentDepth] == 26:
		if z_Mod26 == digitAddDiff {
			newZ = z_Div26
		}
	}

	if currentDepth == 14 {
		return newZ == 0
	}

	// < 14 still digging
	depthZ[currentDepth] = newZ
	newDepth := currentDepth + 1

	switch valueToFind {
	case "max":
		for i := 9; i >= 1; i-- {
			solutionFound := dfs(newDepth, i, _1or26, add, booster, depthZ, currentSolution, valueToFind)
			if solutionFound {
				return true
			}
		}
	case "min":
		for i := 1; i <= 9; i++ {
			solutionFound := dfs(newDepth, i, _1or26, add, booster, depthZ, currentSolution, valueToFind)
			if solutionFound {
				return true
			}
		}
	}

	return
}

// MONAD into GO -> inlined in dfs
func instructionStep(z int64, digit int,
	_1or26 int, diff int, booster int) (newZ int64) {

	z_Mod26 := z % 26
	z_Div26 := z / 26
	z_Mul26 := z * 26
	digitAddDiff := int64(digit - diff)
	digitBooster := int64(digit + booster)

	switch {
	case _1or26 == 26:
		//1. DIV 26  -> <original
		if z_Mod26 == digitAddDiff {
			// ONLY CASE when it can become 0  if z<26
			newZ = z_Div26 // reverts case 3
		}

		//2. TRUNC + inputbooster -> all three cases but not 0
		// need to alwasy actually delete in inline ABOVE see comments there
		if z_Mod26 != digitAddDiff {
			newZ = z + (digitBooster - z_Mod26)
		}

	// from Input just ONE case here, since z_Mod26 cannnot be negative and digit
	// diff is always negative since diff is double-digit
	case _1or26 == 1: // 3. MUL 26 + inputBooster > original
		newZ = z_Mul26 + digitBooster // digitBooster = new z_Mod26
	}

	return
}
