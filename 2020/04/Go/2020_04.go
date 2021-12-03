package main

import (
	"bufio"
	"errors"
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

	result = task1(lines)
	fmt.Println(result)

	result = task2(lines)
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

func (p passport) hasAllRequiredFields() bool {
	if p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		// cid is optional
		p.pid != "" {
		return true
	}
	return false
}

func (p *passport) setField(key string, value string, withFullValidation bool) error {
	if key == "" || value == "" {
		return errors.New("Key or value is empty")
	}

	switch key {
	case "byr":
		if !withFullValidation {
			p.byr = value
			return nil
		}
		if conv, err := strconv.Atoi(value); err == nil &&
			1920 <= conv && conv <= 2002 {
			p.byr = value
		} else if err != nil {
			log.Panic(err)
		}
	case "iyr":
		if !withFullValidation {
			p.iyr = value
			return nil
		}
		if conv, err := strconv.Atoi(value); err == nil &&
			2010 <= conv && conv <= 2020 {
			p.iyr = value
		} else if err != nil {
			log.Panic(err)
		}
	case "eyr":
		if !withFullValidation {
			p.eyr = value
			return nil
		}
		if conv, err := strconv.Atoi(value); err == nil &&
			2020 <= conv && conv <= 2030 {
			p.eyr = value
		} else if err != nil {
			log.Panic(err)
		}
	case "hgt":
		if !withFullValidation {
			p.hgt = value
			return nil
		}
		if len(value) < 3 {
			break
		}
		unit := value[len(value)-2:]
		measure := value[:len(value)-2]
		switch unit {
		case "cm":
			if conv, err := strconv.Atoi(measure); err == nil &&
				150 <= conv && conv <= 193 {
				p.hgt = value
			} else if err != nil {
				log.Panic(err)
			}
		case "in":
			if conv, err := strconv.Atoi(measure); err == nil &&
				59 <= conv && conv <= 76 {
				p.hgt = value
			} else if err != nil {
				log.Panic(err)
			}
		}
	case "hcl":
		if !withFullValidation {
			p.hcl = value
			return nil
		}
		if value[0] == '#' {
			validHex := true
			for _, r := range value[1:] {
				if !('0' <= r && r <= '9' || 'a' <= r && r <= 'f') {
					validHex = false
					break
				}
			}
			if validHex {
				p.hcl = value

			}
		}
	case "ecl":
		if !withFullValidation {
			p.ecl = value
			return nil
		}
		switch value {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			p.ecl = value
		}
	case "pid":
		if !withFullValidation {
			p.pid = value
			return nil
		}
		if len(value) == 9 {
			p.pid = value
		}
	case "cid": // ignored
		p.cid = value
	default:
		panic("no such field")
	}
	return nil
}

func task1(lines []string) (valid int) {
	return parseAllLines(lines, false)
}

func task2(lines []string) (valid int) {
	return parseAllLines(lines, true)
}

// passports separated by blank lines
// kvp separated by \n or " "
// cid is optional
func parseAllLines(lines []string, withFullValidation bool) (valid int) {

	// issue with readfile - eat last line if empty
	if len(lines) > 0 && lines[len(lines)-1] != "" {
		lines = append(lines, "")
	}

	currentPassport := passport{}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			if currentPassport.hasAllRequiredFields() {
				valid++
			}
			currentPassport = passport{}
			continue
		}
		parseSingleLine(lines[i], &currentPassport, withFullValidation)
	}

	return valid
}

// TODO Proper Error-handling
func parseSingleLine(line string, pass *passport, withFullValidation bool) {
	tokens := strings.Split(line, " ")
	for _, t := range tokens {
		keyValuePair := strings.Split(t, ":")
		pass.setField(keyValuePair[0], keyValuePair[1], withFullValidation)
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
