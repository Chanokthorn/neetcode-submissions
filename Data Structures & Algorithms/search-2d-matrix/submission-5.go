func searchMatrix(matrix [][]int, target int) bool {
    row := findRow(matrix, target)
    if row == -1 {
        return false
    }
    return bs(matrix[row], target)
}

func findRow(matrix [][]int, target int) int {
    for i, row := range matrix {
        if row[0] <= target && row[len(row) - 1] >= target {
            return i
        }
    }
    return -1
}

func bs(nums []int, target int) bool {
    l := 0
    r := len(nums) - 1
    for l <= r {
        m := (l + r) / 2
        if nums[m] == target {
            return true
        }
        if nums[m] < target {
            l = m + 1
            continue
        }
        r = m - 1
    }
    return false
}
