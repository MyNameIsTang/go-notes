package demo

import (
	"os"
	"strconv"
)

type Stringer interface {
	String() string
}

type Celsius1 float64

func (c Celsius1) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " °C"
}

type Day1 int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (d Day1) String() string {
	return dayName[d]
}

func print(args ...Any) {
	for i, v := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
		}
		switch a := v.(type) {
		case Stringer:
			os.Stdout.WriteString(a.String())
		case int:
			os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			os.Stdout.WriteString(a)
		default:
			os.Stdout.WriteString("???")
		}
	}
}

func InitPrint() {
	print(Day1(1), "was", Celsius1(82.32))
}
