package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T){
	cases := []struct{
		input int
		res bool
	}{
		{1221, true},
		{121, true},
		{-121, false},
		{10, false},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, isPalindrome(cas.input))
	}
}
