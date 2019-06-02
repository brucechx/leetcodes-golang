package _07_用两个栈实现队列

import (
	"testing"
	"fmt"
)

func TestStackQueue(t *testing.T) {
	t.Run("enqueue", func(t *testing.T) {
		sq := initStackQueue()
		sq.Enqueue(1)
		sq.Enqueue(2)
		fmt.Println(sq.stack1.Items)
	})

	t.Run("dequeue", func(t *testing.T) {
		sq := initStackQueue()
		sq.Enqueue(1)
		sq.Enqueue(2)

		got := sq.Dequeue()
		want := 1
		if got != want{
			t.Errorf("got %v want %v", got, want)
		}

		sq.Enqueue(3)
		got = sq.Dequeue()
		want = 2
		if got != want{
			t.Errorf("got %v want %v", got, want)
		}

	})
}
