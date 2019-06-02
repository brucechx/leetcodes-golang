package _04_替换空格

import "strings"

func replaceSpace(old string) string{
	return strings.Replace(old, " ", "%20", -1)
}

func replaceSpaceMethod2(old string) string{
	replacedRune := []rune("")
	repSlice := []rune("%20")
	for _, s := range old{
		if string(s) == " "{
			replacedRune = append(replacedRune, repSlice...)
		}else {
			replacedRune = append(replacedRune, s)
		}
	}
	return string(replacedRune)
}
