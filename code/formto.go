// // package code

// // import "strconv"

// // func pad(n int) string {
// // 	if n < 10 {
// // 		return "0" + strconv.Itoa(n)
// // 	}
// // 	return strconv.Itoa(n)
// // }

// // func FromTo(from, to int) string {
// // 	if from > 99 || from < 0 || to > 99 || to < 0 {
// // 		return "Invalid\n"
// // 	}
// // 	result := ""
// // 	if from <= to {
// // 		for i := from; i <= to; i++ {
// // 			if i != from {
// // 				result += ", "
// // 			}
// // 			result += pad(i)
// // 		}
// // 	} else {
// // 		for i := from; i >= to; i-- {
// // 			if i != from {
// // 				result += ", "
// // 			}
// // 			result += pad(i)
// // 		}
// // 	}
// // 	return result + "\n"
// // }

// package code

// import "strconv"

// func Pad(n int)string{
// 	if n<10{
// 		return "0"+strconv.Itoa(n)
// 	}
// 	return strconv.Itoa(n)
// }

// func FromTo(from, to int)string{
// 	if from>99 || from<0 || to>99 || to<0{
// 		return "Invalid\n"
// 	}

// 	result:=""
// 	if from<to{
// 		for i:=from;i<=to;i++{
// 			if i!=from{
// 				result+=","
// 			}
// 			result+=Pad(i)
// 		}
// 	}else{
// 		for i:=from;i>=to;i--{
// 			if i!=from{
// 				result+=","
// 			}
// 			result+=Pad(i)
// 		}
// 	}
// 	return result+"\n"
// }

package code

import "strconv"

func Pad(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}

func FromTo(from, to int) string {
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}
	result := ""
	if from < to {
		for i := from; i < to; i++ {
			if i != from {
				result += ","
			}
			result += Pad(i)
		}
	} else {
		for i := from; i >= to; i-- {
			if i != from {
				result += ","
			}
			result += Pad(i)
		}
	}
	return result + "\n"
}
