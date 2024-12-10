package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getScore(board [][]int, point [2]int, expectedHeight int) (summitsPart1 map[[2]int]bool, trailsPart2 int) {
	summitsPart1 = make(map[[2]int]bool)
	trailsPart2 = 0
	posY := point[0]
	posX := point[1]
	if posY < 0 || posY >= len(board) || posX < 0 || posX >= len(board[0]) {
		return
	}

	directions := map[string][2]int{
		"up":    {-1, 0},
		"down":  {1, 0},
		"left":  {0, -1},
		"right": {0, 1},
	}

	height := board[posY][posX]
	if height == expectedHeight && height == 9 {
		summitsPart1[point] = true
		trailsPart2++
	} else if height == expectedHeight {
		for _, dir := range directions {
			summits, trails := getScore(board, [2]int{posY + dir[0], posX + dir[1]}, height+1)
			for summit := range summits {
				summitsPart1[summit] = true
			}
			trailsPart2 += trails
		}
	}
	return
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
		summits, trails := getScore(board, trailhead, 0)
		part1Sum += len(summits)
		part2Sum += trails
	}

	fmt.Printf("Part 1 sum: %d\n", part1Sum)
	fmt.Printf("Part 2 sum: %d\n", part2Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
