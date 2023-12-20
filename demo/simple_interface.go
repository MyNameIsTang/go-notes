package demo

import "fmt"

type Simpler interface {
	Get() int
	Set(v int)
}

type Simple struct {
	v int
}

func (s *Simple) Get() int {
	return s.v
}

func (s *Simple) Set(v int) {
	s.v = v
}

func InitSimpleInterface() {
	var s Simpler = &Simple{7}
	fmt.Printf("Test Get: %d\n", s.Get())
	s.Set(10)
	fmt.Printf("Test Get: %d\n", s.Get())
}
