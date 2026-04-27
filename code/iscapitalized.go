// package code

// func IsCapitalized(s string) bool {
// 	if len(s) == 0 {
// 		return false
// 	}
// 	newWord := true
// 	for i := 0; i < len(s); i++ {
// 		c := s[i]
// 		if c == ' ' {
// 			newWord = true
// 			continue
// 		}
// 		if newWord {
// 			if c >= 'a' && c <= 'z' {
// 				return false
// 			}
// 			if c >= 'A' && c <= 'Z' {
// 				newWord = false
// 			}
// 		}
// 	}
// 	return true
// }

// package code

// func IsCapitalized(s string)bool{
// 	if len(s)==0{
// 		return false
// 	}

// 	newWord:=true
// 	for i:=0;i<len(s);i++{
// 		c:=s[i]
// 		if c==' '{
// 			newWord=true
// 			continue
// 		}
// 		if newWord{
// 			if c>='a' && c<='z'{
// 				return false
// 			}
// 			if c>='A' && c<='Z'{
// 				newWord=false
// 			}
// 		}
// 	}
// 	return true
// }

package code

func IsCapitalized(str string) bool {
	if len(str) == 0 {
		return false
	}

	newWord := true
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c == ' ' {
			newWord = true
			continue
		}
		if newWord {
			if c >= 'a' && c <= 'z' {
				return false
			}
			if c >= 'A' && c <= 'Z' {
				return false
			}
		}
	}
	return true
}
