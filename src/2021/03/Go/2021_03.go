package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
	"log"
)

const inputPath = "../input.txt"

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	var result int

	result = task01(lines)
	fmt.Println(result)

	result = task02(lines) // possible conversion stuff
	fmt.Println(result)
}

func task01(lines []string) (result int) {
	totalLines := len(lines)
	lineLength := len(lines[0]) // assumption equal line-length

	gamma := ""   // most common bit
	epsilon := "" // least common bit
	for x := 0; x < lineLength; x++ {
		column := make([]rune, 0)
		for y := 0; y < totalLines; y++ {
			column = append(column, rune(lines[y][x]))
		}
		// no equally common ... so assumption is that epsilon is the inverted gamma

		count0, count1 := returnSliceCountsForOnesAndZeroes(column)
		if count1 >= count0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	// TODO: check conversion
	result = convertFromBinary(gamma) * convertFromBinary(epsilon)
	return result
}

func task02(lines []string) (result int) {

	totalLines := len(lines)
	lineLength := len(lines[0]) // assumption equal line-length
	hashmapIndecesOxygen := make(map[int]int)
	hashmapIndecesCo2 := make(map[int]int)
	for i := 0; i < totalLines; i++ {
		hashmapIndecesOxygen[i] = i
		hashmapIndecesCo2[i] = i
	}

	for x := 0; x < lineLength; x++ {
		runeLineOxygen := make(map[int]rune)
		runeLineCo2 := make(map[int]rune)
		for y := 0; y < totalLines; y++ {
			if _, ok := hashmapIndecesOxygen[y]; ok {
				runeLineOxygen[y] = rune(lines[y][x])
			}
			if _, ok := hashmapIndecesCo2[y]; ok {
				runeLineCo2[y] = rune(lines[y][x])
			}
		}

		if len(hashmapIndecesOxygen) > 1 {
			filterLogic(runeLineOxygen, &hashmapIndecesOxygen, "oxygen")
		}
		if len(hashmapIndecesCo2) > 1 {
			filterLogic(runeLineCo2, &hashmapIndecesCo2, "co2")
		}
	}

	if len(hashmapIndecesOxygen) != 1 {
		log.Panic("Something's wrong - Oxygen value should be just one number")
	}
	var oxygen string
	for k := range hashmapIndecesOxygen {
		oxygen = lines[k]
	}

	if len(hashmapIndecesCo2) != 1 {
		log.Panic("Something's wrong - CO2 Value should be just one number")
	}
	var co2 string
	for k := range hashmapIndecesCo2 {
		co2 = lines[k]
	}

	// TODO: check conversion
	result = convertFromBinary(oxygen) * convertFromBinary(co2)
	return result
}

func determineFilterRune(count0 int, count1 int, filterStrategy string) (result rune) {
	switch filterStrategy {
	case "oxygen":
		if count1 >= count0 {
			result = '1'
		} else {
			result = '0'
		}
	case "co2":
		if count0 <= count1 {
			result = '0'
		} else {
			result = '1'
		}
	default:
		log.Panic("Unsupported strategy")
	}

	return result
}

func filterLogic(lineRuneMap map[int]rune, hashMapIndeces *map[int]int, filterStrategy string) {
	count0, count1 := returnMapCountsForOnesAndZeroes(lineRuneMap)

	filterRune := determineFilterRune(count0, count1, filterStrategy)

	filterHashmap(lineRuneMap, hashMapIndeces, filterRune)
}

func returnSliceCountsForOnesAndZeroes(lines []rune) (count0 int, count1 int) {
	for _, v := range lines {
		switch v {
		case '1':
			count1++
		case '0':
			count0++
		}
	}
	return
}

func returnMapCountsForOnesAndZeroes(lines map[int]rune) (count0 int, count1 int) {
	for _, v := range lines {
		switch v {
		case '1':
			count1++
		case '0':
			count0++
		}
	}
	return
}

func filterHashmap(lineRuneMap map[int]rune, hashmapIndeces *map[int]int, filter rune) {
	for k, v := range lineRuneMap {
		if v != filter {
			delete(*hashmapIndeces, k)
		}
	}
}

// ToDo Overflow Handling
func convertFromBinary(binaryString string) (result int) {
	for i := 0; i < len(binaryString); i++ {
		if !(binaryString[i] == '0' || binaryString[i] == '1') {
			log.Panic("Only accept 0 and 1 in string")
		}
		result <<= 1 // make place to the left
		if binaryString[i] == '1' {
			result |= 1
		}
	}
	return
}
