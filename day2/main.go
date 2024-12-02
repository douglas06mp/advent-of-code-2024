package main

import (
	"adventofcode2024/utils"
	"fmt"
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
	var reports [][]int

	for _, input := range inputs {
		report := getIntArray(input)
		reports = append(reports, report)
	}

	var count int
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}

	return count
}

func getIntArray(input string) []int {
	var arr []int

	split := strings.Split(input, " ")
	for _, i := range split {
		num, _ := strconv.Atoi(i)
		arr = append(arr, num)
	}

	return arr
}

func isSafe(report []int) bool {
	maxDiff := 3
	increasing := report[0] < report[1]

	for i := 1; i < len(report); i++ {
		curr := report[i]
		prev := report[i-1]

		isSame := curr == prev
		isIncreasing := increasing && (curr < prev || curr-prev > maxDiff)
		isDecreasing := !increasing && (curr > prev || prev-curr > maxDiff)

		if isSame || isIncreasing || isDecreasing {
			return false
		}
	}

	return true
}

// ####################
// #      Part 2      #
// ####################
func part2(inputs []string) int {
	var reports [][]int

	for _, input := range inputs {
		report := getIntArray(input)
		reports = append(reports, report)
	}

	var count int
	for _, report := range reports {
		if isSafe(report) || isSafeAfterRemoval(report) {
			count++
		}
	}

	return count
}

func isSafeAfterRemoval(report []int) bool {
	for i := 0; i < len(report); i++ {
		copyReport := make([]int, len(report))
		copy(copyReport, report)

		if isSafe(append(copyReport[:i], copyReport[i+1:]...)) {
			return true
		}
	}

	return false
}
