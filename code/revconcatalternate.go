package code

import "fmt"

// func RevConcatAlternate(slice1, slice2 []int) []int {
// 	len1, len2 := len(slice1), len(slice2)
// 	result := []int{}

// 	// Determine which slice is larger
// 	var larger []int
// 	var largerLen, smallerLen int

// 	if len1 > len2 {
// 		larger = slice1
// 		largerLen = len1
// 		smallerLen = len2
// 	} else if len2 > len1 {
// 		larger = slice2
// 		largerLen = len2
// 		smallerLen = len1
// 	} else {
// 		// Equal length
// 		largerLen = len1
// 		smallerLen = len2
// 	}

// 	// Add extra elements from larger slice (in reverse)
// 	for i := largerLen - 1; i >= smallerLen; i-- {
// 		result = append(result, larger[i])
// 	}

// 	// Alternate from the remaining elements (in reverse), always starting with slice1
// 	for i := smallerLen - 1; i >= 0; i-- {
// 		result = append(result, slice1[i])
// 		result = append(result, slice2[i])
// 	}

// 	return result
// }

func Dis() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	len1, len2 := len(slice1), len(slice2)

	//Determine which slice is larger
	var larger []int
	var largerLen, smallerLen int

	result := []int{}

	if len1 > len2 {
		larger = slice1
		largerLen = len1
		smallerLen = len2
	} else if len2 > len1 {
		larger = slice2
		largerLen = len2
		smallerLen = len1
	}else{
		largerLen=len1
		smallerLen=len2
	}

	for i := largerLen - 1; i >= smallerLen; i-- {
		result = append(result, larger[i])
	}

	for i := smallerLen - 1; i >= 0; i-- {
		result = append(result, slice1[i])
		result = append(result, slice2[i])
	}
	return result
}
