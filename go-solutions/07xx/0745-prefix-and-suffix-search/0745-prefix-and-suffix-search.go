package _745_prefix_and_suffix_search

/**
https://leetcode.com/problems/prefix-and-suffix-search/
745. Prefix and Suffix Search
Hard
*/
type WordFilter struct {
	trie      *TrieNode
	separator rune
}

func Constructor(words []string) WordFilter {
	trie := NewTrieNode()
	separator := 'a' - 'a' + 26 + 'a'

	for weight := 0; weight < len(words); weight++ {
		word := []rune(words[weight])
		word = append(word, separator)
		for i := 0; i < len(word); i++ {
			cur := trie
			cur.weight = weight

			for j := i; j < 2*len(word)-1; j++ {
				k := word[j%len(word)] - 'a'
				if cur.children[k] == nil {
					cur.children[k] = NewTrieNode()
				}
				cur = cur.children[k]
				cur.weight = weight
			}
		}
	}

	return WordFilter{trie: trie, separator: separator}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	cur := this.trie
	word := []rune(suffix)
	word = append(word, this.separator)
	word = append(word, []rune(prefix)...)
	for _, letter := range word {
		k := letter - 'a'
		if cur.children[k] == nil {
			return -1
		}
		cur = cur.children[k]
	}
	return cur.weight
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */

type TrieNode struct {
	children []*TrieNode
	weight   int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 27),
	}
}
