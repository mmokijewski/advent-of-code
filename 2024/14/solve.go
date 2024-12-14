package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func part1(boardWidth int, boardHeight int, robots [][]int) int {
	result := 1
	loopsQuantity := 100
	quarters := make(map[int]int)

	for _, robot := range robots {
		robot[0] = (robot[0] + (loopsQuantity * robot[2])) % boardWidth
		robot[1] = (robot[1] + (loopsQuantity * robot[3])) % boardHeight
		if robot[0] < 0 {
			robot[0] = boardWidth + robot[0]
		}
		if robot[1] < 0 {
			robot[1] = boardHeight + robot[1]
		}

		if robot[0] < boardWidth/2 && robot[1] < boardHeight/2 {
			quarters[1]++
		} else if robot[0] > boardWidth/2 && robot[1] < boardHeight/2 {
			quarters[2]++
		} else if robot[0] < boardWidth/2 && robot[1] > boardHeight/2 {
			quarters[3]++
		} else if robot[0] > boardWidth/2 && robot[1] > boardHeight/2 {
			quarters[4]++
		}
	}

	for _, robotCount := range quarters {
		result *= robotCount
	}
	return result
}

func part2(boardWidth int, boardHeight int, robots [][]int) int {
	var result int
	for i := 1; i < 10000; i++ {
		grid := make([][]string, boardHeight)
		for j := 0; j < boardHeight; j++ {
			grid[j] = make([]string, boardWidth)
		}
		robotLocations := make(map[[2]int]bool)
		twoRobotsOnOneLocation := false
		for _, robot := range robots {
			robot[0] = (robot[0] + robot[2]) % boardWidth
			robot[1] = (robot[1] + robot[3]) % boardHeight
			if robot[0] < 0 {
				robot[0] = boardWidth + robot[0]
			}
			if robot[1] < 0 {
				robot[1] = boardHeight + robot[1]
			}
			if robotLocations[[2]int{robot[1], robot[0]}] {
				twoRobotsOnOneLocation = true
				continue
			}
			robotLocations[[2]int{robot[1], robot[0]}] = true
			grid[robot[1]][robot[0]] = "#"
		}
		if twoRobotsOnOneLocation {
			continue
		}
		for y, line := range grid {
			for x, sign := range line {
				if sign != "#" {
					grid[y][x] = "."
				}
			}
		}
		for _, row := range grid {
			fmt.Println(row)
		}
		result = i
		break
	}
	return result
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	numRegex := regexp.MustCompile(`\d+|-\d+`)
	var robots [][]int
	boardWidth := 101
	boardHeight := 103

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		var robot []int
		strNums := numRegex.FindAllString(line, 4)
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			robot = append(robot, num)
		}
		robots = append(robots, robot)
	}

	//result := part1(boardWidth, boardHeight, robots)
	result := part2(boardWidth, boardHeight, robots)

	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
