## 函数(function)

函数是 Go 里面的基本代码块：Go 函数的功能非常强大，以至于被认为拥有函数式编程语言的多种特性。

1.  介绍

    - 每一个程序都包含很多的函数：函数是基本的代码块。
    - Go 是编译型语言，所以函数编写的顺序是无关紧要的；鉴于可读性的需求，最好把 main() 函数写在文件的前面，其他函数按照一定逻辑顺序进行编写（例如函数被调用的顺序）。
    - 编写多个函数的主要目的是将一个需要很多行代码的复杂问题分解为一系列简单的任务（那就是函数）来解决。而且，同一个任务（函数）可以被调用多次，有助于代码重用。
    - 当函数执行到代码块最后一行（} 之前）或者 return 语句的时候会退出，其中 return 语句可以带有零个或多个参数；这些参数将作为返回值供调用者使用。简单的 return 语句也可以用来结束 for 死循环，或者结束一个协程 (goroutine)。
    - Go 里面有三种类型的函数：
      - 普通的带有名字的函数
      - 匿名函数或者 lambda 函数
      - 方法（Methods）
    - 除了 main()、init() 函数外，其它所有类型的函数都可以有参数与返回值。函数参数、返回值以及它们的类型被统称为函数签名。
    - 函数可以将其他函数调用作为它的参数，只要这个被调用函数的返回值个数、返回值类型和返回值的顺序与调用函数所需求的实参是一致的，例如：`f1(f2(a, b))`。
    - 如果需要申明一个在外部定义的函数，你只需要给出函数名与函数签名，不需要给出函数体：`func flushICache(begin, end uintptr) // implemented externally`。
    - 函数也可以以申明的方式被使用，作为一个函数类型，就像：`type binOp func(int, int) int`。
    - 函数是一等值 (first-class value)：它们可以赋值给变量，就像 `add := binOp` 一样。
    - 函数值 (functions value) 之间可以相互比较：如果它们引用的是相同的函数或者都是 nil 的话，则认为它们是相同的函数。函数不能在其它函数里面声明（不能嵌套），不过我们可以通过使用匿名函数来破除这个限制。
    - 目前 Go 没有泛型 (generic) 的概念，也就是说它不支持那种支持多种类型的函数。不过在大部分情况下可以通过接口 (interface)，特别是空接口与类型选择（type switch）与/或者通过使用反射（reflection）来实现相似的功能。

2.  函数参数与返回值

    - 函数能够接收参数供自己使用，也可以返回零个或多个值（我们通常把返回多个值称为返回一组值）。相比与 C、C++、Java 和 C#，多值返回是 Go 的一大特性，为我们判断一个函数是否正常执行提供了方便。
    - 通过 return 关键字返回一组值。事实上，任何一个有返回值（单个或多个）的函数都必须以 return 或 panic 结尾。
    - 在函数块里面，return 之后的语句都不会执行。如果一个函数需要返回值，那么这个函数里面的每一个代码分支 (code-path) 都要有 return 语句。
    - 函数定义时，它的形参一般是有名字的，不过我们也可以定义没有形参名的函数，只有相应的形参类型，就像这样：`func f(int, int, float64)`。
    - 没有参数的函数通常被称为 niladic 函数 (niladic function)，就像 `main.main()`。

    1. 按值传递（call by value）按引用传递（call by reference）
       - Go 默认使用按值传递来传递参数，也就是传递参数的副本。函数接收参数副本之后，在使用变量的过程中可能对副本的值进行更改，但不会影响到原来的变量，比如 `Function(arg1)`。
       - 希望函数可以直接修改参数的值，而不是对参数的副本进行操作，需要将参数的地址（变量名前面添加 & 符号，比如 &variable）传递给函数，这就是按引用传递，比如 `Function(&arg1)`，此时传递给函数的是一个指针。
       - 如果传递给函数的是一个指针，指针的值（一个地址）会被复制，但指针的值所指向的地址上的值不会被复制；我们可以通过这个指针的值来修改这个值所指向的地址上的值。
       - 指针也是变量类型，有自己的地址和值，通常指针的值指向一个变量的地址。所以，按引用传递也是按值传递。
       - 几乎在任何情况下，传递指针（一个 32 位或者 64 位的值）的消耗都比传递副本来得少。
       - 在函数调用时，像切片 (slice)、字典 (map)、接口 (interface)、通道 (channel) 这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。
       - 有些函数只是完成一个任务，并没有返回值。仅仅是利用了这种函数的副作用 (side-effect)，就像输出文本到终端，发送一个邮件或者是记录一个错误等。但是绝大部分的函数还是带有返回值的。
       - 如果一个函数需要返回四到五个值，我们可以传递一个切片给函数（如果返回值具有相同类型）或者是传递一个结构体（如果返回值具有不同的类型）。因为传递一个指针允许直接修改变量的值，消耗也更少。
    2. 命名的返回（named return variables）
       - 当需要返回多个非命名返回值时，需要使用 () 把它们括起来，比如 `(int, int)`。
       - 命名返回值作为结果形参 (result parameters) 被初始化为相应类型的零值，当需要返回的时候，我们只需要一条简单的不带参数的 return 语句。
       - 需要注意的是，即使只有一个命名返回值，也需要使用 () 括起来
         ```
            func getX2AndX3_2(input int) (x2, x3 int) {
               x2 = input * 2
               x3 = input * 3
               return
            }
         ```
       - 即使函数使用了命名返回值，你依旧可以无视它而返回明确的值。
       - 任何一个非命名返回值（使用非命名返回值是很糟的编程习惯）在 return 语句里面都要明确指出包含返回值的变量或是一个可计算的值（就像上面警告所指出的那样）。
       - 尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。
    3. 空白符（blank identifier）
       - 空白符用来匹配一些不需要的值，然后丢弃掉：`a, _ = func()`
    4. 改变外部变量（outside variable）
       - 传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 return 返回。
         ```
            func Multiply(a, b int, reply *int) {
               *reply = a + b
            }
         ```
       - 然而，如果不小心使用的话，传递一个指针很容易引发一些不确定的事，所以，我们要十分小心那些可以改变外部变量的函数，在必要时，需要添加注释以便其他人能够更加清楚的知道函数里面到底发生了什么。

3.  传递变长参数

    - 如果函数的最后一个参数是采用 ...type 的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为 0，这样的函数称为变参函数。例如：`func myFunc(a, b, arg ...int) {}`。
    - 函数接受一个类似于切片 (slice) 的参数，该参数可以通过 for 循环结构迭代。
    - 如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 `slice...` 的形式来传递参数，调用变参函数。
    - 一个接受变长参数的函数可以将这个参数作为其它函数的参数进行传递：

      ```
         func F1(s ...string) {
            F2(s...)
            F3(s)
         }

         func F2(s ...string) { }
         func F3(s []string) { }
      ```

    - 变长参数可以作为对应类型的 slice 进行二次传递。
    - 但是如果变长参数的类型并不是都相同的呢？解决方法：

      - 使用结构

        ```
           type Options struct {
              par1 type1,
              par2 type2,
              ...
           }

        ```

        - TODO:函数 F1() 可以使用正常的参数 a 和 b，以及一个没有任何初始化的 Options 结构： F1(a, b, Options {})。如果需要对选项进行初始化，则可以使用 F1(a, b, Options {par1:val1, par2:val2})。

      - 使用空接口
        - 如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}，这样就可以接受任何类型的参数。
        - 不仅可以用于长度未知的参数，还可以用于任何不确定类型的参数。
        - 一般而言我们会使用一个 for-range 循环以及 switch 结构对每个参数的类型进行判断：
          ```
             func typecheck(..,..,values … interface{}) {
                for _, value := range values {
                   switch v := value.(type) {
                      case int: …
                      case float: …
                      case string: …
                      case bool: …
                      default: …
                   }
                }
             }
          ```

4.  defer 和追踪

    - 关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数。
    - 关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。
    - 使用 defer 的语句同样可以接受参数。
    - 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）。
    - 关键字 defer 允许我们进行一些函数执行完成后的收尾工作
      - 关闭文件流： `defer file.Close()`。
      - 解锁一个加锁的资源：`defer mu.Unlock()`。
      - 打印最终报告：`defer printFooter()`。
      - 关闭数据库链接：`defer disconnectFromDB()`。
    - 使用 defer 语句实现代码追踪

      ```
         func trace(s string) string {
            fmt.Println("entering:", s)
            return s
         }

         func un(s string) {
            fmt.Println("leaving:", s)
         }

         func a() {
            defer un(trace("a"))
            fmt.Println("in a")
         }

         func b() {
            defer un(trace("b"))
            fmt.Println("in b")
            a()
         }
      ```

    - 使用 defer 语句来记录函数的参数与返回值
      ```
         func func1(s string) (n int, err error) {
            defer func() {
               log.Printf("func1(%q) = %d, %v", s, n, err)
            }()
            return 7, io.EOF
         }
      ```

5.  内置函数

    - Go 语言拥有一些不需要进行导入操作就可以使用的内置函数。
    - 可以针对不同的类型进行操作，例如：len()、cap() 和 append()，或必须用于系统级的操作，例如：panic()。

      | 名称                       | 说明                                                                                                                                                                                                                                                                                                                                                                                                         |
      | -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
      | close()                    | 用于管道通信                                                                                                                                                                                                                                                                                                                                                                                                 |
      | len()、cap()               | len() 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap() 是容量的意思，用于返回某个类型的最大容量（只能用于数组、切片和管道，不能用于 map）                                                                                                                                                                                                                                               |
      | new()、make()              | new() 和 make() 均是用于分配内存：new() 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针。它也可以被用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，因此它比 new() 进行更多的工作。new() 是一个函数，不要忘记它的括号。 |
      | copy()、append()           | 用于复制和连接切片                                                                                                                                                                                                                                                                                                                                                                                           |
      | panic()、recover()         | 两者均用于错误处理机制                                                                                                                                                                                                                                                                                                                                                                                       |
      | print()、println()         | 底层打印函数，在部署环境中建议使用 fmt 包                                                                                                                                                                                                                                                                                                                                                                    |
      | complex()、real ()、imag() | 用于创建和操作复数                                                                                                                                                                                                                                                                                                                                                                                           |

6.  递归函数

    - 当一个函数在其函数体内调用自身，则称之为递归。最经典的例子便是计算斐波那契数列，即前两个数为 1，从第三个数开始每个数均为前两个数之和。
    - 许多问题都可以使用优雅的递归来解决，比如说著名的快速排序算法。
    - 在使用递归函数时经常会遇到的一个重要问题就是栈溢出：一般出现在大量的递归调用导致的程序栈内存分配耗尽。
    - 这个问题可以通过一个名为 懒惰求值 的技术解决，在 Go 语言中，我们可以使用管道 (channel) 和 goroutine 来实现。
    - Go 语言中也可以使用相互调用的递归函数：多个函数之间相互调用形成闭环。因为 Go 语言编译器的特殊性，这些函数的声明顺序可以是任意的。

7.  将函数作为参数

    - 函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调。
    - 将函数作为参数的最好的例子是函数 `strings.IndexFunc()`，该函数的签名是 `func IndexFunc(s string, f func(c rune) bool) int`，它的返回值是字符串 s 中第一个使函数 f(c) 返回 true 的 Unicode 字符的索引值。如果找不到，则返回 -1。

8.  闭包

    - 不希望给函数起名字的时候，可以使用匿名函数，例如：`func(x, y int) int { return x + y }`。
    - 这样的一个函数不能够独立存在（编译器会返回错误：non-declaration statement outside function body），但可以被赋值于某个变量，即保存函数的地址到变量中：`fplus := func(x, y int) int { return x + y }`，然后通过变量名对函数进行调用：`fplus(3,4)`。
    - 也可以直接对匿名函数进行调用：`func(x, y int) int { return x + y } (3, 4)`。
    - 表示参数列表的第一对括号必须紧挨着关键字 func，因为匿名函数没有名称。花括号 {} 涵盖着函数体，最后的一对括号表示对该匿名函数的调用。
    - 实际上拥有的是一个函数值：匿名函数可以被赋值给变量并作为值使用。
    - 匿名函数像所有函数一样可以接受或不接受参数。
      ```
         func (u string) {
            fmt.Println(u)
            …
         }(v)
      ```
    - defer 语句和匿名函数，关键字 defer 经常配合匿名函数使用，它可以用于改变函数的命名返回值。
    - 匿名函数还可以配合 go 关键字来作为 goroutine 使用
    - 匿名函数同样被称之为闭包（函数式语言的术语）：它们被允许调用定义在其它环境下的变量。
      - 闭包可使得某个函数捕捉到一些外部状态，例如：函数被创建时的状态。
      - 另一种表示方式为：一个闭包继承了函数所声明时的作用域。
      - 这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁。
      - 闭包经常被用作包装函数：它们会预先定义好 1 个或多个参数以用于包装。
      - 另一个不错的应用就是使用闭包来完成更加简洁的错误检查

9.  应用闭包：将函数作为返回值

    ```
       func Add2() (func(b int) int)
       func Adder(a int) (func(b int) int)
    ```

    - 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
    - 这些局部变量同样可以是参数。
    - 在闭包中使用到的变量可以是在闭包函数体内声明的，也可以是在外部函数声明的
      ```
         var g int
         go func(i int) {
            s := 0
            for j := 0; j < i; j++ { s += j }
            g = s
         }(1000)
      ```
    - 这样闭包函数就能够被应用到整个集合的元素上，并修改它们的值。然后这些变量就可以用于表示或计算全局或平均值。
      ```
         func fi() func() int {
            pre1, pre2 := 0, 1
            return func() int {
               pre1, pre2 = pre2, pre1+pre2
               return pre1
            }
         }
      ```
    - 一个返回值为另一个函数的函数可以被称之为工厂函数，这在您需要创建一系列相似的函数的时候非常有用。
      ```
         func MakeAddSuffix(suffix string) func(string) string {
            return func(name string) string {
               if !strings.HasSuffix(name, suffix) {
                  return name + suffix
               }
               return name
            }
         }
      ```
    - 可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数，是函数式语言的特点。
    - 函数也是一种值，因此很显然 Go 语言具有一些函数式语言的特性。
    - 闭包在 Go 语言中非常常见，常用于 goroutine 和管道操作。

10. 使用闭包调试

    - 在分析和调试复杂的程序时，无数个函数在不同的代码文件中相互调用，如果这时候能够准确地知道哪个文件中的具体哪个函数正在执行，对于调试是十分有帮助的。
    - 可以使用 runtime 或 log 包中的特殊函数来实现这样的功能。
    - 包 runtime 中的函数 Caller() 提供了相应的信息，因此可以在需要的时候实现一个 where() 闭包函数来打印函数执行的位置：
      ```
         where := func() {
            _, file, line, _ := runtime.Caller(1)
            log.Printf("%s:%d", file, line)
         }
         where()
         // some code
         where()
         // some more code
         where()
      ```
    - 可以设置 log 包中的 flag 参数来实现：
      ```
         log.SetFlags(log.Llongfile)
         log.Print("")
      ```
    - 或使用一个更加简短版本的 where() 函数：
      ```
         var where = log.Print
      ```