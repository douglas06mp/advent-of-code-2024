package main

import (
	"adventofcode2024/utils"
	"fmt"
	"regexp"
	"strconv"
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

	r, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	for _, input := range inputs {
		matches := r.FindAllStringSubmatch(input, -1)

		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}

	return sum
}

func part2(inputs []string) int {
	var sum int
	enabled := true

	r, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`)

	for _, input := range inputs {
		matches := r.FindAllStringSubmatch(input, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			}

			if enabled {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				sum += num1 * num2
			}
		}
	}

	return sum
}
