## 错误处理与测试

- Go 没有像 Java 和 .NET 那样的 try/catch 异常机制：不能执行抛异常操作。但是有一套 defer-panic-and-recover 机制。
- 通过在函数和方法中返回错误对象作为它们的唯一或最后一个返回值——如果返回 nil，则没有错误发生——并且主调 (calling) 函数总是应该检查收到的错误。
- 永远不要忽略错误，否则可能会导致程序崩溃！！
- 处理错误并且在函数发生错误的地方给用户返回错误信息：照这样处理就算真的出了问题，程序也能继续运行并且通知给用户。`panic()` 和 `recover()` 是用来处理真正的异常（无法预测的错误）而不是普通的错误。
- 库函数通常必须返回某种错误提示给主调函数。
- Go 检查和报告错误条件的惯有方式：
  - 产生错误的函数会返回两个变量，一个值和一个错误码；如果后者是 nil 就是成功，非 nil 就是发生了错误。
  - 为了防止发生错误时正在执行的函数（如果有必要的话甚至会是整个程序）被中止，在调用函数后必须检查错误。

1. 错误处理

   - Go 有一个预先定义的 error 接口类型
     ```
       type error interface {
         Error() string
       }
     ```
   - 错误值用来表示异常状态；errors 包中有一个 errorString 结构体实现了 error 接口。当程序处于错误状态时可以用 `os.Exit(1)` 来中止运行。

   1. 定义错误

      - 任何时候当需要一个新的错误类型时，都可以用 errors 包（必须先 import）的 `errors.New()` 函数接收合适的错误信息来创建，像下面这样：`err := errors.New("math - square root of negative number")`。
      - 在大部分情况下自定义错误结构类型很有意义的，可以包含除了（低层级的）错误信息以外的其它有用信息。

        ```
          type PathError struct {
            Op string    // "open", "unlink", etc.
            Path string  // The associated file.
            Err error  // Returned by the system call.
          }

          func (e *PathError) Error() string {
            return e.Op + " " + e.Path + ": "+ e.Err.Error()
          }
        ```

      - 如果有不同错误条件可能发生，那么对实际的错误使用类型断言或类型判断（type-switch）是很有用的，并且可以根据错误场景做一些补救和恢复操作。
        ```
          switch err := err.(type) {
            case ParseError:
              PrintParseError(err)
            case PathError:
              PrintPathError(err)
            ...
            default:
              fmt.Printf("Not a special error, just %s\n", err)
          }
        ```
      - 用 json 包的情况。当 `json.Decode()` 在解析 JSON 文档发生语法错误时，指定返回一个 SyntaxError 类型的错误：

        ```
          type SyntaxError struct {
            msg    string // description of error
          // error occurred after reading Offset bytes, from which line and columnnr can be obtained
            Offset int64
          }

          func (e *SyntaxError) Error() string { return e.msg }
        ```

        - 在调用代码中可以像这样用类型断言测试错误是不是上面的类型：
          ```
            if serr, ok := err.(*json.SyntaxError); ok {
              line, col := findLine(f, serr.Offset)
              return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
            }
          ```

      - 包也可以用额外的方法 (methods)定义特定的错误，比如 `net.Error`：
        ```
          package net
          type Error interface {
            Timeout() bool   // Is the error a timeout?
            Temporary() bool // Is the error temporary?
          }
        ```
      - 遵循同一种命名规范：错误类型以 `...Error` 结尾，错误变量以 `err...` 或 `Err...` 开头或者直接叫 err 或 Err。
      - syscall 是低阶外部包，用来提供系统基本调用的原始接口。它们返回封装整数类型错误码的 `syscall.Errno`；类型 `syscall.Errno` 实现了 Error 接口。
      - 大部分 syscall 函数都返回一个结果和可能的错误，比如：
        ```
          r, err := syscall.Open(name, mode, perm)
          if err != nil {
            fmt.Println(err.Error())
          }
        ```
      - os 包也提供了一套像 `os.EINAL` 这样的标准错误，它们基于 syscall 错误：
        ```
          var (
            EPERM		Error = Errno(syscall.EPERM)
            ENOENT	Error = Errno(syscall.ENOENT)
            ESRCH		Error = Errno(syscall.ESRCH)
            EINTR		Error = Errno(syscall.EINTR)
            EIO			Error = Errno(syscall.EIO)
            ...
          )
        ```

   2. 用 fmt 创建错误对象
      - 通常想要返回包含错误参数的更有信息量的字符串，例如：可以用 `fmt.Errorf()` 来实现：它和 `fmt.Printf()` 完全一样，接收一个或多个格式占位符的格式化字符串和相应数量的占位变量。和打印信息不同的是它用信息生成错误对象。
      - 比如在平方根例子中使用：
        ```
          if f < 0 {
            return 0, fmt.Errorf("math: square root of negative number %g", f)
          }
        ```
      - 从命令行读取输入时，如果加了 --help 或 -h 标志，我们可以用有用的信息产生一个错误：
        ```
          if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
            err = fmt.Errorf("usage: %s infile.txt outfile.txt", filepath.Base(os.Args[0]))
            return
          }
        ```

2. 运行时异常和 panic

   - 当发生像数组下标越界或类型断言失败这样的运行错误时，Go 运行时会触发运行时 **panic**，伴随着程序的崩溃抛出一个 `runtime.Error` 接口类型的值。这个错误值有个 `RuntimeError()` 方法用于区别普通错误。
   - `panic()` 可以直接从代码初始化：当错误条件（我们所测试的代码）很严苛且不可恢复，程序不能继续运行时，可以使用 `panic()` 函数产生一个中止程序的运行时错误。`panic()` 接收一个做任意类型的参数，通常是字符串，在程序死亡时被打印出来。Go 运行时负责中止程序并给出调试信息。
   - 一个检查程序是否被已知用户启动的具体例子：

     ```
       var user = os.Getenv("USER")

       func check() {
         if user == "" {
           panic("Unknown user: no value for $USER")
         }
       }
     ```

   - 可以在导入包的 `init()` 函数中检查这些。
   - 当发生错误必须中止程序时，`panic()` 可以用于错误处理模式：
     ```
       if err != nil {
         panic("ERROR occurred:" + err.Error())
       }
     ```
   - Go panicking：
     - 在多层嵌套的函数调用中调用 `panic()`，可以马上中止当前函数的执行，所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者。这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 `panic()` 的值报告错误情况：这个终止过程就是 panicking。
     - 标准库中有许多包含 Must 前缀的函数，像 `regexp.MustComplie()` 和 `template.Must()`；当正则表达式或模板中转入的转换字符串导致错误时，这些函数会 `panic()`。
     - 不能随意地用 `panic()` 中止程序，必须尽力补救错误让程序能继续执行。

3. 从 panic 中恢复 (recover)

   - 正如名字一样，这个 (`recover()`) 内建函数被用于从 panic 或错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。
   - recover 只能在 defer 修饰的函数中使用：用于取得 `panic()` 调用中传
     递过来的错误值，如果是正常执行，调用 `recover()` 会返回 nil，且没有其它效果。
   - `panic()` 会导致栈被展开直到 defer 修饰的 `recover()` 被调用或者程序中止。
   - 这跟 Java 和 .NET 这样的语言中的 catch 块类似。
   - log 包实现了简单的日志功能：默认的 log 对象向标准错误输出中写入并打印每条日志信息的日期和时间。除了 Println 和 Printf 函数，其它的致命性函数都会在写完日志信息后调用 `os.Exit(1)`，那些退出函数也是如此。而 Panic 效果的函数会在写完日志信息后调用 `panic()`；可以在程序必须中止或发生了临界错误时使用它们，就像当 web 服务器不能启动时那样。
   - log 包用那些方法 (methods) 定义了一个 Logger 接口类型，如果你想自定义日志系统的话可以参考 http://golang.org/pkg/log/#Logger 。
   - `defer-panic()-recover()` 在某种意义上也是一种像 if，for 这样的控制流机制。
   - Go 标准库中许多地方都用了这个机制，例如，json 包中的解码和 regexp 包中的 `Complie()` 函数。Go 库的原则是即使在包的内部使用了 `panic()`，在它的对外接口 (API) 中也必须用 `recover()` 处理成显式返回的错误。

4. 自定义包中的错误处理和 panicking

   - 这是所有自定义包实现者应该遵守的最佳实践：
     - 在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 `panic()`
     - 向包的调用者返回错误值（而不是 panic）。
   - 在包内部，特别是在非导出函数中有很深层次的嵌套调用时，将 panic 转换成 error 来告诉调用方为何出错，是很实用的（且提高了代码可读性）。

5. 一种用闭包处理错误的模式
   - 结合 defer/panic/recover 机制和闭包可以得到一个更加优雅的模式。不过这个模式只有当所有的函数都是同一种签名时可用，这样就有相当大的限制。
   - 一个很好的使用它的例子是 web 应用，所有的处理函数都是下面这样：`func handler1(w http.ResponseWriter, r *http.Request) { ... }`。
   - 例如：
     - 假设所有的函数都有这样的签名：`func f(a type1, b type2)`。
     - 参数的数量和类型是不相关的。给这个类型一个名字：`fType1 = func f(a type1, b type2)`。
     - 使用了两个帮助函数：
       - `check()`：这是用来检查是否有错误和 panic 发生的函数：`func check(err error) { if err != nil { panic(err) } }`。
       - `errorhandler()`：这是一个包装函数。接收一个 fType1 类型的函数 fn 并返回一个调用 fn 的函数。里面就包含有 defer/recover 机制：
         ```
           func errorHandler(fn fType1) fType1 {
              return func(a type1, b type2) {
                defer func() {
                  if err, ok := recover().(error); ok {
                    log.Printf("run time panic: %v", err)
                  }
                }()
                fn(a, b)
              }
            }
         ```
       - 当错误发生时会 recover 并打印在日志中；除了简单的打印，应用也可以用 template 包为用户生成自定义的输出。
       - 通过这种机制，所有的错误都会被 recover，并且调用函数后的错误检查代码也被简化为调用 `check(err)` 即可。在这种模式下，不同的错误处理必须对应不同的函数类型；它们（错误处理）可能被隐藏在错误处理包内部。可选的更加通用的方式是用一个空接口类型的切片作为参数和返回值。
       - 
