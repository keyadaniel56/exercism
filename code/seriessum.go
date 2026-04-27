// package code

// import "fmt"

// func SeriesSum(n int) string {
// 	sum := 0.0
// 	for i := 0; i < n; i++ {
// 		sum += 1.0 / float64(1+i*3)
// 	}
// 	return fmt.Sprintf("%.2f", sum)
// }

package code

import "fmt"

func SeriesSum(n int) string {
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += 1.0 / float64(i+i*3)
	}
	return fmt.Sprintf("%.2f", sum)
}
