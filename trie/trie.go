// TRIE: Prefix tree
// variant of a n-ary tree where chharacters are stored at each node
// each path down the tree represents a word
// we need to have a terminating indicator

package trie

import "strings"

const (
	SIZE  = 26
	UPPER = 'A'
	LOWER = 'a'
)

type TrieNode struct {
	Links       []*TrieNode
	IsEndofWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Links: make([]*TrieNode, SIZE), IsEndofWord: false}
}

func (tn *TrieNode) Insert(word string) {

	if len(word) <= 0 {
		return
	}
	crawler := tn

	var first rune
	if word == strings.ToUpper(word) {
		first = UPPER
	} else {
		first = LOWER
	}

	for _, character := range word {
		index := character - first
		if crawler.Links[index] == nil {
			crawler.Links[index] = NewTrieNode()
		}
		crawler = crawler.Links[index]
	}
	crawler.IsEndofWord = true

}
