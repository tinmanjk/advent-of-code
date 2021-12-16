package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int64
	byteArray := parseInput(lines)

	// part 1
	result = findResult(byteArray, false)
	fmt.Println(result)

	// part 2
	result = findResult(byteArray, true)
	fmt.Println(result)
}

func parseInput(slicesOfLines []string) (binaryArray []rune) {

	line := slicesOfLines[0]

	binaryArray = []rune{}

	for i := 0; i < len(line); i++ {
		number := 0
		hex := line[i]

		switch {
		case '0' <= hex && hex <= '9':
			number = int(hex - '0')
		case 'A' <= hex && hex <= 'F':
			number = int(hex-'A') + 10
		}
		binaryConvertedHex := []rune(fmt.Sprintf("%04b", number))
		binaryArray = append(binaryArray, binaryConvertedHex...)
	}

	return
}

const inputPath = "../input.txt"
const invalidEvaluation = -1

func findResult(binaryArray []rune, partTwo bool) (result int64) {
	rootPacket, _ := createPacketFromBinary(binaryArray, 0, &result)

	if partTwo {
		result = rootPacket.Evaluation
	}
	return
}

type packet struct {
	Version           int64
	Type              int64
	LengthTypeId      int64
	SubPackets        []packet
	TotalLengthInBits int64 // length id = 0
	NumberSubpackets  int64 // length id = 1
	Evaluation        int64
	TypeFunc          func([]int64) int64
}

func createPacketFromBinary(binaryArray []rune, startIndex int, versionSumAccumulator *int64) (pack packet, consumedBits int) {

	pack = packet{}
	pack.Evaluation = invalidEvaluation

	lowBound := startIndex
	upperBound := startIndex + 3
	versionBinary := binaryArray[lowBound:upperBound]
	consumedBits += upperBound - lowBound
	pack.Version, _ = strconv.ParseInt(string(versionBinary), 2, 64)
	*versionSumAccumulator += pack.Version

	lowBound = startIndex + int(consumedBits)
	upperBound = lowBound + 3
	packTypeBinary := binaryArray[lowBound:upperBound]
	consumedBits += upperBound - lowBound
	pack.Type, _ = strconv.ParseInt(string(packTypeBinary), 2, 64)

	switch pack.Type {
	case 4:
		consumedBits += consumeLiteralPacket(&pack, binaryArray, startIndex+consumedBits)
		return // leaf
	case 0:
		pack.TypeFunc = sum
	case 1:
		pack.TypeFunc = product
	case 2:
		pack.TypeFunc = min
	case 3:
		pack.TypeFunc = max
	case 5:
		pack.TypeFunc = greaterThan
	case 6:
		pack.TypeFunc = lessThan
	case 7:
		pack.TypeFunc = equalTo
	}

	pack.LengthTypeId, _ = strconv.ParseInt(string(binaryArray[startIndex+consumedBits]), 2, 64)
	consumedBits++

	var numberBinarySlice []rune
	switch {
	case pack.LengthTypeId == 0:
		lowBound = startIndex + consumedBits
		upperBound = lowBound + 15
		numberBinarySlice = binaryArray[lowBound:upperBound]
		consumedBits += 15 // should be 22 bits
		pack.TotalLengthInBits, _ = strconv.ParseInt(string(numberBinarySlice), 2, 64)

		for totalChildrenConsumption := 0; totalChildrenConsumption < int(pack.TotalLengthInBits); {
			childPack, childConsumedBits := createPacketFromBinary(binaryArray, startIndex+consumedBits, versionSumAccumulator)
			consumedBits += childConsumedBits
			pack.SubPackets = append(pack.SubPackets, childPack)
			totalChildrenConsumption += childConsumedBits
		}
		// we are back in the recursion, so values have been calculated for all subpackets already
		evaluatePacket(&pack)

	case pack.LengthTypeId == 1:
		lowBound = consumedBits + startIndex
		upperBound = lowBound + 11
		numberBinarySlice = binaryArray[lowBound:upperBound]
		consumedBits += 11 // should be 18 bits
		pack.NumberSubpackets, _ = strconv.ParseInt(string(numberBinarySlice), 2, 64)

		for len(pack.SubPackets) < int(pack.NumberSubpackets) {
			childPack, childConsumedBits := createPacketFromBinary(binaryArray, startIndex+consumedBits, versionSumAccumulator)
			consumedBits += childConsumedBits
			pack.SubPackets = append(pack.SubPackets, childPack)
		}
		// we are back in the recursion, so values have been calculated for all subpackets already
		evaluatePacket(&pack)
	}

	return
}

func consumeLiteralPacket(pack *packet, binaryArray []rune, startIndex int) (consumedBits int) {

	number := []rune{}
	for i := 0; ; i++ {
		lowBound := startIndex + i*5
		lastFourBitGroup := binaryArray[lowBound] == '0'
		// first bit reading
		consumedBits++
		lowBound++
		number = append(number, binaryArray[lowBound:lowBound+4]...)
		consumedBits += 4
		if lastFourBitGroup {
			break
		}
	}
	pack.Evaluation, _ = strconv.ParseInt(string(number), 2, 64)
	return
}

func evaluatePacket(pack *packet) {
	arrayOfEvaluations := []int64{}
	for _, sp := range pack.SubPackets {
		arrayOfEvaluations = append(arrayOfEvaluations, sp.Evaluation)
	}
	pack.Evaluation = pack.TypeFunc(arrayOfEvaluations)
}

func sum(nums []int64) (result int64) {

	for _, v := range nums {
		result += v
	}
	return
}

func product(nums []int64) (result int64) {

	if len(nums) == 1 {
		return nums[0]
	}

	result = 1
	for _, v := range nums {
		result *= v
	}
	return
}

func min(nums []int64) (result int64) {

	result = math.MaxInt64
	for _, v := range nums {
		if result >= v {
			result = v
		}
	}
	return
}

func max(nums []int64) (result int64) {

	result = math.MinInt64
	for _, v := range nums {
		if result <= v {
			result = v
		}
	}
	return
}

func greaterThan(nums []int64) (result int64) {

	if nums[0] > nums[1] {
		return 1
	}
	return 0
}

func lessThan(nums []int64) (result int64) {

	if nums[0] < nums[1] {
		return 1
	}
	return 0
}

func equalTo(nums []int64) (result int64) {

	if nums[0] == nums[1] {
		return 1
	}
	return 0
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
