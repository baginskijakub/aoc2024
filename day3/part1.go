package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetResults() {
	input, err := ReadFile()

	if err != nil {
		fmt.Errorf("failed to read input: %w", err)
		return
	}

	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, match := range matches {
		if len(match) == 3 {
			n, _ := strconv.Atoi(match[1])
			m, _ := strconv.Atoi(match[2])
			sum += n * m
		}
	}

	fmt.Println("Total:", sum)
}
