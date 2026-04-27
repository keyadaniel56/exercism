// package code

// func IsPrefix(str, prefix string) bool {
// 	if len(prefix) > len(str) {
// 		return false
// 	}

// 	s := []rune(str)
// 	p := []rune(prefix)

// 	for i := 0; i < len(p); i++ {
// 		if s[i] != p[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func IsSuffix(str, suffix string) bool {
// 	if len(suffix) > len(str) {
// 		return false
// 	}
// 	s := []rune(str)
// 	sfx := []rune(suffix)

// 	start := len(s) - len(sfx)
// 	for i := 0; i < len(sfx); i++ {
// 		if s[start+i] != sfx[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

package code

func IsPrefix(str, prefix string) bool {
	if len(prefix) > len(str) {
		return false
	}

	s := []rune(str)
	p := []rune(prefix)

	for i := 0; i < len(p); i++ {
		if s[i] != p[i] {
			return false
		}
	}
	return true
}

func IsSuffix(str, suffix string) bool {
	if len(suffix) > len(str) {
		return false
	}

	s := []rune(str)
	sfx := []rune(suffix)

	start := len(s) - len(sfx)

	for i := 0; i < len(sfx); i++ {
		if s[start+i] != sfx[i] {
			return false
		}
	}
	return true
}

func WordAnatomy(words []string, prefix, suffix string) string {
	result := ""
	for _, word := range words {
		if IsPrefix(word, prefix) {
			result += "prefix:" + prefix + "\n"
		}
		if IsSuffix(word, suffix) {
			result += "suffix:" + suffix
		}
	}
	return result
}
