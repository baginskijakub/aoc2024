package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile() ([]int, []int, error) {
	filename := "day1/input.txt"

	var leftColumn []int
	var rightColumn []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in line: %s", line)
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in line: %s", line)
		}

		leftColumn = append(leftColumn, left)
		rightColumn = append(rightColumn, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return leftColumn, rightColumn, nil
}
