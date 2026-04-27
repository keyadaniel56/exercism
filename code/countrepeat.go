package code

import "strconv"

func CountRepeat(str string) string {
	if len(str) == 0 {
		return ""
	}

	count := 1
	result := ""
	b := []rune(str)

	for i := 1; i < len(b); i++ {
		if b[i] == b[i-1] {
			count++
		} else {
			if count > 1 {
				result += string(b[i-1]) + strconv.Itoa(count)
			} else {
				result += string(b[i-1])
			}
			count = 1
		}
	}

	if count > 1 {
		result += string(b[len(b)-1]) + strconv.Itoa(count)
	} else {
		result += string(b[len(b)-1])
	}

	return result
}
