package day4

import "fmt"

func Find2() {
	matrix, err := ReadFile()

	if err != nil {
		panic(err)
	}

	sum := 0

	for i, row := range matrix {
		for j, col := range row {
			if col == "A" {
				if check(matrix, i, j) {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func check(matrix [][]string, i int, j int) bool {
	return safelyCheck(func() bool {
		topLeft := matrix[i-1][j-1]
		topRight := matrix[i-1][j+1]
		bottomLeft := matrix[i+1][j-1]
		bottomRight := matrix[i+1][j+1]

		fmt.Println(topLeft, topRight, bottomLeft, bottomRight)

		if topLeft == "M" && topRight == "M" && bottomLeft == "S" && bottomRight == "S" {
			return true
		}

		if topLeft == "S" && topRight == "S" && bottomLeft == "M" && bottomRight == "M" {
			return true
		}

		if topRight == "M" && bottomRight == "M" && topLeft == "S" && bottomLeft == "S" {
			fmt.Println("here")
			return true
		}

		if topRight == "S" && bottomRight == "S" && topLeft == "M" && bottomLeft == "M" {
			return true
		}

		return false
	})
}
