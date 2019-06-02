package _05

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPrintReverse(t *testing.T){
	l := newSingeList()
	l.append(1)
	l.append(2)
	l.append(3)

	res := l.printReverse()
	want := []interface{}{3, 2, 1}
	assert.Equal(t, want, res)
}
