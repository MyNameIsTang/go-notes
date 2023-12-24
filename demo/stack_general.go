package demo

import "fmt"

type Stack2 struct {
	emIndex int
	arr     []Any
}

func (s Stack2) Len() int {
	return len(s.arr)
}

func (s Stack2) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack2) Push(v Any) {
	s.arr = append(s.arr, v)
	s.emIndex += 1
}
func (s *Stack2) Pop() (interface{}, error) {
	s.emIndex--
	if s.emIndex < 0 {
		return 0, fmt.Errorf("无")
	}
	last := s.arr[s.emIndex]
	s.arr = s.arr[:s.emIndex]
	return last, nil
}

func InitStackGeneral() {
	stack2 := new(Stack2)
	stack2.Push(2)
	stack2.Push(3)
	stack2.Push(4)
	fmt.Println(stack2.arr)
	l, _ := stack2.Pop()
	fmt.Println(l)
	fmt.Println(stack2.arr)
}
