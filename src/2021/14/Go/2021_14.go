package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int
	mapInsertionRules, initialTemplate := parseInput(lines)

	// part 1
	result = findResult(mapInsertionRules, initialTemplate, 10)
	fmt.Println(result)

	// part 2
	result = findResult(mapInsertionRules, initialTemplate, 40)
	fmt.Println(result)
}

type pair struct {
	first  rune
	second rune
}

// polymer template first line
// pairs
// pair insertion rules
func parseInput(slicesOfLines []string) (mapInsertionRules map[pair]rune, initialTemplate []rune) {

	// points
	initialTemplate = []rune(slicesOfLines[0])
	mapInsertionRules = map[pair]rune{}
	i := 2
	// memory
	for ; i < len(slicesOfLines); i++ {
		line := slicesOfLines[i]
		splitted := strings.Split(line, " -> ")
		firstSecond := splitted[0]
		between := rune(splitted[1][0])
		first := rune(firstSecond[0])
		second := rune(firstSecond[1])

		pairChen := pair{first, second}
		mapInsertionRules[pairChen] = between
	}
	return
}

const inputPath = "../input.txt"

func findResult(mapInsertionRules map[pair]rune,
	initialTemplate []rune, steps int) (result int) {

	// Initial state
	elementCounts := map[rune]int{}
	for i := 0; i < len(initialTemplate); i++ {
		elementCounts[initialTemplate[i]]++
	}

	pairCounts := map[pair]int{}
	for i := 0; i < len(initialTemplate)-1; i++ {
		first := initialTemplate[i]
		second := initialTemplate[i+1]
		newPair := pair{first, second}
		pairCounts[newPair]++
	}

	// Step iteration
	for i := 0; i < steps; i++ {
		pairCountsNew := map[pair]int{}
		for p, count := range pairCounts {
			elementToBeInserted := mapInsertionRules[p]
			firstNewPair := pair{p.first, elementToBeInserted}
			secondNewWPair := pair{elementToBeInserted, p.second}
			pairCountsNew[firstNewPair] += count
			pairCountsNew[secondNewWPair] += count
			elementCounts[elementToBeInserted] += count
		}
		pairCounts = pairCountsNew
	}

	// Sort and Result
	arrayInts := []int{}
	for _, v := range elementCounts {
		arrayInts = append(arrayInts, v)
	}
	sort.Ints(arrayInts)
	max := arrayInts[len(arrayInts)-1]
	min := arrayInts[0]
	result = max - min
	return
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
