package code

func RemoveTrailing(word string) string {
	if len(word) == 0 {
		return ""
	}
	start := 0
	end := len(word) - 1

	for start <= end && word[start] == ' ' {
		start++
	}
	for end >= start && word[end] == ' ' {
		end--
	}
	return word[start : end+1]
}

func Firstword(word string) string {
	str := RemoveTrailing(word)
	start := 0
	end := len(word) - 1
	for start <= end && str[start] != ' ' {
		start++
	}
	for end >= start && str[end] != ' ' {
		end--
	}
	return str[:start]
}
