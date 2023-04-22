package main

import "fmt"

func minInsertions(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for t := 2; t <= n; t++ {
		for l := 0; l+t-1 < n; l++ {
			r := l + t - 1
			if s[l] == s[r] {
				dp[l][r] = dp[l+1][r-1]
			} else {
				dp[l][r] = 2 * n
			}
			dp[l][r] = min(dp[l][r], dp[l+1][r]+1)
			dp[l][r] = min(dp[l][r], dp[l][r-1]+1)
		}
	}
	return dp[0][n-1]
}

func main() {
	fmt.Println(minInsertions("mbadm"))
}
