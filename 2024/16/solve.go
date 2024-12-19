package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"time"
)

type Point struct {
	y, x int
}

type Path struct {
	position      Point
	lastDirection int
	score         int
	visitedFields []Point
}

var directions = []Point{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func findCheapestPath(board [][]string, start Point) (int, int) {
	paths := []Path{{start, 1, 0, []Point{start}}}
	visitedFields := make(map[[3]int]int)
	placesToSit := make(map[Point]bool)
	var lowestScore int

	for len(paths) > 0 {
		sort.Slice(paths, func(i, j int) bool {
			return paths[i].score <= paths[j].score
		})
		currentPath := paths[0]
		if len(paths) > 1 {
			paths = append([]Path{}, paths[1:]...)
		} else {
			paths = []Path{}
		}

		if lowestScore != 0 && currentPath.score > lowestScore {
			continue
		}

		posY := currentPath.position.y
		posX := currentPath.position.x
		currentPosWithDir := [3]int{posY, posX, currentPath.lastDirection}

		if board[posY][posX] == "E" {
			lowestScore = currentPath.score
			for _, place := range currentPath.visitedFields {
				placesToSit[place] = true
			}
			continue
		}

		if visitedFields[currentPosWithDir] != 0 && currentPath.score > visitedFields[currentPosWithDir] {
			continue
		}
		visitedFields[currentPosWithDir] = currentPath.score

		for dirIndex, dir := range directions {
			newPos := Point{posY + dir.y, posX + dir.x}

			if slices.Contains(currentPath.visitedFields, newPos) {
				continue
			}

			if board[newPos.y][newPos.x] == "#" {
				continue
			}

			newVisitedFields := slices.Clone(currentPath.visitedFields)
			newVisitedFields = append(newVisitedFields, newPos)
			nextPath := Path{newPos, dirIndex, currentPath.score, newVisitedFields}
			if currentPath.lastDirection != dirIndex {
				nextPath.score += 1001
			} else {
				nextPath.score += 1
			}
			paths = append(paths, nextPath)
		}
	}
	return lowestScore, len(placesToSit)
}

func main() {
	timeStart := time.Now()
	inputFile, _ := os.Open("input")

	var start Point
	var board [][]string

	scanner := bufio.NewScanner(inputFile)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		var stringArray []string
		for j, str := range line {
			stringArray = append(stringArray, string(str))
			if string(str) == "S" {
				start = Point{i, j}
			}
		}
		board = append(board, stringArray)
		i++
	}

	score, placesToSit := findCheapestPath(board, start)

	fmt.Printf("Part 1: %d\n", score)
	fmt.Printf("Part 2: %d\n", placesToSit)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
