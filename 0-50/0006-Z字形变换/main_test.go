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
		{"LEETCODEISHIRING", "LCIRETOESIIGEDHN", 3},
		{"LEETCODEISHIRING", "LDREOEIIECIHNTSG", 4},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, convert(cas.input, cas.num))
	}
}
