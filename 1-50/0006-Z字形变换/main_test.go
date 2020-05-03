package _006_Z字形变换

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T){
	cases := []struct{
		input, res string
		num int
	}{
		{"AB", "AB", 1},
		{"LEETCODEISHIRING", "LCIRETOESIIGEDHN", 3},
		{"LEETCODEISHIRING", "LDREOEIIECIHNTSG", 4},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, convert2(cas.input, cas.num))
	}
}
