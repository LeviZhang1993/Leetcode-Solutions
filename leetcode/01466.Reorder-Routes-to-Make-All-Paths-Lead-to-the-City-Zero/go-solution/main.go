package main

func dfs(curr int, pre, nxt [][]int, ans *int, vis []int) {
	vis[curr] = 1
	for _, v := range nxt[curr] {
		if vis[v] == 0 {
			*ans++
			dfs(v, pre, nxt, ans, vis)
		}
	}
	for _, v := range pre[curr] {
		if vis[v] == 0 {
			dfs(v, pre, nxt, ans, vis)
		}
	}

}

func minReorder(n int, connections [][]int) int {
	pre := make([][]int, n)
	nxt := make([][]int, n)
	for _, connection := range connections {
		u, v := connection[0], connection[1]
		pre[v] = append(pre[v], u)
		nxt[u] = append(nxt[u], v)
	}
	vis := make([]int, n)
	ans := 0
	dfs(0, pre, nxt, &ans, vis)
	return ans
}

func main() {
}
