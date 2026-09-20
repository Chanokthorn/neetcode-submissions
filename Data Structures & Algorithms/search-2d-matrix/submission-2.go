func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    nums := make([]int, m * n)

    for i, row := range matrix {
        for j, cell := range row {
            nums[i * n + j] = cell
        }
    }

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
