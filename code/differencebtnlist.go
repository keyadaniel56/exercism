package code

func Findunique(a, b []int) []int {
	seenA := make(map[int]bool)
	seenb := make(map[int]bool)

	result := []int{}

	for _, num := range a {
		seenA[num] = true
	}

	for _, num := range b {
		seenb[num] = true
	}

	for _, num := range a {
		if seenA[num] && !seenb[num] {
			result = append(result, num)

		}
	}
	return result
}
