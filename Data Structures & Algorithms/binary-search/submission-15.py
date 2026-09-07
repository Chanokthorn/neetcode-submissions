class Solution:
    def search(self, nums: List[int], target: int) -> int:
        [l, r] = [0, len(nums) - 1]
        while l < r:
            m = (l + r) // 2
            if target == nums[m]:
                return m
            if nums[m] < target:
                l = m + 1
                continue
            r = m - 1
        if nums[l] == target:
            return l
        return -1