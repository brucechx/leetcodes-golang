package _211_添加与搜索单词_数据结构设计

type WordDictionary struct {
	next [26]*WordDictionary
	end bool
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		next: [26]*WordDictionary{},
		end:  false,
	}
}


/** Adds a word into the data structure. */
func (t *WordDictionary) AddWord(word string)  {
	root := t
	for _, w := range word{
		if root.next[w - 'a'] == nil{
			root.next[w - 'a'] = &WordDictionary{
				next: [26]*WordDictionary{},
				end:  false,
			}
		}
		root = root.next[w - 'a']
	}
	root.end = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (t *WordDictionary) Search(word string) bool {
	root := t
	for i, w := range word{
		if w == '.'{
			for _, c := range root.next{
				if c != nil && c.Search(word[i+1:]){
					return true
				}
			}
			return false
		}else if root.next[w - 'a'] == nil{
			return false
		}else {
			root = root.next[w - 'a']
		}
	}
	return root.end
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
