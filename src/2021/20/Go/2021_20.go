package main

import (
	"aoc/libs/go/inputParse"
	"aoc/libs/go/matrixHelpers"
	"fmt"
	"log"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	imageEnhancement, inputImage := parseInput(lines)
	var result int
	// part 1
	result = findResult(imageEnhancement, inputImage, 2)
	fmt.Println(result)

	// part 2
	result = findResult(imageEnhancement, inputImage, 50)
	fmt.Println(result)
}

func parseInput(lines []string) (imageEnhancement []rune, inputImage [][]rune) {

	imageEnhancement = []rune(lines[0])
	for i := 2; i < len(lines); i++ {
		lineToRune := []rune(lines[i])
		inputImage = append(inputImage, lineToRune)
	}
	return
}

const inputPath = "../input.txt"

func findResult(imageEnhancement []rune, inputImage [][]rune, times int) (result int) {

	outputImage := [][]rune{}
	for i := 0; i < times; i++ {
		outputImage = enhanceImage(imageEnhancement, inputImage, i)
		inputImage = outputImage
	}

	counterLight := 0
	for i := 0; i < len(outputImage); i++ {
		for j := 0; j < len(outputImage[i]); j++ {
			if outputImage[i][j] == '#' {
				counterLight++
			}
		}
	}
	return counterLight
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

	// to write result here
	outputImage = make([][]rune, len(adjacentInputImage))

	paddedAdjacentInputImage := matrixHelpers.AddPaddingsWithRunes(adjacentInputImage, restOfInfinite)
	for i := 1; i < len(paddedAdjacentInputImage)-1; i++ {
		outputImage[i-1] = make([]rune, len(paddedAdjacentInputImage[i])-2) // -2  because of padding
		for j := 1; j < len(paddedAdjacentInputImage[i])-1; j++ {
			// determine number
			outPutPixel := []rune{}
			for r := i - 1; r <= i+1; r++ {
				for c := j - 1; c <= j+1; c++ {
					digit := '0'
					if paddedAdjacentInputImage[r][c] == '#' {
						digit = '1'
					}
					outPutPixel = append(outPutPixel, digit)
				}
			}

			indexFromEnhancer := convertFromBinary(outPutPixel)
			newPixel := imageEnhancement[indexFromEnhancer]
			outputImage[i-1][j-1] = newPixel
		}

	}

	return outputImage
}

func convertFromBinary(binaryString []rune) (result int) {
	for i := 0; i < len(binaryString); i++ {
		if !(binaryString[i] == '0' || binaryString[i] == '1') {
			log.Panic("Only accept 0 and 1 in string")
		}
		result <<= 1 // make place to the left
		if binaryString[i] == '1' {
			result |= 1
		}
	}
	return
}
