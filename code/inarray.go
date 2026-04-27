// package code

// import (
// 	"sort"
// 	"strings"
// )

// func InArray(a1 []string, a2 []string) []string {
// 	seen := map[string]bool{}
// 	result := []string{}

// 	for _, s1 := range a1 {
// 		for _, s2 := range a2 {
// 			if strings.Contains(s2, s1) && !seen[s1] {
// 				seen[s1] = true
// 				result = append(result, s1)
// 				break
// 			}
// 		}
// 	}

// 	sort.Strings(result)
// 	return result
// }

package code

import (
	"sort"
	"strings"
)

func InArray(a1 []string, a2 []string) []string {
	seen := map[string]bool{}
	result := []string{}

	for _, s1 := range a1 {
		for _, s2 := range a2 {
			if strings.Contains(s2, s1) && !seen[s1] {
				seen[s1] = true
				result = append(result, s1)
				break
			}
		}
	}
	sort.Strings(result)
	return result
}
