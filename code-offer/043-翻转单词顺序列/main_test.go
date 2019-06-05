package _43_翻转单词顺序列

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverseSentence(t *testing.T) {
	cases := []struct{
		input string
		res string
	}{
		{"I am a student ", " student a am I"},
		{"I am a student", "student a am I"},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, ReverseSentence(cas.input))
		assert.Equal(t, cas.res, ReverseSentence2(cas.input))
		assert.Equal(t, cas.res, ReverseSentence3(cas.input))
	}
}
