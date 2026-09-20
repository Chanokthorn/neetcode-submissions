/*
3,4,5,6
^
look for 4
{3: 0}
  ^
look for 3
{}
*/
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		if j, ok := seen[target - num]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return []int{-1, -1}
}
