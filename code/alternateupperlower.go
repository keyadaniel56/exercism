package code

import "fmt"

func Display() {
	for i := 'a'; i <= 'z'; i++ {
		if i%2 != 0 {
			fmt.Printf("%c ", i-('a'-'A')) // Convert i to uppercase
		} else {
			fmt.Printf("%c", i) // Keep i lowercase
		}
	}
	fmt.Println()
}
