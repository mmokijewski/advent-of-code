package main

import (
	"bufio"
	"fmt"
	"os"
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

	part1sum := 0
	var board [][]string

	// Prepare board
	scanner := bufio.NewScanner(inputFile)
	i := 0
	for scanner.Scan() {
		board = append(board, make([]string, 0))
		line := scanner.Text()
		split := strings.Split(line, "")
		for _, char := range split {
			board[i] = append(board[i], char)
		}
		i++
	}

	// Go through the prepared board
	for i, line := range board {
		if i == len(board)-1 {
			continue
		}
		for j, char := range line {
			if char == "S" || char == "|" {
				if board[i+1][j] == "^" {
					board[i+1][j-1] = "|"
					board[i+1][j+1] = "|"
					part1sum++
				} else {
					board[i+1][j] = "|"
				}
			}
		}
	}

	fmt.Printf("Part1 : %d\n", part1sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
