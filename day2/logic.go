package day2

func GetSafeReports() int {
	matrix, err := ReadFile()

	if err != nil {
		return 0
	}

	safeCount := 0
	for _, arr := range matrix {
		if checkVariations(arr) {
			safeCount++
		}
	}
	return safeCount
}

func checkVariations(arr []int) bool {
	if check(arr) {
		return true
	}

	for i := 0; i < len(arr); i++ {
		temp := remove(arr, i)

		if check(temp) {
			return true
		}
	}

	return false
}

func check(arr []int) bool {
	delta := getDeltas(arr)
	isDesc := func(i int) bool { return i <= -1 && i >= -3 }
	isAsc := func(i int) bool { return i >= 1 && i <= 3 }

	desc := all(delta, isDesc)
	asc := all(delta, isAsc)

	return desc || asc
}

func getDeltas(arr []int) []int {
	deltas := make([]int, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		deltas[i] = arr[i] - arr[i+1]
	}
	return deltas
}

func all(s []int, pred func(int) bool) bool {
	for _, t := range s {
		if !pred(t) {
			return false
		}
	}

	return true
}

func remove(s []int, i int) []int {
	r := make([]int, 0)
	r = append(r, s[:i]...)
	return append(r, s[i+1:]...)
}
