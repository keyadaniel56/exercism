package code

import "fmt"

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	dotPos := -1
	for i := 0; i < len(dec); i++ {
		if dec[i] == '.' {
			if dotPos != -1 {
				return dec + "\n"
			}
			dotPos = i
		} else if dec[i] == '-' && i != 0 {
			return dec + "\n"
		} else if dec[i] < '0' || dec[i] > '9' {
			return dec + "\n"
		}
	}

	if dotPos == -1 {
		return dec + "\n"
	}

	allZeros := true
	for i := dotPos + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			allZeros = false
			break
		}
	}
	if allZeros {
		return dec[:dotPos] + "\n"
	}

	result := dec[:dotPos] + dec[dotPos+1:]
	start := 0
	if result[0] == '-' {
		start = 1
	}
	for start < len(result) && result[start] == '0' {
		start++
	}
	if start >= len(result) {
		return "0\n"
	}
	if result[0] == '-' {
		return "-" + result[start:] + "\n"
	}
	return result[start:] + "\n"
}

func Rub() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
