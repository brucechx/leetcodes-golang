package _058_最后一个单词的长度

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T){
	cases := []struct{
		input string
		res int
	}{
		{
			input: "a",
			res: 1,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, lengthOfLastWord(cas.input))
	}
}
