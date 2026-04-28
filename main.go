package main

import (
	"fmt"

	"checkin/arrays"
)

func main() {
	// Test MaxProfit function
Input:= []int{1, 4, 45, 6, 10, 8}
target:= 13

fmt.Println(arrays.HasTripletSum(Input,target))
}
