package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"time"
)

type Point struct {
	y, x int
}

type PointWithDir struct {
	pos Point
	dir Point
}

type Path struct {
	position      Point
	visitedFields []Point
	moves         map[Point]int
	lastMove      Point
}

var directions = map[string]Point{
	"<": {0, -1},
	"v": {1, 0},
	"^": {-1, 0},
	">": {0, 1},
}

func findCheapestPathOnKeypad(keypad [][]string, start string, end string) string {
	var startPos Point
	for i, line := range keypad {
		for j, key := range line {
			if key == start {
				startPos = Point{i, j}
			}
		}
	}
	visitedFields := make(map[PointWithDir]bool)
	var lowestScore int

	moves := make(map[Point]int)
	paths := []Path{{startPos, []Point{}, moves, Point{-1, -1}}}

	for {
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i].visitedFields) <= len(paths[j].visitedFields)
		})
		currentPath := paths[0]
		//fmt.Println(start, startPos, end, currentPath)
		if len(paths) > 1 {
			paths = append([]Path{}, paths[1:]...)
		}
		if lowestScore != 0 && len(currentPath.visitedFields) > lowestScore {
			continue
		}

		posY := currentPath.position.y
		posX := currentPath.position.x
		currentPosWithDir := PointWithDir{currentPath.position, currentPath.lastMove}

		if keypad[posY][posX] == end {
			lowestScore = len(currentPath.visitedFields)
			for key, value := range currentPath.moves {
				moves[key] = value
			}
			break
		}

		if visitedFields[currentPosWithDir] {
			continue
		}
		visitedFields[currentPosWithDir] = true

		for _, dir := range directions {
			newPos := Point{posY + dir.y, posX + dir.x}

			if newPos.y < 0 || newPos.y >= len(keypad) || newPos.x < 0 || newPos.x >= len(keypad[0]) || keypad[newPos.y][newPos.x] == "X" {
				continue
			}

			if slices.Contains(currentPath.visitedFields, newPos) {
				continue
			}

			newVisitedFields := slices.Clone(currentPath.visitedFields)
			newVisitedFields = append(newVisitedFields, newPos)
			newMoves := make(map[Point]int)
			for key, value := range currentPath.moves {
				newMoves[key] = value
			}
			newMoves[dir]++
			nextPath := Path{newPos, newVisitedFields, newMoves, dir}
			paths = append(paths, nextPath)
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
						result += dirKey
					}
					nextPos = tempNextPos
					continue
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

	fmt.Println(firstRobotMoves)

	secondRobotMoves := ""
	firstRobotMoves = "A" + firstRobotMoves
	for i, sign := range firstRobotMoves {
		if i == len(firstRobotMoves)-1 {
			break
		}
		secondRobotMoves = secondRobotMoves + findCheapestPathOnKeypad(dirKeypad, string(sign), string(firstRobotMoves[i+1]))
	}
	fmt.Println(secondRobotMoves)

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
