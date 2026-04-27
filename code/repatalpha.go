package code

func RepeatAlpha(str string) string {
	if len(str) == 0 {
		return str
	}

	result := ""
	b := []rune(str)
	for i := 0; i < len(b); i++ {
		if IsLower(b[i]) {
			count := int(b[i] - 'a' + 1)
			for j := 0; j < count; j++ {
				result += string(b[i])
			}
		} else if IsUpper(b[i]) {
			count := int(b[i] - 'A' + 1)
			for j := 0; j < count; j++ {
				result += string(b[i])
			}
		}
	}
	return result
}

func IsLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
