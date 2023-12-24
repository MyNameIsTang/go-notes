package demo

import (
	"fmt"
	"reflect"
)

type T2 struct {
	A int
	B string
}

func InitReflectStruct2() {
	t := T2{23, "skidnd"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(89)
	s.Field(1).SetString("sunset strip")
	fmt.Println("t is now", t)
}
