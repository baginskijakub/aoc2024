package day5

import "fmt"

func Part2() {
	rules, updates, err := ReadFile()

	if err != nil {
		panic("Couldnt read file")
	}

	sum := 0

	for _, update := range updates {
		isValid := true
		tempUpdate := update

		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				for _, rule := range rules {
					preceeding := rule[0]
					proceeding := rule[1]

					if preceeding == update[j] && proceeding == update[i] {
						isValid = false
						temp := update[j]
						update[j] = update[i]
						update[i] = temp
					}

				}
			}
		}

		if !isValid {
			sum += tempUpdate[len(tempUpdate)/2]
		}
	}

	fmt.Println(sum)
}
