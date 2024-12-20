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
	position        Point
	lastDirection   int
	score           int
	cheats          int
	visitedFields   []Point
	placesWithScore map[Point]int
}

var directions = []Point{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func findCheapestPath(board [][]string, start Point) (int, map[Point]int) {
	initialPlacesWithScore := make(map[Point]int)
	paths := []Path{{start, 1, 0, 0, []Point{start}, initialPlacesWithScore}}
	visitedFields := make(map[[3]int]int)
	var lowestScore int

	for {
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
			initialPlacesWithScore = make(map[Point]int)
			for key, value := range currentPath.placesWithScore {
				initialPlacesWithScore[key] = value
			}
			break
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
			newPlacesWithScore := make(map[Point]int)
			for key, value := range currentPath.placesWithScore {
				newPlacesWithScore[key] = value
			}
			newPlacesWithScore[newPos] = currentPath.score + 1
			nextPath := Path{newPos, dirIndex, currentPath.score + 1, 0, newVisitedFields, newPlacesWithScore}
			paths = append(paths, nextPath)
		}
	}
	return lowestScore, initialPlacesWithScore
}

func findCheatedPaths(board [][]string, start Point, minDiff int) int {
	lowestScore, placesWithScore := findCheapestPath(board, start)
	fmt.Println("Lowest: ", lowestScore)
	initialPlacesWithScore := make(map[Point]int)
	paths := []Path{{start, 1, 0, 0, []Point{start}, initialPlacesWithScore}}
	result := 0

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

		posY := currentPath.position.y
		posX := currentPath.position.x

		for dirIndex, dir := range directions {
			newPos := Point{posY + dir.y, posX + dir.x}

			if slices.Contains(currentPath.visitedFields, newPos) {
				continue
			}

			if currentPath.cheats == 1 && board[newPos.y][newPos.x] != "#" {
				currentPathFinalScore := lowestScore - (placesWithScore[newPos] - currentPath.score)
				if currentPathFinalScore <= lowestScore-minDiff {
					result++
				}
				continue
			}

			newCheats := currentPath.cheats
			if board[newPos.y][newPos.x] == "#" {
				if currentPath.cheats == 1 || newPos.y == 0 || newPos.y == len(board)-1 || newPos.x == 0 || newPos.x == len(board[0])-1 {
					continue
				}
				newCheats++
			}

			newVisitedFields := slices.Clone(currentPath.visitedFields)
			newVisitedFields = append(newVisitedFields, newPos)
			newPlacesWithScore := make(map[Point]int)
			for key, value := range currentPath.placesWithScore {
				newPlacesWithScore[key] = value
			}
			newPlacesWithScore[newPos] = currentPath.score + 1
			nextPath := Path{newPos, dirIndex, currentPath.score + 1, newCheats, newVisitedFields, newPlacesWithScore}
			paths = append(paths, nextPath)
		}
	}
	return result
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

	score := findCheatedPaths(board, start, 100)
	fmt.Printf("Part 1: %d\n", score)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
