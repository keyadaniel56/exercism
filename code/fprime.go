// package code

// import (
// 	"fmt"
// 	"strconv"
// )

// func Fprime(args []string) {
// 	if len(args) != 1 {
// 		return
// 	}
// 	n, err := strconv.Atoi(args[0])
// 	if err != nil || n <= 1 {
// 		return
// 	}

// 	factors := []int{}
// 	for i := 2; i <= n; i++ {
// 		for n%i == 0 {
// 			factors = append(factors, i)
// 			n /= i
// 		}
// 	}

// 	if len(factors) == 0 {
// 		return
// 	}
// 	for i, factor := range factors {
// 		if i > 0 {
// 			fmt.Print("*")
// 		}
// 		fmt.Print(factor)
// 	}
// 	fmt.Println()
// }

package code

import (
	"fmt"
	"strconv"
)



func Fprime(arrgs []string){
	if len(arrgs)!=1{
		return
	}
	n,err:=strconv.Atoi(arrgs[0])
	if err!=nil || n<=1{
		return
	}

	factors:=[]int{}

	for i:=2;i<=n;i++{
		for n%i==0{
			factors = append(factors, i)
			n/=i
		}
	}

	if len(factors)==0{
		return
	}

	for i,factor:=range factors{
		if i>0{
			fmt.Print("*")
		}
		fmt.Print(factor)
	}
	fmt.Println()
}

