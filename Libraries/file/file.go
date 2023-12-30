package file

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
)

// ReadLines reads the given location and returns a slice of strings or an error."
func ReadLines(location string) (lines []string, err error) {

	file, err := os.Open(location)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil

}

// ReadBuffer reads the given location and returns a slice of bytes or an error."
func ReadBuffer(location string) (bytes []byte, err error) {

	bytes, err = ioutil.ReadFile(location)
	if err != nil {
		return nil, err
	}

	return bytes, err

}

// ReadWords reads the given location and returns a slice of strings or an error
// Words are split by the separator provided
func ReadWords(location string, separator string) (words []string, err error) {

	file, err := os.Open(location)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	switch separator {
	case ",":
		scanner.Split(splitComma)
	case " ":
		scanner.Split(bufio.ScanWords)
	}

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil

}

// ReadNumbers reads the given location and returns a slice of strings or an error
// Words are split by the separator provided
func ReadNumbers(location string, separator string) (numbers []int, err error) {

	words, err := ReadWords(location, separator)
	if err != nil {
		return nil, err
	}

	numbers = make([]int, len(words))
	for i, w := range words {
		numbers[i], err = strconv.Atoi(w)
		if err != nil {
			return nil, err
		}
	}

	return numbers, nil

}

func splitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {

	commaidx := bytes.IndexByte(data, ',')
	if commaidx > 0 {
		buffer := data[:commaidx]
		return commaidx + 1, bytes.TrimSpace(buffer), nil
	}

	if atEOF && len(data) > 0 {
		return len(data), bytes.TrimSpace(data), nil
	}

	return 0, nil, nil

}
