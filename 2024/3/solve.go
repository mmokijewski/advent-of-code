package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	inputFile, err := os.Open("input")
	checkError(err)

	totalSumPart1 := 0

	reg := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	numReg := regexp.MustCompile(`[0-9]{1,3}`)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		matches := reg.FindAllString(line, -1)

		for i := range matches {
			numbers := numReg.FindAllString(matches[i], -1)
			leftNum, err := strconv.Atoi(numbers[0])
			checkError(err)
			rightNum, err := strconv.Atoi(numbers[1])
			checkError(err)
			totalSumPart1 += leftNum * rightNum
		}
	}
	fmt.Printf("Total sum part1: %d\n", totalSumPart1)
}
