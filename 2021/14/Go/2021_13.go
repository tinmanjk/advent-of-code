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
	first  string
	second string
}

// polymer template first line
// pairs
// pair insertion rules
func parseInput(slicesOfLines []string) (mapInsertionRules map[pair]string, initialTemplate []rune) {

	// points
	initialTemplate = []rune(slicesOfLines[0])
	mapInsertionRules = map[pair]string{}
	i := 2
	// memory
	for ; i < len(slicesOfLines); i++ {
		line := slicesOfLines[i]
		splitted := strings.Split(line, " -> ")
		firstSecond := splitted[0]
		between := string(splitted[1][0])
		first := string(firstSecond[0])
		second := string(firstSecond[1])

		pairChen := pair{first, second}
		mapInsertionRules[pairChen] = between
	}
	return
}

const inputPath = "../input.txt"

func findResult(mapInsertionRules map[pair]string,
	initialTemplate []rune, steps int) (result int) {

	elementCounts := map[rune]int{}
	for i := 0; i < len(initialTemplate); i++ {
		elementCounts[initialTemplate[i]]++
	}

	pairCounts := map[pair]int{}
	for i := 0; i < len(initialTemplate)-1; i++ {
		first := string(initialTemplate[i])
		second := string(initialTemplate[i+1])
		newPair := pair{first, second}
		pairCounts[newPair]++
	}

	for i := 0; i < steps; i++ {
		pairCountsNew := map[pair]int{}
		for k, count := range pairCounts {
			elementToBeInserted := mapInsertionRules[k]
			firstNewPair := pair{k.first, elementToBeInserted}
			secondNewWPair := pair{elementToBeInserted, k.second}
			pairCountsNew[firstNewPair] += count
			pairCountsNew[secondNewWPair] += count
			elementCounts[rune(elementToBeInserted[0])] += count
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
