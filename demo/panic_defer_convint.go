package demo

import (
	"fmt"
	"log"
	"math"
)

func convertInt64ToInt(i int64) int {
	if i > math.MaxInt32 || i < math.MinInt32 {
		panic(fmt.Sprintf("%v 不符合int32范围内", i))
	}
	return int(i)
}

func intFormInt64(i int64) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = convertInt64ToInt(i)
	return
}

func InitPanicDeferConvint() {
	l := int64(math.MaxInt32 + 15000)
	v, err := intFormInt64(l)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(v)
}
