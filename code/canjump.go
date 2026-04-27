package code

func CanJump(step []uint) bool {
	if len(step) == 0 {
		return false
	}
	if len(step) == 1 {
		return true
	}
	visited := make(map[int]bool)
	pos := 0
	for {
		if pos == len(step)-1 {
			return true
		}
		if pos < 0 || pos >= len(step) {
			return false
		}
		if visited[pos] {
			return false
		}
		visited[pos] = true
		pos += int(step[pos])
	}
}
