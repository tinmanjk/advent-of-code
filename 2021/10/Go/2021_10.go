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

// line - 1 or more chunks, chunks - zero or more chunks
// no delimiter separation
// open close - delimiter basically
// ( ) [ ] { } <>
// () empty chunk

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

	incorectCharMap := make(map[rune]int, 4)
	incorectCharMap[')'] = 3
	incorectCharMap[']'] = 57
	incorectCharMap['}'] = 1197
	incorectCharMap['>'] = 25137

	incompleteCharMap := make(map[rune]int, 4)
	incompleteCharMap[')'] = 1
	incompleteCharMap[']'] = 2
	incompleteCharMap['}'] = 3
	incompleteCharMap['>'] = 4

	// open and closing
	openingRunes := "([{<"
	closingRunes := make(map[rune]rune, 4)
	closingRunes['('] = ')'
	closingRunes['['] = ']'
	closingRunes['{'] = '}'
	closingRunes['<'] = '>'

	scores := make([]int, 0)
	for _, line := range inputData {
		lineCorrupt := false
		s := make(stack, 0)
		for _, r := range line {
			if strings.ContainsRune(openingRunes, r) {
				s.Push(closingRunes[r])
				continue
			}
			p, _ := s.Pop()
			if p != r {
				if partOne {
					result += incorectCharMap[r]
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
			totalScore += incompleteCharMap[p]
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
