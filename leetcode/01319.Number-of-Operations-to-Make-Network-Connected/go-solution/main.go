package main

func find(x int, f []int) int {
	if f[x] != x {
		f[x] = find(f[x], f)
	}
	return f[x]
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = i
	}
	for _, connection := range connections {
		u, v := connection[0], connection[1]
		f[find(u, f)] = find(v, f)
	}
	d := make(map[int]bool)
	for i := 0; i < n; i++ {
		d[find(i, f)] = true
	}
	return len(d) - 1
}

func main() {
}
