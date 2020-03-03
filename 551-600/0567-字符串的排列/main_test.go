package _567_字符串的排列

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T){
	cases := []struct{
		s1 string
		s2 string
		res bool
	}{
		{
			s1: "ab",
			s2: "eidbaooo",
			res: true,
		},
		{
			s1: "ab",
			s2: "eidboaoo",
			res: false,
		},
		{
			s1: "hello",
			s2: "ooolleoooleh",
			res: false,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, checkInclusion(cas.s1, cas.s2))
	}
}
