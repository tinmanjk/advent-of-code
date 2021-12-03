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

	// result = task01(lines)
	// fmt.Println(result)

	result = task02(lines)
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

	// issue with readfile - eat last line if empty
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

		parseLine1(lines[i], &currentPassport)

	}

	return valid
}

func task02(lines []string) (valid int) {

	// issue with readfile - eat last line if empty
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

		parseLine2(lines[i], &currentPassport)

	}

	return valid
}

func parseLine1(line string, pass *passport) {
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
		case "cid": // ignored
			pass.cid = value
		default:
			panic("no such field")
		}

	}
}

func parseLine2(line string, pass *passport) {
	tokens := strings.Split(line, " ")
	for _, t := range tokens {
		keyValuePair := strings.Split(t, ":")
		key := keyValuePair[0]
		value := keyValuePair[1]
		switch key {
		case "byr":
			if conv, err := strconv.Atoi(value); err == nil &&
				1920 <= conv && conv <= 2002 {
				pass.byr = value
			} else if err != nil {
				log.Panic(err)
			}
		case "iyr":
			if conv, err := strconv.Atoi(value); err == nil &&
				2010 <= conv && conv <= 2020 {
				pass.iyr = value
			} else if err != nil {
				log.Panic(err)
			}
		case "eyr":
			if conv, err := strconv.Atoi(value); err == nil &&
				2020 <= conv && conv <= 2030 {
				pass.eyr = value
			} else if err != nil {
				log.Panic(err)
			}
		case "hgt":
			if len(value) < 3 {
				break
			}
			unit := value[len(value)-2:]
			measure := value[:len(value)-2]
			switch unit {
			case "cm":
				if conv, err := strconv.Atoi(measure); err == nil &&
					150 <= conv && conv <= 193 {
					pass.hgt = value
				} else if err != nil {
					log.Panic(err)
				}
			case "in":
				if conv, err := strconv.Atoi(measure); err == nil &&
					59 <= conv && conv <= 76 {
					pass.hgt = value
				} else if err != nil {
					log.Panic(err)
				}
			}
		case "hcl":
			if value[0] == '#' {
				validHex := true
				for _, r := range value[1:] {
					if !('0' <= r && r <= '9' || 'a' <= r && r <= 'f') {
						validHex = false
						break
					}
				}
				if validHex {
					pass.hcl = value

				}
			}
		case "ecl":
			switch value {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
				pass.ecl = value
			}
		case "pid":
			if len(value) == 9 {
				pass.pid = value
			}
		case "cid": // ignored
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
