package demo

import "fmt"

type Car1 struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

type Cars []*Car1

func (cs Cars) Process(f func(c *Car1, ix int)) {
	for i, v := range cs {
		f(v, i)
	}
}

func (cs Cars) FindAll(f func(car *Car1) bool) Cars {
	cars := make(Cars, 0)
	cs.Process(func(c *Car1, _ int) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func (cs Cars) Map(f func(car *Car1) Any) []Any {
	result := make([]Any, len(cs))
	cs.Process(func(c *Car1, ix int) {
		result[ix] = f(c)
	})
	return result
}

func MakeSortedAppender(manufacturers []string) (func(car *Car1, ix int), map[string]Cars) {
	sortedCars := make(map[string]Cars)

	for _, v := range manufacturers {
		sortedCars[v] = make(Cars, 0)
	}

	sortedCars["Default"] = make(Cars, 0)
	appender := func(c *Car1, ix int) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}

func InitCars() {
	ford := &Car1{"Fiesta", "Ford", 2008}
	bmw := &Car1{"XL 450", "BMW", 2011}
	merc := &Car1{"D600", "Mercedes", 2009}
	bmw2 := &Car1{"X 800", "BMW", 2008}
	allCars := Cars([]*Car1{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car1) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)
	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}
