package code

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	// Convert number to string to get individual digits
	digits := strconv.Itoa(n)

	// Calculate sum of digits raised to consecutive powers
	sum := 0
	for i, digitChar := range digits {
		digit, _ := strconv.Atoi(string(digitChar))
		power := p + i
		sum += int(math.Pow(float64(digit), float64(power)))
	}

	// Check if sum is divisible by n
	if sum%n == 0 {
		return sum / n
	}

	return -1
}

