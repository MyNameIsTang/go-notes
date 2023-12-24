package demo

var i = 5
var str = "ABV"

type Any interface{}

type specicalString string

var whatIsThis specicalString = "hello"

type AbsInterface interface {
	Abs() float32
}
type SqrInterface interface {
	Sqr() float32
}

var ai AbsInterface
var si SqrInterface

func InitEmptyInterface() {
	// var val Any
	// val = 5
	// fmt.Printf("val has the value: %v\n", val)
	// val = str
	// fmt.Printf("val has the value: %v\n", val)
	// val = Person{"Tom", "jk"}
	// fmt.Printf("val has the value: %v\n", val)

	// testFunc := func(any Any) {
	// 	switch v := any.(type) {
	// 	case bool:
	// 		fmt.Printf("any %v is a bool type", v)
	// 	case int:
	// 		fmt.Printf("any %v is an int type", v)
	// 	case float32:
	// 		fmt.Printf("any %v is a float32 type", v)
	// 	case string:
	// 		fmt.Printf("any %v is a string type", v)
	// 	case specicalString:
	// 		fmt.Printf("any %v is a special String!", v)
	// 	default:
	// 		fmt.Println("unknown type!")
	// 	}
	// }
	// testFunc(whatIsThis)

	// var empty Any
	// pp := new(Point)

	// empty = pp
	// ai = empty.(AbsInterface)
	// si = ai.(SqrInterface)
	// empty = si
}
