package main

type Node struct {
	nxt    map[rune]*Node
	IsWord bool
	Cnt    int
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{&Node{
		map[rune]*Node{},
		false,
		0,
	}}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for _, v := range word {
		curr.Cnt += 1
		if _, ok := curr.nxt[v]; !ok {
			curr.nxt[v] = &Node{map[rune]*Node{}, false, 0}
		}
		curr = curr.nxt[v]
	}
	curr.Cnt += 1
	curr.IsWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, v := range word {
		if _, ok := curr.nxt[v]; !ok {
			return false
		}
		curr = curr.nxt[v]
	}
	return curr.IsWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, v := range prefix {
		if _, ok := curr.nxt[v]; !ok {
			return false
		}
		curr = curr.nxt[v]
	}
	return curr.Cnt > 0

}

func main() {
}
