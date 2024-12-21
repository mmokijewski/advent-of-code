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

type Path struct {
	position Point
	moves    map[Point]int
}

var directions = []Point{
	{0, -1},
	{1, 0},
	{-1, 0},
	{0, 1},
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
	moves := make(map[Point]int)

	newPos := startPos
	for keypad[newPos.y][newPos.x] != end {
		for dirIndex, dir := range directions {
			tempPos := Point{newPos.y + dir.y, newPos.x + dir.x}
			if tempPos.y < 0 || tempPos.y >= len(keypad) || tempPos.x < 0 || tempPos.x >= len(keypad[0]) {
				continue
			}

			if keypad[tempPos.y][tempPos.x] == "X" {
				continue
			}
			if newPos.x == endPos.x && (dirIndex == 0 || dirIndex == 3) {
				continue
			}
			if newPos.y == endPos.y && (dirIndex == 1 || dirIndex == 2) {
				continue
			}
			if newPos.x < endPos.x && dirIndex == 0 {
				continue
			}
			if newPos.y > endPos.y && dirIndex == 1 {
				continue
			}
			moves[dir]++
			newPos = tempPos
			break
		}
	}

	result := ""
	nextPos := startPos
	for keypad[nextPos.y][nextPos.x] != end {
		for dirKey, dir := range directions {
			if moves[directions[dirKey]] > 0 {
				movesAmount := moves[directions[dirKey]]
				tempNextPos := Point{nextPos.y + movesAmount*dir.y, nextPos.x + movesAmount*dir.x}
				if keypad[tempNextPos.y][tempNextPos.x] != "X" {
					moves[directions[dirKey]] = 0
					for range movesAmount {
						switch dirKey {
						case 0:
							result += "<"
						case 1:
							result += "v"
						case 2:
							result += "^"
						case 3:
							result += ">"
						}
					}
					nextPos = tempNextPos
				}
			}
		}
	}
	return result + "A"
}

func countCodeResult(code string) int {
	numericKeypad := make([][]string, 4)
	numericKeypad[0] = append(numericKeypad[0], "7", "8", "9")
	numericKeypad[1] = append(numericKeypad[1], "4", "5", "6")
	numericKeypad[2] = append(numericKeypad[2], "1", "2", "3")
	numericKeypad[3] = append(numericKeypad[3], "X", "0", "A")

	dirKeypad := make([][]string, 2)
	dirKeypad[0] = append(dirKeypad[0], "X", "^", "A")
	dirKeypad[1] = append(dirKeypad[1], "<", "v", ">")

	firstRobotMoves := ""

	codeNumPart, _ := strconv.Atoi(code[:len(code)-1])

	code = "A" + code
	for i, sign := range code {
		if i == len(code)-1 {
			break
		}
		firstRobotMoves = firstRobotMoves + findCheapestPathOnKeypad(numericKeypad, string(sign), string(code[i+1]))
	}

	secondRobotMoves := ""
	firstRobotMoves = "A" + firstRobotMoves
	for i, sign := range firstRobotMoves {
		if i == len(firstRobotMoves)-1 {
			break
		}
		secondRobotMoves = secondRobotMoves + findCheapestPathOnKeypad(dirKeypad, string(sign), string(firstRobotMoves[i+1]))
	}

	thirdRobotMoves := ""
	secondRobotMoves = "A" + secondRobotMoves
	for i, sign := range secondRobotMoves {
		if i == len(secondRobotMoves)-1 {
			break
		}
		thirdRobotMoves = thirdRobotMoves + findCheapestPathOnKeypad(dirKeypad, string(sign), string(secondRobotMoves[i+1]))
	}
	return len(thirdRobotMoves) * codeNumPart
}

func main() {
	timeStart := time.Now()
	part1Result := 0
	var codes []string

	inputFile, _ := os.Open("input")
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		codes = append(codes, line)
	}

	for _, code := range codes {
		part1Result += countCodeResult(code)
	}

	fmt.Printf("Part 1 result: %d \n", part1Result)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
