// package code

// import (
// 	"strings"
// 	"unicode"
// )

// // func Longestword(str string) string {
// // 	if len(str) == 0 {
// // 		return str
// // 	}
// // 	longword := ""
// // 	for _, word := range strings.Fields(str) {
// // 		if IsLong(word) && len(word) > len(longword) {
// // 			longword = word
// // 		}
// // 	}
// // 	return longword
// // }

// func Longestword(str string) string {
// 	if len(str) == 0 {
// 		return str
// 	}
// 	longword := ""
// 	for _, word := range strings.Fields(str) {
// 		if IsLong(word) && len(word) > len(longword) {
// 			longword = word
// 		}
// 	}
// 	return longword
// }
// func IsLong(str string) bool {
// 	if len(str) == 0 {
// 		return false
// 	}
// 	for _, char := range str {
// 		if unicode.IsNumber(char) {
// 			return false
// 		}
// 	}
// 	return true
// }

package code

import (
	"strings"
	"unicode"
)

func IsLong(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, char := range str {
		if unicode.IsNumber(char) {
			return false
		}
	}
	return true
}

func Longestword(str string) string {
	if len(str) == 0 {
		return ""
	}
	longword := ""
	for _, word := range strings.Fields(str) {
		if IsLong(word) && len(word) > len(longword) {
			longword = word
		}
	}
	return longword
}
