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
		signalWireToSegmentMap := createSignalWireToSegmentMap(solutionData[i].tenUniqueSignalPattern)

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

func createSignalWireToSegmentMap(tenPattern []string) (decodeTemplate map[byte]byte) {
	var cf, acf, bcdf string // 1, 7, 4
	twoThreeFive := make([]string, 0)

	for _, s := range tenPattern {
		switch len(s) {
		case 2:
			cf = s
		case 3:
			acf = s
		case 4:
			bcdf = s
		case 5:
			twoThreeFive = append(twoThreeFive, s)
		}
	}

	a := setDiff(cf, acf)
	abcdf := a + bcdf
	// abcdf -> "acdeg": 2, -> eg
	// abcdf -> "acdfg": 3, -> g
	// abcdf -> "abdfg": 5, -> g
	var acdeg, eg, g string
	for _, s := range twoThreeFive {
		diff := setDiff(abcdf, s)
		if len(diff) == 2 {
			acdeg = s // 2
			eg = diff
		} else {
			g = diff
		}
	}

	e := setDiff(g, eg)
	aeg := a + eg
	cd := setDiff(aeg, acdeg)
	bd := setDiff(cf, bcdf)

	// cf - cd - bd
	b := setDiff(cd, bd)
	c := setDiff(bd, cd)
	d := setDiff(cf, cd)
	f := setDiff(cd, cf)

	decodeTemplate = make(map[byte]byte, 0)
	decodeTemplate[a[0]] = 'a'
	decodeTemplate[b[0]] = 'b'
	decodeTemplate[c[0]] = 'c'
	decodeTemplate[d[0]] = 'd'
	decodeTemplate[e[0]] = 'e'
	decodeTemplate[f[0]] = 'f'
	decodeTemplate[g[0]] = 'g'

	return
}

func decodeDigit(code string, segmentsDecodeMap map[byte]byte) (digit int) {

	decoded := make([]byte, 0)
	for i := 0; i < len(code); i++ {
		decoded = append(decoded, segmentsDecodeMap[code[i]])
	}

	sort.Slice(decoded, func(i int, j int) bool { return decoded[i] < decoded[j] })
	sortedDecoded := string(decoded)

	mapResults := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

	if val, ok := mapResults[sortedDecoded]; ok {
		digit = val
	} else {
		// some error handling TODO
	}

	return
}

// https://en.wikipedia.org/wiki/Complement_(set_theory)#Relative_complement
// returns the elements of right which are not part of left
func setDiff(left string, right string) string {
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
