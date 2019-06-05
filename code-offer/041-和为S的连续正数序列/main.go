package _41_和为S的连续正数序列


/**
滑动窗口法
链接：https://www.nowcoder.com/questionTerminal/c451a3fd84b64cb19485dad758a55ebe
来源：牛客网
用两个数字begin和end分别表示序列的最大值和最小值，
首先将begin初始化为1，end初始化为2.
如果从begin到end的和大于s，我们就从序列中去掉较小的值(即增大begin),
相反，只需要增大end。
如果和等于s，则记录begin到end的数组，begin+=1,end=begin+1,在开始
终止条件为：一直增加begin到(1+sum)/2并且end小于sum为止
*/
func FindContinuousSequence(sum int) (result [][]int) {
	begin := 1
	end := 2
	for begin < (1+sum)/2 && end < sum {
		tmpSum := (begin + end) * (end - begin + 1) / 2
		if tmpSum == sum {
			tmpArr := make([]int, end-begin+1)
			for i := range tmpArr {
				tmpArr[i] = begin + i
			}
			result = append(result, tmpArr)
			begin++
			end = begin + 1
		} else if tmpSum > sum {
			begin++
		} else {
			end++
		}
	}
	return
}
