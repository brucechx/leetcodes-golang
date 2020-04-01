package _006_Z字形变换

/*
题目理解：
- 字符串s是以Z字形为顺序存储的字符串，目标是按行打印。
- 设 numRows 行字符串分别为s1, s2, s3, ..., sn
  则容易发现：按顺序遍历字符串s时，每个字符c在Z字形中对应的行索引,
  先从s1增大至sn, 再从sn减小至s1 如此反复。
- 因此，解决方案为：模拟这个行索引的变化，在遍历s中把每个字符填到正确的行res[i]

算法流程：
- 按顺序遍历字符串s；
- res[i] += c; 把每个字符c填入对应行si
- i += flag; 更新当前字符c对应的行索引；
- flag = - flag; 在达到Z字形转折点时，执行反向。

复杂度分析：
时间复杂度O(N): 遍历一遍字符串s；
空间复杂度O(N): 各行字符串共占用O(N)额外空间。
*/

func convert2(s string, numRows int) string {
	if numRows < 2{
		return s
	}
	rows := make([][]rune, 0)
	for i:=0; i<numRows; i++{
		// 每一行
		rows = append(rows, []rune{})
	}
	index, flag := 0, -1
	for _, v := range s{
		rows[index] = append(rows[index], v)
		// 先s1,s2...sn 然后 sn,sn-1,...,s1
		if index == 0 || index == numRows - 1{
			flag = -flag
		}
		index += flag
	}
	result := ""
	for _, row := range rows{
		result += string(row)
	}
	return result
}
