package _208_实现_Trie_前缀树

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd bool
}

func NewTrieNode()  *TrieNode{
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isEnd:    false,
	}
}

type Trie struct {
	root *TrieNode
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root:NewTrieNode()}
}


/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	node := t.root
	for i := range word{
		if _, ok := node.children[word[i]]; !ok {
			node.children[word[i]] = NewTrieNode()
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.root
	for i := range word{
		if _, ok := node.children[word[i]]; !ok {
			return false
		}
		node = node.children[word[i]]
	}
	return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for i := range prefix{
		if _, ok := node.children[prefix[i]]; !ok{
			return false
		}
		node = node.children[prefix[i]]
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