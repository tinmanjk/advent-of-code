package main

import (
	"aoc/libs/go/inputParse"
	"aoc/libs/go/matrixHelpers"
	"fmt"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	imageEnhancement, inputImage := parseInput(lines)
	var result int
	// part 1
	result = findResult(imageEnhancement, inputImage)
	fmt.Println(result)

	// // part 2
	// result = findResult(lines)
	// fmt.Println(result)
}

func parseInput(lines []string) (imageEnhancement []rune, inputImage [][]rune) {

	imageEnhancement = []rune(lines[0])
	for i := 2; i < len(lines); i++ {
		lineToRune := []rune(lines[i])
		inputImage = append(inputImage, lineToRune)
	}
	return
}

const inputPath = "../input0.txt"

func findResult(imageEnhancement []rune, inputImage [][]rune) (result int) {

	outputImage := [][]rune{}
	for i := 0; i < 10; i++ {
		outputImage = enhanceImage(imageEnhancement, inputImage, i)
	}

	return len(outputImage)
}

func determineRestOfInfinite(time int, imageEnhancement []rune) (restOfInfinite rune) {
	if time == 0 {
		restOfInfinite = '.'
	}

	if time == 1 {
		restOfInfinite = imageEnhancement[0] // sum of ... ... .. = 0
	}

	if time >= 2 {
		if imageEnhancement[0] == '.' {
			restOfInfinite = '.' // stays there forever
		}

		if imageEnhancement[0] == '#' && imageEnhancement[len(imageEnhancement)-1] == '#' {
			restOfInfinite = '#' // stays there forever
		}

		// alternating
		if imageEnhancement[0] == '#' && imageEnhancement[len(imageEnhancement)-1] == '.' {
			if time%2 == 0 {
				restOfInfinite = imageEnhancement[len(imageEnhancement)-1] // '.'
			} else {
				restOfInfinite = imageEnhancement[0] // '#'
			}
		}
	}

	return
}

func enhanceImage(imageEnhancement []rune, inputImage [][]rune, time int) (outputImage [][]rune) {

	restOfInfinite := determineRestOfInfinite(time, imageEnhancement)

	adjacentInputImage := matrixHelpers.AddPaddingsWithRunes(inputImage, restOfInfinite)

	return adjacentInputImage
}
