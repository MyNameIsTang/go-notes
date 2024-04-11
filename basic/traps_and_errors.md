## 常见的陷阱与错误

- 常见陷阱

  - 永远不要使用形如 `var p*a` 声明变量，这会混淆指针声明和乘法运算
  - 永远不要在 for 循环自身中改变计数器变量
  - 永远不要在 for-range 循环中使用一个值去改变自身的值
  - 永远不要将 goto 和前置标签一起使用
  - 永远不要忘记在函数名后加括号 ()，尤其是调用一个对象的方法或者使用匿名函数启动一个协程时
  - 永远不要使用 `new()` 一个 map，一直使用 `make()`
  - 当为一个类型定义一个 `String()` 方法时，不要使用 `fmt.Print` 或者类似的代码
  - 永远不要忘记当终止缓存写入时，使用 `Flush()` 函数
  - 永远不要忽略错误提示，忽略错误会导致程序崩溃
  - 不要使用全局变量或者共享内存，这会使并发执行的代码变得不安全
  - `println()` 函数仅仅是用于调试的目的

- 最佳实践：对比以下使用方式：
  - 使用正确的方式初始化一个元素是切片的映射，例如 `map[type]slice`
  - 一直使用逗号 ok 模式或者 checked 形式作为类型断言
  - 使用一个工厂函数创建并初始化自己定义类型
  - 仅当一个结构体的方法想改变结构体时，使用结构体指针作为方法的接受者，否则使用一个结构体值类型

1. 误用短声明导致变量覆盖

   - if 语句
     ```
       var remember bool = false
       if something {
           remember := true //错误
       }
     ```
   - 此类错误也容易在 for 循环中出现，尤其当函数返回一个具名变量时难于察觉，例如：
     ```
       func shadow() (err error) {
         x, err := check1() // x 是新创建变量，err 是被赋值
         if err != nil {
           return // 正确返回 err
         }
         if y, err := check2(x); err != nil { // y 和 if 语句中 err 被创建
           return // if 语句中的 err 覆盖外面的 err，所以错误的返回 nil ！
         } else {
           fmt.Println(y)
         }
         return
       }
     ```

2. 误用字符串

   - 当需要对一个字符串进行频繁的操作时，谨记在 go 语言中字符串是不可变的（类似 Java 和 C#）。
   - 使用诸如 `a += b`形式连接字符串效率低下，尤其在一个循环内部使用这种形式。这会导致大量的内存开销和拷贝。
   - 应该使用一个字符数组代替字符串，将字符串内容写入一个缓存中：
     ```
       var b bytes.Buffer
       ...
       for condition {
           b.WriteString(str) // 将字符串str写入缓存buffer
       }
       return b.String()
     ```
   - 由于编译优化和依赖于使用缓存操作的字符串大小，当循环次数大于 15 时，效率才会更佳。

3. 发生错误时使用 defer 关闭一个文件

   - 如果在一个 for 循环内部处理一系列文件，需要使用 defer 确保文件在处理完毕后被关闭，例如：
     ```
       for _, file := range files {
         if f, err = os.Open(file); err != nil {
             return
         }
         // 这是错误的方式，当循环结束时文件没有关闭
         defer f.Close()
         // 对文件进行操作
         f.Process(data)
       }
     ```
   - 但是在循环内结尾处的 defer 没有执行，所以文件一直没有关闭！垃圾回收机制可能会自动关闭文件，但是这会产生一个错误，更好的做法是：
     ```
       for _, file := range files {
         if f, err = os.Open(file); err != nil {
             return
         }
         // 对文件进行操作
         f.Process(data)
         // 关闭文件
         f.Close()
       }
     ```
   - defer 仅在函数返回时才会执行，在循环内的结尾或其他一些有限范围的代码内不会执行。

4. 何时使用 new()和 make()

   - 切片、映射和通道，使用 `make()`
   - 数组、结构体和所有的值类型，使用 `new()`

5. 不需要将一个指向切片的指针传递给函数

   - 切片实际是一个指向潜在数组的指针。
   - 常常需要把切片作为一个参数传递给函数是因为：实际就是传递一个指向变量的指针，在函数内可以改变这个变量，而不是传递数据的拷贝。
   - 因此应该这样做：`func findBiggest( listOfNumbers []int ) int {}`
   - **当切片作为参数传递时，切记不要解引用切片**。

6. 使用指针指向接口类型

   ```
     type nexter interface {
        next() byte
     }
     func nextFew1(n nexter, num int) []byte {
        var b []byte
        for i:=0; i < num; i++ {
            b[i] = n.next()
        }
        return b
     }
     func nextFew2(n *nexter, num int) []byte {
        var b []byte
        for i:=0; i < num; i++ {
            b[i] = n.next() // 编译错误：n.next 未定义（*nexter 类型没有 next 成员或 next 方法）
        }
        return b
     }
   ```

   - **永远不要使用一个指针指向一个接口类型，因为它已经是一个指针**。

7. 使用值类型时误用指针

   - 将一个值类型作为一个参数传递给函数或者作为一个方法的接收者，似乎是对内存的滥用，因为值类型一直是传递拷贝。
   - 但是另一方面，值类型的内存是在栈上分配，内存分配快速且开销不大。
   - 如果传递一个指针，而不是一个值类型，Go 编译器大多数情况下会认为需要创建一个对象，并将对象移动到堆上，所以会导致额外的内存分配：因此当使用指针代替值类型作为参数传递时，没有任何收获。

8. 使用协程和通道

   - 在实际应用中，不需要并发执行，或者不需要关注协程和通道的开销，在大多数情况下，通过栈传递参数会更有效率。
   - 但是，如果使用 break、return 或者 `panic()` 去跳出一个循环，很有可能会导致内存溢出，因为协程正处理某些事情而被阻塞。
   - 在实际代码中，通常仅需写一个简单的过程式循环即可。**当且仅当代码中并发执行非常重要，才使用协程和通道**。

9. 闭包和协程的使用

   - 常规闭包
     ```
       for ix := range values { // ix 是索引值
         func() {
            fmt.Print(ix, " ")
         }() // 调用闭包打印每个索引值
       }
       // 输出 0 1 2 3 4
     ```
   - 使用协程
     ```
        for ix := range values {
          go func() {
            fmt.Print(ix, " ")
          }()
        }
        // 输出 4 4 4 4 4
        // 协程师并发执行的，ix变量是一个单变量，这些闭包斗志绑定到一个变量。协程可能在循环结束后还没开始执行，所以ix为4
     ```
   - 使用协程，并给闭包传值
     ```
       for ix := range values {
         go func(ix interface{}) {
            fmt.Print(ix, " ")
         }(ix)
       }
       // 输出 1 0 3 4 2
       // 调用每个闭包时将 ix 作为参数传递给闭包。ix 在每次循环时都被重新赋值，并将每个协程的 ix 放置在栈中，所以当协程最终被执行时，每个索引值对协程都是可用的。
     ```
   - 在循环内初始化新变量
     ```
      for ix := range values {
        val := values[ix]
        go func() {
          fmt.Print(val, " ")
        }()
      }
      // 输出 10 11 12 13 14
      // 变量声明是在循环体内部，所以在每次循环时，这些变量相互之间是不共享的，所以这些变量可以单独的被每个闭包使用。
     ```

10. 糟糕的错误处理，而非创建额外布尔型变量

    - 直接判断 error 是否为 nil
      ```
        ... err1 := api.Func1()
        if err1 != nil { … }
      ```
    - 避免错误检测使代码变得混乱，解决此问题的好办法是尽可能以闭包的形式封装错误检测，例如：

      ```
        func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
          err := func () error {
              if req.Method != "GET" {
                  return errors.New("expected GET")
              }
              if input := parseInput(req); input != "command" {
                  return errors.New("malformed command")
              }
              // 可以在此进行其他的错误检测
          } ()

          if err != nil {
              w.WriteHeader(400)
              io.WriteString(w, err)
              return
          }
          doSomething()
          ...
        }
      ```

11. 关于 , ok 模式的情况

    - map 查找：`m := map[string]int{"one": 1};v, ok := m["one"]`
    - 类型断言：`str, ok := i.(string)`
    - 通道接收：`v, ok := <-ch`
    - range 循环与通道：`for v, ok range ch {}`
    - 自定义函数返回多个值：`v, ok := func1()`

12. 关于 defer 模式的情况
    - 资源释放：`defer file.Close()`
    - 解锁互斥锁：`defer mu.Unlock()`
    - 错误处理：`defer releaseResource(resource)`
    - 延迟执行函数：`defer fmt.Println("world")`
    - 修改返回值：
      ```
        defer func() {
          if r := recover(); r != nil {
            // 处理恐慌，例如返回一个错误
            fmt.Println("Recovered in calculateSum")
          }
        }()
      ```
    - 改变函数参数：
      ```
        func printSlice(s []int) {
          defer fmt.Println(s) // 延迟打印切片
          s = append(s, 42)    // 修改切片
        }
      ```
