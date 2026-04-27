package code

func HashCode(dec string) string {
	size := len(dec)
	result := []byte{}
	for i := 0; i < size; i++ {
		c := int(dec[i]) + size
		if c < 32 {
			c += 32
		}
		result = append(result, byte(c))
	}
	return string(result)

}
