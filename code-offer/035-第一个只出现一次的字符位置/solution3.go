package _35_第一个只出现一次的字符位置

func FirstNotRepeatingChar(str string) int{
	if len(str) < 1 || len(str) > 10000{
		return  -1
	}
	// 小写字母
	x := make([]int, 26)
	// 大写字母
	y := make([]int, 26)
	for i, s := range str{
		if 'a' <=s && s <= 'z'{
			if x[s-'a'] == 0{ //  首次出现保存出现位置
				x[s-'a'] = i+1
			}else {   //  出现多次, 就置标识-1
				x[s-'a'] = -1
			}
		}
		if 'A' <=s && s <= 'Z'{
			if y[s-'A'] == 0{
				y[s-'A'] = i + 1
			}else {
				y[s-'A'] = -1
			}
		}
	}

	// 随意指定个大数
	res := len(str) + 1
	for i:=0; i<26;i++{
		if x[i] !=0 && x[i] != -1{
			res = min(res, x[i])
		}
		if y[i] !=0 && y[i] != -1{
			res = min(res, y[i])
		}
	}
	if res == len(str) + 1{
		return -1
	}
	return res - 1
}

func min(x, y int) int{
	if x > y{
		return y
	}
	return x
}
