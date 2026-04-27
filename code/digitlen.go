package code

func Digitlen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	count := 0
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return 1
	}
	for n > 0 {
		n /= base
		count++
	}
	return count
}
