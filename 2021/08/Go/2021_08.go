package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	solutionData := parseInput(lines)

	var result int

	result = findResult(solutionData, true)
	fmt.Println(result)

	result = findResult(solutionData, false)
	fmt.Println(result)
}

type inputLine struct {
	tenUniqueSignalPattern []string
	fourDigitOutputValue   []string
}

func parseInput(slicesOfLines []string) (sliceOfInputLines []inputLine) {

	sliceOfInputLines = make([]inputLine, len(slicesOfLines))
	for i := 0; i < len(slicesOfLines); i++ {
		line := slicesOfLines[i]
		splitted := strings.Split(line, " | ")
		// 10 string
		tenDigitPatterns := strings.Split(splitted[0], " ")
		// 4 string
		fourDigits := strings.Split(splitted[1], " ")
		inputL := inputLine{tenDigitPatterns, fourDigits}
		sliceOfInputLines[i] = inputL

	}
	return
}

func findResult(solutionData []inputLine, partOne bool) (result int) {

	for i := 0; i < len(solutionData); i++ {
		digitToSignalPatternMap := createDigitToSignalMap(solutionData[i].tenUniqueSignalPattern)
		// signalwire is e.g. 'b' -> segment is also 'b' - do not match
		signalWireToSegmentMap := createSignalWireToSegmentMap(digitToSignalPatternMap)

		multiplier := 1000
		for k := 0; k < 4; k++ {
			digit := decodeDigit(solutionData[i].fourDigitOutputValue[k], signalWireToSegmentMap)
			if partOne {
				switch digit {
				case 1, 4, 7, 8:
					result++
				}
			} else {
				result += multiplier * digit
				multiplier = multiplier / 10
			}

		}
	}
	return
}

func decodeDigit(code string, decodeMap map[rune]rune) (digit int) {

	decodedRunes := make([]rune, 0)
	for _, v := range code {
		decodedRunes = append(decodedRunes, decodeMap[v])
	}

	sort.Slice(decodedRunes, func(i int, j int) bool { return decodedRunes[i] < decodedRunes[j] })
	sortedDecoded := string(decodedRunes)

	zero := "abcefg"
	one := "cf"
	two := "acdeg"
	three := "acdfg"
	four := "bcdf"
	five := "abdfg"
	six := "abdefg"
	seven := "acf"
	eight := "abcdefg"
	nine := "abcdfg"

	switch sortedDecoded {
	case zero:
		return 0
	case one:
		return 1
	case two:
		return 2
	case three:
		return 3
	case four:
		return 4
	case five:
		return 5
	case six:
		return 6
	case seven:
		return 7
	case eight:
		return 8
	case nine:
		return 9
	}

	return
}

func createSignalWireToSegmentMap(digitToUniqueMap map[int]string) (segmentsDecodeMap map[rune]rune) {

	segmentsDecodeMap = make(map[rune]rune, 7)
	// 1 vs 7 = a
	a := rune(diff(digitToUniqueMap[1], digitToUniqueMap[7])[0])
	segmentsDecodeMap[a] = 'a'
	// 0 vs 8 = d
	d := rune(diff(digitToUniqueMap[0], digitToUniqueMap[8])[0])
	segmentsDecodeMap[d] = 'd'
	// 6 vs 8 = c
	c := rune(diff(digitToUniqueMap[6], digitToUniqueMap[8])[0])
	segmentsDecodeMap[c] = 'c'
	// 9 vs 8 = e
	e := rune(diff(digitToUniqueMap[9], digitToUniqueMap[8])[0])
	segmentsDecodeMap[e] = 'e'
	// 2 vs 3 = f
	f := rune(diff(digitToUniqueMap[2], digitToUniqueMap[3])[0])
	segmentsDecodeMap[f] = 'f'
	// 3 vs 5 = b
	b := rune(diff(digitToUniqueMap[3], digitToUniqueMap[5])[0])
	segmentsDecodeMap[b] = 'b'

	// g - the one left
	for _, r := range "abcdefg" {
		if _, ok := segmentsDecodeMap[r]; !ok {
			segmentsDecodeMap[r] = 'g'
			break
		}
	}

	return
}

func createDigitToSignalMap(tenPattern []string) (digitToCode map[int]string) {
	digitToCode = make(map[int]string, 0)

	lengthPattern := make(map[int][]string, 0)
	lengthPattern[5] = make([]string, 0)
	lengthPattern[6] = make([]string, 0)
	for _, v := range tenPattern {
		length := len(v)
		switch length {
		case 2:
			digitToCode[1] = v
		case 3:
			digitToCode[7] = v
		case 4:
			digitToCode[4] = v
		case 5: // 2 3 5
			lengthPattern[5] = append(lengthPattern[5], v)
		case 6: // 0, 6, 9
			lengthPattern[6] = append(lengthPattern[6], v)
		case 7:
			digitToCode[8] = v
		}
	}
	// 0 vs 8 = d
	// 6 vs 8 = c
	// 9 vs 8 = e
	// -> Find Dif
	dce := ""
	for i := 0; i < len(lengthPattern[6]); i++ {
		difference := diff(lengthPattern[6][i], digitToCode[8])
		dce += difference
	}

	// find 2, 3, 5
	// 2 vs 5 -> 2
	// 2 vs 3 -> 1
	// 3 vs 5 -> 1
	case0v1 := diff(lengthPattern[5][0], lengthPattern[5][1])
	case1v2 := diff(lengthPattern[5][1], lengthPattern[5][2])
	case0v2 := diff(lengthPattern[5][0], lengthPattern[5][2])
	switch {
	case len(case0v1) == 2: // 2 i 5
		digitToCode[3] = lengthPattern[5][2]
		// v dce tarsim .. ako containva vsichkite - znachi e 2, inache e 5
		if containsAllRunesFromPattern(lengthPattern[5][0], dce) {
			digitToCode[2] = lengthPattern[5][0]
			digitToCode[5] = lengthPattern[5][1]
		} else {
			digitToCode[2] = lengthPattern[5][1]
			digitToCode[5] = lengthPattern[5][0]
		}

	case len(case1v2) == 2: // 2 i5
		digitToCode[3] = lengthPattern[5][0]
		if containsAllRunesFromPattern(lengthPattern[5][1], dce) {
			digitToCode[2] = lengthPattern[5][1]
			digitToCode[5] = lengthPattern[5][2]
		} else {
			digitToCode[2] = lengthPattern[5][2]
			digitToCode[5] = lengthPattern[5][1]
		}
	case len(case0v2) == 2: // 2 i 5
		digitToCode[3] = lengthPattern[5][1]
		if containsAllRunesFromPattern(lengthPattern[5][0], dce) {
			digitToCode[2] = lengthPattern[5][0]
			digitToCode[5] = lengthPattern[5][2]
		} else {
			digitToCode[2] = lengthPattern[5][2]
			digitToCode[5] = lengthPattern[5][0]
		}
	}

	// 5 vs 0 vs 6 vs 9 -> if 2 razliki

	for i := 0; i < len(lengthPattern[6]); i++ {
		difference := diff(digitToCode[5], lengthPattern[6][i])
		if len(difference) == 2 {
			digitToCode[0] = lengthPattern[6][i]
		}
	}

	// 7miciata se vklucva izqcalo v 9kata, dokato ne v 6
	for i := 0; i < len(lengthPattern[6]); i++ {
		// we do have 0
		if lengthPattern[6][i] == digitToCode[0] {
			continue
		}

		// namirame 9kata
		contains := true
		for _, r := range digitToCode[7] {
			if !strings.Contains(lengthPattern[6][i], string(r)) {
				contains = false
			}
		}
		if contains {
			digitToCode[9] = lengthPattern[6][i]
		}
	}

	// namirame 6ca
	for i := 0; i < len(lengthPattern[6]); i++ {
		found := false
		for _, v := range digitToCode {
			if v == lengthPattern[6][i] {
				found = true
			}
		}
		if !found {
			digitToCode[6] = lengthPattern[6][i]
			break
		}
	}

	return
}

func containsAllRunesFromPattern(str string, pattern string) bool {

	for _, r := range pattern {
		if !strings.ContainsRune(str, r) {
			return false
		}
	}
	return true
}

func diff(left string, right string) string {
	temp := make([]rune, 0)
	for _, r := range right {
		if !strings.ContainsRune(left, r) {
			temp = append(temp, r)
		}
	}
	return string(temp)
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
