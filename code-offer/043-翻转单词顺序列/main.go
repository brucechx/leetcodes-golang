package _43_翻转单词顺序列

import "strings"

func ReverseSentence(str string) string{
	words := strings.Split(str, " ")
	var res string
	for _, w := range words{
		res = w + " "+ res
	}
	return res[:len(res)-1]
}

// 递归
func ReverseSentence2(str string) (result string) {
	i := strings.Index(str, " ")
	if i >= 0 {
		return ReverseSentence2(str[i+1:]) + " " + str[:i]
	}
	return str
}

// 栈
func ReverseSentence3(str string) (result string) {
	if len(str) == 0{
		return str
	}
	s1 := NewStack() // 入栈
	s2 := NewStack() // 出栈
	for _, s := range str{
		s1.Push(string(s))
	}
	for ! s1.IsEmpty(){
		tmp := s1.Pop()
		if string(tmp) != " "{
			s2.Push(tmp)
		}else {
			for !s2.IsEmpty(){
				result += s2.Pop()
			}
			result += " "
		}
	}

	for ! s2.IsEmpty(){
		result += s2.Pop()
	}
	return
}


type Stack struct {
	Val []string
}

func NewStack() *Stack{
	return &Stack{}
}

func (s *Stack) Push(b string){
	s.Val = append(s.Val, b)
}

func (s *Stack) Pop() string{
	tmp := s.Val[len(s.Val) - 1]
	s.Val = s.Val[: len(s.Val) - 1]
	return tmp
}

func (s *Stack) IsEmpty() bool{
	return len(s.Val) == 0
}
