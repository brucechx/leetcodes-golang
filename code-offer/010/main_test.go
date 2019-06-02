package _10

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestV1(t *testing.T) {
	cases := []struct{
		input int64
		output int
	}{
		{int64(3), 2},
		{int64(2), 1},
	}

	for _, cas := range cases{
		want1 := V1(cas.input)
		want2 := V2(cas.input)
		want3 := V3(cas.input)
		assert.Equal(t, cas.output, want1)
		assert.Equal(t, cas.output, want2)
		assert.Equal(t, cas.output, want3)
	}
}
