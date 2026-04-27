

package code

func ToCamelCase(str string) string {
	if len(str) == 0 {
		return str
	}

	b := []rune(str)
	result:=[]rune{}
	for i := 0; i < len(b); i++ {
		if b[i] == '_' || b[i] == '-' {
			if i+1 < len(b) && b[i+1] >= 'a' && b[i+1] <= 'z' {
				i++
				result = append(result, b[i]-32)
			}
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)
}
