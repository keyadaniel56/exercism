package code

func Trimspace(str string) string {
	if len(str) == 0 {
		return ""
	}

	result := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' && i > 0 && str[i-1] == ' ' {
			continue
		}
		result += string(str[i])
	}
	return result
}

func RemoveTrailingSpace(str string) string {
	start := 0
	end := len(str) - 1

	for start <= end && str[start] == ' ' {
		start++
	}

	for end >= start && str[end] == ' ' {
		end--
	}

	return str[start : end+1]
}

func Cleanstr(str string) string {
	notrail := RemoveTrailingSpace(str)
	cleanded := Trimspace(notrail)
	return cleanded
}
