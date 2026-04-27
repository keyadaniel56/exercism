package code

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(n int) int {
	for i := n - 1; i >= 2; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return -1
}
