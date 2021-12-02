package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int

	result = task01(lines)
	fmt.Println(result)

	// result = task02(lines, slopes)
	fmt.Println(result)
}

type passport struct {
	byr string //(Birth Year)
	iyr string //(Issue Year)
	eyr string //(Expiration Year)
	hgt string //(Height)
	hcl string //(Hair Color)
	ecl string //(Eye Color)
	pid string //(Passport ID)
	cid string //(Country ID)
}

// passports separated by blank lines
// kvp separated by \n or " "
// cid is optional
// number of valid passports
func task01(lines []string) (valid int) {

	// issue with readfile - eat last line if empty ""
	if len(lines) > 0 && lines[len(lines)-1] != "" {
		lines = append(lines, "")
	}

	currentPassport := passport{}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			// valid++
			if currentPassport.byr != "" &&
				currentPassport.iyr != "" &&
				currentPassport.eyr != "" &&
				currentPassport.hgt != "" &&
				currentPassport.hcl != "" &&
				currentPassport.ecl != "" &&
				// cid is optional
				currentPassport.pid != "" {

				valid++
			}
			currentPassport = passport{}
			continue
		}

		parseLine(lines[i], &currentPassport)

	}

	return valid
}

func parseLine(line string, pass *passport) {
	tokens := strings.Split(line, " ")
	for _, t := range tokens {
		keyValuePair := strings.Split(t, ":")
		key := keyValuePair[0]
		value := keyValuePair[1]
		switch key {
		case "byr":
			pass.byr = value
		case "iyr":
			pass.iyr = value
		case "eyr":
			pass.eyr = value
		case "hgt":
			pass.hgt = value
		case "hcl":
			pass.hcl = value
		case "ecl":
			pass.ecl = value
		case "pid":
			pass.pid = value
		case "cid":
			pass.cid = value
		default:
			panic("no such field")
		}

	}
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

func splitLine(line string) (firstNumber int, secondNumber int,
	char rune, password string) {

	// Example line: 5-6 v: hvvgvrm
	lineSplit := strings.Split(line, " ") // should be 3
	numbers := strings.Split(lineSplit[0], "-")
	// TODO: Better Error handling
	firstNumber, err := strconv.Atoi(numbers[0])
	if err != nil {
		log.Panic(err)
	}
	secondNumber, err = strconv.Atoi(numbers[1])
	if err != nil {
		log.Panic(err)
	}

	// use if there are multi-byte unicode chars
	for _, r := range lineSplit[1] {
		char = r
		break
	}

	password = lineSplit[2]
	return
}
