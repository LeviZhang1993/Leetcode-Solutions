package main

func vowelStrings(words []string, left int, right int) int {
	d := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	ans := 0
	for i := left; i <= right; i++ {
		if d[words[i][0]] && d[words[i][len(words[i])-1]] {
			ans++
		}
	}
	return ans
}

func main() {
}
