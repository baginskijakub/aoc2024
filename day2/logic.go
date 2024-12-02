package day2

func GetSafeReports() int {
	matrix, err := ReadFile()

	if err != nil {
		return 0
	}

	safeCount := 0

	for _, arr := range matrix {
		if computeRowVariations(arr) {
			safeCount++
		}
	}

	return safeCount
}

func computeRowVariations(arr []int) bool {
	if computeRow(arr) {
		return true
	}

	for i := range arr {
		variation := remove(arr, i)

		if computeRow(variation) {
			return true
		}
	}

	return false
}

func computeRow(arr []int) bool {
	var shouldIncrease *bool

	prev := arr[0]

	for i := 1; i < len(arr); i++ {
		curr := arr[i]

		// precondition to figure out if should increase
		if shouldIncrease == nil {
			if prev < curr {
				shouldIncrease = new(bool)
				*shouldIncrease = true
			}

			if prev > curr {
				shouldIncrease = new(bool)
				*shouldIncrease = false
			}

			if prev == curr {
				return false
			}
		}

		if *shouldIncrease && prev > curr {
			return false
		}

		if !*shouldIncrease && prev < curr {
			return false
		}

		diff := abs(prev - curr)

		if diff > 3 || diff == 0 {
			return false
		}

		prev = arr[i]
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
