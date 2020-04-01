package _379_电话目录管理系统

import "fmt"

type PhoneDirectory struct {
	allPhones []int
	usedPhones []bool
	index int
	size int
}


/** Initialize your data structure here
  @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func Constructor(maxNumbers int) PhoneDirectory {
	p := PhoneDirectory{
		allPhones: make([]int, maxNumbers, maxNumbers),
		usedPhones: make([]bool, maxNumbers, maxNumbers),
		index:0,
		size:maxNumbers,
	}
	for i:=0; i<maxNumbers; i++{
		p.allPhones[i] = i
	}
	return p
}


/** Provide a number which is not assigned to anyone.
  @return - Return an available number. Return -1 if none is available. */
func (p *PhoneDirectory) Get() int {
	if p.index == p.size {
		return -1
	}
	val := p.allPhones[p.index]
	fmt.Println("get, ", p.index, " ", val)
	p.usedPhones[val] = true
	p.index ++
	fmt.Println("get,,,", p.index)
	return val
}


/** Check if a number is available or not. */
func (p *PhoneDirectory) Check(number int) bool {
	if number > p.size{
		return false
	}
	return ! p.usedPhones[number]
}


/** Recycle or release a number. */
func (p *PhoneDirectory) Release(number int)  {
	if number < p.size && p.usedPhones[number]{
		p.usedPhones[number] = false
		p.index --
		p.allPhones[p.index] = number
	}
}


/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */
