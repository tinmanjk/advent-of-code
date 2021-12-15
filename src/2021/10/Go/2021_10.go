package main

import (
	"aoc/libs/go/ds"
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
	var result int

	// part 1
	result = findResult(lines, true)
	fmt.Println(result)

	// // part 2
	result = findResult(lines, false)
	fmt.Println(result)
}

// corrupted vs incomplete
// corrupted -> closes with WRONG char (}
// WHOLE line corrupted if ONE

// [({(<(())[]>[[{[]{<()<>>
// [(()[<>])]({[<{<<[]>>(
// {([(<{}[<>[]}>{[]{[(<()> // corrupted
// (((({<>}<{<{<>}{[]{[]{}
// [[<[([]))<([[{}[[()]]] // corrupted
// [{[{({}]{}}([{[{{{}}([] // corrupted
// {<[[]]>}<{[{[{[]{()[[[]
// [<(<(<(<{}))><([]([]() // corrupted
// <{([([[(<>()){}]>(<<{{ // corrupted
// <{([{{}}[<[[[<>{}]]]>[]]

func findResult(inputData []string, partOne bool) (result int) {

	incorrectRunetoScore := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	incompleteRuneToScore := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	// open and closing
	openingRunes := "([{<"
	openToCloseRune := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	scores := []int{}
	for _, line := range inputData {
		lineCorrupt := false
		closingRunes := ds.Stack{}
		for _, r := range line {
			if strings.ContainsRune(openingRunes, r) {
				closingRunes.Push(openToCloseRune[r])
				continue
			}
			expectedClosing, err := closingRunes.Pop()
			if expectedClosing != r || err != nil {
				if partOne {
					result += incorrectRunetoScore[r]
				}
				lineCorrupt = true
				break
			}
		}

		if partOne || lineCorrupt {
			continue
		}

		totalScore := 0
		for !closingRunes.IsEmpty() {
			closingRune, _ := closingRunes.Pop()
			totalScore = 5*totalScore + incompleteRuneToScore[closingRune]
		}
		scores = append(scores, totalScore)
	}

	if partOne {
		return
	}
	sort.Ints(scores)
	// odd number of scores always
	return scores[len(scores)/2]
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
