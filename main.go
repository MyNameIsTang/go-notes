package main

import (
	"basic/demo"
)

// _ "basic/urlshortener"

// type Person struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Email string `json:"email"`
// }

func main() {

	demo.InitRemove3Till5Char()
	// var f float64 = 3.4
	// v := reflect.TypeOf(f)
	// v1 := reflect.ValueOf(&f)
	// fmt.Println(reflect.Float64 == v.Kind())
	// fmt.Println(v1.Type())
	// v1 = v1.Elem()
	// fmt.Println(v1.CanSet())
	// v1.SetFloat(3.14159)
	// fmt.Println(v1.Interface())
	// demo.InitMinInterface()

	// obj := []int{1, 32}

	// runtime.SetFinalizer(&obj, func(o interface{}) {
	// 	fmt.Println("Object has been freed!")
	// })

	// time.Sleep(2 * time.Second)
	// fmt.Println("123")
	// runtime.GC()

	// var m runtime.MemStats
	// runtime.ReadMemStats(&m)
	// fmt.Printf("%d Kb\n", m.Alloc/1024)

	// demo.InitStackArr()
	// p := new(demo.Person)
	// p.SetFirstName("Tom")
	// fmt.Println(p.FirstName())
	// demo.InitMethodSet1()
	// demo.InitAmonymousStruct()
	// demo.InitEmbeddStruct()
	// demo.InitStructAnonymouse()
	// demo.InitStructTag()

	// demo.InitVCard()
	// demo.InitRectangle()

	// struct1 := new(demo.ExpStruct)
	// struct1.Mf1 = 1.1
	// struct1.Mi1 = 1
	// fmt.Println(struct1)

	// urlshortener.Server()

	// person := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	// jsonData, err := json.Marshal(person)
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// fmt.Print(string(jsonData))

	// var intP *int
	// var i1 = 5
	// fmt.Printf("memory is %p\n", &i1)
	// intP = &i1
	// fmt.Printf("intP memory is %p\n, intP value is %v", intP, *intP)
	// s := "hhhhh"
	// var ptr *string = &s
	// *ptr = "2333"
	// fmt.Printf("md is %p, ptr value is %s, s value is %s", ptr, *ptr, s)
	// var p *int = nil
	// *p = 0
	// test()
	// switch1(3)
	// fmt.Print(Season(2))
	// for i := 0; i < 15; i++ {
	// 	fmt.Printf("number is %d\n", i)
	// }
	// 	i := 0

	// HEAR:
	// 	if i < 15 {
	// 		fmt.Printf("number is %d\n", i)
	// 		i++
	// 		goto HEAR
	// 	}

	// for i, str := 0, "G"; i < 25; i, str = i+1, str+"G" {
	// 	fmt.Println(str)
	// }
	// for i := 0; i <= 10; i++ {
	// 	fmt.Printf("序号是 %d, 补码 是 %b\n", i, ^i)
	// }
	// for i := 1; i <= 100; i++ {
	// if i%3 == 0 && i%5 == 0 {
	// 	fmt.Println("FizzBuzz")
	// 	continue
	// }
	// if i%3 == 0 {
	// 	fmt.Println("Fizz")
	// 	continue
	// }
	// if i%5 == 0 {
	// 	fmt.Println("Buzz")
	// 	continue
	// }
	// fmt.Println(i)
	// 	switch {
	// 	case i%3 == 0 && i%5 == 0:
	// 		fmt.Println("FizzBuzz")
	// 	case i%3 == 0:
	// 		fmt.Println("Fizz")
	// 	case i%5 == 0:
	// 		fmt.Println("Buzz")
	// 	default:
	// 		fmt.Println(i)
	// 	}
	// }

	// for i := 1; i <= 200; i++ {
	// 	if i%20 == 0 {
	// 		fmt.Println("*")
	// 		continue
	// 	}
	// 	fmt.Print("* ")
	// }
	// var i int = 5
	// for i >= 0 {
	// 	i--
	// 	fmt.Println(i)
	// }

	// str := "123dnas你打开撒娇你大街上"
	// fmt.Println("index int(rune) rune    char bytes")
	// for i, v := range str {
	// 	fmt.Printf("%-2d    %d 	%U  '%c'  % X \n", i, v, v, v, []byte(string(v)))
	// }
	// for i := 0; i < 5; i++ {
	// 	var v int
	// 	fmt.Printf("%d ", v)
	// 	v = 5
	// }

	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }
	// for i := 0; i < 3; {
	// 	fmt.Println("Value of i:", i)
	// }

	// s := ""
	// for s != "aaaaa" {
	// 	fmt.Println("Value of s:", s)
	// 	s = s + "a"
	// }
	// for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
	// 	s = i+1, j+1, s+"a" {
	// 	fmt.Println("Value of i, j, s:", i, j, s)
	// }
	// var i int = 5
	// for {
	// 	i = i - 1
	// 	fmt.Printf("The variable i is now: %d\n", i)
	// 	if i < 0 {
	// 		break
	// 	}
	// }
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j > 5 {
	// 			break
	// 		}
	// 		fmt.Print(j)
	// 	}
	// 	fmt.Print("-")
	// }

	// // LABEL1:
	// for i := 0; i <= 5; i++ {
	// 	for j := 0; j <= 5; j++ {
	// 		if j == 4 {
	// 			// continue LABEL1
	// 			break
	// 		}
	// 		fmt.Printf("i is: %d, and j is: %d\n", i, j)
	// 	}
	// }
	// 0,1,2
	// i := 0
	// for { //since there are no checks, this is an infinite loop
	// 	if i >= 3 {
	// 		break
	// 	}
	// 	//break out of this for loop when this condition is met
	// 	fmt.Println("Value of i is:", i)
	// 	i++
	// }
	// fmt.Println("A statement just after for loop.")
	// 1,3,5
	// for i := 0; i < 7; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd:", i)
	// }
	// i := 0

	// Multiply(1, 3, &i)
	// fmt.Print(i)

	// s := []string{"dsa", "122", "dcc"}

	// Greeting("hdsai", s...)
	// Varargs(s...)
	// F1(s, "23", 2)

	// function1()'
	// a()

	// func1("Go")
	// a := []string{"1", "2"}
	// fmt.Print(cap(a))
	// a := make([]int, 2)
	// fmt.Printf("%#v", a)

	// result := 0
	// for i := 0; i <= 10; i++ {
	// 	ii, result := fibonacci(i)
	// 	fmt.Printf("%d , %d\n", result, ii)
	// }
	// recursive(10)
	// for i := uint64(0); i < uint64(30); i++ {
	// 	fmt.Printf("原始值是%d, 阶乘是%d\n", i, factorial(i))
	// }
	// callback(2, Add)

	// fmt.Printf("index is %d", strings.IndexFunc("23 2dsa", unicode.IsSpace))
	// s := "21d年少多金看12&*@#&导出s"
	// strings.IndexFunc()
	// fmt.Print(strings.Map(IsAscii, s))
	// fv := func() {
	// 	fmt.Println("hello world")
	// }
	// fv()
	// fmt.Printf("%#v", reflect.TypeOf(fv))
	// fmt.Print(f())
	// f := Adder()
	// fmt.Print(f(1), "-")
	// fmt.Print(f(30), "-")
	// fmt.Print(f(53))
	// fi1 := fi()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(fi1())
	// }
	// start := time.Now()

	// where := func() {
	// 	_, file, line, _ := runtime.Caller(1)
	// 	log.Printf("%s:%d\n", file, line)
	// 	// log.SetFlags(log.Llongfile)
	// 	// log.Print("---")
	// }

	// var result uint64
	// for i := 0; i < 41; i++ {
	// 	result = fibonacci(i)
	// 	fmt.Println(result)
	// }

	// // fmt.Println("213123")
	// // where()

	// // fmt.Println("snakdnka")
	// // where()

	// end := time.Now()
	// delta := end.Sub(start)
	// fmt.Print(delta)

	// a := [...]string{"1", "2", "3", "4"}
	// for i := range a {
	// 	fmt.Println("Array item", i, "is", a[i])
	// }

	// var arr2 [5]int
	// arr1 := new([5]int)
	// fmt.Printf("%#v, %v", arr2, *arr1)
	// arr1 := [3]int{1, 2, 3}
	// arr2 := arr1
	// arr2[1] = 34
	// fmt.Printf("arr1 is %v, arr2 is %v", arr1, arr2)
	// fmt.Printf("%p, %p", &arr1, &arr2)
	// fmt.Print(arr1 == arr2)

	// var arr1 [16]int
	// for i := 0; i <= 15; i++ {
	// 	arr1[i] = i
	// 	fmt.Printf("index is %d, value is %d\n", i, i)
	// }

	// for i := 0; i < 50; i++ {
	// 	fibonacci(i)
	// }
	// for i := 0; i < len(fib); i++ {
	// 	fmt.Println(fib[i])
	// }
	// for i := 0; i < 3; i++ {
	// 	fp(&[3]int{i, i * i, i * i * i})
	// }

	// const (
	// 	WIDTH  = 1920
	// 	HEIGHT = 1080
	// )

	// var screen [WIDTH][HEIGHT]int

	// for i := 0; i < HEIGHT; i++ {
	// 	for j := 0; j < WIDTH; j++ {
	// 		screen[j][i] = 0
	// 	}
	// }

	// fmt.Println(screen)
	// array := [3]float64{7.3, 1.2, 34.2}
	// x := Sum(&array)
	// fmt.Print(x)
	// var slice1 []int
	// var arr1 = [10]int{1}
	// slice1 = arr1[3:5]
	// for i := 0; i < len(arr1); i++ {
	// 	arr1[i] = i
	// }
	// fmt.Printf("slice1:%v,len is: %d, cap is %d\n", slice1, len(slice1), cap(slice1))
	// slice1 = slice1[0:3]
	// fmt.Printf("slice1:%v,len is: %d, cap is %d\n", slice1, len(slice1), cap(slice1))
	// slice1[2] = 10
	// fmt.Printf("slice1:%v,len is: %d, cap is %d\n", slice1, len(slice1), cap(slice1))
	// fmt.Printf("arr1:%v,arr1 is: %d, arr1 is %d\n", arr1, len(arr1), cap(arr1))
	// slice1 = slice1[0:7]
	// fmt.Printf("slice1:%v,len is: %d, cap is %d\n", slice1, len(slice1), cap(slice1))
	// [o,l,a],[g,o],[l,...],[...]
	// b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	// fmt.Printf("b[1:4]:%c、b[:2]:%c、b[2:]:%c、b[:]:%c", b[1:4], b[:2], b[2:], b[:])
	// var slice1 []int = make([]int, 10, 20)
	// for i := 0; i < len(slice1); i++ {
	// 	slice1[i] = i * 3
	// }
	// fmt.Print(slice1, len(slice1), cap(slice1))
	// fmt.Print(FibonacciFuncarray([]int{0, 1, 2, 3, 4, 5, 6}))

	// s := make([]byte, 5)
	// fmt.Println(len(s), cap(s))
	// s = s[2:4]
	// fmt.Println(len(s), cap(s))
	// s1 := []byte{'p', 'o', 'e', 'm'}
	// s2 := s1[2:]
	// fmt.Println(s2)
	// s2[1] = 't'
	// fmt.Println(s1, s2)
	// var buffer bytes.Buffer
	// for i := 0; i < 10; i++ {
	// 	buffer.WriteString(fmt.Sprint(i))
	// }
	// fmt.Print(buffer.String())
	// sl1 := make([]byte, 3, 5)
	// sl1 = Append(sl1, make([]byte, 10))
	// // sl1[4] = 4
	// fmt.Print(sl1)

	// sl1 := make([]byte, 10)

	// sl2, sl3 := sl1[:5], sl1[5:]
	// fmt.Println(sl1)
	// fmt.Println(sl2)
	// fmt.Println(sl3)
	// items := [...]int{10, 20, 30, 40, 50}

	// for i := range items {
	// 	// v *= 2
	// 	items[i] *= 2
	// }
	// fmt.Print(items)
	// var numbers = []int{1, 2, 4, 5, 6}
	// fmt.Print(Sum(numbers...))

	// var numbers = []float32{1.2, 2, 4, 5.2, 6}
	// fmt.Print(Sum(numbers))

	// sum, aver := SumAndAverage([]int{1, 2, 4, 5, 6})
	// fmt.Print(sum, aver)

	// fmt.Println(minSlice([]int{1, 2, 4, 5, 6}))
	// fmt.Println(maxSlice([]int{1, 2, 4, 5, 6}))

	// var slice1 = make([]int, 4)
	// slice1 = slice1[2:3]
	// fmt.Print(len(slice1))
	// slFrom := []int{1, 2, 3}
	// slTo := make([]int, 10)
	// n := copy(slTo, slFrom)
	// fmt.Println(slTo)
	// fmt.Println(n)

	// sl3 := []int{1, 2, 3}
	// sl3 = append(sl3, 4, 5, 6)
	// fmt.Println(sl3)
	// sl3 := []int{1, 2, 3}
	// sl4 := MagnifySlice(sl3, 2)
	// fmt.Print(sl3, '-', sl4)

	// sl3 := []int{1, 2, 3, 4, 5, 6}
	// fmt.Print(Filter(sl3, func(v int) bool {
	// 	if v%2 == 0 {
	// 		return true
	// 	}
	// 	return false
	// }))

	// sl3 := []int{1, 2, 3, 4, 5, 6}
	// sl4 := []int{7, 8, 9}

	// fmt.Print(InsertStringSlice(sl3, sl4, 2))

	// sl3 := []int{1, 2, 3, 4, 5, 6}
	// fmt.Print(RemoveStringSlice(sl3, 2, 4))
	// s := "dsancja1"
	// a := []byte(s)
	// r := []rune(s)
	// fmt.Printf("%c, %c", a, r)

	// sl3 := []int{1, 4, 3, 7, 5, 2}
	// fmt.Println(sort.IntsAreSorted(sl3))
	// sort.Ints(sl3)
	// fmt.Println(sl3)
	// fmt.Println(sort.IntsAreSorted(sl3))
	// fmt.Print(sort.SearchInts(sl3, 3))

	// str := "dsnajk你打加快1231"
	// f, l := SplitString(str, 3)
	// fmt.Print(f, "--", l)

	// str := "ddsnajj221231"
	// fmt.Print(str[len(str)/2:] + str[:len(str)/2])

	// fmt.Print(StringReverse(str))
	// fmt.Printf("%c", Uniq([]byte(str)))

	// var arr []byte = []byte{'a', 'b', 'a', 'a', 'a', 'c', 'd', 'e', 'f', 'g'}
	// arru := make([]byte, len(arr)) // this will contain the unique items
	// ixu := 0                       // index in arru
	// tmp := byte(0)
	// for _, val := range arr {
	// 	if val != tmp {
	// 		arru[ixu] = val
	// 		fmt.Printf("%c ", arru[ixu])
	// 		ixu++
	// 	}
	// 	tmp = val
	// }
	// var arr []int = []int{1, 32, 4, 2, 2, 6, 20, 5}
	// fmt.Print(bubblesort(arr))

	// fmt.Print(MapFunction(func(v int) int {
	// 	return v * 10
	// }, []int{1, 3, 5, 2, 6}))

	// var map1 = map[string]int{"1": 1, "2": 2}
	// for key, val := range map1 {
	// 	fmt.Println(key, "-", val)
	// }

	// var week = map[string]int{"Monday": 1, "Tuesday": 2, "Wednesday": 3, "Thursday": 4, "Friday": 5, "Saturday": 6, "Sunday": 7}
	// for key, val := range week {
	// 	fmt.Printf("the day is %s, num is %d\n", key, val)
	// }
	// if _, ok := week["Tuesday"]; ok {
	// 	fmt.Printf("Tuesday\n")
	// }
	// if _, ok := week["Hollyday"]; ok {
	// 	fmt.Printf("Hollyday\n")
	// }

	// items := make([]map[int]int, 5)
	// for i := range items {
	// 	items[i] = make(map[int]int, 1)
	// 	items[i][1] = 2
	// }
	// fmt.Printf("Version A: Value of items: %v\n", items)

	// var barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
	// 	"delta": 87, "echo": 56, "foxtrot": 12,
	// 	"golf": 34, "hotel": 16, "indio": 87,
	// 	"juliet": 65, "kili": 43, "lima": 98}
	// keys := make([]string, len(barVal))
	// i := 0
	// for k := range barVal {
	// 	keys[i] = k
	// 	i++
	// }
	// sort.Strings(keys)
	// fmt.Println(keys)
	// for _, k := range keys {
	// 	fmt.Printf("key: %v, Value is %v\n", k, barVal[k])
	// }

	// var barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
	// 	"delta": 87, "echo": 56, "foxtrot": 12,
	// 	"golf": 34, "hotel": 16, "indio": 87,
	// 	"juliet": 65, "kili": 43, "lima": 98}
	// invMap := make(map[int]string, len(barVal))
	// for key, val := range barVal {
	// 	invMap[val] = key
	// }
	// fmt.Print(invMap)

	// x := 42
	// var p *int = &x
	// uintPrtValue := uintptr(unsafe.Pointer(p))
	// fmt.Print(uintPrtValue)

	// listener, err := net.Listen("tcp", "localhost:9000")
	// if err != nil {
	// 	fmt.Println("Error", err.Error())
	// 	return
	// }

	// defer listener.Close()
	// fmt.Println("Server listening on localhost:9000")

	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		fmt.Println("Error", err.Error())
	// 		return
	// 	}
	// 	go handleConnection(conn)
	// }

	// li := list.New()
	// li.PushBack(101)
	// li.PushBack(102)
	// li.PushBack(103)
	// for n := li.Front(); n != nil; n = n.Next() {
	// 	fmt.Printf("node is %v, value is %v\n", n, n.Value)
	// }

	// i := 100
	// fmt.Print(unsafe.Sizeof(i))

	// str := "123ndsau12i"
	// ok1, _ := regexp.Match("/\\d/", []byte(str))
	// fmt.Println(ok1)
	// str2 := "123"
	// ok2, _ := regexp.MatchString("/\\d/g", str2)
	// fmt.Println(ok2)

	// searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	// pat := "[0-9]+.[0-9]+"
	// f := func(s string) string {
	// 	v, _ := strconv.ParseFloat(s, 32)
	// 	return strconv.FormatFloat(v*2, 'f', 2, 32)
	// }
	// if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
	// 	fmt.Println("match found")
	// }
	// re, _ := regexp.Compile(pat)
	// str := re.ReplaceAllString(searchIn, "##.#")
	// fmt.Println(str)
	// str2 := re.ReplaceAllStringFunc(searchIn, f)
	// fmt.Println(str2)

	// im := big.NewInt(math.MaxInt64)
	// in := im
	// io := big.NewInt(1956)
	// ip := big.NewInt(1)
	// ip.Mul(im, in).Add(ip, im).Div(ip, io)
	// fmt.Printf("Big Int: %v\n", ip)
	// rm := big.NewRat(math.MaxInt64, 1956)
	// rn := big.NewRat(-1956, math.MaxInt64)
	// ro := big.NewRat(19, 56)
	// rp := big.NewRat(1111, 2222)
	// rq := big.NewRat(1, 1)
	// rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	// fmt.Printf("Big Rat:%v\n", rq)

	// test1 := pack.ReturnStr()
	// fmt.Print(test1)
	// fmt.Print(pack.pack1Float)

	// if ok := greetings.IsEvening(); ok {
	// 	greetings.GoodNight()
	// }
	// even.MainOddven(100)

	// result := fibo.Fibonacci("*", 10)
	// fmt.Print(result)
}

// func handleConnection(conn net.Conn) {
// 	buffer := make([]byte, 1024)
// 	n, err := conn.Read(buffer)
// 	if err != nil {
// 		fmt.Println("Error", err.Error())
// 		return
// 	}
// 	message := string(buffer[:n])
// 	fmt.Println("Received message:", message)
// 	conn.Write([]byte("Message Received"))
// }

// func MapFunction(fn func(v int) int, arr []int) (res []int) {
// 	res = make([]int, len(arr))
// 	for i, v := range arr {
// 		res[i] = fn(v)
// 	}
// 	return
// }

// func bubblesort(arr []int) []int {
// 	for i := 1; i < len(arr); i++ {
// 		for j := 0; j < len(arr)-i; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// 	return arr
// }

// func Uniq(str []byte) (res []byte) {
// 	for i := range str {
// 		if i > 0 && str[i] != str[i-1] {
// 			res = append(res, str[i])
// 		}
// 	}
// 	return
// }

// func StringReverse(str string) string {
// 	as := []byte(str)
// 	l := len(as)
// 	for i, v := range as {
// 		if i >= l/2 {
// 			break
// 		}
// 		as[i], as[l-i-1] = as[l-i-1], v
// 	}
// 	return string(as)
// }

// func SplitString(str string, i int) (string, string) {
// 	r := []byte(str)
// 	f, l := r[:i], r[i:]
// 	return string(f), string(l)
// }

// func RemoveStringSlice(src []int, start, end int) []int {
// 	result := make([]int, len(src)-(end-start+1))
// 	at := copy(result, src[:start])
// 	copy(result[at:], src[end+1:])
// 	return result
// }

// func InsertStringSlice(src []int, det []int, start int) []int {
// 	result := make([]int, len(src)+len(det))
// 	at := copy(result, src[:start])
// 	at += copy(result[at:], det)
// 	copy(result[at:], src[start:])
// 	return result
// }

// func Filter(slice1 []int, fn func(int) bool) (res []int) {
// 	res = make([]int, 0)
// 	for _, v := range slice1 {
// 		if fn(v) {
// 			res = append(res, v)
// 		}
// 	}
// 	return
// }

// func MagnifySlice(slice1 []int, factor int) (res []int) {
// 	res = make([]int, len(slice1)*factor)
// 	copy(res, slice1)
// 	return
// }

// 用于切分切片
func MinSlice(slice1 []int) (min int) {
	min = slice1[0]
	for _, v := range slice1 {
		if min > v {
			min = v
		}
	}
	return
}

// func maxSlice(slice1 []int) (max int) {
// 	max = slice1[0]
// 	for _, v := range slice1 {
// 		if max < v {
// 			max = v
// 		}
// 	}
// 	return
// }

// func SumAndAverage(slice1 []int) (int, float32) {
// 	var sum int
// 	for _, v := range slice1 {
// 		sum += v
// 	}
// 	return sum, float32(sum / len(slice1))
// }

// func Sum(arrF []float32) (res float32) {
// 	for _, v := range arrF {
// 		res += v
// 	}
// 	return
// }

// func Sum(numbers ...int) (res int) {
// 	for _, v := range numbers {
// 		res += v
// 	}
// 	return
// }

// func Split(){

// }

// func Append(slice, data []byte) []byte {
// 	// res = make([]byte, cap(slice)+cap(data))
// 	var buffer bytes.Buffer
// 	buffer.Write(slice)
// 	if (buffer.Cap() - len(data)) < 0 {
// 		buffer.Grow(buffer.Len() + len(data))
// 	}
// 	buffer.Write(data)
// 	return buffer.Bytes()
// }

// func Sum(a *[3]float64) (sum float64) {
// 	for _, v := range a {
// 		sum += v
// 	}
// 	return
// }

// func fp(a *[3]int) {
// 	fmt.Println(a)
// }

// func fi() func() int {
// 	pre1, pre2 := 0, 1
// 	return func() int {
// 		pre1, pre2 = pre2, pre1+pre2
// 		return pre1
// 	}
// }

// func Adder() func(int) int {
// 	var x int
// 	return func(n int) int {
// 		x += n
// 		return x
// 	}
// }

// func f() (ret int) {
// 	defer func() {
// 		ret++
// 	}()
// 	return 1
// }

// func IsAscii(c rune) rune {
// 	if c > 127 {
// 		return ' '
// 	}
// 	return c
// }

// func callback(y int, f func(int, int)) {
// 	f(y, 4)
// }

// func Add(a, b int) {
// 	fmt.Printf("sum is, %d", a+b)
// }

// func factorial(n uint64) (fac uint64) {
// 	fac = 1
// 	if n > 0 {
// 		fac = n * factorial(n-1)
// 	}
// 	return
// }

// func recursive(n int) {
// 	if n > 0 {
// 		fmt.Println(n)
// 		recursive(n - 1)
// 	}
// }

// func FibonacciFuncarray(arr []int) (res []int) {
// 	res = make([]int, len(arr), cap(arr))
// 	for i, v := range arr {
// 		res[i] = fibonacci(v)
// 	}
// 	return
// }

// var fib [50]int

// func fibonacci(i int) (res int) {
// 	if fib[i] != 0 {
// 		res = fib[i]
// 		return
// 	}
// 	if i <= 1 {
// 		res = 1
// 	} else {
// 		v1 := fibonacci(i - 1)
// 		v2 := fibonacci(i - 2)
// 		res = v1 + v2
// 	}
// 	fib[i] = res
// 	return
// }

// func func1(s string) (n int, err error) {
// 	defer func() {
// 		log.Printf("func1(%q) = %d, %v", s, n, err)
// 	}()
// 	return 7, io.EOF
// }

// func a() {
// 	// i := 0
// 	// defer fmt.Print(i)
// 	// i++
// 	// return
// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println(i)
// 	}
// }

// func function1() {
// 	fmt.Println("in dsanjd dsakj")
// 	defer function2()
// 	fmt.Println("dsnandkjsa dsa")
// }

// func function2() {
// 	fmt.Println("function2 dshads")
// }

// func Varargs(args ...string) {
// 	for i := 0; i < len(args); i++ {
// 		fmt.Println(args[i])
// 	}
// }

// type Options struct {
// 	par1 int
// 	par2 string
// }

// func F1(values ...interface{}) {
// 	fmt.Printf("%v", values...)
// }

// func Greeting(prefix string, who ...string) {
// 	fmt.Printf("prefix is %s, other is %#v", prefix, who)
// }

// func Multiply(a, b int, reply *int) {
// 	*reply = a + b
// }

// func switch1(num int) {
// 	switch {
// 	case num < 1:
// 		fmt.Printf("%d < 1", num)
// 	case num > 2 && num < 4:
// 		fmt.Printf("%d > 2 && %d < 4", num, num)
// 	default:
// 		fmt.Printf("num == %d", num)
// 	}
// }

// func Season(month int) string {
// 	switch {
// 	case month >= 1 && month <= 3:
// 		return "春季"
// 	case month >= 4 && month <= 6:
// 		return "夏季"
// 	case month >= 7 && month <= 9:
// 		return "秋季"
// 	case month >= 10 && month <= 12:
// 		fallthrough
// 	default:
// 		return "冬季"
// 	}
// }

// func test() error {

// 	f, err := os.Open("./tools.md")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("filename is %s", f.Name())
// 	return nil
// }

// func flushICache(begin, end uintptr)

// type binOp func(int, int) int

// func getX2AndX3_2(input int) (x2, x3 int) {
// 	x2 = input * 2
// 	x3 = input * 3
// 	return
// }

// func multReturnVal(a, b int) (int, int, int) {
// 	return a + b, a * b, a - b
// }

// func multReturnVal_2(a, b int) (x1, x2, x3 int) {
// 	return a + b, a * b, a - b
// }

// func MySqrt(a float64) (float64, error) {
// 	if a < 0 {
// 		return float64(math.NaN()), errors.New("值小于0")
// 	}
// 	return math.Sqrt(a), nil
// }

// func MySqrt_2(a float64) (x float64, e error) {
// 	if a < 0 {
// 		x = float64(math.NaN())
// 		e = errors.New("值小于0")
// 		return
// 	}
// 	x = math.Sqrt(a)
// 	return
// }
