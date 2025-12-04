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

type Point struct {
	y, x int
}

var directions = []Point{
	{-1, 0},  // up
	{0, 1},   // right
	{1, 0},   // down
	{0, -1},  // left
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	var board [][]string

	part1 := 0
	scanner := bufio.NewScanner(inputFile)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		var currentLineArray []string
		for _, sign := range strings.Split(line, "") {
			currentLineArray = append(currentLineArray, sign)
		}
		board = append(board, currentLineArray)
		y++
	}
	maxY := len(board) - 1
	maxX := len(board[0]) - 1

	for y := range board {
		for x, sign := range board[y] {
			if sign == "." {
				continue
			}
			dirCount := 0
			for _, dir := range directions {
				newPoint := Point{y + dir.y, x + dir.x}
				if newPoint.x < 0 || newPoint.x > maxX || newPoint.y < 0 || newPoint.y > maxY {
					continue
				}
				if board[newPoint.y][newPoint.x] == "@" {
					dirCount++
				}
				if dirCount >= 4 {
					break
				}
			}
			if dirCount < 4 {
				part1++
			}
		}
	}

	fmt.Printf("Part1 : %d\n", part1)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
