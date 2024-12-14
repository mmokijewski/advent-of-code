package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func flood(garden [][]string, point [2]int, visitedFields map[[2]int]bool) int {
	area := 0
	score := 0
	newRegion := [][2]int{point}
	visitedFields[point] = true

	gardenHeight := len(garden)
	gardenWidth := len(garden[0])
	plantType := garden[point[0]][point[1]]

	directions := map[string][2]int{
		"up":    {-1, 0},
		"down":  {1, 0},
		"left":  {0, -1},
		"right": {0, 1},
	}

	for len(newRegion) > 0 {
		currentField := newRegion[0]
		newRegion = newRegion[1:]
		area++

		for _, dir := range directions {
			y := currentField[0] + dir[0]
			x := currentField[1] + dir[1]
			newPos := [2]int{y, x}
			if y < 0 || y >= gardenHeight || x < 0 || x >= gardenWidth {
				score++
			} else {
				nearFieldPlantType := garden[y][x]
				if nearFieldPlantType != plantType {
					score++
				} else if !visitedFields[newPos] {
					visitedFields[newPos] = true
					newRegion = append(newRegion, newPos)
				}
			}
		}
	}
	fmt.Printf("Plant: %s, area: %d, score: %d\n", plantType, area, score)
	return area * score
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	var garden [][]string
	visitedFields := make(map[[2]int]bool)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		var gardenLine []string
		for _, field := range strings.Split(line, "") {
			gardenLine = append(gardenLine, field)
		}
		garden = append(garden, gardenLine)
	}

	part1Sum := 0
	for y, line := range garden {
		for x := range line {
			point := [2]int{y, x}
			if !visitedFields[point] {
				part1Sum += flood(garden, point, visitedFields)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
