package _820_short_encoding_of_words

/**
https://leetcode.com/problems/short-encoding-of-words/
820. Short Encoding of Words
Medium
*/

/**

 */
func minimumLengthEncoding_map(words []string) int {
	good := make(map[string]struct{}, len(words))
	for _, w := range words {
		good[w] = struct{}{}
	}
	for _, w := range words {
		// просто удаляем из мапы все другие подстроки начиная с [1:]. И так далее [2:], [3:]
		for k := 1; k < len(w); k++ {
			key := w[k:]
			if _, ok := good[key]; ok {
				delete(good, key)
			}
		}
	}
	ans := 0
	for w := range good {
		ans += len(w) + 1
	}

	return ans
}

/**

 */
func minimumLengthEncoding(words []string) int {
	trie := NewTrieNode()            // тут храним все концы слов. См for j:=len()
	nodes := make(map[*TrieNode]int) // а тут храним все начала
	for i := 0; i < len(words); i++ {
		cur := trie
		for j := len(words[i]) - 1; j >= 0; j-- {
			cur = cur.get(words[i][j])
		}
		nodes[cur] = i
	}

	ans := 0
	for node, i := range nodes {
		// Почему нужны только с node.count == 0?
		// Так как node - это начало слова, то по получается что node.count == 0 - самое длинное слово из всех что нам нужны
		// Другими словами time, me. Так как me - наша подстрока. То мы в nodes будем видеть указание на t. Но у самого t еще будут дети - Перед t в слове time есть i.
		if node.count == 0 {
			ans += len(words[i]) + 1
		}
	}
	return ans
}

type TrieNode struct {
	children []*TrieNode
	count    int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 26),
	}
}

func (t *TrieNode) get(letter byte) *TrieNode {
	idx := letter - 'a'
	if t.children[idx] == nil {
		t.children[idx] = NewTrieNode()
		t.count++
	}
	return t.children[idx]
}
