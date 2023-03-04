package main

import "fmt"

func countSubarrays_solution1(nums []int, minK int, maxK int) int64 {
	var ans int64 = 0
	idx := 0
	for idx < len(nums) {
		if nums[idx] < minK || nums[idx] > maxK {
			idx += 1
			continue
		}
		i, j, curr := -1, -1, idx
		for ; curr < len(nums) && minK <= nums[curr] && nums[curr] <= maxK; curr++ {
			if nums[curr] == minK {
				i = curr
			}
			if nums[curr] == maxK {
				j = curr
			}
			if i != -1 && j != -1 {
				ans += int64(func(a, b int) int {
					if a < b {
						return a
					}
					return b
				}(i, j) - idx + 1)
			}
		}
		idx = curr
	}
	return ans
}

func countSubarrays_solution2(nums []int, minK int, maxK int) int64 {
	idx := 0
	var ans int64
	for idx < len(nums) {
		if nums[idx] < minK || nums[idx] > maxK {
			idx++
			continue
		}
		j := idx
		for j < len(nums) && minK <= nums[j] && nums[j] <= maxK {
			j++
		}
		// get the subarray
		func() {
			t := 0
			if minK != maxK {
				t = 1
			}
			d := map[int]bool{}
			cnt := map[int]int{}
			d[maxK], d[minK] = true, true
			res := int64(j - idx)
			res = (res * (res + 1)) / 2
			l := idx // sliding window left side
			for i := idx; i < j; i++ {
				if _, ok := d[nums[i]]; ok {
					cnt[nums[i]]++
					if cnt[nums[i]] == 1 {
						t -= 1
					}
				}
				for t < 0 {
					_, ok := d[nums[l]]
					if ok {
						cnt[nums[l]]--
						if cnt[nums[l]] == 0 {
							t++
						}
					}
					l++
				}
				res -= int64(i - l + 1)
			}
			ans += res
		}()
		idx = j
	}
	return ans
}

func main() {
	fmt.Println(countSubarrays_solution1([]int{1, 3, 5, 2, 7, 5}, 1, 5) == 2)
	fmt.Println(countSubarrays_solution2([]int{1, 3, 5, 2, 7, 5}, 1, 5) == 2)
}
