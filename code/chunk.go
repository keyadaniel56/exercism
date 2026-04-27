package code

func Chunk(a []string, size int) [][]string {
	var result [][]string
	for i := 0; i < len(a); i += size {
		end := i + size
		if end > len(a) {
			end = len(a)
		}
		result = append(result, a[i:end])
	}
	return result
}
