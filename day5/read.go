package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile() ([][]int, [][]int, error) {
	filename := "day5/input.txt"

	var rules [][]int
	var upadates [][]int

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		if strings.Contains(line, "|") {
			s := strings.Split(line, "|")

			v1, e1 := strconv.Atoi(s[0])

			v2, e2 := strconv.Atoi(s[1])

			if e1 != nil || e2 != nil {
				panic("Error reading file")
			}

			vals := []int{v1, v2}

			rules = append(rules, vals)
		} else if strings.Contains(line, ",") {
			s := strings.Split(line, ",")

			var slice []int

			for _, v := range s {
				parsedV, e := strconv.Atoi(v)

				if e == nil {
					slice = append(slice, parsedV)
				}
			}

			upadates = append(upadates, slice)
		}
	}

	return rules, upadates, nil
}
