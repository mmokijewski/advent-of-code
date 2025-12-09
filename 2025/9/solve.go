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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type point struct {
	x, y int
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	maxField := 0
	var redTiles []point

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		box := scanner.Text()
		split := strings.Split(box, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		redTiles = append(redTiles, point{x, y})
	}

	for i, tile := range redTiles {
		for j := i + 1; j < len(redTiles); j++ {
			field := int((math.Abs(float64(tile.x-redTiles[j].x)) + 1) * (math.Abs(float64(tile.y-redTiles[j].y)) + 1))
			if field > maxField {
				maxField = field
			}
		}
	}

	fmt.Printf("Part1 : %d\n", maxField)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
