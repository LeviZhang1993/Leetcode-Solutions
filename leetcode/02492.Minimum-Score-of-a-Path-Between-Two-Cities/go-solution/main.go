package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minScore(n int, roads [][]int) int {
	nxt := make([][][2]int, n+1)
	for _, road := range roads {
		nxt[road[0]] = append(nxt[road[0]], [2]int{road[1], road[2]})
		nxt[road[1]] = append(nxt[road[1]], [2]int{road[0], road[2]})
	}
	dis := make([]int, n+1)
	for idx := range dis {
		dis[idx] = -1
	}
	dis[1] = 1e9
	q := make([]int, n+1)
	q = append(q, 1)
	var u int
	for len(q) != 0 {
		u = q[0]
		q = q[1:]
		for _, route := range nxt[u] {
			v, d := route[0], route[1]
			if dis[v] == -1 || dis[v] > min(dis[u], d) {
				dis[v] = min(dis[u], d)
				q = append(q, v)
			}
		}
	}
	return dis[n]
}

func main() {
}
