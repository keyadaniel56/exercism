package main

import (
	"fmt"

	"checkin/arrays"
)

func main() {
	// Test MaxProfit function
	fmt.Println("Testing MaxProfit function:")

	// Test case 1: [7, 10, 1, 3, 6, 9, 2] -> Expected: 8
	prices1 := []int{7, 10, 1, 3, 6, 9, 2}
	result1 := arrays.MaxProfit(prices1)
	fmt.Printf("Input: %v, Output: %d (Expected: 8)\n", prices1, result1)

	// Test case 2: [7, 6, 4, 3, 1] -> Expected: 0
	prices2 := []int{7, 6, 4, 3, 1}
	result2 := arrays.MaxProfit(prices2)
	fmt.Printf("Input: %v, Output: %d (Expected: 0)\n", prices2, result2)

	// Test case 3: [1, 3, 6, 9, 11] -> Expected: 10
	prices3 := []int{1, 3, 6, 9, 11}
	result3 := arrays.MaxProfit(prices3)
	fmt.Printf("Input: %v, Output: %d (Expected: 10)\n", prices3, result3)

	// Original test
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	fmt.Println("MoveAllZeroes result:", arrays.MoveAllZeroes(arr))
}
