package main

import "sort"

func findMinimumTime(tasks [][]int) int {
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	curr := make([]int, 2001)
	for _, task := range tasks {
		l, r, d := task[0], task[1], task[2]
		have := 0
		for i := l; i <= r; i++ {
			if curr[i] != 0 {
				have += 1
			}
		}
		idx := r
		for have < d {
			if curr[idx] == 1 {
				idx--
				continue
			}
			curr[idx] = 1
			have++
			idx--
		}
	}
	ans := 0
	for i := 0; i <= 2000; i++ {
		if curr[i] == 1 {
			ans++
		}
	}
	return ans
}

func main() {
}
