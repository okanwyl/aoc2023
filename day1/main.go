package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const FILE_PATH = "input.txt"

var DIGITMAP = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

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

func FirstDigit(num int) int {
	numStr := strconv.Itoa(num)
	firstDigit := numStr[0] - '0'
	return int(firstDigit)
}

func LastDigit(num int) int {
	numStr := strconv.Itoa(num)
	lastDigit := numStr[len(numStr)-1] - '0'
	return int(lastDigit)
}

func ReplaceSpelledDigits(line string) string {
	re := regexp.MustCompile(`[a-zA-Z0-9]+`)
	return re.ReplaceAllStringFunc(line, func(match string) string {
		for i := 0; i < len(match); i++ {
			for spelledDigit, digit := range DIGITMAP {
				if strings.HasPrefix(match[i:], spelledDigit) {
					before := match[:i]
					after := match[i+len(spelledDigit):]
					match = before + strconv.Itoa(digit) + after
					// Update the index to skip past the newly replaced number
					i += len(strconv.Itoa(digit)) - 1
					break
				}
			}
		}
		return match
	})
}

func main() {
	lines, err := ReadFile()
	if err != nil {
		return
	}

	var numbers []int
	for _, line := range lines {
		regex := regexp.MustCompile(`\d+`)
		fmt.Println("This is the line:", line)
		replacedLine := ReplaceSpelledDigits(line)
		fmt.Println("This is the replaced line:", replacedLine)
		matches := regex.FindAllString(replacedLine, -1)
		if len(matches) >= 1 {
			firstMatch, _ := strconv.Atoi(matches[0])
			lastMatch, _ := strconv.Atoi(matches[len(matches)-1])
			merged := strconv.Itoa(FirstDigit(firstMatch)) + strconv.Itoa(LastDigit(lastMatch))

			mergedInt, err := strconv.Atoi(merged)
			if err != nil {
				return
			}

			numbers = append(numbers, mergedInt)

		}
	}

	var sum int
	for _, num := range numbers {
		sum += num
	}

	fmt.Println("This is the sum:", sum)
}
