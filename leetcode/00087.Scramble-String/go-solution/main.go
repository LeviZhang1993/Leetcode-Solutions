package main

func isScramble_(s1, s2 []rune, d map[string]bool) (ans bool) {
	key := string(s1) + "_" + string(s2)
	if len(s1) != len(s2) {
		ans = false
		return
	}
	if string(s1) == string(s2) {
		ans = true
		return
	}
	if r, ok := d[key]; ok {
		return r
	}
	c := make([]int, 26)
	n := len(s1)
	for i := 0; i < n; i++ {
		c[s1[i]-'a']++
		c[s2[i]-'a']--
	}
	for i := range c {
		if c[i] != 0 {
			d[key] = false
			return false
		}
	}
	for i := 1; i < n; i++ {
		if isScramble_(s1[:i], s2[:i], d) && isScramble_(s1[i:], s2[i:], d) {
			d[key] = true
			return true
		}
		if isScramble_(s1[:i], s2[n-i:], d) && isScramble_(s1[i:], s2[:n-i], d) {
			d[key] = true
			return true
		}
	}
	d[key] = false
	return false
}

func isScramble(s1 string, s2 string) bool {
	return isScramble_([]rune(s1), []rune(s2), make(map[string]bool))
}

func main() {
}
