package main

func beautifulSubarrays(nums []int) int64 {
	var ans int64 = 0
	d := map[int]int64{}
	d[0] = 1
	curr := 0
	for _, v := range nums {
		for j := 0; j < 20; j++ {
			if v&(1<<j) > 0 {
				curr ^= (1 << j)
			}
		}
		ans += d[curr]
		d[curr] += 1
	}
	return ans
}

func main() {
}
