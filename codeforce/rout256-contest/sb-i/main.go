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

		res := problemI(word, dict)
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

func problemI(word string, dict []string) string {
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

func NewTrie() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string, idx int) {
	cur := this
	for i, c := range word {
		n := c - 'a'

		if cur.children[n] == nil {
			cur.children[n] = &Trie{}
		}
		if i == 0 {
			cur.indeces = append(cur.indeces, idx)
		}
		cur = cur.children[n]
		if i == len(word)-1 {
			cur.isWord = true
		}

	}
}

func (this *Trie) Search(word string) (bool, []int) {
	var indeces []int
	cur := this
	for i, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			return false, indeces
		}
		if i == 0 {
			indeces = cur.indeces
		}
		cur = cur.children[n]
	}
	return cur.isWord, indeces
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		cur = cur.children[n]
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
