package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	timeStart := time.Now()
	part1Result := 0
	inputFile, _ := os.Open("input")
	scanner := bufio.NewScanner(inputFile)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	var registerA, registerB, registerC int
	fmt.Sscanf(lines[0], "Register A: %d", &registerA)
	fmt.Sscanf(lines[1], "Register B: %d", &registerB)
	fmt.Sscanf(lines[2], "Register C: %d", &registerC)
	var program []int
	programLine := strings.TrimPrefix(lines[4], "Program: ")
	for _, programStr := range strings.Split(programLine, ",") {
		singleCommand, _ := strconv.Atoi(programStr)
		program = append(program, singleCommand)
	}

	for i, command := range program {
		switch command {
		case 1:
			fmt.Print(registerA/int(math.Pow(float64(2), float64(program[i+1]))), ",")
		case 2:
			fmt.Print(registerB^2, ",")
		case 3:
			fmt.Print(program[i+1] % 8)
		case 4:
		case 5:
		case 7:

		}
	}

	fmt.Println(registerA, registerB, registerC, program)

	fmt.Printf("Part 1 result: %d \n", part1Result)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
