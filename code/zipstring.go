package code

import "strconv"

func ZipString(s string) string {
	if len(s) == 0 {
		return s
	}
	result := ""
	i := 0
	for i < len(s) {
		count := 1
		for i+count < len(s) && s[i+count] == s[i] {
			count++
		}
		result += strconv.Itoa(count) + string(s[i])
		i += count
	}
	return result
}
