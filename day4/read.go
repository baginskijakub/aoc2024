package day4

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile() ([][]string, error) {
	filename := "day4/input.txt"

	var matrix [][]string

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var arr []string

		line := scanner.Text()

		for _, char := range line {
			arr = append(arr, string(char))
		}

		matrix = append(matrix, arr)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return matrix, nil
}
