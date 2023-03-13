package main

import "sort"

func maxScore(nums []int) int {
	ans := 0
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	curr := 0
	for _, v := range nums {
		curr += v
		if curr <= 0 {
			break
		}
		ans++
	}
	return ans
}

func main() {
}
