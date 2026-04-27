package code

// func MinSubArrayLen(target int, nums []int) int {
//     left := 0
//     sum := 0
//     minLen := len(nums) + 1  // start with an impossible large value

//     for right := 0; right < len(nums); right++ {
//         sum += nums[right]  // expand window to the right

//         for sum >= target {
//             window := nums[left : right+1]  // current subarray
//             if len(window) < minLen {
//                 minLen = len(window)
//             }
//             sum -= nums[left]  // shrink window from the left
//             left++
//         }
//     }

//     if minLen == len(nums)+1 {
//         return 0  // no valid subarray found
//     }
//     return minLen
// }

func MinSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := len(nums) + 1

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= right {
			window := nums[left : right+1]
			if len(window) < minLen {
				minLen = len(window)
			}
			sum -= nums[left]
			left--
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
