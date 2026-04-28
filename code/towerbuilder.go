package code

func TowerBuilder(nFloors int) []string {
	result := []string{}

	for i := 1; i <= nFloors; i++ {
		// Number of stars for this floor: 2*i - 1
		stars := 2*i - 1

		// Number of spaces on each side: (maxWidth - stars) / 2
		// maxWidth is the width of the bottom floor: 2*nFloors - 1
		maxWidth := 2*nFloors - 1
		spaces := (maxWidth - stars) / 2

		// Build the floor string
		floor := ""
		for j := 0; j < spaces; j++ {
			floor += " "
		}
		for j := 0; j < stars; j++ {
			floor += "*"
		}
		for j := 0; j < spaces; j++ {
			floor += " "
		}

		result = append(result, floor)
	}

	return result
}
