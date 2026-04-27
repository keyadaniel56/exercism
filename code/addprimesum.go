package code

import "os"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func AddprimeSum(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			count += i
		}
	}
	return count
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int(c-'0')
	}
	return n
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(rune('0'+n%10)) + result
		n /= 10
	}
	return result
}

func AddprimeSumCLI() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("0\n")
		return
	}

	n := atoi(os.Args[1])
	if n < 0 {
		os.Stdout.WriteString("0\n")
		return
	}

	os.Stdout.WriteString(itoa(AddprimeSum(n)) + "\n")
}
