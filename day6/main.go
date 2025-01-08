package main

import (
	"adventofcode2024/utils"
	"fmt"
)

const (
	start    = '^'
	obstacle = '#'
)

func main() {
	inputs := utils.ParseInput("input.txt")

	part1Result := part1(inputs)
	fmt.Println("Part 1:", part1Result)

	// part2Result := part2(inputs)
	// fmt.Println("Part 2:", part2Result)
}

// ####################
// #      Part 1      #
// ####################
func part1(inputs []string) int {
	x, y := findStart(inputs)
	fmt.Println(x, y)

	seen := make(map[[2]int]bool)
	seen[[2]int{x, y}] = true

	var xDir, yDir int
	dir := 0
	for {
		switch dir % 4 {
		case 0:
			xDir, yDir = -1, 0
		case 1:
			xDir, yDir = 0, 1
		case 2:
			xDir, yDir = 1, 0
		case 3:
			xDir, yDir = 0, -1
		}

		if reachEdge(inputs, x+xDir, y+yDir) {
			break
		}

		if reachObstacle(inputs, x+xDir, y+yDir) {
			dir++
			continue
		}

		x += xDir
		y += yDir

		seen[[2]int{x, y}] = true
	}

	return len(seen)
}

func reachEdge(inputs []string, x, y int) bool {
	return x < 0 || y < 0 || x >= len(inputs) || y >= len(inputs[0])
}

func reachObstacle(inputs []string, x, y int) bool {
	return inputs[x][y] == obstacle
}

func findStart(inputs []string) (int, int) {
	for i, row := range inputs {
		for j, char := range row {
			if char == start {
				return i, j
			}
		}
	}
	return -1, -1
}
