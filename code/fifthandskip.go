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

func FifthAndSkip(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	noSpaces := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			noSpaces += string(str[i])
		}
	}

	// check if less than 5 characters
	if len(noSpaces) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	i := 0

	for i < len(noSpaces) {
		count := 0
		for count < 5 && i < len(noSpaces) {
			result += string(noSpaces[i])
			i++
			count++
		}
		if i < len(noSpaces) {
			i++
			if i < len(noSpaces) {
				result += " "
			}
		}
	}
	return result + "\n"
}
