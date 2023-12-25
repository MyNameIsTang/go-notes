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
      - 