package _127_单词接龙

/*
	BFS or 双向BFS
*/

// 双向BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dic := make(map[string]bool, len(wordList))
	for _, val := range wordList{
		dic[val] = true
	}

	if _, ok := dic[endWord]; !ok{
		return 0
	}

	q1 := make(map[string]bool)
	q2 := make(map[string]bool)
	q1[beginWord] = true
	q2[endWord] = true

	l := len(beginWord)
	steps := 0
	for len(q1) > 0 && len(q2) > 0{
		steps++

		if len(q1) > len(q2){
			q1, q2 = q2, q1
		}
		q := make(map[string]bool) // 下一层
		for k := range q1{
			chs := []rune(k)
			for i:=0; i<l; i++{
				ch := chs[i] // 为了后续复位
				for c:='a'; c<='z'; c++{ // 遍历26个字母予以尝试
					if ch == c{ // 自身忽略
						continue
					}
					chs[i] = c // 只改变一位
					tmp := string(chs) // 变成新的单次
					if _, ok := q2[tmp]; ok { // q2 结合中是否有新单次，有则退出
						return steps + 1
					}
					if _, ok := dic[tmp]; !ok{ // 候选单词列表中是有新单词，如果没有则继续尝试
						continue
					}
					// 如果有，则加入新的集合；并删除候选单词列表
					q[tmp] = true
					delete(dic, tmp)
				}
				chs[i] = ch // 复位
			}
		}
		q1 = q  // q1修改为新扩展的q
	}
	return 0
}
