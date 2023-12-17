package demo

import (
	"fmt"
	"strconv"
)

type Celsius float64

func (c Celsius) String() string {
	return "the value is: " + strconv.FormatFloat(float64(c), 'f', 1, 32) + "°C"
}

func InitCelsius() {
	var c Celsius = 19.3
	fmt.Println(c)
}
