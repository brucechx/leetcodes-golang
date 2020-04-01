package _017_电话号码的字母组合

func letterCombinations2(digits string) []string {
	if digits == "" {
		return nil
	}

	ret := []string{""}

	// 让digits中所有的数字都有机会进行迭代。
	for i := 0; i < len(digits); i++ {
		var temp []string
		// 让ret中的每个元素都能添加新的字母。
		for j := 0; j < len(ret); j++ {
			// 把digits[i]所对应的字母，多次添加到ret[j]的末尾
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
		}

		ret = temp
	}

	return ret
}
