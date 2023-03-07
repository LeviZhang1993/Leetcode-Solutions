package main

func minimumTime(time []int, totalTrips int) int64 {
	var l, r int64 = 0, 0
	minTime := time[0]
	for _, v := range time {
		if minTime > v {
			minTime = v
		}
	}
	r = int64(minTime) * int64(totalTrips)
	for l <= r {
		mid := (l + r) / 2
		have := int64(0)
		for _, v := range time {
			have += mid / int64(v)
		}
		if have >= int64(totalTrips) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	println(minimumTime([]int{1, 2, 3}, 5) == 3)
	println(minimumTime([]int{2}, 1) == 2)
}
