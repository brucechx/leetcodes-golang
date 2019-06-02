package _04

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReplaceSpace(t *testing.T){
	cases := []struct{
		input string
		output string
	}{
		{"We Are Happy", "We%20Are%20Happy"},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.output, replaceSpace(cas.input))
		assert.Equal(t, cas.output, replaceSpaceMethod2(cas.input))
	}

}
