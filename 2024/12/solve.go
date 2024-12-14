package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func flood(garden [][]string, point [2]int, visitedFields map[[2]int]bool) (int, int) {
	area := 0
	score := 0
	newRegion := [][2]int{point}
	visitedFields[point] = true
	edges := make(map[[3]int]bool)

	gardenHeight := len(garden)
	gardenWidth := len(garden[0])
	plantType := garden[point[0]][point[1]]

	directions := map[int][3]int{
		0: {-1, 0}, //up
		1: {1, 0},  //down
		2: {0, -1}, //left
		3: {0, 1},  //right
	}

	for len(newRegion) > 0 {
		currentField := newRegion[0]
		newRegion = newRegion[1:]
		area++

		for dirIndex, dir := range directions {
			y := currentField[0] + dir[0]
			x := currentField[1] + dir[1]
			newPos := [2]int{y, x}
			if y < 0 || y >= gardenHeight || x < 0 || x >= gardenWidth {
				score++
				edges[[3]int{currentField[0], currentField[1], dirIndex}] = true
			} else {
				nearFieldPlantType := garden[y][x]
				if nearFieldPlantType != plantType {
					score++
					edges[[3]int{currentField[0], currentField[1], dirIndex}] = true
				} else if !visitedFields[newPos] {
					visitedFields[newPos] = true
					newRegion = append(newRegion, newPos)
				}
			}
		}
	}
	for edge := range edges {
		dir := edge[2]
		if dir == 0 || dir == 1 { //horizontal edge
			_, rightExists := edges[[3]int{edge[0], edge[1] + 1, dir}]
			if rightExists {
				edges[[3]int{edge[0], edge[1] + 1, dir}] = false
			}
		} else { //vertical edge
			_, downExists := edges[[3]int{edge[0] - 1, edge[1], dir}]
			if downExists {
				edges[[3]int{edge[0] - 1, edge[1], dir}] = false
			}
		}
	}

	edgesCount := 0
	for _, edge := range edges {
		if edge == true {
			edgesCount++
		}
	}

	return area * score, area * edgesCount
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
	part2Sum := 0
	for y, line := range garden {
		for x := range line {
			point := [2]int{y, x}
			if !visitedFields[point] {
				part1, part2 := flood(garden, point, visitedFields)
				part1Sum += part1
				part2Sum += part2
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Part 2: %d\n", part2Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
