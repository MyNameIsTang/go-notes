package demo

import "fmt"

func mapFunc(mf func(Any) Any, list []Any) []Any {
	result := make([]Any, len(list))
	for ix, v := range list {
		result[ix] = mf(v)
	}
	return result
}

func mapFunc2(mf func(Any) Any, list ...Any) []Any {
	result := make([]Any, len(list))
	for ix, v := range list {
		result[ix] = mf(v)
	}
	return result
}

func InitMapFunctionInterface() {
	list := []Any{0, 1, 2, 3, 4, 5, 6, 7}
	mf := func(v Any) Any {
		switch t := v.(type) {
		case string:
			return t + t
		case int:
			return t * 2
		default:
			return t
		}
	}
	fmt.Printf("%v\n", mapFunc(mf, list))
	list2 := []Any{"123", "456", "789"}
	fmt.Printf("%v\n", mapFunc(mf, list2))
	fmt.Printf("%v\n", mapFunc2(mf, list2...))
}
