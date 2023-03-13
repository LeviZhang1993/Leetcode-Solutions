package main

import "fmt"

func waysToReachTarget(target int, types [][]int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	const mod int = 1e9 + 7
	for _, e := range types {
		cnt, m := e[0], e[1]
		curr := make([]int, target+1)
		copy(curr, dp)
		for i := 1; i <= cnt; i++ {
			if i*m > target {
				break
			}
			for j := target; j >= i*m; j-- {
				curr[j] = (curr[j] + dp[j-i*m]) % mod
			}
		}
		copy(dp, curr)
	}
	return dp[target]
}

func main() {
	fmt.Println(waysToReachTarget(6, [][]int{
		{6, 1},
		{3, 2},
		{2, 3},
	}) == 7)
}
