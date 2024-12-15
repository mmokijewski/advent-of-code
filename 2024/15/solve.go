package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func checkIfBoxCanBeMoved(board [][]string, boxPos [2]int, dir [2]int) (bool, [2]int) {
	currY := boxPos[0]
	currX := boxPos[1]
	nextY := currY + dir[0]
	nextX := currX + dir[1]

	if board[nextY][nextX] == "#" {
		return false, [2]int{-1, -1}
	} else if board[nextY][nextX] == "O" {
		return checkIfBoxCanBeMoved(board, [2]int{nextY, nextX}, dir)
	} else {
		return true, [2]int{nextY, nextX}
	}
}

func makeMove(board [][]string, currentPos [2]int, move string) ([][]string, [2]int) {
	directions := map[string][2]int{
		"^": {-1, 0}, //up
		"v": {1, 0},  //down
		"<": {0, -1}, //left
		">": {0, 1},  //right
	}
	currY := currentPos[0]
	currX := currentPos[1]
	nextY := currY + directions[move][0]
	nextX := currX + directions[move][1]
	nextPos := [2]int{nextY, nextX}

	if board[nextY][nextX] == "." {
		board[nextY][nextX] = "@"
		board[currY][currX] = "."
		return board, nextPos
	} else if board[nextY][nextX] == "O" {
		movePossible, whereEmpty := checkIfBoxCanBeMoved(board, nextPos, directions[move])
		if movePossible {
			board[nextY][nextX] = "@"
			board[currY][currX] = "."
			board[whereEmpty[0]][whereEmpty[1]] = "O"
			return board, nextPos
		}
	}
	return board, currentPos
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	var board [][]string
	var moves []string
	var currentPos [2]int

	scanner := bufio.NewScanner(inputFile)
	y := 0
	scanMoves := false
	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []string
		for x, field := range line {
			lineArray = append(lineArray, string(field))
			if string(field) == "@" {
				currentPos = [2]int{y, x}
			}
		}
		if len(lineArray) == 0 {
			scanMoves = true
			continue
		}
		if scanMoves {
			moves = append(moves, lineArray...)
		} else {
			board = append(board, lineArray)
		}
		y++
	}

	for _, move := range moves {
		board, currentPos = makeMove(board, currentPos, move)
	}

	part1Sum := 0
	for y, line := range board {
		for x, field := range line {
			if field == "O" {
				part1Sum += 100*y + x
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
