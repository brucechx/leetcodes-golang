package _379_电话目录管理系统

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	phone := Constructor(3)
	phone.Release(2)
	t.Log(phone.Get()) // 0
	phone.Release(2)
	phone.Release(0)
	t.Log(phone.Get()) // 1
	t.Log(phone.Get()) // 0 ?
	phone.Check(1) // false
	t.Log(phone.Get()) // 2
	phone.Release(0)
	t.Log(phone.Get()) // 0
	phone.Release(0)
	phone.Release(0)
	t.Log(phone.Get()) // 0
	phone.Check(1) // false
	phone.Release(1)
}
