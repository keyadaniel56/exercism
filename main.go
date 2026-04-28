package main

import (
	"checkin/arrays"
	"fmt"
)

func main() {
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	fmt.Println(arrays.MoveAllZeroes(arr))
}
