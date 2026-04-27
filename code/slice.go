package code

func Slice(a []string, nbrs ...int) []string {
	n := len(a)

	normalize := func(i int) int {
		if i < 0 {
			i = n + i
		}
		if i < 0 {
			i = 0
		}
		if i > n {
			i = n
		}
		return i
	}

	if len(nbrs) == 1 {
		return a[normalize(nbrs[0]):]
	}
	start := normalize(nbrs[0])
	end := normalize(nbrs[1])
	if start >= end {
		return []string{}
	}
	return a[start:end]
}
