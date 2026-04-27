// package code

// func Solution(A []int) int {
// 	seen := make(map[int]bool)
// 	duplicate := make(map[int]bool)
// 	for _, num := range A {
// 		if seen[num] {
// 			duplicate[num] = true
// 		}
// 		seen[num] = true
// 	}
// 	for _, num := range A {
// 		if !duplicate[num] {
// 			return num
// 		}
// 	}
// 	return -1
// }

package code

// func Solution(A []int)int{
// 	seen:=make(map[int]bool)
// 	duplicate:=make(map[int]bool)

// 	for _,num:=range A{
// 		if seen[num]{
// 			duplicate[num]=true
// 		}
// 		seen[num]=true
// 	}

// 	for _,num:=range A{
// 		if !duplicate[num]{
// 			return num
// 		}
// 	}
// 	return -1
// }

func Solution(A []int) int {
	seen := make(map[int]bool)
	duplicate := make(map[int]bool)

	for _, num := range A {
		if seen[num] {
			duplicate[num] = true
		}
		seen[num] = true
	}

	for _, num := range A {
		if !duplicate[num] {
			return num
		}
	}
	return -1
}
