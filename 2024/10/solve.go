package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func directions() map[string][2]int {
	directions := map[string][2]int{
		"up":    {-1, 0},
		"down":  {1, 0},
		"left":  {0, -1},
		"right": {0, 1},
	}
	return directions
}

func getScore(board [][]int, point [2]int, expectedHeight int) map[[2]int]bool {
	summits := make(map[[2]int]bool)
	posY := point[0]
	posX := point[1]
	if posY < 0 || posY >= len(board) || posX < 0 || posX >= len(board[0]) {
		return summits
	}

	height := board[posY][posX]
	if height != expectedHeight {
		return summits
	} else if height == 9 {
		summits[point] = true
	} else {
		for _, dir := range directions() {
			for summit := range getScore(board, [2]int{posY + dir[0], posX + dir[1]}, height+1) {
				summits[summit] = true
			}
		}
	}
	return summits
}

func getScorePart2(board [][]int, point [2]int, expectedHeight int) int {
	trails := 0
	posY := point[0]
	posX := point[1]
	if posY < 0 || posY >= len(board) || posX < 0 || posX >= len(board[0]) {
		return trails
	}

	height := board[posY][posX]
	if height != expectedHeight {
		return trails
	} else if height == 9 {
		trails++
	} else {
		for _, dir := range directions() {
			trails += getScorePart2(board, [2]int{posY + dir[0], posX + dir[1]}, height+1)
		}
	}
	return trails
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	var board [][]int
	var trailheads [][2]int
	part1Sum := 0
	part2Sum := 0

	scanner := bufio.NewScanner(inputFile)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		var currentLineArray []int
		for j, sign := range strings.Split(line, "") {
			currentNum, _ := strconv.Atoi(sign)
			currentLineArray = append(currentLineArray, currentNum)
			if currentNum == 0 {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}
		board = append(board, currentLineArray)
		i++
	}

	for _, trailhead := range trailheads {
		part1Sum += len(getScore(board, trailhead, 0))
		part2Sum += getScorePart2(board, trailhead, 0)
	}

	fmt.Printf("Part 1 sum: %d\n", part1Sum)
	fmt.Printf("Part 2 sum: %d\n", part2Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
