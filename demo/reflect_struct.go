package demo

import (
	"fmt"
	"reflect"
)

type NotKnownType struct {
	s1, s2, s3 string
}

func (n NotKnownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

var secret Any = NotKnownType{"adn", "go", "Oberon"}

func InitReflectStruct() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	results := value.Method(0).Call(nil)
	fmt.Println(results)
}
