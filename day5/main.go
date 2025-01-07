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

	// part2Result := part2(inputs)
	// fmt.Println("Part 2:", part2Result)
}

// ####################
// #      Part 1      #
// ####################
func part1(inputs []string) int {
	var sum int

	rules, numbers := getRulesAndNumbers(inputs)

	for _, number := range numbers {
		isValid := true
		middle := number[len(number)/2]

	OuterLoop:
		for i := 0; i < len(number); i++ {
			for j := i + 1; j < len(number); j++ {
				key := strconv.Itoa(number[j]) + "|" + strconv.Itoa(number[i])
				if _, exists := rules[key]; exists {
					isValid = false
					break OuterLoop
				}
			}
		}

		if isValid {
			sum += middle
		}
	}

	return sum
}

func getRulesAndNumbers(inputs []string) (map[string]bool, [][]int) {
	rules := make(map[string]bool)
	numbers := make([][]int, 0)

	isRules := true

	for _, input := range inputs {
		if input == "" {
			isRules = false
			continue
		}

		if isRules {
			rules[input] = true
		} else {
			number := getNumArray(input)
			numbers = append(numbers, number)
		}

	}

	return rules, numbers
}

func getNumArray(input string) []int {
	numsStr := strings.Split(input, ",")
	numsInt := make([]int, 0)

	for _, num := range numsStr {
		n, _ := strconv.Atoi(num)
		numsInt = append(numsInt, n)
	}

	return numsInt
}
