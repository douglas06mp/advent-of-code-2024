package main

import (
	"adventofcode2024/utils"
	"fmt"
	"math"
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
	var sum int
	funcs := []func(int, int) int{add, multiply}

	for _, equations := range inputs {
		split := strings.Split(equations, ": ")
		target, _ := strconv.Atoi(split[0])
		nums := toNums(strings.Split(split[1], " "))

		if isValid(target, nums, funcs) {
			sum += target
		}
	}

	return sum
}

// ####################
// #      Part 2      #
// ####################
func part2(inputs []string) int {
	var sum int
	funcs := []func(int, int) int{add, multiply, concatenate}

	for _, equations := range inputs {
		split := strings.Split(equations, ": ")
		target, _ := strconv.Atoi(split[0])
		nums := toNums(strings.Split(split[1], " "))

		if isValid(target, nums, funcs) {
			sum += target
		}
	}

	return sum
}

func isValid(target int, nums []int, funcs []func(int, int) int) bool {
	return isValidHelper(nums[0], target, nums, 1, funcs)
}

func isValidHelper(initital int, target int, nums []int, idx int, funcs []func(int, int) int) bool {
	if idx == len(nums)-1 {
		for _, f := range funcs {
			if f(initital, nums[idx]) == target {
				return true
			}
		}

		return false
	}

	for _, f := range funcs {
		if isValidHelper(f(initital, nums[idx]), target, nums, idx+1, funcs) {
			return true
		}
	}

	return false
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func concatenate(a, b int) int {
	return a*int(math.Pow(10, float64(countDigits(b)))) + b
}

func countDigits(num int) int {
	count := 0

	for num != 0 {
		num /= 10
		count++
	}

	return count
}

func toNums(input []string) []int {
	nums := make([]int, len(input))

	for i, num := range input {
		nums[i], _ = strconv.Atoi(num)
	}

	return nums
}
