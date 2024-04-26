## 出于性能考虑的实用代码片段

1. 字符串

   - 如何修改字符串中的一个字符
     ```
       str := "hello"
       c := []byte(str)
       c[0] = 'c'
       s2 := string(c)
     ```
   - 如何获取字符串的子串
     ```
       substr := str[n:m]
     ```
   - 如何使用 for 或者 for-range 遍历一个字符串
     ```
       for i:=0; l < len(str); i++{
          s := str[i]
       }
       for ix, ch := range str {}
     ```
   - 如何获取一个字符串的字节数：`len(str)`
     - 获取一个字符串的字符数：
       - 最快速：`utf8.RuneCountInString(str)`，
       - `len([]rune(str))`
   - 如何连接字符串
     - 最快速：`bytes.Buffer`
     - `Strings.Join()`
     - 使用 `+=`
   - 如何解析命令行参数：使用 os 或者 flag 包

2. 数组和切片
   - 定义方式：
     - 创建：
       - `arr1 := new([len]type)`
       - `slice1 := make([]type, len)`
     - 初始化：
       - `arr1 := [...]type{i1, i2, i3}`
       - `arrKeyValue := [len]type{i1: val1, i2: val2}`
       - `var slice1 []type = arr1[start:end]`
   - 如何截断数组或者切片的最后一个元素：`line := line[:len(line)-1]`
   - 如何使用 for 或者 for-range 遍历一个数组（或者切片）
     ```
       for i:=0; i < len(arr); i++ {
         v := arr[i]
       }
       for i, v := range arr {}
     ```
   - 如何在一个二维数组或者切片 arr2Dim 中查找一个指定值 V：
     ```
       found := false
       Found: for row := range arr2Dim{
        for column := range arr2Dim[row] {
          if arr2Dim[row][column] == V {
            found = true
            break Found
          }
        }
       }
     ```
3. 映射

   - 定义方式
     - 创建：`map1 := make(map[keytype]valuetype)`
     - 初始化：`map1 := map[string]int{"one": 1, "two": 2}`
   - 如何使用 for 或者 for-range 遍历一个映射
     ```
       for key, value := range map1 { }
     ```
   - 如何在一个映射中检测键 key1 是否存在：`val1, isPresent := map1[key1]`
   - 如何在映射中删除一个键：delete(map1, key1)

4. 结构体
   - 定义方式
     - 创建：
       ```
         type struct1 struct {
           field1 type1
         }
         ms := new(struct1)
       ```
     - 初始化：`ms := &struct1{10, "23"}`
   - 当结构体的命名以大写字母开头时，该结构体在包外可见。 通常情况下，为每个结构体定义一个构建函数，并推荐使用构建函数初始化结构体：
     ```
       ms := NewStruct11{10, 14.1, "Chris"}
       func NewStruct11(n int, f float32, name string) *struct1 {
        return &struct1{n, f, name}
       }
     ```
5. 接口

   - 如何检测一个值 v 是否实现了接口 Stringer
     ```
       if v, ok := v.(Stringer); ok {
         fmt.Printf("implements String(): %s\n", v.String())
       }
     ```
   - 如何使用接口实现一个类型分类函数：
     ```
       func classifier (items ...interface{}){
         for i, x := range items {
           switch x.(type) {
             case bool:
                fmt.Printf("param #%d is a bool\n", i)
             case float64:
                fmt.Printf("param #%d is a float64\n", i)
             case int, int64:
                fmt.Printf("param #%d is an int\n", i)
             case nil:
                fmt.Printf("param #%d is nil\n", i)
             case string:
                fmt.Printf("param #%d is a string\n", i)
             default:
                fmt.Printf("param #%d’s type is unknown\n", i)
             }
         }
       }
     ```

6. 函数
   - 如何使用内建函数 `recover()`终止 `panic()`过程
     ```
       func protect(g func()) {
         defer func(){
           log.Println("done")
           if x := recover(); x != nil {
             log.Printf("run time panic: %v", x)
           }
         }()
         log.Println("start")
         g()
       }
     ```
7. 文件
   - 如何打开一个文件并读取
     ```
       file, err := os.Open("input.dat")
       if err != nil {
         fmt.Printf("An error occurred on opening the inputfile\n" +
            "Does the file exist?\n" +
            "Have you got acces to it?\n")
         return
       }
       defer file.Close()
       iReader := bufio.NewReader(file)
       for {
        str, err := iReader.ReadString('\n')
        if err != nil {
          return
        }
         fmt.Printf("The input was: %s", str)
       }
     ```
   - 如何通过切片读写文件
     ```
       func cat(f *file.File) {
         const NBUF = 512
         var buf [NBUF]byte
         for {
          switch nr, er := f.Read(buf[:]); true {
            case nr < 0:
              fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n",
              f.String(), er.String())
              os.Exit(1)
            case nr == 0:
              return
            case nr > 0:
              if nw, ew := file.Stdout.Write(buf[0:nr]); nw != nr {
                fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n",
                f.String(), ew.String())
              }
          }
         }
       }
     ```
