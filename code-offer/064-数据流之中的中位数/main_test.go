package _64_数据流之中的中位数

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestDynamicArray(t *testing.T) {
	d := NewDynamicArray()
	d.insert(1)
	d.insert(2)
	d.insert(3)
	assert.Equal(t, float32(2), d.getMedian())
	d.insert(4)
	d.insert(5)
	d.insert(6)
	assert.Equal(t, float32(3.5), d.getMedian())
}

func TestHello(t *testing.T){
	a := float32((3 +4)/2.0)
	fmt.Println(a)
}
