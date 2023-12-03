package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const FILE_PATH = "input.txt"

func ReadFile() ([]string, error) {
	file, err := os.Open(FILE_PATH)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func firstDigit(num int) int {
	numStr := strconv.Itoa(num)
	firstDigit := numStr[0] - '0'
	// lastDigit := numStr[len(numStr)-1] - '0'
	return int(firstDigit)
}

func lastDigit(num int) int {
	numStr := strconv.Itoa(num)
	lastDigit := numStr[len(numStr)-1] - '0'
	return int(lastDigit)
}

func main() {
	lines, err := ReadFile()
	if err != nil {
		return
	}

	var numbers []int
	for _, line := range lines {
		regex := regexp.MustCompile(`\d+`)
		matches := regex.FindAllString(line, -1)
		if len(matches) >= 1 {
			firstMatch, _ := strconv.Atoi(matches[0])
			lastMatch, _ := strconv.Atoi(matches[len(matches)-1])
			merged := strconv.Itoa(firstDigit(firstMatch)) + strconv.Itoa(lastDigit(lastMatch))

			mergedInt, err := strconv.Atoi(merged)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}

			numbers = append(numbers, mergedInt)

		}
	}

	// sum all numbers
	var sum int
	for _, num := range numbers {
		sum += num
	}

	fmt.Println("This is the sum:", sum)

}
