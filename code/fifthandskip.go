package code

// func FifthAndSkip(str string) string {
//     if len(str) == 0 {
//         return "\n"
//     }

//     // Remove spaces from the string
//     noSpaces := ""
//     for i := 0; i < len(str); i++ {
//         if str[i] != ' ' {
//             noSpaces += string(str[i])
//         }
//     }

//     // Check if less than 5 characters
//     if len(noSpaces) < 5 {
//         return "Invalid Input\n"
//     }

//     result := ""
//     i := 0

//     for i < len(noSpaces) {
//         // Take up to 5 characters
//         count := 0
//         for count < 5 && i < len(noSpaces) {
//             result += string(noSpaces[i])
//             i++
//             count++
//         }

//         // Skip the 6th character
//         if i < len(noSpaces) {
//             i++
//             // Add space if there are more characters coming
//             if i < len(noSpaces) {
//                 result += " "
//             }
//         }
//     }

//     return result + "\n"
// }


package main

import "fmt"
func FifthAndSkip(str string)string{
	if len(str)==0{
		return "\n"
	}

	var runes []rune
	for _,r:=range str{
		if str[r]!=' '{
			runes=append(runes,r)
		}
	}

	if len(runes)<5{
		return "Invalid\n"
	}

	var resultRunes []rune
	i:=0
	for i<len(runes){
		end:=i+5
		if end>len(runes){
			end=len(runes)
		}
		resultRunes=append(resultRunes,runes[i:end]...)

		if end-i==5 && end<len(runes){
			resultRunes=append(resultRunes,' ')
		}
		i+=6
	}
	return string(resultRunes)
}

func main(){
	fmt.Println(FifthAndSkip("HelloWorld"))
}
