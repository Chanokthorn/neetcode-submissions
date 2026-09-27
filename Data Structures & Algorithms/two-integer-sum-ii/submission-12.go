/*
1,2,3,4
^     ^
l	  r
sum(l, r) > target -> reduce r
sum(l, r) < target -> increase l
*/

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum > target {
			r -= 1
			continue
		}
		l += 1

	}
	return nil
}
