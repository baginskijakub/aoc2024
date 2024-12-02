package logic

func FindDistance(x []int, y []int) int {
	for i := range x {
		for j := range x {
			if x[i] > x[j] {
				temp := x[i]

				x[i] = x[j]
				x[j] = temp
			}

			if y[i] > y[j] {
				temp := y[i]

				y[i] = y[j]
				y[j] = temp
			}
		}
	}

	diff := 0

	for i := range x {
		diff += abs(x[i] - y[i])
	}

	return diff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
