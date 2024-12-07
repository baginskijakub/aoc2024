package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile() ([][]string, error) {
	filename := "day6/input.txt"

	var matrix [][]string

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		s := strings.Split(line, "")

		var row []string

		for _, c := range s {
			row = append(row, c)
		}

		matrix = append(matrix, row)
	}

	return matrix, nil
}
