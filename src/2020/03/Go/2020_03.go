package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int

	// part 1
	slp := slope{3, 1}
	result = task01(lines, slp)
	fmt.Println(result)

	// part 2
	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result = task02(lines, slopes)
	fmt.Println(result)
}

type slope struct {
	hor int // right is positive
	ver int // down is positive
}

func task01(lines []string, slp slope) (numberTrees int) {
	hor := 0
	ver := 0
	for ver < len(lines) {
		line := lines[ver]
		currentWidth := len(line)

		if hor >= currentWidth {
			hor = hor % currentWidth
		}
		if line[hor] == '#' {
			numberTrees++
		}
		hor += slp.hor
		ver += slp.ver
	}

	return numberTrees
}

func task02(lines []string, slopes []slope) (multiplication int) {
	multiplication = 1
	for _, s := range slopes {
		multiplication *= task01(lines, s)
	}

	return multiplication
}

func returnSliceOfLinesFromFile(filePath string) (sliceOfLines []string) {
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open(filePath)

	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		lines = append(lines, strings.TrimRight(sc.Text(), "\n "))
	}

	if err := sc.Err(); err != nil {
		log.Panic(err)
	}

	return lines
}
