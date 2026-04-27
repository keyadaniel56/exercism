// package code

// import "strings"

// func DuplicateCount(s string) int {
// 	s = strings.ToLower(s)
// 	counts := map[rune]int{}

// 	for _, ch := range s {
// 		counts[ch]++
// 	}

// 	duplicates := 0
// 	for _, count := range counts {
// 		if count > 1 {
// 			duplicates++
// 		}
// 	}

// 	return duplicates
// }

package code

import "strings"

func Duplicate_count(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.ToLower(s)

	counts := map[rune]int{}
	for _, char := range s {
		counts[char]++
	}
	duplicates := 0
	for _, count := range counts {
		if count > 1 {
			duplicates++
		}
	}
	return duplicates
}
