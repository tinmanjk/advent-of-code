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

	// 512 chars long
	// line 1 = image inhancement algorithm
	// can match each 9bit binary number
	// 2 to the 9

	// image
	// two-dimensional array
	// light pixels - #
	// dark pixels - .
	for i := 2; i < len(lines); i++ {
	}
	return
}

const inputPath = "../input0.txt"

func findResult(imageEnhancement []rune, inputImage [][]rune) (result int) {

	outputImage := [][]rune{}
	for i := 0; i < 2; i++ {
		outputImage = enhanceImage(imageEnhancement, inputImage)
	}
	// pad TWICE...i taka
	// padding sas sigurnost
	// za da stane INFINITE IMAGE..bla

	return len(outputImage)
}

func enhanceImage(imageEnhancement []rune, inputImage [][]rune) (outputImage [][]rune) {
	wipImage := matrixHelpers.AddPaddingsWithRunes(inputImage, '.')

	return wipImage
}
