package day4

import "fmt"

func Find() {
	matrix, err := ReadFile()

	if err != nil {
		panic(err)
	}

	sum := 0

	for i, row := range matrix {
		for j, col := range row {
			if col == "X" {
				sum += checkX(matrix, i, j)
			}
		}
	}

	fmt.Println(sum)
}

func checkX(matrix [][]string, i int, j int) int {
	sum := 0

	if checkHorizontalLeftRight(matrix, i, j) {
		sum++
	}

	if checkHorizontalRightLeft(matrix, i, j) {
		sum++
	}

	if checkVerticalTopBottom(matrix, i, j) {
		sum++
	}

	if checkVerticalBottomTop(matrix, i, j) {
		sum++
	}

	if checkDiagonalTopLeftBottomRight(matrix, i, j) {
		sum++
	}

	if checkDiagonalBottomRightTopLeft(matrix, i, j) {
		sum++
	}

	if checkDiagonalTopRightBottomLeft(matrix, i, j) {
		sum++
	}

	if checkDiagonalBottomLeftTopRight(matrix, i, j) {
		sum++
	}

	return sum
}

func safelyCheck(f func() bool) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	return f()
}

func checkHorizontalLeftRight(matrix [][]string, i int, j int) bool {
	return safelyCheck(func() bool {
		x := matrix[i][j]
		m := matrix[i][j+1]
		a := matrix[i][j+2]
		s := matrix[i][j+3]

		return x == "X" && m == "M" && a == "A" && s == "S"
	})
}

func checkHorizontalRightLeft(matrix [][]string, i int, j int) bool {
	return safelyCheck(func() bool {
		x := matrix[i][j]
		m := matrix[i][j-1]
		a := matrix[i][j-2]
		s := matrix[i][j-3]

		return x == "X" && m == "M" && a == "A" && s == "S"
	})
}

func checkVerticalTopBottom(matrix [][]string, i int, j int) bool {
	return safelyCheck(func() bool {
		x := matrix[i][j]
		m := matrix[i+1][j]
		a := matrix[i+2][j]
		s := matrix[i+3][j]

		return x == "X" && m == "M" && a == "A" && s == "S"
	})
}

func checkVerticalBottomTop(matrix [][]string, i int, j int) bool {
	return safelyCheck(func() bool {
		x := matrix[i][j]
		m := matrix[i-1][j]
		a := matrix[i-2][j]
		s := matrix[i-3][j]

		return x == "X" && m == "M" && a == "A" && s == "S"
	})
}

func checkDiagonalTopLeftBottomRight(matrix [][]string, i int, j int) bool {
	return safelyCheck(func() bool {
		x := matrix[i][j]
		m := matrix[i+1][j+1]
		a := matrix[i+2][j+2]
		s := matrix[i+3][j+3]

		return x == "X" && m == "M" && a == "A" && s == "S"
	})
}

func checkDiagonalBottomRightTopLeft(matrix [][]string, i int, j int) bool {
	return safelyCheck(func() bool {
		x := matrix[i][j]
		m := matrix[i-1][j-1]
		a := matrix[i-2][j-2]
		s := matrix[i-3][j-3]

		return x == "X" && m == "M" && a == "A" && s == "S"
	})
}

func checkDiagonalTopRightBottomLeft(matrix [][]string, i int, j int) bool {
	return safelyCheck(func() bool {
		x := matrix[i][j]
		m := matrix[i-1][j+1]
		a := matrix[i-2][j+2]
		s := matrix[i-3][j+3]

		return x == "X" && m == "M" && a == "A" && s == "S"
	})
}

func checkDiagonalBottomLeftTopRight(matrix [][]string, i int, j int) bool {
	return safelyCheck(func() bool {
		x := matrix[i][j]
		m := matrix[i+1][j-1]
		a := matrix[i+2][j-2]
		s := matrix[i+3][j-3]

		return x == "X" && m == "M" && a == "A" && s == "S"
	})
}
