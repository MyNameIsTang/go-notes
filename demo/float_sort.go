package demo

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Float64Array []float64

func (fa Float64Array) Len() int {
	return len(fa)
}

func (fa Float64Array) Less(i, j int) bool {
	return fa[i] < fa[j]
}

func (fa Float64Array) Swap(i, j int) {
	fa[i], fa[j] = fa[j], fa[i]
}

func (_ Float64Array) NewFloat64Array() *Float64Array {
	m := new(Float64Array)
	for i := 0; i < 25; i++ {
		(*m)[i] = 0
	}
	return m
}

func (fa Float64Array) List1(flag string) string {
	var str string
	for i, v := range fa {
		if i != 0 {
			str += flag
		}
		str += strconv.FormatFloat(v, 'f', 2, 32)
	}
	return str
}

func (fa Float64Array) String() string {
	return fa.List1(",")
}

func (fa Float64Array) Fill(n int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < n; i++ {
		fa[i] = 100 * rand.Float64()
	}
}

func InitFloatSort() {
	fa := []float64{1.2, 4, 2.4, 4.3, 80, 10.2, 0.2}
	fa1 := Float64Array(fa)
	Sort(fa1)
	fmt.Printf("fa sort result %v\n", fa1)
	fl1 := make([]float64, 10)
	fa2 := Float64Array(fl1)
	fa2.Fill(10)
	fmt.Printf("fa2 value %s\n", fa2)
}
