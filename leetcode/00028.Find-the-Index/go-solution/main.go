package main

import "fmt"

func strStr(haystack string, needle string) int {

	nxt := func(s string) []int {
		n := len(s)
		nxt := make([]int, n)
		for i := range nxt {
			nxt[i] = -1
		}
		i, j := 0, -1
		for i < n-1 {
			if j == -1 || s[i] == s[j] {
				i, j = i+1, j+1
				nxt[i] = j
			} else {
				j = nxt[j]
			}
		}
		return nxt
	}(needle)
	i, j, m, n := 0, 0, len(haystack), len(needle)
	if n == 0 {
		return -1
	}
	for i < m && j < n {
		if j == -1 || haystack[i] == needle[j] {
			i, j = i+1, j+1
		} else {
			j = nxt[j]
		}
		if j == n {
			return i - n
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("abcde", "bc") == 1)
	fmt.Println(strStr("abcde", "") == -1)
	fmt.Println(strStr("ababcde", "ab") == 0)
}
