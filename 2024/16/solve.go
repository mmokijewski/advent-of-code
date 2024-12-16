package main

import (
	"bufio"
	"fmt"
	"os"
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
}

var directions = []Point{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

//func countPathScore(path [][2]int) int {
//	//fmt.Println(path)
//	start := path[0]
//	score := 0
//	score += len(path) - 1
//	for i, field := range path {
//		if i == 0 {
//			continue
//		}
//		if i == 1 {
//			if field[0] == start[0]-1 {
//				score += 1000
//			}
//			continue
//		}
//		if field[0] != path[i-2][0] && field[1] != path[i-2][1] {
//			score += 1000
//		}
//	}
//	return score
//}

func getPathWithLowestScore(paths []Path) (Path, []Path) {
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].score <= paths[j].score
	})
	minPath := paths[0]
	if len(paths) > 1 {
		paths = paths[1:]
	}
	return minPath, paths
}

func findCheapestPath(board [][]string, start Point) int {
	paths := []Path{{start, 1, 0}}
	visitedFields := make(map[[3]int]bool)
	var output int

	for {
		var currentPath Path
		currentPath, paths = getPathWithLowestScore(paths)
		posY := currentPath.position.y
		posX := currentPath.position.x
		currentPosWithDir := [3]int{posY, posX, currentPath.lastDirection}

		if board[posY][posX] == "E" {
			output = currentPath.score
			break
		}

		if visitedFields[currentPosWithDir] {
			continue
		}
		visitedFields[currentPosWithDir] = true

		for dirIndex, dir := range directions {
			newPos := Point{posY + dir.y, posX + dir.x}

			if board[newPos.y][newPos.x] == "#" {
				continue
			}

			nextPath := Path{currentPath.position, dirIndex, currentPath.score}
			if currentPath.lastDirection != dirIndex {
				nextPath.score += 1000
			} else {
				nextPath.position = newPos
				nextPath.score += 1
			}
			paths = append(paths, nextPath)
		}
	}
	return output
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

	score := findCheapestPath(board, start)

	fmt.Printf("Part 1: %d\n", score)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
