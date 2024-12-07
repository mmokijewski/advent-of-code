package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(expectedResult int, numbers []int, currentIndex int, currentResult int) bool {
	if currentIndex == len(numbers) {
		return currentResult == expectedResult
	}

	tempResult := currentResult + numbers[currentIndex]
	if tempResult <= expectedResult {
		if check(expectedResult, numbers, currentIndex+1, tempResult) {
			return true
		}
	}

	tempResult = currentResult * numbers[currentIndex]
	if tempResult <= expectedResult {
		if check(expectedResult, numbers, currentIndex+1, tempResult) {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	var part1Sum int

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		splitByColon := strings.Split(line, ": ")
		expectedResult, _ := strconv.Atoi(splitByColon[0])
		var numbers []int
		for _, num := range strings.Split(splitByColon[1], " ") {
			number, _ := strconv.Atoi(num)
			numbers = append(numbers, number)
		}
		if check(expectedResult, numbers, 0, 0) {
			part1Sum += expectedResult
		}

	}

	fmt.Printf("Part 1 sum: %d\n", part1Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
