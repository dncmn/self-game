package trie

import (
	"fmt"
	"unicode"
)

type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
}

func New() Trie {
	var r Trie
	r.Root = NewTrieNode()
	return r
}

func NewTrieNode() *TrieNode {
	n := new(TrieNode)
	n.Children = make(map[rune]*TrieNode)
	return n
}

func (this *Trie) Append(txt string) {
	if len(txt) < 1 {
		return
	}
	node := this.Root
	key := []rune(txt)
	for i := 0; i < len(key); i++ {
		if _, exists := node.Children[key[i]]; !exists {
			node.Children[key[i]] = NewTrieNode()
		}
		node = node.Children[key[i]]
	}

	node.End = true
}

func isNoneChar(r rune) bool {
	return (unicode.In(r, unicode.Han) || unicode.IsLetter(r) || unicode.IsDigit(r)) &&
		!unicode.IsPunct(r) && !unicode.IsSpace(r)
}

func replaceChars(words []rune, i, j int) {
	for ; i < j; i++ {
		words[i] = '*'
	}
}

func (this *Trie) Replace(txt string) (string, bool) {
	if txt == "" {
		return "", false
	}
	origin := []rune(txt)
	words := []rune(txt)
	replace := false
	var (
		ok   bool
		node *TrieNode
	)
	for i, word := range origin {
		if node, ok = this.Root.Children[word]; !ok {
			continue
		}
		j := i + 1
		if node.End {
			replaceChars(words, i, j)
		}
		for ; j < len(origin); j++ {
			if !isNoneChar(origin[j]) {
				continue
			}
			if v, ok := node.Children[origin[j]]; !ok {
				break
			} else {
				node = v
			}
		}
		if node.End {
			replace = true
			replaceChars(words, i, j)
		}
	}
	return string(words), replace
}

func (this *Trie) Print() {
	node := this.Root
	for k, v := range node.Children {
		for k1, v1 := range v.Children {
			for k2, _ := range v1.Children {
				fmt.Printf("%s%s%s", string(k), string(k1), string(k2))
			}
		}
	}
}
