package arrays

func MoveAllZeroes(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	result := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			result = append(result, arr[i])
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			result = append(result, arr[i])
		}
	}
	return result
}
