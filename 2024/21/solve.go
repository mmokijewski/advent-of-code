package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Point struct {
	y, x int
}

type Direction struct {
	y, x   int
	symbol string
}

var directions = []Direction{
	{0, -1, "<"},
	{1, 0, "v"},
	{-1, 0, "^"},
	{0, 1, ">"},
}

func findCheapestPathOnKeypad(keypad [][]string, start string, end string) string {
	var startPos Point
	var endPos Point
	for i, line := range keypad {
		for j, key := range line {
			if key == start {
				startPos = Point{i, j}
			} else if key == end {
				endPos = Point{i, j}
			}
		}
	}
	moves := make(map[string]int)

	newPos := startPos
	for keypad[newPos.y][newPos.x] != end {
		for _, dir := range directions {
			tempPos := Point{newPos.y + dir.y, newPos.x + dir.x}
			if tempPos.y < 0 || tempPos.y >= len(keypad) || tempPos.x < 0 || tempPos.x >= len(keypad[0]) {
				continue
			}

			if keypad[tempPos.y][tempPos.x] == "X" {
				continue
			}
			if newPos.x == endPos.x && (dir.symbol == "<" || dir.symbol == ">") {
				continue
			}
			if newPos.y == endPos.y && (dir.symbol == "v" || dir.symbol == "^") {
				continue
			}
			if newPos.x < endPos.x && dir.symbol == "<" {
				continue
			}
			if newPos.y > endPos.y && dir.symbol == "v" {
				continue
			}
			moves[dir.symbol]++
			newPos = tempPos
			break
		}
	}

	result := ""
	nextPos := startPos
	for keypad[nextPos.y][nextPos.x] != end {
		for _, dir := range directions {
			if moves[dir.symbol] > 0 {
				movesAmount := moves[dir.symbol]
				tempNextPos := Point{nextPos.y + movesAmount*dir.y, nextPos.x + movesAmount*dir.x}
				if keypad[tempNextPos.y][tempNextPos.x] != "X" {
					moves[dir.symbol] = 0
					for range movesAmount {
						result += dir.symbol
					}
					nextPos = tempNextPos
				}
			}
		}
	}
	return result + "A"
}

func countCodeResult(code string, robotsCount int) int {
	numericKeypad := make([][]string, 4)
	numericKeypad[0] = append(numericKeypad[0], "7", "8", "9")
	numericKeypad[1] = append(numericKeypad[1], "4", "5", "6")
	numericKeypad[2] = append(numericKeypad[2], "1", "2", "3")
	numericKeypad[3] = append(numericKeypad[3], "X", "0", "A")

	dirKeypad := make([][]string, 2)
	dirKeypad[0] = append(dirKeypad[0], "X", "^", "A")
	dirKeypad[1] = append(dirKeypad[1], "<", "v", ">")

	codeNumPart, _ := strconv.Atoi(code[:len(code)-1])

	lastRobotMoves := make(map[string]int)
	lastRobotMoves[code]++

	for r := range robotsCount + 1 {
		thisRobotMoves := make(map[string]int)

		for move, count := range lastRobotMoves {
			move = "A" + move
			for i, sign := range move {
				if i == len(move)-1 {
					break
				}
				if r == 0 {
					newMove := findCheapestPathOnKeypad(numericKeypad, string(sign), string(move[i+1]))
					thisRobotMoves[newMove] += count
				} else {
					newMove := findCheapestPathOnKeypad(dirKeypad, string(sign), string(move[i+1]))
					thisRobotMoves[newMove] += count
				}
			}
		}
		lastRobotMoves = make(map[string]int)
		for key, value := range thisRobotMoves {
			lastRobotMoves[key] = value
		}
	}

	result := 0

	for move, moveCount := range lastRobotMoves {
		result += len(move) * moveCount * codeNumPart
	}

	return result
}

func main() {
	timeStart := time.Now()
	part1Result := 0
	part2Result := 0
	var codes []string

	inputFile, _ := os.Open("input")
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		codes = append(codes, line)
	}

	for _, code := range codes {
		part1Result += countCodeResult(code, 2)
		part2Result += countCodeResult(code, 25)
	}

	fmt.Printf("Part 1 result: %d \n", part1Result)
	fmt.Printf("Part 2 result: %d \n", part2Result)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
