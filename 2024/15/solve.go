package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

func directions(name string) [2]int {
	var direction [2]int
	switch name {
	case "^":
		direction = [2]int{-1, 0}
	case "v":
		direction = [2]int{1, 0}
	case "<":
		direction = [2]int{0, -1}
	case ">":
		direction = [2]int{0, 1}
	}
	return direction
}

func checkIfBoxCanBeMoved(board [][]string, boxPos [2]int, move string, boxesToMove map[[2]int]bool) (bool, [2]int, map[[2]int]bool) {
	currY := boxPos[0]
	currX := boxPos[1]
	nextY := currY + directions(move)[0]
	nextX := currX + directions(move)[1]
	nextPos := [2]int{nextY, nextX}
	if board[currY][currX] == "." && (move == "^" || move == "v") {
		return true, boxPos, boxesToMove
	} else if board[currY][currX] == "#" || board[nextY][nextX] == "#" {
		return false, boxPos, boxesToMove
	} else if (move == "^" || move == "v") && (board[currY][currX] == "[" || board[currY][currX] == "]") {
		editedNextX := nextX
		if board[currY][currX] == "]" {
			editedNextX -= 1
		}
		movePossibleRight, _, boxesToAppendLeft := checkIfBoxCanBeMoved(board, [2]int{nextY, editedNextX + 1}, move, boxesToMove)
		movePossibleLeft, _, boxesToAppendRight := checkIfBoxCanBeMoved(board, [2]int{nextY, editedNextX}, move, boxesToMove)
		if movePossibleLeft && movePossibleRight {
			for field := range boxesToAppendLeft {
				boxesToMove[field] = true
			}
			for field := range boxesToAppendRight {
				boxesToMove[field] = true
			}
			boxesToMove[[2]int{currY, editedNextX}] = true
			return true, nextPos, boxesToMove
		} else {
			return false, boxPos, boxesToMove
		}
	} else if board[nextY][nextX] == "O" {
		return checkIfBoxCanBeMoved(board, nextPos, move, boxesToMove)
	} else if (move == "<" || move == ">") && (board[nextY][nextX] == "[" || board[nextY][nextX] == "]") {
		return checkIfBoxCanBeMoved(board, nextPos, move, boxesToMove)
	} else {
		return true, nextPos, boxesToMove
	}
}

func makeMove(board [][]string, currentPos [2]int, move string) ([][]string, [2]int) {
	currY := currentPos[0]
	currX := currentPos[1]
	nextY := currY + directions(move)[0]
	nextX := currX + directions(move)[1]
	nextPos := [2]int{nextY, nextX}
	boxesToMove := make(map[[2]int]bool)
	var movePossible bool
	var whereEmpty [2]int

	if board[nextY][nextX] == "." {
		board[nextY][nextX] = "@"
		board[currY][currX] = "."
		return board, nextPos
	} else if board[nextY][nextX] == "O" {
		movePossible, whereEmpty, boxesToMove = checkIfBoxCanBeMoved(board, nextPos, move, boxesToMove)
		if movePossible {
			board[nextY][nextX] = "@"
			board[currY][currX] = "."
			board[whereEmpty[0]][whereEmpty[1]] = "O"
			return board, nextPos
		}
	} else if board[nextY][nextX] == "[" || board[nextY][nextX] == "]" {
		if move == "<" || move == ">" {
			movePossible, whereEmpty, boxesToMove = checkIfBoxCanBeMoved(board, nextPos, move, boxesToMove)
			if movePossible {
				if move == "<" {
					for i := whereEmpty[1]; i < nextX; i++ {
						board[nextY][i] = board[nextY][i+1]
					}
				} else {
					for i := whereEmpty[1]; i > nextX; i-- {
						board[nextY][i] = board[nextY][i-1]
					}
				}
				board[nextY][nextX] = "@"
				board[currY][currX] = "."
				return board, nextPos
			}
		} else {
			movePossible, whereEmpty, boxesToMove = checkIfBoxCanBeMoved(board, nextPos, move, boxesToMove)
			if movePossible {
				var sortedBoxesToMove [][2]int
				for key := range boxesToMove {
					sortedBoxesToMove = append(sortedBoxesToMove, key)
				}
				sort.Slice(sortedBoxesToMove, func(i, j int) bool {
					if move == "^" {
						return sortedBoxesToMove[i][0] <= sortedBoxesToMove[j][0]
					} else {
						return sortedBoxesToMove[i][0] >= sortedBoxesToMove[j][0]
					}
				})

				for _, boxToMove := range sortedBoxesToMove {
					board[boxToMove[0]+directions(move)[0]][boxToMove[1]] = "["
					board[boxToMove[0]+directions(move)[0]][boxToMove[1]+1] = "]"
					board[boxToMove[0]][boxToMove[1]] = "."
					board[boxToMove[0]][boxToMove[1]+1] = "."
				}
				board[nextY][nextX] = "@"
				board[currY][currX] = "."

				return board, nextPos
			}
		}
	}
	return board, currentPos
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	var board [][]string
	var boardPart2 [][]string
	var moves []string
	var currentPos [2]int
	var currentPosPart2 [2]int

	scanner := bufio.NewScanner(inputFile)
	y := 0
	scanMoves := false
	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []string
		var lineArrayPart2 []string
		for x, field := range line {
			lineArray = append(lineArray, string(field))
			if string(field) == "@" {
				currentPos = [2]int{y, x}
				currentPosPart2 = [2]int{y, 2 * x}
				lineArrayPart2 = append(lineArrayPart2, "@", ".")
			} else if string(field) == "O" {
				lineArrayPart2 = append(lineArrayPart2, "[", "]")
			} else {
				lineArrayPart2 = append(lineArrayPart2, string(field), string(field))
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
			boardPart2 = append(boardPart2, lineArrayPart2)
		}
		y++
	}

	for _, move := range moves {
		board, currentPos = makeMove(board, currentPos, move)
		boardPart2, currentPosPart2 = makeMove(boardPart2, currentPosPart2, move)
	}

	part1Sum := 0
	for y, line := range board {
		for x, field := range line {
			if field == "O" {
				part1Sum += 100*y + x
			}
		}
	}
	part2Sum := 0
	for y, line := range boardPart2 {
		for x, field := range line {
			if field == "[" {
				part2Sum += 100*y + x
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Part 2: %d\n", part2Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
