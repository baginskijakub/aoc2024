package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile() ([][]int, error) {
	filename := "day2/input.txt"

	var matrix [][]int

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var arr []int

		line := scanner.Text()

		parts := strings.Fields(line)

		for _, num := range parts {
			number, err := strconv.Atoi(num)

			if err != nil {
				return nil, fmt.Errorf("invalid number in line: %s", line)
			}

			arr = append(arr, number)
		}

		matrix = append(matrix, arr)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return matrix, nil
}
