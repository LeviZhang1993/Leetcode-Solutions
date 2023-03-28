package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func mincostTickets(days []int, costs []int) int {
	arr := append([]int{-1000}, days...)
	n := len(days)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + costs[0]
		for _, c := range [][]int{{7, costs[1]}, {30, costs[2]}} {
			d, cost := c[0], c[1]
			j := i - 1
			for arr[i]-arr[j] < d {
				j--
			}
			dp[i] = min(dp[i], cost+dp[j])
		}
	}
	return dp[n]
}
