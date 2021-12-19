package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	var result int
	// part 1
	result = findResult(lines, false)
	fmt.Println(result)

	result = findResult(lines, true)
	fmt.Println(result)
}

const inputPath = "../input.txt"

func findResult(lines []string, partTwo bool) (result int) {

	currentNumber := lines[0]
	for i := 1; i < len(lines); i++ {
		currentNumber = "[" + currentNumber + "," + lines[i] + "]"
		currentNumber = reduceNumber(currentNumber)
	}

	if partTwo {
		maxSum := 0
		for i := 0; i < len(lines); i++ {
			firstNumber := lines[i]
			for j := 0; j < len(lines); j++ {
				if firstNumber == lines[j] {
					continue
				}
				secondNumber := lines[j]
				sumNumber := "[" + firstNumber + "," + secondNumber + "]"
				sumReduced := reduceNumber(sumNumber)
				magnitude := calcMagnitude(sumReduced)
				if magnitude >= maxSum {
					maxSum = magnitude
				}
			}
		}
		return maxSum
	}

	result = calcMagnitude(currentNumber)
	return
}

func splitPairLeftRight(number string) (left string, right string) {

	countOpen := 1 // always open
	splitIndex := -1
	for i := 1; i < len(number); i++ {
		if countOpen == 1 && number[i] == ',' {
			splitIndex = i
			break
		}
		switch number[i] {
		case '[':
			countOpen++
		case ']':
			countOpen--
		}
	}

	// without the [ and ]
	left = number[1:splitIndex]
	right = number[splitIndex+1 : len(number)-1]

	return
}

func calcMagnitude(number string) (result int) {

	var leftNumber, rightNumber int
	isOnePair := strings.Count(number, "[") == 1
	if isOnePair {
		numbersToSplit := number[1 : len(number)-1]
		splittedNumbers := strings.Split(numbersToSplit, ",")
		leftNumber, _ = strconv.Atoi(splittedNumbers[0])
		rightNumber, _ = strconv.Atoi(splittedNumbers[1])
		result = 3*leftNumber + 2*rightNumber
		return
	}
	left, right := splitPairLeftRight(number)

	if left[0] == '[' {
		leftNumber = calcMagnitude(left)
	} else {
		leftNumber, _ = strconv.Atoi(left)
	}

	if right[0] == '[' {
		rightNumber = calcMagnitude(right)
	} else {
		rightNumber, _ = strconv.Atoi(right)
	}

	result = 3*leftNumber + 2*rightNumber
	return
}

func reduceNumber(number string) (reduced string) {

	// first just explode over ALL of the number
explodeStrategy:
	openPairTagCounter := 0
	// traverse the whole number to find at least one explode
	for i := 0; i < len(number); i++ {

		switch number[i] {
		case '[':
			openPairTagCounter++
		case ']':
			openPairTagCounter--

		}
		if openPairTagCounter == 5 {
			number = explodePair(number, i)
			goto explodeStrategy
		}
	}

	// in case no explode was found
	for i := 0; i < len(number); i++ {
		foundNumber := false
		numberFound := 0
		firstDigitIndex := 0
		foundNumberAsString := ""

		if '0' <= number[i] && number[i] <= '9' {
			foundNumber = true
			firstDigitIndex = i
			additionalDigitsCount := 0
			for j := i + 1; j < len(number); j++ {
				if '0' <= number[j] && number[j] <= '9' {
					additionalDigitsCount++
				} else {
					break
				}
			}
			lastDigitIndex := firstDigitIndex + additionalDigitsCount
			foundNumberAsString = number[firstDigitIndex : lastDigitIndex+1]
			numberFound, _ = strconv.Atoi(foundNumberAsString)
		}

		if foundNumber && numberFound >= 10 {
			number = splitBigIntoPair(number, firstDigitIndex, foundNumberAsString)
			goto explodeStrategy
		}
	}

	reduced = number
	return
}

func splitBigIntoPair(number string, startIndexNumber int, tobeSplitNumber string) (splitted string) {

	leftPart := number[:startIndexNumber]
	numberSize := len(tobeSplitNumber)
	rightPart := number[startIndexNumber+numberSize:]
	tobeSplitNumberInt, _ := strconv.Atoi(tobeSplitNumber)
	leftPair := tobeSplitNumberInt / 2 // truncates floor by default
	rightPair := leftPair + (tobeSplitNumberInt % 2)

	pair := fmt.Sprintf("[%v,%v]", leftPair, rightPair)

	splitted = leftPart + pair + rightPart
	return
}

func explodePair(number string, startIndexPair int) (exploded string) {

	leftSide := number[:startIndexPair]
	// pair
	// startIndex is the "["
	pairSize := 0
	for i := startIndexPair; ; i++ {
		pairSize++
		if number[i] == ']' {
			break
		}
	}
	pair := number[startIndexPair : startIndexPair+pairSize]
	rightSide := number[startIndexPair+pairSize:]
	pairJustNumbers := pair[1 : len(pair)-1]
	pairJustNumbersSpitted := strings.Split(pairJustNumbers, ",")
	leftNumber, _ := strconv.Atoi(pairJustNumbersSpitted[0])
	rightNumber, _ := strconv.Atoi(pairJustNumbersSpitted[1])

	foundLeftNumber := false
	for i := len(leftSide) - 1; i >= 0; i-- {
		if '0' <= leftSide[i] && leftSide[i] <= '9' {
			foundLeftNumber = true
			lastDigitIndex := i
			additionalDigitsCount := 0
			for j := i - 1; j >= 0; j-- {
				if '0' <= leftSide[j] && leftSide[j] <= '9' {
					additionalDigitsCount++
				} else {
					break
				}
			}
			firstDigitIndex := lastDigitIndex - additionalDigitsCount

			// imame number Count
			numberToBeReplaced := leftSide[firstDigitIndex : lastDigitIndex+1]
			numbetoBeReplacedInt, _ := strconv.Atoi(numberToBeReplaced)
			newNumber := numbetoBeReplacedInt + leftNumber
			leftBeforeNumber := leftSide[:firstDigitIndex]
			leftAfterNumber := leftSide[lastDigitIndex+1:]
			// TODO sprintF v ??
			leftSide = leftBeforeNumber + fmt.Sprintf("%v", newNumber) + leftAfterNumber
		}

		if foundLeftNumber {
			break
		}
	}

	// TODO refactor to eliminate duplication with above
	foundRightNumber := false
	for i := 0; i < len(rightSide); i++ {
		if '0' <= rightSide[i] && rightSide[i] <= '9' {
			foundRightNumber = true
			firstDigitIndex := i
			additionalDigitsCount := 0
			for j := i + 1; j < len(rightSide); j++ {
				if '0' <= rightSide[j] && rightSide[j] <= '9' {
					additionalDigitsCount++
				} else {
					break
				}
			}
			lastDigitIndex := firstDigitIndex + additionalDigitsCount

			// imame number Count
			numberToBeReplaced := rightSide[firstDigitIndex : lastDigitIndex+1]
			numbetoBeReplacedInt, _ := strconv.Atoi(numberToBeReplaced)
			newNumber := numbetoBeReplacedInt + rightNumber
			rightBeforeNumber := rightSide[:firstDigitIndex]
			rightAfterNumber := rightSide[lastDigitIndex+1:]
			// TODO sprintF v ??
			rightSide = rightBeforeNumber + fmt.Sprintf("%v", newNumber) + rightAfterNumber
		}

		if foundRightNumber {
			break
		}
	}

	exploded = leftSide + "0" + rightSide
	converted := exploded
	return converted
}

// DEBUGGING MANUALLY
// [[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]
// [[[[4,0],[5,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[0,[7,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,0],[15,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,0],[[7,8],5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[0,13]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[0,[6,7]]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[14,[[[3,7],[4,3]],[[6,3],[8,8]]]]]

// prefer 3 7 over 14 split
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[14,[[[3,7],[4,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[17,[[0,[11,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[8,9],[[0,[11,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[8,9],[[0,[5,6],3]],[[6,3],[8,8]]]]]

// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[7,7],[[[3,7],[4,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[7,10],[[0,[11,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[7,[5,5]],[[0,[11,3]],[[6,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[7,[5,5]],[[11,0],[[9,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[7,[5,5]],[[[5,5],0],[[9,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[7,[5,10]],[[0,5],[[9,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[7,[5,[5,5]]],[[0,5],[[9,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[7,[10,0]],[[5,5],[[9,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[7,[[5,5],0]],[[5,5],[[9,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[12,[0,5]],[[5,5],[[9,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[0,5]],[[5,5],[[9,3],[8,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[0,5]],[[5,14],[0,[11,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[0,5]],[[5,[7,7]],[0,[11,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[0,5]],[[12,0],[7,[11,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[0,5]],[[[6,6],0],[7,[11,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[0,11]],[[0,6],[7,[11,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[0,[5,6]]],[[0,6],[7,[11,8]]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,0]],[[6,6],[7,[11,8]]]]] ->> interesting 11-8
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,0]],[[6,6],[18,0]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,0]],[[6,6],[[9,9],0]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,0]],[[6,15],[0,9]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,0]],[[6,[7,8]],[0,9]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,0]],[[13,0],[8,9]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,0]],[[[6,7],0],[8,9]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,6]],[[0,7],[8,9]]]] --> interesting 0-7
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,6]],[0,[15,9]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,6]],[0,[[7,8],9]]]]
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,6]],[7,[0,17]]]] --> interesting LAST
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,6]],[7,[0,[8,9]]]]] -> interesting last pair
// [[ [ [4,0],[5,4] ],[[7,7],[6,0]]],[[[6,6],[5,6]],[7,[8,0]]]]

// [ 14,	[ [[3,7],[4,3]],  [[6,3],[8,8]] ] ]
