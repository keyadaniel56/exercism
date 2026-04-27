// package code

// import "fmt"

// func BinaryGap(N int) int {
// 	bits := fmt.Sprintf("%b", N)
// 	max, count := 0, 0
// 	for _, b := range bits {
// 		if b == '1' {
// 			if count > max {
// 				max = count
// 			}
// 			count = 0
// 		} else {
// 			count++
// 		}
// 	}
// 	return max
// }

package code

import "fmt"

func BinaryGap(N int) int {
	bits := fmt.Sprintf("%b", N)
	max, count := 0, 0
	for _, b := range bits {
		if b == '1' {
			if count > max {
				max = count
			}
			count = 0
		} else {
			count++
		}
	}
	return max
}
