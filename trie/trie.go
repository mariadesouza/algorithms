package trie

const SIZE = 26

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

	for _, character := range word {
		index := character - 'A'
		if crawler.Links[index] == nil {
			crawler.Links[index] = NewTrieNode()
		}
		crawler = crawler.Links[index]
	}
	crawler.IsEndofWord = true

}
