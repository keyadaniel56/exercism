// package code

// func Deadfish(program string) []int {
// 	value := 0
// 	result := []int{}

// 	for _, cmd := range program {
// 		switch cmd {
// 		case 'i':
// 			value++
// 		case 'd':
// 			value--
// 		case 's':
// 			value = value * value
// 		case 'o':
// 			result = append(result, value)
// 		}
// 	}

// 	return result
// }

package code

func Deadfish(program string) []int {
	value := 0
	result := []int{}

	for _, cmd := range program {
		switch cmd {
		case 'i':
			value++
		case 'd':
			value--
		case 's':
			value *= value
		case 'o':
			result = append(result, value)
		}
	}
	return result
}
