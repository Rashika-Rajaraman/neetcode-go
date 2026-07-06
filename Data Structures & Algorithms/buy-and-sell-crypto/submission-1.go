func maxProfit(prices []int) int {
    l, r := 0, 1
    maxP := 0

    // Time Complexity - O(n)
    // Space Complexity - O(1)
    for r < len(prices) {
        if prices[l] < prices[r] {
            profit := prices[r] - prices[l]
            if profit > maxP {
                maxP = profit
            }
        } else {
            l=r
        }
    r++
    }
    return maxP
}
