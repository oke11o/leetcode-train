package _2xx

// 0208. Implement Trie (Prefix Tree)
type Trie struct {
	root *node
}

type node struct {
	children []*node
	val      rune
	isFinal  bool
}

func (n *node) childrenExists(v rune) (*node, bool) {
	for _, c := range n.children {
		if c.val == v {
			return c, true
		}
	}
	return nil, false
}

func Constructor() Trie {
	return Trie{root: &node{}}
}

func (this *Trie) Insert(word string) {
	parent := this.root
	var j int
	for i, s := range word {
		j = i
		if c, ok := parent.childrenExists(s); ok {
			parent = c
		} else {
			n := &node{val: rune(word[j])}
			parent.children = append(parent.children, n)
			parent = n
		}
	}

	parent.isFinal = true

}

func (this *Trie) Search(word string) bool {
	n := this.root
	for _, s := range word {
		if c, ok := n.childrenExists(s); ok {
			n = c
		} else {
			return false
		}
	}
	return n.isFinal
}

func (this *Trie) StartsWith(prefix string) bool {
	n := this.root
	for _, s := range prefix {
		if c, ok := n.childrenExists(s); ok {
			n = c
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
