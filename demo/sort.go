package demo

import "fmt"

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort(data Sorter) {
	for i := 0; i < data.Len(); i++ {
		for j := i; j < data.Len(); j++ {
			if j+1 < data.Len() && data.Less(j+1, i) {
				data.Swap(j+1, i)
			}
		}
	}
	// for i := 1; i < data.Len(); i++ {
	// 	for j := 0; j < data.Len()-i; j++ {
	// 		if data.Less(j+1, j) {
	// 			data.Swap(j, j+1)
	// 		}
	// 	}
	// }
}

func isSorted(data Sorter) bool {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		if data.Less(i+1, i) {
			return false
		}
	}
	return true
}

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	date []*day
}

func (da *dayArray) Len() int {
	return len(da.date)
}
func (da *dayArray) Less(i, j int) bool {
	return da.date[i].num < da.date[j].num
}
func (da *dayArray) Swap(i, j int) {
	da.date[i], da.date[j] = da.date[j], da.date[i]
}

func InitSort() {
	data := []int{4, 2, 3, 100, 1, 10, 9, 5}
	a := IntArray(data)
	Sort(a)
	fmt.Println(isSorted(a))

	dayArr := &dayArray{date: []*day{{5, "FRI", "Friday"}, {1, "MON", "Monday"}, {6, "SAT", "Saturday"}, {0, "SUN", "Sunday"}, {2, "TUE", "Tuesday"}, {3, "WED", "Wednesday"}, {4, "THU", "Thursday"}}}
	Sort(dayArr)
	for _, v := range dayArr.date {
		fmt.Printf("%s ", v.longName)
	}
}
