package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

func goFurther(playground []string, posY int, posX int, direction int) []int {
	switch direction {
	case 1: //^
		if posY == 0 {
			return []int{-1, -1, -1}
		}
		next := string(playground[posY-1][posX])
		if next == "#" {
			return []int{posY, posX + 1, 3}
		} else {
			return []int{posY - 1, posX, 1}
		}
	case 2: //v
		if posY == len(playground)-1 {
			return []int{-1, -1, -1}
		}
		next := string(playground[posY+1][posX])
		if next == "#" {
			return []int{posY, posX - 1, 4}
		} else {
			return []int{posY + 1, posX, 2}
		}
	case 3: //>
		if posX == len(playground[1])-1 {
			return []int{-1, -1, -1}
		}
		next := string(playground[posY][posX+1])
		if next == "#" {
			return []int{posY + 1, posX, 2}
		} else {
			return []int{posY, posX + 1, 3}
		}
	case 4: //<
		if posX == 0 {
			return []int{-1, -1, -1}
		}
		next := string(playground[posY][posX-1])
		if next == "#" {
			return []int{posY - 1, posX, 1}
		} else {
			return []int{posY, posX - 1, 4}
		}
	}
	return []int{-1, -1, -1}
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	var playground []string
	var direction int
	var posX int
	var posY int
	var visitedPositions []string

	i := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		playground = append(playground, line)
		for j, sign := range line {
			if string(sign) != "." && string(sign) != "#" {
				posX = j
				posY = i
				strDirection := string(sign)
				if strDirection == "^" {
					direction = 1
				} else if strDirection == "v" {
					direction = 2
				} else if strDirection == ">" {
					direction = 3
				} else if strDirection == "<" {
					direction = 4
				}
			}
		}
		i++
	}
	for direction != -1 {
		stringPosition := fmt.Sprintf("%d,%d", posY, posX)
		if !slices.Contains(visitedPositions, stringPosition) {
			visitedPositions = append(visitedPositions, stringPosition)
		}
		move := goFurther(playground, posY, posX, direction)
		posY = move[0]
		posX = move[1]
		direction = move[2]

	}

	fmt.Printf("Part 1 sum: %d\n", len(visitedPositions))
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
