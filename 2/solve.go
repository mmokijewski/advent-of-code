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

func main() {
	inputFile, err := os.Open("input")
	checkError(err)

	safeLines := 0

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		arrayType := "undefined"
		for i := range items {
			if i == len(items)-1 {
				safeLines++
				continue
			}
			leftNum, err := strconv.Atoi(items[i])
			checkError(err)
			rightNum, err := strconv.Atoi(items[i+1])
			checkError(err)

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
	}
	fmt.Printf("Safe lines count: %d\n", safeLines)
}
