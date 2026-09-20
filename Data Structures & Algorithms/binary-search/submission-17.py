class Solution:
    def search(self, nums: List[int], target: int) -> int:
        [l, r] = [0, len(nums)-1]
        while l <= r:
            m = (l + r) // 2
            print(l,m,r)
            if nums[m] == target:
                return m
            if nums[m] < target:
                l = m + 1
                continue
            r = m - 1
        return -1
        