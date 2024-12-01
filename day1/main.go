package main

import (
	"adventofcode2024/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
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
func part1(inputs []string) int {
	left, right := getTwoArrays(inputs)

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += getDiff(left[i], right[i])
	}

	return sum
}

func getTwoArrays(inputs []string) ([]int, []int) {
	var left []int
	var right []int

	for _, input := range inputs {
		split := strings.Split(input, " ")

		leftNum, _ := strconv.Atoi(split[0])
		rightNum, _ := strconv.Atoi(split[3])

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	return left, right
}

func getDiff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

// ####################
// #      Part 2      #
// ####################
func part2(inputs []string) int {
	left, right := getTwoArrays(inputs)

	freqMap := make(map[int]int)
	for _, num := range right {
		freqMap[num]++
	}

	sum := 0
	for _, num := range left {
		sum += num * freqMap[num]
	}

	return sum
}
