package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetConditionalResults() error {
	input, err := ReadFile()

	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	pattern := `mul\(([0-9]{1,3}),([0-9]{1,3})\)|don't\(\)|do\(\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0
	enabled := true

	for _, match := range matches {
		switch match[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				n, _ := strconv.Atoi(match[1])
				m, _ := strconv.Atoi(match[2])
				sum += n * m
			}
		}
	}

	fmt.Println("Total:", sum)

	return nil
}
