package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var err error
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var dictLen int
	_, err = fmt.Fscan(in, &dictLen)
	panicErr(err)

	dict := make([]string, dictLen)
	trie := NewTrie()
	for i := 0; i < dictLen; i++ {
		var dictWord string
		_, err = fmt.Fscan(in, &dictWord)
		panicErr(err)
		dict[i] = dictWord
		trie.Insert(reverseString(dictWord), i)
	}

	var inLen int
	_, err = fmt.Fscan(in, &inLen)
	panicErr(err)
	for i := 0; i < inLen; i++ {
		var word string
		_, err = fmt.Fscan(in, &word)
		panicErr(err)

		res := problemI(word, dict, trie)
		fmt.Print(res, "\r\n")
	}

}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func reverseString(in string) string {
	out := []rune(in)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}

func problemI(word string, dict []string, trie *Trie) string {
	rev := reverseString(word)
	var resIdx []int
	for i := 1; i <= len(rev); i++ {
		if ok, idxs := trie.StartsWith(rev[:i]); ok {
			if i == len(rev) { // Слова не должны совпадать полностью
				var tmp = make([]int, 0, len(idxs))
				for _, idx := range idxs {
					if dict[idx] != word {
						tmp = append(tmp, idx)
					}
				}
				if len(tmp) > 0 {
					resIdx = tmp
				}
			} else {
				resIdx = idxs
			}
		} else {
			break
		}
	}
	var tmp = make([]int, 0, len(resIdx))
	for _, idx := range resIdx {
		if dict[idx] != word {
			tmp = append(tmp, idx)
		}
	}
	if len(tmp) > 0 {
		resIdx = tmp
	}

	if len(tmp) > 0 {
		return dict[tmp[0]]
	}

	return dict[0]
}

func problemI2(word string, dict []string) string {
	result := dict[0]
	max := 0
	for _, d := range dict {
		if word == d {
			continue
		}
		zarif := findZarif(word, d)
		if max < zarif {
			max = zarif
			result = d
		}
	}

	return result
}

func findZarif(word string, d string) int {
	result := 0
	i := len(word) - 1
	j := len(d) - 1
	for i >= 0 && j >= 0 {
		if word[i] != d[j] {
			break
		}
		result++
		i--
		j--
	}
	return result
}

type Trie struct {
	isWord   bool
	children [26]*Trie
	indeces  []int
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string, idx int) {
	cur := t
	for i, c := range word {
		n := c - 'a'

		if cur.children[n] == nil {
			cur.children[n] = &Trie{}
		}
		cur = cur.children[n]
		cur.indeces = append(cur.indeces, idx)
		//if i == 0 {
		//	cur.indeces = append(cur.indeces, idx)
		//}
		if i == len(word)-1 {
			cur.isWord = true
		}
	}
}

func (t *Trie) Search(word string) (bool, []int) {
	var indices []int
	cur := t
	for i, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			return false, indices
		}
		cur = cur.children[n]
		if i == 0 {
			indices = cur.indeces
		}
	}
	return cur.isWord, indices
}

func (t *Trie) StartsWith(prefix string) (bool, []int) {
	var indices []int
	cur := t
	for i, c := range prefix {
		n := c - 'a'
		if cur.children[n] == nil {
			return false, indices
		}
		cur = cur.children[n]
		if i == 0 {
			indices = cur.indeces
		}
	}
	return true, cur.indeces
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
