package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 1
	for _, pile := range piles {
		if r < pile {
			r = pile
		}
	}
	for l <= r {
		mid := l + (r-l)/2
		eat := 0
		for _, pile := range piles {
			eat += (pile-1)/mid + 1
		}
		if eat <= h {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8) == 4)
}
