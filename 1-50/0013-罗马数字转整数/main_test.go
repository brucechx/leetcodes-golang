package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T){
	cases := []struct{
		input string
		res int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
	}

	for _, cas := range cases{
		assert.Equal(t, romanToInt(cas.input), cas.res)
	}
}
