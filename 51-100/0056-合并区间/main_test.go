package _056_合并区间

import (
	"fmt"
	"testing"
	"time"
)

func TestMerge(t *testing.T){
	msgs := make(chan int, 10)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-msgs:
				fmt.Println("in", time.Now())
				time.Sleep(3 *time.Second)
				fmt.Println("msg=", msg, " ", time.Now())
			case <- quit:
				fmt.Println("exit ", time.Now())
				return
			}
		}
	}()
	msgs <- 1
	close(quit)
	time.Sleep(5 *time.Second)
	fmt.Println("over", time.Now())
}
