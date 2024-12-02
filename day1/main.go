package main

import (
	"day1/logic"
	read "day1/read"
	"fmt"
)

func main() {
	x, y, err := read.ReadFile()

	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	fmt.Println(logic.FindDistance(x, y))
}
