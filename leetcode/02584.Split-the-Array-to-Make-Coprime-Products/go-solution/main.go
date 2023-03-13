package main

import "fmt"

func findValidSplit(nums []int) int {
	pre := [1000001]int{}
	for i := 2; i <= 1000000; i++ {
		if pre[i] == 0 {
			pre[i] = i
			for j := i * i; j <= 1000000; j += i {
				pre[j] = i
			}
		}
	}
	curr := 0
	cnt := map[int]int{}
	for _, i := range nums {
		for i != 1 {
			cnt[pre[i]] += 1
			i /= pre[i]
		}
	}
	now := map[int]int{}
	for idx, i := range nums {
		for j := i; j != 1; {
			p := pre[j]
			now[pre[j]] += 1
			cnt[pre[j]] -= 1
			if now[p] == 1 && cnt[p] > 0 {
				curr += 1
			}
			if now[p] > 1 && cnt[p] == 0 {
				curr -= 1
			}
			j /= p
		}
		if curr == 0 && idx < len(nums)-1 {
			return idx
		}
	}
	return -1
}

func main() {
	fmt.Println(findValidSplit([]int{4, 7, 8, 15, 3, 5}) == 2)
}
