package arrays

func FindUnion(arra []int, arrb []int) []int {
	arra = RemoveDuplicates(arra)
	arrb = RemoveDuplicates(arrb)

	result := []int{}
	seenA := make(map[int]bool)
	seenB := make(map[int]bool)

	for i := 0; i < len(arra); i++ {
		seenA[arra[i]] = true
	}

	for i := 0; i < len(arrb); i++ {
		seenB[arrb[i]] = true
	}

	// Add all elements from first array
	for _, char := range arra {
		result = append(result, char)
	}

	// Add elements from second array that aren't already in first array
	for _, char := range arrb {
		if !seenA[char] {
			result = append(result, char)
		}
	}
	return result
}

func RemoveDuplicates(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	seen := make(map[int]bool)
	result := []int{}

	for i := 0; i < len(arr); i++ {
		if !seen[arr[i]] {
			result = append(result, arr[i])
			seen[arr[i]] = true
		}
	}
	return result
}
