## 读写数据

除了 fmt 和 os 包，还需要用到 bufio 包来处理缓冲的输入和输出。

1. 读取用户的输入

   - 从键盘和标准输入 os.Stdin 读取输入，最简单的办法是使用 fmt 包提供的 Scan... 和 Sscan... 开头的函数。
   - `Scanln()` 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
   - `Scanf()` 与其类似，除了 `Scanf()` 的第一个参数用作格式字符串，用来决定如何读取。
   - `Sscan...` 和以 `Sscan...` 开头的函数则是从字符串读取，除此之外，与 `Scanf()` 相同。
   - 如果这些函数读取到的结果与预想的不同，可以检查成功读入数据的个数和返回的错误。
   - 可以使用 bufio 包提供的缓冲读取器 (buffered reader) 来读取数据。
   - inputReader 是一个指向 `bufio.Reader` 的指针。`inputReader := bufio.NewReader(os.Stdin)` 这行代码，将会创建一个读取器，并将其与标准输入绑定。
   - `bufio.NewReader()` 构造函数的签名为：`func NewReader(rd io.Reader) *Reader`。
   - 该函数的实参可以是满足 `io.Reader` 接口的任意对象（任意包含有适当的 `Read()` 方法的对象），函数返回一个新的带缓冲的 `io.Reader` 对象，它将从指定读取器（例如 os.Stdin）读取内容。
   - 返回的读取器对象提供一个方法 `ReadString(delim byte)`，该方法从输入中读取内容，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区。
   - ReadString 返回读取到的字符串，如果碰到错误则返回 nil。如果它一直读到文件结束，则返回读取到的字符串和 `io.EOF`。如果读取过程中没有碰到 delim 字符，将返回错误 `err != nil`。
   - 屏幕是标准输出 `os.Stdout`；`os.Stderr` 用于显示错误信息，大多数情况下等同于 `os.Stdout`。

2. 文件读写
   1. 读文件
      - 在 Go 语言中，文件使用指向 `os.File` 类型的指针来表示的，也叫做文件句柄。
      - 标准输入 `os.Stdin` 和标准输出 `os.Stdout`，他们的类型都是 `*os.File`。
      - 其他类似函数：
        - 将整个文件的内容读到一个字符串里，可以使用 io/ioutil 包里的 `ioutil.ReadFile()` 方法，该方法第一个返回值的类型是 `[]byte`，里面存放读取到的内容，第二个返回值是错误，如果没有错误发生，第二个返回值为 nil。函数 `WriteFile()` 可以将 `[]byte` 的值写入文件。
        - 带缓冲的读取，在很多情况下，文件的内容是不按行划分的，或者干脆就是一个二进制文件。在这种情况下，`ReadString()` 就无法使用了，我们可以使用 `bufio.Reader` 的 `Read()`，它只接收一个参数：
          ```
             buf := make([]byte, 1024)
             ...
             n, err := inputReader.Read(buf)
             if (n == 0) { break}
          ```
          - 变量 n 的值表示读取到的字节数.
        - 按列读取文件中的数据，如果数据是按列排列并用空格分隔的，可以使用 fmt 包提供的以 `FScan...` 开头的一系列函数来读取他们。
   2. compress 包：读取压缩文件
      - compress 包提供了读取压缩文件的功能，支持的压缩文件格式为：bzip2、flate、gzip、lzw 和 zlib。
   3. 写文件
      - 除了文件句柄，还需要 bufio 的 Writer。以只写模式打开文件 `output.dat`，如果文件不存在则自动创建：`outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)`。
      - OpenFile 函数有三个参数：文件名、一个或多个标志（使用逻辑运算符 | 连接）、使用的文件权限。
      - 通常会用到以下标志：
        - os.O_RDONLY：只读
        - os.O_WRONLY：只写
        - os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
        - os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为 0 。
      - 在读文件的时候，文件的权限是被忽略的，所以在使用 `OpenFile()` 时传入的第三个参数可以用 0 。而在写文件时，不管是 Unix 还是 Windows，都需要使用 0666。
      - 然后，创建一个写入器（缓冲区）对象： `outputWriter := bufio.NewWriter(outputFile)`。
      - 将字符串写入缓冲区：`outputWriter.WriteString(outputString)`。
      - 缓冲区的内容紧接着被完全写入文件：`outputWriter.Flush()`。
      - 如果写入的东西很简单，可以使用 `fmt.Fprintf(outputFile, "Some test data.\n")` 直接将内容写入文件。fmt 包里的 `F...` 开头的 `Print()` 函数可以直接写入任何 `io.Writer`，包括文件
      - 使用 `os.Stdout.WriteString("hello, world\n")`，我们可以输出到屏幕。
3. 文件拷贝

   - 最简单的方式就是使用 io 包：`io.Copy(dst, src)`。
   - 注意 defer 的使用：当打开 dst 文件时发生了错误，那么 defer 仍然能够确保 `src.Close()` 执行。如果不这么做，src 文件会一直保持打开状态并占用资源。

4. 从命令行读取参数

   1. os 包
      - os 包中有一个 string 类型的切片变量 `os.Args`，用来处理一些基本的命令行参数，它在程序启动后读取命令行输入的参数。
      - 命令行参数会放置在切片 `os.Args[]` 中（以空格分隔），从索引 1 开始（os.Args[0] 放的是程序本身的名字，在本例中是 os_args）。
   2. flag 包
      - flag 包有一个扩展功能用来解析命令行选项。但是通常被用来替换基本常量，例如，在某些情况下我们希望在命令行给常量一些不一样的值。
      - 在 flag 包中有一个 Flag 是被定义成一个含有如下字段的结构体：
        ```
           type Flag struct {
              Name     string // name as it appears on command line
              Usage    string // help message
              Value    Value  // value as set
              DefValue string // default value (as text); for usage message
           }
        ```
      - `flag.Parse()` 扫描参数列表（或者常量列表）并设置 flag，`flag.Arg(i)` 表示第 i 个参数。`Parse()` 之后 `flag.Arg(i)` 全部可用，`flag.Arg(0)` 就是第一个真实的 flag，而不是像 `os.Args(0)` 放置程序的名字。
      - `flag.Narg()` 返回参数的数量。
      - `flag.PrintDefaults()` 打印 flag 的使用帮助信息。
      - `flag.VisitAll(fn func(*Flag))` 是另一个有用的功能：按照字典顺序遍历 flag，并且对每个标签调用 fn
      - 要给 flag 定义其它类型，可以使用 `flag.Int()`，`flag.Float64()`，`flag.String()`。

5. 用 buffer 读取文件
6. 用切片读写文件
   - 切片提供了 Go 中处理 I/O 缓冲的标准方式。
7. 用 defer 关闭文件
   - defer 关键字对于在函数结束时关闭打开的文件非常有用。
8. 使用接口的实际例子：fmt.Fprintf
   - `fmt.Fprintf()` 函数的实际签名：`func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)`。
   - 不是写入一个文件，而是写入一个 `io.Writer` 接口类型的变量，下面是 Writer 接口在 io 包中的定义：
     ```
        type Writer interface {
           Write(p []byte) (n int, err error)
        }
     ```
   - `fmt.Fprintf()` 依据指定的格式向第一个参数内写入字符串，第一个参数必须实现了 `io.Writer` 接口。
   - `Fprintf()` 能够写入任何类型，只要其实现了 Write 方法，包括 `os.Stdout`，文件（例如 `os.File`），管道，网络连接，通道等等。同样地，也可以使用 bufio 包中缓冲写入。bufio 包中定义了 `type Writer struct{...}` 。
   - `bufio.Writer` 实现了 `Write()` 方法：`func (b *Writer) Write(p []byte) (nn int, err error)`。
   - 它还有一个工厂函数：传给它一个 `io.Writer` 类型的参数，它会返回一个带缓冲的 `bufio.Writer` 类型的 `io.Writer` ：`func NewWriter(wr io.Writer) (b *Writer)`，适合任何形式的缓冲写入。
   - 在缓冲写入的最后千万不要忘了使用 `Flush()`，否则最后的输出不会被写入。
