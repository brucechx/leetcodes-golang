package _14_调整数组顺序使奇数位于偶数前面

import (
	"testing"
	"reflect"
)

func TestEvenOddArray(t *testing.T) {
	got := EvenOddArray([]int{12, 34, 45, 9, 8, 90, 3})
	want := []int{45, 9, 3, 12, 34, 8, 90}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestEvenOddArray2(t *testing.T) {
	got := EvenOddArray2([]int{12, 34, 45, 9, 8, 90, 3})
	want := []int{3, 9, 45, 34, 8, 90, 12}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
