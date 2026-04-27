package code

func Expandstr(str string) string {
	cleaned := Cleanstr(str)

	space := "   "
	result := ""

	for i := 0; i < len(cleaned); i++ {
		if cleaned[i] == ' ' {
			result += space
		}
		result += string(cleaned[i])
	}
	return result
}
