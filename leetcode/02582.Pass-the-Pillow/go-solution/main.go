package main

func passThePillow(n int, t int) int {
	t = t % (2*n - 2)
	if t == 0 {
		return 1
	}
	if t <= n-1 {
		return 1 + t
	}
	return 2*n - t - 1
}

func main() {
	println(passThePillow(3, 2) == 3)
	println(passThePillow(4, 5) == 2)
}
