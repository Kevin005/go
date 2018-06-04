package main

import "fmt"

type Slice []int

func NewSlice() Slice{
	return make(Slice,0)
}

func (s *Slice) Add(elem int) *Slice{
	*s = append(*s,elem)
	fmt.Print(elem)
	return s
}

func main(){
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
	//a1(&s)
}

//defer识别的是最后一个函数 所以这段代码等效于
func a2(s *Slice){
	s1 := s.Add(1)
	defer s1.Add(2)
	s.Add(3)
}

//可以用下面的代码验证
func a1(s *Slice){
	defer s.Add(1).Add(5).Add(6).Add(7).Add(2)
	s.Add(3)
	//输出 156732
}