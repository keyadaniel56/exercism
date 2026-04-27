// package code

// func SearchReplace(word, str1, str2 string) string {
// 	if len(word) == 0 {
// 		return word
// 	}
// 	result := ""
// 	for _, char := range word {
// 		if string(char) == str1 {
// 			result += str2
// 		} else {
// 			result += string(char)
// 		}
// 	}
// 	return result
// }

package code

func SearchReplace(word, str1, str2 string) string {
	if len(word) == 0 {
		return word
	}

	result := ""
	for _, char := range word {
		if string(char) == str1 {
			result += string(str2)
		} else {
			result += string(char)
		}
	}
	return result
}
