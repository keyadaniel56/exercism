package code



// func SaveAndMiss(arg string, num int) string {
//     if num <= 0 {
//         return arg
//     }
//     result := ""
//     for i := 0; i < len(arg); i++ {
//         groupIndex := i / num  // Which group (0, 1, 2, ...)
//         if groupIndex%2 == 0 { // Even groups = save
//             result += string(arg[i])
//         }
//         // Odd groups = miss (do nothing)
//     }
//     return result
// }


func SaveAndMiss(arg string, num int)string{
	if num<=0{
		return arg
	}

	result:=""
	for i:=0;i<len(arg);i++{
		groupIndex:=i/num
		if groupIndex%2==0{
			result+=string(arg[i])
		}
	}
	return result
}