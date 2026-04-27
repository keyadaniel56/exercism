package code

func Lastword(str string) string {
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			return str[i+1:]
		}
	}
	return str
}
