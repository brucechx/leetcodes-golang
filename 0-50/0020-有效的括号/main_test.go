package _020_有效的括号

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T){
	cases := []struct{
		input string
		res bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"([)]", false},
		{"(]", false},
		{"{[]}", true},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, isValid(cas.input))
	}
}
