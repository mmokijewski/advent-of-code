package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type Point struct {
	y, x int
}

type Path struct {
	position Point
	score    int
}

var directions = []Point{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func findCheapestPath(board [][]string, start Point) int {
	paths := []Path{{start, 0}}
	visitedFields := make(map[Point]bool)
	var output int

	for {
		sort.Slice(paths, func(i, j int) bool {
			return paths[i].score <= paths[j].score
		})
		currentPath := paths[0]
		if len(paths) > 1 {
			paths = paths[1:]
		}

		posY := currentPath.position.y
		posX := currentPath.position.x

		if board[posY][posX] == "E" {
			output = currentPath.score
			break
		}

		if visitedFields[currentPath.position] {
			continue
		}
		visitedFields[currentPath.position] = true

		for _, dir := range directions {
			newPos := Point{posY + dir.y, posX + dir.x}
			if board[newPos.y][newPos.x] != "#" {
				nextPath := Path{newPos, currentPath.score + 1}
				paths = append(paths, nextPath)
			}
		}
	}
	return output
}

func main() {
	timeStart := time.Now()
	inputFile, _ := os.Open("input")

	start := Point{1, 1}
	boardSize := 73
	bytesToDrop := 1024
	board := make([][]string, boardSize)
	for i := range board {
		board[i] = make([]string, boardSize)
	}
	numReg := regexp.MustCompile(`\d+`)

	for i, line := range board {
		for j := range line {
			if j == 0 || i == 0 || j == boardSize-1 || i == boardSize-1 {
				board[i][j] = "#"
			} else if j == boardSize-2 && i == boardSize-2 {
				board[i][j] = "E"
			} else {
				board[i][j] = "."
			}
		}
	}

	scanner := bufio.NewScanner(inputFile)
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		strNums := numReg.FindAllString(line, 2)
		x, _ := strconv.Atoi(strNums[0])
		y, _ := strconv.Atoi(strNums[1])
		if i > bytesToDrop {
			break
		}
		board[y+1][x+1] = "#"
		i++
	}

	score := findCheapestPath(board, start)

	fmt.Printf("Part 1: %d\n", score)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
