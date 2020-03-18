package _155_最小栈

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-1)
	fmt.Println(minStack.GetMin()) // -2
	fmt.Println(minStack.Top()) // -1
	minStack.Pop() //
	fmt.Println(minStack.GetMin()) // -2
}
