func maxProfit(prices []int) int {
 var l, r, max int
 for i, _ := range prices {
    if prices[i] < prices[l] {
        l = i
        r = l
    }
    if prices[i] > prices[r] {
        r = i
    }
    diff := prices[r] - prices[l]
    if diff > max {
        max = diff
    }
 }

 return max
}
