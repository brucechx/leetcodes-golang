package _165_比较版本号

import "strings"

/*
	if version1 > version2{
		return 1
	}else if version1 < version2{
		return -1
	}else{
		return 0
	}
*/

func compareVersion(version1 string, version2 string) int {
	l1 := strings.Split(version1, ".")
	l2 := strings.Split(version2, ".")
	for len(l1) < len(l2){
		l1 = append(l1, "0")
	}
	for len(l2) < len(l1){
		l2 = append(l2, "0")
	}
	for i:=0; i<len(l1); i++{
		val1 := l1[i]
		val2 := l2[i]
		for len(val1) < len(val2){
			val1 = "0"+ val1
		}
		for len(val2) < len(val1){
			val2 = "0"+ val2
		}
		l := len(val1)
		for j:=0; j<l; j++{
			v1 := val1[j]
			v2 := val2[j]
			if v1 == v2{
				continue
			}
			if v1 > v2{
				return 1
			}
			return -1
		}
	}
	return 0
}



