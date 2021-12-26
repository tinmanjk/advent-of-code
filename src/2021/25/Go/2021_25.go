package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	matrix := parseInput(lines)

	// part 1
	result := findResult(matrix)
	fmt.Println(result)
	// apparently no part 2 until 49 stars
	// part 2
	// result := findResult(cuboids)
	// fmt.Println("Part 1:", result)
}

func parseInput(lines []string) (matrix [][]rune) {

	matrix = make([][]rune, len(lines))
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		matrix[i] = []rune(line)
	}
	return
}

const inputPath = "../input.txt"

func findResult(matrix [][]rune) (result int) {

	steps := 0
	columnWidth := len(matrix[0])
	rowHeight := len(matrix)

	for {
		moves := 0
		for i := 0; i < rowHeight; i++ {

			firstMoveMarker := false
			if matrix[i][0] == '>' {
				if matrix[i][1] == '.' {
					firstMoveMarker = true
				}
			}

			for j := 1; j < columnWidth; j++ {
				if matrix[i][j] != '>' {
					continue
				}
				if j != columnWidth-1 {
					if matrix[i][j+1] == '.' {
						matrix[i][j+1] = '>'
						matrix[i][j] = '.'
						moves++
						j++
					}
				} else {
					if matrix[i][0] == '.' {
						matrix[i][0] = '>'
						matrix[i][columnWidth-1] = '.'
						moves++
					}
				}
			}

			if firstMoveMarker {
				matrix[i][0] = '.'
				matrix[i][1] = '>'
				moves++
			}
		}

		for j := 0; j < columnWidth; j++ {

			firstMoveMarker := false
			if matrix[0][j] == 'v' {
				if matrix[1][j] == '.' {
					firstMoveMarker = true
				}
			}

			for i := 1; i < rowHeight; i++ {
				if matrix[i][j] != 'v' {
					continue
				}
				if i != rowHeight-1 {
					if matrix[i+1][j] == '.' {
						matrix[i+1][j] = 'v'
						matrix[i][j] = '.'
						moves++
						i++
					}
				} else {
					if matrix[0][j] == '.' {
						matrix[0][j] = 'v'
						matrix[i][j] = '.'
						moves++
					}
				}
			}

			if firstMoveMarker {
				matrix[0][j] = '.'
				matrix[1][j] = 'v'
				moves++
			}
		}

		steps++
		if moves == 0 {
			break
		}
	}

	result = steps
	return
}
