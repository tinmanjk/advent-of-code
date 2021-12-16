package inputParse

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReturnSliceOfLinesFromFile(filePath string) (sliceOfLines []string) {
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

func ReturnSliceOfIntsFromFile(filePath string) (sliceOfLines []int) {
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open(filePath)

	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]int, 0)
	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		// TODO better Error handling Atoi
		number, err := strconv.Atoi(strings.TrimRight(sc.Text(), "\n "))
		if err != nil {
			log.Panic(err)
		}

		lines = append(lines, number)
	}

	if err := sc.Err(); err != nil {
		log.Panic(err)
	}

	return lines
}
