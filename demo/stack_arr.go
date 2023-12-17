package demo

import (
	"fmt"
	"strconv"
)

const LIMIT = 4

type Stack1 struct {
	emIndex int
	arr     [LIMIT]int
}

func (s *Stack1) Push(v int) {
	if s.emIndex+1 > LIMIT {
		return // stack is full!
	}
	s.arr[s.emIndex] = v
	s.emIndex += 1
}

func (s *Stack1) Pop() (res int) {
	// if s.emIndex == 0 {
	// 	res = s.arr[s.emIndex]
	// 	s.arr[0] = 0
	// 	return
	// }
	// res = s.arr[s.emIndex-1]
	// s.arr[s.emIndex-1] = 0
	// s.emIndex -= 1
	// return
	s.emIndex--
	return s.arr[s.emIndex]
}

type Stack [LIMIT]int

func (s *Stack) Push(v int) {
	for i, v := range s {
		if v == 0 {
			s[i] = v
			return
		}
	}
}

func (s *Stack) Pop() int {
	v := 0
	for i := len(*s) - 1; i >= 0; i-- {
		if v = s[i]; v != 0 {
			s[i] = 0
			return v
		}
	}
	return v
}

func (s Stack) String() string {
	var str string
	for i, v := range s {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(v) + "]" + " "
	}
	return str
}

func InitStackArr() {
	arr := new(Stack)
	fmt.Println(arr)
	arr.Push(2)
	arr.Push(39)
	l := arr.Pop()
	fmt.Printf("the last one is %v\n", l)
	fmt.Println(arr)
}
