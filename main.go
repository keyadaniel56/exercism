package main

import (
	"fmt"

	"checkin/arrays"
)

func main() {
	
a:= []int{1, 2, 3, 2, 1}
 b:=[]int {3, 2, 2, 3, 3, 2}
fmt.Println(arrays.FindUnion(a,b))
}
