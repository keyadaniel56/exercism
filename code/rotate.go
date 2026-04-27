package code

func Rotate(a []string, n int) []string {
	if len(a) == 0 {
		return a
	}
	n = n % len(a)
	if n < 0 {
		n = n + len(a)
	}
	return append(a[n:], a[:n]...)
}
