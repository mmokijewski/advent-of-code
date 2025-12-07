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

type point struct {
	x, y int
}

func findTimelines(board [][]string, timelines map[point]int) (map[point]int, bool) {
	newTimelines := make(map[point]int)
	var finished bool
	for currentPos, count := range timelines {
		finished = currentPos.y >= len(board)-1
		if finished {
			return timelines, true
		}
		nextPos := board[currentPos.y+1][currentPos.x]
		if nextPos == "|" {
			newTimelines[point{currentPos.x, currentPos.y + 1}] += count
		} else if nextPos == "^" {
			newTimelines[point{currentPos.x - 1, currentPos.y + 1}] += count
			newTimelines[point{currentPos.x + 1, currentPos.y + 1}] += count
		}
	}
	return newTimelines, finished
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	part1sum := 0
	var board [][]string
	timelines := make(map[point]int)

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
			if char == "S" {
				// Add initial timeline for part 2
				timelines[point{j, i}] = 1
			}
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

	// Part 2
	finished := false
	for !finished {
		timelines, finished = findTimelines(board, timelines)
	}

	timelinesCount := 0
	for _, count := range timelines {
		timelinesCount += count
	}

	fmt.Printf("Part1 : %d\n", part1sum)
	fmt.Printf("Part2 : %d\n", timelinesCount)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
