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
				multiplier /= 10
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
	a := rune(diffAdditions(digitToUniqueMap[1], digitToUniqueMap[7])[0])
	segmentsDecodeMap[a] = 'a'
	// 3 vs 5 = b
	b := rune(diffAdditions(digitToUniqueMap[3], digitToUniqueMap[5])[0])
	segmentsDecodeMap[b] = 'b'
	// 6 vs 8 = c
	c := rune(diffAdditions(digitToUniqueMap[6], digitToUniqueMap[8])[0])
	segmentsDecodeMap[c] = 'c'
	// 0 vs 8 = d
	d := rune(diffAdditions(digitToUniqueMap[0], digitToUniqueMap[8])[0])
	segmentsDecodeMap[d] = 'd'
	// 9 vs 8 = e
	e := rune(diffAdditions(digitToUniqueMap[9], digitToUniqueMap[8])[0])
	segmentsDecodeMap[e] = 'e'
	// 2 vs 3 = f
	f := rune(diffAdditions(digitToUniqueMap[2], digitToUniqueMap[3])[0])
	segmentsDecodeMap[f] = 'f'

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

	twoThreeFive := make([]string, 0)
	zeroSixNine := make([]string, 0)
	for _, v := range tenPattern {
		length := len(v)
		switch length {
		case 2:
			digitToCode[1] = v
		case 3:
			digitToCode[7] = v
		case 4:
			digitToCode[4] = v
		case 5:
			twoThreeFive = append(twoThreeFive, v)
		case 6:
			zeroSixNine = append(zeroSixNine, v)
		case 7:
			digitToCode[8] = v
		}
	}
	// 0 vs 8 = d
	// 6 vs 8 = c
	// 9 vs 8 = e
	// -> Find Dif 8 vs 6-length ones gives us dce
	dce := ""
	for i := 0; i < len(zeroSixNine); i++ {
		difference := diffAdditions(zeroSixNine[i], digitToCode[8])
		dce += difference
	}

	// find 2, 3, 5
	// 2 vs 5 -> 2 !!!
	// 2 vs 3 -> 1
	// 3 vs 5 -> 1
	for i := 0; i < len(twoThreeFive); i++ {
		for k := i + 1; k < len(twoThreeFive); k++ {
			differences := len(diffAdditions(twoThreeFive[i], twoThreeFive[k]))
			if differences == 2 {
				// 2 contains all of dce
				// other part of the pair is 5
				if len(diffAdditions(twoThreeFive[i], dce)) == 0 {
					digitToCode[2] = twoThreeFive[i]
					digitToCode[5] = twoThreeFive[k]
				} else {
					digitToCode[2] = twoThreeFive[k]
					digitToCode[5] = twoThreeFive[i]
				}
				// three is the remaining
				for index := range twoThreeFive {
					if index != i && index != k {
						digitToCode[3] = twoThreeFive[index]
					}
				}
			}
		}
	}

	for _, v := range zeroSixNine {
		// 5 vs 6 -> 1
		// 5 vs 9 -> 1
		// 5 vs 0 -> 2 !!!
		difference := diffAdditions(digitToCode[5], v)
		if len(difference) == 2 {
			digitToCode[0] = v
			continue
		}

		// 9 vs 7 -> 0 additions
		// 6 vs 7 -> 1
		difference = diffAdditions(v, digitToCode[7])
		if len(difference) == 0 {
			digitToCode[9] = v
		} else {
			digitToCode[6] = v
		}
	}

	return
}

func diffAdditions(left string, right string) string {
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
