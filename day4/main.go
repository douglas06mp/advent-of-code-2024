package main

import (
	"adventofcode2024/utils"
	"fmt"
)

func main() {
	inputs := utils.ParseInput("input.txt")

	part1Result := part1(inputs)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(inputs)
	fmt.Println("Part 2:", part2Result)
}

// ####################
// #      Part 1      #
// ####################
var DIRECTIONS = [][2]int{
	{-1, 0},  // top
	{-1, 1},  // top right
	{0, 1},   // right
	{1, 1},   // bottom right
	{1, 0},   // bottom
	{1, -1},  // bottom left
	{0, -1},  // left
	{-1, -1}, // top left
}

func part1(inputs []string) int {
	var sum int
	xMax := len(inputs[0])
	yMax := len(inputs)

	for i := 0; i < yMax; i++ {
		for j := 0; j < xMax; j++ {
			if inputs[i][j] != 'X' {
				continue
			}

			for _, dir := range DIRECTIONS {
				xM := i + dir[0]
				yM := j + dir[1]

				if !isValidStep(xM, yM, yMax, xMax) || inputs[xM][yM] != 'M' {
					continue
				}

				xA, yA := i+dir[0]*2, j+dir[1]*2
				xS, yS := i+dir[0]*3, j+dir[1]*3

				hasA := isValidStep(xA, yA, yMax, xMax) && inputs[xA][yA] == 'A'
				hasS := isValidStep(xS, yS, yMax, xMax) && inputs[xS][yS] == 'S'

				if hasA && hasS {
					sum++
				}
			}
		}
	}

	return sum
}

func isValidStep(x, y, xMax, yMax int) bool {
	return x < xMax && y < yMax && x >= 0 && y >= 0
}

// ####################
// #      Part 2      #
// ####################
func part2(inputs []string) int {
	var sum int
	xMax := len(inputs[0])
	yMax := len(inputs)

	corners := [4][2]int{
		DIRECTIONS[7], // top left
		DIRECTIONS[1], // top right
		DIRECTIONS[5], // bottom left
		DIRECTIONS[3], // bottom right
	}

	for i := 0; i < yMax; i++ {
		for j := 0; j < xMax; j++ {
			if inputs[i][j] != 'A' {
				continue
			}

			shouldContinue := false
			for _, corner := range corners {
				if !isValidStep(i+corner[0], j+corner[1], yMax, xMax) {
					shouldContinue = true
				}
			}
			if shouldContinue {
				continue
			}

			xTL, yTL := i+corners[0][0], j+corners[0][1]
			xTR, yTR := i+corners[1][0], j+corners[1][1]
			xBL, yBL := i+corners[2][0], j+corners[2][1]
			xBR, yBR := i+corners[3][0], j+corners[3][1]

			if (inputs[xTL][yTL] == 'M' && inputs[xTR][yTR] == 'M' && inputs[xBL][yBL] == 'S' && inputs[xBR][yBR] == 'S') ||
				(inputs[xTR][yTR] == 'M' && inputs[xBR][yBR] == 'M' && inputs[xTL][yTL] == 'S' && inputs[xBL][yBL] == 'S') ||
				(inputs[xBR][yBR] == 'M' && inputs[xBL][yBL] == 'M' && inputs[xTR][yTR] == 'S' && inputs[xTL][yTL] == 'S') ||
				(inputs[xBL][yBL] == 'M' && inputs[xTL][yTL] == 'M' && inputs[xBR][yBR] == 'S' && inputs[xTR][yTR] == 'S') {
				sum++
			}
		}
	}

	return sum
}
