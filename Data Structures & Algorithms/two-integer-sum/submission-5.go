/*
nums[i] + nums[j] == target
nums[j] == target - nums[i]
*/
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
		targetMinus := target - nums[i]
		for j := i+1; j < len(nums); j++ {
			if nums[j] == targetMinus {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
