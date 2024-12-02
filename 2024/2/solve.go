package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkIfDescending(leftNum int, rightNum int) bool {
	if leftNum > rightNum && leftNum <= rightNum+3 {
		return true
	} else {
		return false
	}
}

func checkIfAscending(leftNum int, rightNum int) bool {
	if leftNum < rightNum && leftNum+3 >= rightNum {
		return true
	} else {
		return false
	}
}

func checkIfSafe(numbers []int) bool {
	arrayType := "undefined"
	safeLine := false
	for i := range numbers {
		if i == len(numbers)-1 {
			safeLine = true
			continue
		}
		leftNum := numbers[i]
		rightNum := numbers[i+1]

		if arrayType == "undefined" {
			if checkIfDescending(leftNum, rightNum) {
				arrayType = "descending"
			} else if checkIfAscending(leftNum, rightNum) {
				arrayType = "ascending"
			} else {
				break
			}
		}

		if arrayType == "descending" {
			if !checkIfDescending(leftNum, rightNum) {
				break
			}
		} else if arrayType == "ascending" {
			if !checkIfAscending(leftNum, rightNum) {
				break
			}
		}
	}
	return safeLine
}

func main() {
	inputFile, err := os.Open("input")
	checkError(err)

	safeLines := 0
	safeLinesPart2 := 0

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		var numbers []int
		for i := range items {
			num, err := strconv.Atoi(items[i])
			checkError(err)
			numbers = append(numbers, num)
		}
		if checkIfSafe(numbers) {
			safeLines++
			safeLinesPart2++
		} else {
			for i := range numbers {
				var newNumbers []int
				newNumbers = append(newNumbers, numbers...)
				newNumbers = append(newNumbers[:i], newNumbers[i+1:]...)
				if checkIfSafe(newNumbers) {
					safeLinesPart2++
					break
				}
			}
		}
	}
	fmt.Printf("Safe lines count: %d\n", safeLines)
	fmt.Printf("Safe lines count part2: %d\n", safeLinesPart2)
}
