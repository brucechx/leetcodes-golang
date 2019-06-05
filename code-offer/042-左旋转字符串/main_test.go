package _42_左旋转字符串

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLeftRotateString(t *testing.T) {
	cases := []struct{
		input string
		num int
		res string
	}{
		{"abcXYZdef", 3, "XYZdefabc"},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, LeftRotateString(cas.input, cas.num))
		assert.Equal(t, cas.res, LeftRotateString2(cas.input, cas.num))
		assert.Equal(t, cas.res, LeftRotateString3(cas.input, cas.num))
	}
}
