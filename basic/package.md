## 包（package）

1. 标准库

   - 像 fmt、os 等这样具有常用功能的内置包在 Go 语言中有 150 个以上，它们被称为标准库，大部分(一些底层的除外)内置于 Go 本身。完整列表可以在 https://gowalker.org/search?q=gorepos 查看。
   - 常用标准库：

     - unsafe: 包含了一些打破 Go 语言“类型安全”的命令，一般的程序中不会被使用，可用在 C/C++ 程序的调用中。
     - syscall-os-os/exec:

       - os: 提供给我们一个平台无关性的操作系统功能接口，采用类 Unix 设计，隐藏了不同操作系统间的差异，让不同的文件系统和操作系统对象表现一致。
       - os/exec: 提供我们运行外部操作系统命令和程序的方式。
       - syscall: 底层的外部包，提供了操作系统底层调用的基本接口。
       - 让 Linux 重启：

         ```
           const LINUX_REBOOT_MAGIC1 uintptr = 0xfee1dead
           const LINUX_REBOOT_MAGIC2 uintptr = 672274793
           const LINUX_REBOOT_CMD_RESTART uintptr = 0x1234567

           func main() {
             syscall.Syscall(syscall.SYS_REBOOT,
               LINUX_REBOOT_MAGIC1,
               LINUX_REBOOT_MAGIC2,
               LINUX_REBOOT_CMD_RESTART)
           }
         ```

     - archive/tar 和 /zip-compress：压缩（解压缩）文件功能。
     - fmt-io-bufio-path/filepath-flag：
       - fmt: 提供了格式化输入输出功能。
       - io: 提供了基本输入输出功能，大多数是围绕系统功能的封装。
       - bufio: 缓冲输入输出功能的封装。
       - path/filepath: 用来操作在当前系统中的目标文件名路径。
       - flag: 对命令行参数的操作。
     - strings-strconv-unicode-regexp-bytes:
       - strings: 提供对字符串的操作。
       - strconv: 提供将字符串转换为基础类型的功能。
       - unicode: 为 unicode 型的字符串提供特殊的功能。
       - regexp: 正则表达式功能。
       - bytes: 提供对字符型分片的操作。
       - index/suffixarray: 子字符串快速查询。
     - math-math/cmath-math/big-math/rand-sort:
       - math: 基本的数学函数。
       - math/cmath: 对复数的操作。
       - math/rand: 伪随机数生成。
       - sort: 为数组排序和自定义集合。
       - math/big: 大数的实现和计算。
     - container-/list-ring-heap: 实现对集合的操作。
       - list: 双链表。
       - ring: 环形链表。
       - 如何遍历一个链表(当 l 是 \*List)：
         ```
           for e := l.Front(); e != nil; e = e.Next() {
             //do something with e.Value
           }
         ```
     - time-log:
       - time: 日期和时间的基本操作。
       - log: 记录程序运行时产生的日志。
     - encoding/json-encoding/xml-text/template:
       - encoding/json: 读取并解码和写入并编码 JSON 数据。
       - encoding/xml: 简单的 XML1.0 解析器。
       - text/template:生成像 HTML 一样的数据与文本混合的数据驱动模板。
     - net-net/http-html:
       - net: 网络数据的基本操作。
       - http: 提供了一个可扩展的 HTTP 服务器和基础客户端，解析 HTTP 请求和回复。
       - html: HTML5 解析器。
     - runtime: Go 程序运行时的交互操作，例如垃圾回收和协程创建。
     - reflect: 实现通过程序运行时反射，让程序操作任意类型的变量。

   - exp 包中有许多将被编译为新包的实验性的包。在下次稳定版本发布的时候，它们将成为独立的包。如果前一个版本已经存在了，它们将被作为过时的包被回收。然而 Go1.0 发布的时候并没有包含过时或者实验性的包。

2. regexp 包

   - 在字符串中对正则表达式模式 (pattern) 进行匹配。简单模式，使用 `Match()` 方法便可：`ok, _ := regexp.Match(pat, []byte(searchIn))`。也可以使用 MatchString()：`ok, _ := regexp.MatchString(pat, searchIn)`。
   - 更多方法中，必须先将正则模式通过 `Compile()` 方法返回一个 Regexp 对象。然后我们将掌握一些匹配，查找，替换相关的功能。
   - `Compile()` 函数也可能返回一个错误，我们在使用时忽略对错误的判断是因为我们确信自己正则表达式是有效的。当用户输入或从数据中获取正则表达式的时候，我们有必要去检验它的正确性。另外我们也可以使用 `MustCompile()` 方法，它可以像 `Compile()` 方法一样检验正则的有效性，但是当正则不合法时程序将 `panic()`。

3. 锁和 sync 包

   - 在一些复杂的程序中，通常通过不同线程执行不同应用来实现程序的并发。当不同线程要使用同一个变量时，经常会出现一个问题：无法预知变量被不同线程修改的顺序！（这通常被称为资源竞争，指不同线程对同一变量使用的竞争）
   - 经典的做法是一次只能让一个线程对共享变量进行操作。当变量被一个线程改变时（临界区），我们为它上锁，直到这个线程执行完成并解锁后，其他线程才能访问它。
   - map 类型是不存在锁的机制来实现这种效果（出于对性能的考虑），所以 map 类型是非线程安全的。当并行访问一个共享的 map 类型的数据，map 数据将会出错。
   - 在 Go 语言中这种锁的机制是通过 sync 包中 Mutex 来实现的。sync 来源于 "synchronized" 一词，这意味着线程将有序的对同一变量进行访问。
   - sync.Mutex 是一个互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区。
   - 假设 info 是一个需要上锁的放在共享内存中的变量。通过包含 Mutex 来实现的一个典型例子如下：

     ```
       import  "sync"

       type Info struct {
         mu sync.Mutex
         // ... other fields, e.g.: Str string
       }

        func Update(info *Info) {
          info.mu.Lock()
          // critical section:
          info.Str = // new value
          // end critical section
          info.mu.Unlock()
        }

     ```

   - 还有一个很有用的例子是通过 Mutex 来实现一个可以上锁的共享缓冲器:
     ```
       type SyncedBuffer struct {
         lock 	sync.Mutex
         buffer  bytes.Buffer
       }
     ```
   - 在 sync 包中还有一个 RWMutex 锁：它能通过 `RLock()` 来允许同一时间多个线程对变量进行读操作，但是只能一个线程进行写操作。如果使用 `Lock()` 将和普通的 Mutex 作用相同。包中还有一个方便的 Once 类型变量的方法 `once.Do(call)`，这个方法确保被调用函数只能被调用一次。
   - 相对简单的情况下，通过使用 sync 包可以解决同一时间只能一个线程访问变量或 map 类型数据的问题。如果这种方式导致程序明显变慢或者引起其他问题，我们要重新思考来通过 goroutines 和 channels 来解决问题，这是在 Go 语言中所提倡用来实现并发的技术。

4. 精密计算和 big 包

   - 使用 Go 语言中的 float64 类型进行浮点运算，返回结果将精确到 15 位，足以满足大多数的任务。
   - 当对超出 int64 或者 uint64 类型这样的大数进行计算时，如果对精度没有要求，float32 或者 float64 可以胜任，但如果对精度有严格要求的时候，我们不能使用浮点数，在内存中它们只能被近似的表示。
   - 对于整数的高精度计算 Go 语言中提供了 big 包，被包含在 math 包下：有用来表示大整数的 `big.Int` 和表示大有理数的 `big.Rat` 类型（可以表示为 2/5 或 3.1416 这样的分数，而不是无理数或 π）。这些类型可以实现任意位类型的数字，只要内存足够大。缺点是更大的内存和处理开销使它们使用起来要比内置的数字类型慢很多。
   - 大的整型数字是通过 `big.NewInt(n)` 来构造的，其中 n 为 int64 类型整数。而大有理数是通过 `big.NewRat(n, d)` 方法构造。n（分子）和 d（分母）都是 int64 型整数。
   - 因为 Go 语言不支持运算符重载，所以所有大数字类型都有像是 `Add()` 和 `Mul()` 这样的方法。它们作用于作为 receiver 的整数和有理数，大多数情况下它们修改 receiver 并以 receiver 作为返回结果。因为没有必要创建 big.Int 类型的临时变量来存放中间结果，所以运算可以被链式地调用，并节省内存。
   -

5. 自定义包和可见性
   - 包是 Go 语言中代码组织和代码编译的主要方式。
   - 当写自己包的时候，要使用短小的不含有 \_（下划线）的小写单词来为文件命名。
   - 主程序利用的包必须在主程序编写之前被编译。因此，按照惯例，子目录和包之间有着密切的联系：为了区分，不同包存放在不同的目录下，每个包（所有属于这个包中的 go 文件）都存放在和包名相同的子目录下。
   - Import with . : `import . "./pack1"`。当使用 . 作为包的别名时，可以不通过包名来使用其中的项目。例如：`test := ReturnStr()`。
   - Import with _ : `import _ "./pack1/pack1"`。pack1 包只导入其副作用，也就是说，只执行它的 `init()` 函数并初始化其中的全局变量。
   - 导入外部安装包:
     - 如果要在应用中使用一个或多个外部包，首先必须使用 `go install` 在本地机器上安装它们。
       - 使用 http://codesite.ext/author/goExample/goex 这种托管在 Google Code、GitHub 和 Launchpad 等代码网站上的包：`go install codesite.ext/author/goExample/goex`。
       - 将一个名为 codesite.ext/author/goExample/goex 的 map 安装在 `$GOROOT/src/` 目录下。
       - 通过以下方式，一次性安装，并导入到你的代码中：`import goex "codesite.ext/author/goExample/goex"`。因此该包的 URL 将用作导入路径。
     - 在 http://golang.org/cmd/goinstall/ 的 `go install` 文档中列出了一些广泛被使用的托管在网络代码仓库的包的导入路径。
   - 包的初始化:
     - 程序的执行开始于导入包，初始化 main 包然后调用 `main()` 函数。
     - 一个没有导入的包将通过分配初始值给所有的包级变量和调用源码中定义的包级 `init()` 函数来初始化。一个包可能有多个 `init()` 函数甚至在一个源码文件中。它们的执行是无序的。这是最好的例子来测定包的值是否只依赖于相同包下的其他值或者函数。
     - `init()` 函数是不能被调用的。
     - 导入的包在包自身初始化前被初始化，而一个包在程序执行中只能初始化一次。
