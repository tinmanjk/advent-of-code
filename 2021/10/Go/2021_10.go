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

	scores := make([]int, 0)
	for _, line := range inputData {
		lineCorrupt := false
		s := make(stack, 0)
		for _, r := range line {
			if strings.ContainsRune(openingRunes, r) {
				s.Push(openToCloseRune[r])
				continue
			}
			p, _ := s.Pop()
			if p != r {
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
		for {
			p, ok := s.Pop()
			if !ok {
				break
			}
			totalScore *= 5
			totalScore += incompleteRuneToScore[p]
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

type stack []rune

func (s *stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *stack) Pop() (result rune, returned bool) {

	if len(*s) < 1 {
		return result, false
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, true
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
