// package code

// func MaxSubSum(arr []int) int {
// 	if len(arr) == 0 {
// 		return 0
// 	}

// 	maxSoFar := 0
// 	maxEndingHere := 0

// 	for _, num := range arr {
// 		maxEndingHere = max(0, maxEndingHere+num)
// 		maxSoFar = max(maxSoFar, maxEndingHere)
// 	}

// 	return maxSoFar
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

package code

func MaxSubSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxSoFar := 0
	maxEndingHere := 0

	for _, num := range arr {
		maxEndingHere = max(0, maxEndingHere+num)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
