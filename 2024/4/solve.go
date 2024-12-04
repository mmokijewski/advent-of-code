package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func checkHorizontal(allLines []string, i int, j int, lineLength int) bool {
	if j > lineLength-4 {
		return false
	} else if string(allLines[i][j]) == "X" && string(allLines[i][j+1]) == "M" && string(allLines[i][j+2]) == "A" && string(allLines[i][j+3]) == "S" {
		return true
	} else if string(allLines[i][j]) == "S" && string(allLines[i][j+1]) == "A" && string(allLines[i][j+2]) == "M" && string(allLines[i][j+3]) == "X" {
		return true
	}
	return false
}

func checkVertical(allLines []string, i int, j int, linesCount int) bool {
	if i > linesCount-4 {
		return false
	} else if string(allLines[i][j]) == "X" && string(allLines[i+1][j]) == "M" && string(allLines[i+2][j]) == "A" && string(allLines[i+3][j]) == "S" {
		return true
	} else if string(allLines[i][j]) == "S" && string(allLines[i+1][j]) == "A" && string(allLines[i+2][j]) == "M" && string(allLines[i+3][j]) == "X" {
		return true
	}
	return false
}

func checkDiagonalDown(allLines []string, i int, j int, lineLength int, linesCount int) bool {
	if j > lineLength-4 || i > linesCount-4 {
		return false
	} else if string(allLines[i][j]) == "X" && string(allLines[i+1][j+1]) == "M" && string(allLines[i+2][j+2]) == "A" && string(allLines[i+3][j+3]) == "S" {
		return true
	} else if string(allLines[i][j]) == "S" && string(allLines[i+1][j+1]) == "A" && string(allLines[i+2][j+2]) == "M" && string(allLines[i+3][j+3]) == "X" {
		return true
	}
	return false
}

func checkDiagonalUp(allLines []string, i int, j int, lineLength int, linesCount int) bool {
	if j > lineLength-4 || i < 3 {
		return false
	} else if string(allLines[i][j]) == "X" && string(allLines[i-1][j+1]) == "M" && string(allLines[i-2][j+2]) == "A" && string(allLines[i-3][j+3]) == "S" {
		return true
	} else if string(allLines[i][j]) == "S" && string(allLines[i-1][j+1]) == "A" && string(allLines[i-2][j+2]) == "M" && string(allLines[i-3][j+3]) == "X" {
		return true
	}
	return false
}

func checkPart2(allLines []string, i int, j int, lineLength int, linesCount int) bool {
	if j > lineLength-3 || i > linesCount-3 {
		return false
	}

	leftUp := string(allLines[i][j])
	leftDown := string(allLines[i+2][j])
	center := string(allLines[i+1][j+1])
	rightUp := string(allLines[i][j+2])
	rightDown := string(allLines[i+2][j+2])

	if center != "A" {
		return false
	} else if (leftUp == "M" && rightDown == "S") || (leftUp == "S" && rightDown == "M") {
		if (leftDown == "S" && rightUp == "M") || (leftDown == "M" && rightUp == "S") {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	matchesCount := 0
	matchesCountPart2 := 0

	lineLength := 0
	var allLines []string

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		if lineLength == 0 {
			lineLength = len(scanner.Text())
		}
		allLines = append(allLines, scanner.Text())
	}
	linesCount := len(allLines)

	for i := range allLines {
		for j := range allLines[i] {
			if checkHorizontal(allLines, i, j, lineLength) {
				matchesCount++
			}
			if checkVertical(allLines, i, j, linesCount) {
				matchesCount++
			}
			if checkDiagonalDown(allLines, i, j, lineLength, linesCount) {
				matchesCount++
			}
			if checkDiagonalUp(allLines, i, j, lineLength, linesCount) {
				matchesCount++
			}
			if checkPart2(allLines, i, j, lineLength, linesCount) {
				matchesCountPart2++
			}
		}
	}
	fmt.Printf("Matches count part1: %d\n", matchesCount)
	fmt.Printf("Matches count part2: %d\n", matchesCountPart2)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
