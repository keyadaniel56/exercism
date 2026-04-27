package code

import "fmt"

// func ConcatAlternate(arr1, arr2 []int) []int {
// 	len1, len2 := len(arr1), len(arr2)
// 	minLen := len1
// 	if len2 < minLen {
// 		minLen = len2
// 	}

// 	result := []int{}

// 	for i := 0; i < minLen; i++ {
// 		result = append(result, arr2[i])
// 		result = append(result, arr1[i])
// 	}

// 	// Append remaining elements from the longer array
// 	if len1 > minLen {
// 		result = append(result, arr1[minLen:]...)
// 	} else if len2 > minLen {
// 		result = append(result, arr2[minLen:]...)
// 	}

// 	return result
// }

func Output() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(arr1, arr2 []int) []int {
	len1, len2 := len(arr1), len(arr2)
	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	result := []int{}

	for i := 0; i < minLen; i++ {
		result = append(result, arr2[i])
		result = append(result, arr1[i])
	}

	if len1 > minLen {
		result = append(result, arr1[minLen:]...)
	} else if len2 > minLen {
		result = append(result, arr2[minLen:]...)
	}
	return result
}
