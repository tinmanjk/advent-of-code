package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

const inputPath = "../input0.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	// lines := returnSliceOfLinesFromFile()
	result := int(task02(lines)) // possible conversion stuff
	fmt.Println(result)

	// result = task02(numbers)
	// fmt.Println(result)
}

func task02(lines []string) (result float64) {

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
			oxygenFilterLogic(runeLineOxygen, &hashmapIndecesOxygen)
		}
		if len(hashmapIndecesCo2) > 1 {
			co2FilterLogic(runeLineCo2, &hashmapIndecesCo2)
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

	var oxygenFloat, co2Float float64
	oxygenFloat = convertFromBinary(oxygen)
	co2Float = convertFromBinary(co2)

	return oxygenFloat * co2Float
}

func returnCountsForOnesAndZeroes(lines map[int]rune) (count0 int, count1 int) {
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

func filterHashmap(lineRuneMap map[int]rune, hashmapIndeces *map[int]int, filterOut rune) {
	for k, v := range lineRuneMap {
		if v != filterOut {
			delete(*hashmapIndeces, k)
		}
	}
}

func oxygenFilterLogic(lineRuneMap map[int]rune, hashMapIndeces *map[int]int) {
	count0, count1 := returnCountsForOnesAndZeroes(lineRuneMap)

	var oxygenFilter rune
	if count1 >= count0 {
		oxygenFilter = '1'
	} else {
		oxygenFilter = '0'
	}

	filterHashmap(lineRuneMap, hashMapIndeces, oxygenFilter)
}

func co2FilterLogic(lineRuneMap map[int]rune, hashMapIndeces *map[int]int) {
	count0, count1 := returnCountsForOnesAndZeroes(lineRuneMap)

	var co2Filter rune
	if count0 <= count1 {
		co2Filter = '0'
	} else {
		co2Filter = '1'
	}

	filterHashmap(lineRuneMap, hashMapIndeces, co2Filter)
}

func convertFromBinary(binaryString string) (result float64) {
	for i := len(binaryString) - 1; i >= 0; i-- {
		power := float64(len(binaryString) - i - 1)

		if binaryString[i] == '1' {
			result += math.Pow(2, power)
		}
	}
	return result
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
