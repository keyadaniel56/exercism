package code

func Comp(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	if len(arr1) == 0 || len(arr2) == 0 {
		return false
	}

	digit := make(map[int]int)

	for _, num := range arr1 {
		digit[num*num]++
	}

	for _, num := range arr2 {
		if digit[num] == 0 {
			return false
		}
		digit[num]++
	}
	return true

}
