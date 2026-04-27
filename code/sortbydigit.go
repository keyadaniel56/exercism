// package code

// import (
// 	"sort"
// 	"strings"
// )

// func SortByDigit(s string) string {
// 	if s == "" {
// 		return ""
// 	}

// 	words := strings.Split(s, " ")

// 	sort.Slice(words, func(i, j int) bool {
// 		return digitIn(words[i]) < digitIn(words[j])
// 	})

// 	return strings.Join(words, " ")
// }

// func digitIn(word string) byte {
// 	for i := 0; i < len(word); i++ {
// 		if word[i] >= '1' && word[i] <= '9' {
// 			return word[i]
// 		}
// 	}
// 	return 0
// }

package code

import (
	"sort"
	"strings"
)

func digitIn(word string) byte {
	for i := 0; i < len(word); i++ {
		if word[i] >= '1' && word[i] <= '9' {
			return word[i]
		}
	}
	return 0
}

func SortByDigit(s string) string {
	if len(s) == 0 {
		return s
	}

	words := strings.Split(s, " ")

	sort.Slice(words, func(i, j int) bool {
		return digitIn(words[i]) < digitIn(words[j])
	})
	return strings.Join(words, " ")
}
