package day3

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile() (string, error) {
	filename := "day3/input.txt"

	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	acc := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		acc += line
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return acc, nil
}
