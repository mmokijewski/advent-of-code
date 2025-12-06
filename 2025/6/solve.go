package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	var operations []string
	var startIndexes []int

	var part1Results []int
	var part2StringNumbers [][]string

	// Go through the file 1 time to collect data about operations and start indexes
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "*") && !strings.HasPrefix(line, "+") {
			continue
		}
		for i, char := range line {
			if string(char) == "*" || string(char) == "+" {
				operations = append(operations, string(char))
				startIndexes = append(startIndexes, i)
			}
		}
	}

	for range len(operations) {
		part2StringNumbers = append(part2StringNumbers, make([]string, len(operations)))
	}

	// Go through the file second time to count part 1 and prepare numbers for part 2
	inputFile, _ = os.Open("input")
	scanner = bufio.NewScanner(inputFile)
	reg := regexp.MustCompile("\\d+")
	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "*") || strings.HasPrefix(line, "+") {
			continue
		}

		for i, startIndex := range startIndexes {
			var numString string
			if i == len(startIndexes)-1 {
				numString = line[startIndex:]
			} else {
				numString = line[startIndex : startIndexes[i+1]-1]
			}
			num, _ := strconv.Atoi(reg.FindString(numString))
			if len(part1Results) < len(operations) {
				part1Results = append(part1Results, num)
			} else {
				if operations[i] == "+" {
					part1Results[i] = part1Results[i] + num
				} else if operations[i] == "*" {
					part1Results[i] = part1Results[i] * num
				}
			}

			// Prepare numbers for part 2
			numLength := len(numString)
			for j := 0; j < numLength; j++ {
				currentChar := string(numString[numLength-1-j])
				if currentChar == " " {
					continue
				} else {
					part2StringNumbers[i][j] = part2StringNumbers[i][j] + currentChar
				}
			}
		}
		lineIndex++
	}

	// Count part 1 result
	part1sum := 0
	for _, result := range part1Results {
		part1sum += result
	}

	// Go through collected numbers for part 2
	part2sum := 0
	for i, column := range part2StringNumbers {
		columnResult := 0
		for _, stringNumber := range column {
			num, _ := strconv.Atoi(stringNumber)
			if num == 0 {
				continue
			}
			if columnResult == 0 {
				columnResult = num
			} else {
				if operations[i] == "+" {
					columnResult = columnResult + num
				} else if operations[i] == "*" {
					columnResult = columnResult * num
				}
			}
		}
		part2sum += columnResult
	}

	fmt.Printf("Part1 : %d\n", part1sum)
	fmt.Printf("Part2 : %d\n", part2sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
