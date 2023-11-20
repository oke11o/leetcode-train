package _2xx

// 0208. Implement Trie (Prefix Tree)
type node struct {
	children [26]*node
	end      bool
}
type Trie struct {
	root *node
}

func Constructor() Trie {
	return Trie{root: &node{}}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, ch := range word {
		v := ch - 'a'
		n := cur.children[v]
		if n == nil {
			cur.children[v] = &node{}
			n = cur.children[v]
		}
		cur = n
	}
	cur.end = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, ch := range word {
		v := ch - 'a'
		n := cur.children[v]
		if n == nil {
			return false
		}
		cur = n
	}
	return cur.end
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, ch := range prefix {
		v := ch - 'a'
		n := cur.children[v]
		if n == nil {
			return false
		}
		cur = n
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
