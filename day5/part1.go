package day5

import "fmt"

func Part1() {
	rules, updates, err := ReadFile()

	if err != nil {
		panic("Couldnt read file")
	}

	sum := 0

	for _, update := range updates {
		midValue := update[len(update)/2]

		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				for _, rule := range rules {
					preceeding := rule[0]
					proceeding := rule[1]

					if preceeding == update[j] && proceeding == update[i] {
						midValue = 0
					}

				}
			}

		}

		sum += midValue
	}

	fmt.Println(sum)
}
