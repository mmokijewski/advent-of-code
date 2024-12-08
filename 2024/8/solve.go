package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func validatePosition(board [][]string, pos [2]int) bool {
	if pos[0] < 0 || pos[0] >= len(board) || pos[1] < 0 || pos[1] >= len(board[0]) {
		return false
	} else {
		return true
	}
}

func getAntinodesForSinglePair(board [][]string, node1 [2]int, node2 [2]int) map[[2]int]bool {
	antinodes := make(map[[2]int]bool)
	dist := [2]int{node2[0] - node1[0], node2[1] - node1[1]}
	pos1 := [2]int{node2[0] + dist[0], node2[1] + dist[1]}
	pos2 := [2]int{node1[0] - dist[0], node1[1] - dist[1]}

	if validatePosition(board, pos1) {
		antinodes[pos1] = true
	}
	if validatePosition(board, pos2) {
		antinodes[pos2] = true
	}
	return antinodes
}

func checkForAntinodes(board [][]string, posY int, posX int) map[[2]int]bool {
	antinodes := make(map[[2]int]bool)
	currentChar := board[posY][posX]
	for i, line := range board {
		for j, sign := range line {
			if i == posY && j == posX {
				continue
			}
			if sign == currentChar {
				for newNode := range getAntinodesForSinglePair(board, [2]int{posY, posX}, [2]int{i, j}) {
					antinodes[newNode] = true
				}
			}
		}
	}
	return antinodes
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	var board [][]string
	antinodes := make(map[[2]int]bool)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		var charLine []string
		for _, sign := range line {
			charLine = append(charLine, string(sign))
		}
		board = append(board, charLine)
	}

	for i, line := range board {
		for j, char := range line {
			if char != "." {
				for newNode := range checkForAntinodes(board, i, j) {
					antinodes[newNode] = true
				}
			}
		}
	}

	fmt.Printf("Part 1 sum: %d\n", len(antinodes))
	//fmt.Printf("Part 1 sum: %d\n", part2Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
