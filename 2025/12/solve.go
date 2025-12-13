package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Present struct {
	fields    [3][3]string
	hashes    int
	rotations []int
}

type Tree struct {
	width, height int
	presents      []int
}

func collectInputData() ([]Present, []Tree) {
	var presents []Present
	var trees []Tree

	numReg := regexp.MustCompile(`\d+`)

	inputFile, _ := os.Open("input")
	scanner := bufio.NewScanner(inputFile)
	i := -1
	j := 0
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 2 {
			presents = append(presents, Present{})
			i++
			j = 0
		} else if strings.Contains(line, "#") {
			for k, char := range strings.Split(line, "") {
				presents[i].fields[j][k] = char
				if char == "#" {
					presents[i].hashes++
				}
			}
			j++
		} else if strings.Contains(line, "x") {
			split := numReg.FindAllString(line, -1)
			var height, width int
			var presentsAmount []int
			for i, numString := range split {
				num, _ := strconv.Atoi(numString)
				if i == 0 {
					width = num
				} else if i == 1 {
					height = num
				} else {
					presentsAmount = append(presentsAmount, num)
				}
			}
			trees = append(trees, Tree{width, height, presentsAmount})
		}
	}

	for i, present := range presents {
		presents[i].rotations = checkPresentRotations(present)
	}

	return presents, trees
}

func prepareEmptyBoard(tree Tree) [][]string {
	var board [][]string
	for range tree.height {
		var line []string
		for range tree.width {
			line = append(line, ".")
		}
		board = append(board, line)
	}
	return board
}

func copyPresent(present Present) Present {
	var newPresent Present
	for i := range 3 {
		for j := range 3 {
			newPresent.fields[i][j] = present.fields[i][j]
		}
	}
	newPresent.hashes = present.hashes
	return newPresent
}

func rotateRight(present Present, times int) Present {
	var tmp, rotatedPresent Present
	tmp = copyPresent(present)
	for range times {
		rotatedPresent = copyPresent(present)
		rotatedPresent.fields[0][0] = tmp.fields[2][0]
		rotatedPresent.fields[0][1] = tmp.fields[1][0]
		rotatedPresent.fields[0][2] = tmp.fields[0][0]
		rotatedPresent.fields[1][0] = tmp.fields[2][1]
		rotatedPresent.fields[1][1] = tmp.fields[1][1]
		rotatedPresent.fields[1][2] = tmp.fields[0][1]
		rotatedPresent.fields[2][0] = tmp.fields[2][2]
		rotatedPresent.fields[2][1] = tmp.fields[1][2]
		rotatedPresent.fields[2][2] = tmp.fields[0][2]
		tmp = copyPresent(rotatedPresent)
	}
	return tmp
}

func compareFields(fields1, fields2 [3][3]string) bool {
	for i, line := range fields1 {
		for j, field := range line {
			if field != fields2[i][j] {
				return false
			}
		}
	}
	return true
}

func checkPresentRotations(present Present) []int {
	var rotations []int
	rotations = append(rotations, []int{0}...)
	turn1 := rotateRight(present, 1).fields
	turn2 := rotateRight(present, 2).fields
	turn3 := rotateRight(present, 3).fields
	if !compareFields(present.fields, turn1) {
		rotations = append(rotations, 1)
	}
	if !compareFields(present.fields, turn2) && !compareFields(turn1, turn2) {
		rotations = append(rotations, 2)
	}
	if !compareFields(present.fields, turn3) && !compareFields(turn1, turn3) && !compareFields(turn2, turn3) {
		rotations = append(rotations, 3)
	}
	return rotations
}

func copyBoard(board [][]string) [][]string {
	var newBoard [][]string
	for _, line := range board {
		var newLine []string
		for _, field := range line {
			newLine = append(newLine, field)
		}
		newBoard = append(newBoard, newLine)
	}
	return newBoard
}

type MemoEntry struct {
	presentIndex, rotations, y, x int
}

func putIntoBoard(board [][]string, availablePresents []Present, remainingPresents []int) ([][]string, []int, bool, map[MemoEntry]bool) {
	var presentToPut Present
	var tmpRemainingPresents []int

	memory := make(map[MemoEntry]bool)

	presentsToPut := 0
	for _, count := range remainingPresents {
		tmpRemainingPresents = append(tmpRemainingPresents, count)
		presentsToPut += count
	}
	var presentToPutIndex int
	for i, count := range tmpRemainingPresents {
		if count == 0 {
			continue
		}
		presentToPut = availablePresents[i]
		tmpRemainingPresents[i]--
		presentToPutIndex = i
		break
	}
	if presentsToPut == 0 {
		return board, remainingPresents, true, memory
	}

	for i, boardLine := range board {
		for j := range boardLine {
			// Check if present will be put outside the board
			if i+2 >= len(board) || j+2 >= len(board[0]) {
				continue
			}

			for _, rotations := range availablePresents[presentToPutIndex].rotations {
				memoryEntry := MemoEntry{presentToPutIndex, rotations, i, j}
				if memory[memoryEntry] {
					continue
				}
				memory[memoryEntry] = true

				if rotations != 0 {
					presentToPut = rotateRight(presentToPut, rotations)
				}
				tmpBoard := copyBoard(board)
				hashOnHash := false
				for k, presentLine := range presentToPut.fields {
					for l, sign := range presentLine {
						if sign == "#" {
							// To not overwrite hashes with .
							if tmpBoard[i+k][j+l] == "#" {
								hashOnHash = true
								break
							}
							tmpBoard[i+k][j+l] = sign
						}
					}
					if hashOnHash {
						break
					}
				}
				if hashOnHash {
					continue
				}

				newBoard, newRemainingPresents, result, newMemory := putIntoBoard(tmpBoard, availablePresents, tmpRemainingPresents)
				for k, v := range newMemory {
					memory[k] = v
				}

				if result {
					return newBoard, newRemainingPresents, true, memory
				}
			}
		}
	}
	return board, tmpRemainingPresents, false, memory
}

func main() {
	start := time.Now()

	part1 := 0
	presents, trees := collectInputData()

	for _, tree := range trees {
		area := tree.height * tree.width
		var totalPresentHashes, presentsAmount int
		for i, amount := range tree.presents {
			totalPresentHashes += presents[i].hashes * amount
			presentsAmount += amount
		}

		// Physically not possible
		if totalPresentHashes > area {
			continue
		}

		// Easy case where each present has its own 3*3 area
		if presentsAmount*9 < area {
			part1++
			continue
		}

		board := prepareEmptyBoard(tree)
		_, _, possible, _ := putIntoBoard(board, presents, tree.presents)
		if possible {
			part1++
		}
	}

	fmt.Printf("Part1 : %d\n", part1)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
