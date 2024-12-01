package utils

import (
	"bufio"
	"os"
)

func ParseInput(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var input []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
