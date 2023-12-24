package demo

import "fmt"

type Miner interface {
	Len2() int
	Less2(i, j int) bool
	Get2(i int) Any
}

func Min(data Miner) Any {
	var minIndex int = 0
	for i := 1; i < data.Len2(); i++ {
		if data.Less2(i, minIndex) {
			minIndex = i
		}
	}
	return data.Get2(minIndex)
}

func (ia IntArray) Len2() int {
	return len(ia)
}
func (ia IntArray) Less2(i, j int) bool {
	return ia[i] < ia[j]
}
func (ia IntArray) Get2(i int) Any {
	return ia[i]
}

func InitMinInterface() {
	arr := []int{3, 4, 10, 5, 20, 9, 2}
	ia1 := IntArray(arr)
	fmt.Printf("the min value is:%v", Min(ia1))
}
