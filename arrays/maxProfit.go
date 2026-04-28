

package arrays



func MaxProfit(arr []int)int{
	if len(arr)<=1{
		return 0
	}

	MaxProfit:=0
	minPrice:=arr[0]

	for i:=1;i<len(arr);i++{
		if arr[i]<minPrice{
			minPrice=arr[i]
		}
		currentProfit:=arr[i]-minPrice
		if currentProfit>MaxProfit{
			MaxProfit=currentProfit
		}
	}
	return MaxProfit
}