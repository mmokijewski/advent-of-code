package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile, _ := os.Open("input")

	totalSumPart1 := 0
	totalSumPart2 := 0

	reg := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	regPart2 := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	numReg := regexp.MustCompile(`[0-9]{1,3}`)

	scanner := bufio.NewScanner(inputFile)
	do := true
	for scanner.Scan() {
		line := scanner.Text()

		matches := reg.FindAllString(line, -1)
		for i := range matches {
			numbers := numReg.FindAllString(matches[i], -1)
			leftNum, _ := strconv.Atoi(numbers[0])
			rightNum, _ := strconv.Atoi(numbers[1])
			totalSumPart1 += leftNum * rightNum
		}

		matchesPart2 := regPart2.FindAllString(line, -1)
		for i := range matchesPart2 {
			if matchesPart2[i] == "do()" {
				do = true
			} else if matchesPart2[i] == "don't()" {
				do = false
			} else if do {
				numbers := numReg.FindAllString(matchesPart2[i], -1)
				leftNum, _ := strconv.Atoi(numbers[0])
				rightNum, _ := strconv.Atoi(numbers[1])
				totalSumPart2 += leftNum * rightNum
			}
		}
	}
	fmt.Printf("Total sum part1: %d\n", totalSumPart1)
	fmt.Printf("Total sum part2: %d\n", totalSumPart2)
}
