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

9. JSON 数据格式

   - 数据结构要在网络中传输或保存到文件，就必须对其编码和解码；
   - 目前存在很多编码格式：JSON，XML，gob，Google 缓冲协议等等。Go 语言支持所有这些编码格式；
   - 结构可能包含二进制数据，如果将其作为文本打印，那么可读性是很差的。另外结构内部可能包含匿名字段，而不清楚数据的用意。
   - 通过把数据转换成纯文本，使用命名的字段来标注，让其具有可读性。这样的数据格式可以通过网络传输，而且是与平台无关的，任何类型的应用都能够读取和输出，不与操作系统和编程语言的类型相关。
   - 术语说明：
     - 数据结构 --> 指定格式 = 序列化 或 编码（传输之前）
     - 指定格式 --> 数据结构 = 反序列化 或 解码（传输之后）
   - 序列化是在内存中把数据转换成指定格式（数据 -> 字符串），反之亦然（字符串 -> 数据）。
   - 编码也是一样的，只是输出一个数据流（实现了 `io.Writer` 接口）；解码是从一个数据流（实现了 `io.Reader`）输出到一个数据结构。
   - JSON（JavaScript Object Notation，参阅 http://json.org）被作为首选，主要是由于其格式上非常简洁。通常 JSON 被用于 web 后端和浏览器之间的通讯，但是在其它场景也同样的有用。
   - JSON 更加简洁、轻量（占用更少的内存、磁盘及网络带宽）和更好的可读性。
   - Go 语言的 json 包可以在程序中方便的读取和写入 JSON 数据。
   - `json.Marshal()` 的函数签名是： `func Marshal(v interface{}) ([]byte, error)`。
   - 出于安全考虑，在 web 应用中最好使用 `json.MarshalforHTML()` 函数，其对数据执行 HTML 转码，所以文本可以被安全地嵌在 `HTML <script>` 标签中。
   - `json.NewEncoder()` 的函数签名是： `func NewEncoder(w io.Writer) *Encoder`，返回的 Encoder 类型的指针可调用方法 `Encode(v interface{})`，将数据对象 v 的 json 编码写入 `io.Writer w` 中。
   - JSON 与 Go 类型对应如下：
     - bool 对应 JSON 的 boolean
     - float64 对应 JSON 的 number
     - string 对应 JSON 的 string
     - nil 对应 JSON 的 null
   - 不是所有的数据都可以编码为 JSON 类型，只有验证通过的数据结构才能被编码：
     - JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 `map[string]T`（T 是 json 包中支持的任何类型）。
     - Channel，复杂类型和函数类型不能被编码。
     - 不支持循环数据结构；它将引起序列化进入一个无限循环
     - 指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）
   - 反序列化：
     - `json.Unmarshal()` 的函数签名是 `func Unmarshal(data []byte, v interface{}) error` 把 JSON 解码为数据结构。
     - 虽然反射能够让 JSON 字段去尝试匹配目标结构字段；但是只有真正匹配上的字段才会填充数据。字段没有匹配不会报错，而是直接忽略掉。
   - 解码任意的数据：
     - json 包使用 `map[string]interface{}` 和 `[]interface{}` 储存任意的 JSON 对象和数组；其可以被反序列化为任何的 JSON blob 存储到接口值中。
     - 可以通过 for range 语法和 type switch 来访问其实际类型，通过这种方式，可以处理未知的 JSON 数据，同时可以确保类型安全。
   - 解码数据到结构

     - 如果事先知道 JSON 数据，可以定义一个适当的结构并对 JSON 数据反序列化。

       ```
          type FamilyMember struct {
             Name    string
             Age     int
             Parents []string
          }
          var m FamilyMember
          err := json.Unmarshal(b, &m)
       ```

     - 程序实际上是分配了一个新的切片。这是一个典型的反序列化引用类型（指针、切片和 map）的例子。

   - 编码和解码流
     - json 包提供 Decoder 和 Encoder 类型来支持常用 JSON 数据流读写。`NewDecoder()` 和 `NewEncoder()` 函数分别封装了 `io.Reader` 和 `io.Writer` 接口。
       ```
          func NewDecoder(r io.Reader) *Decoder
          func NewEncoder(w io.Writer) *Encoder
       ```
     - 把 JSON 直接写入文件，可以使用 `json.NewEncoder` 初始化文件（或者任何实现 `io.Writer` 的类型），并调用 `Encode()`；反过来与其对应的是使用 `json.NewDecoder` 和 `Decode()` 函数：
       ```
          func NewDecoder(r io.Reader) *Decoder
          func (dec *Decoder) Decode(v interface{}) error
       ```
     - 接口是如何对实现进行抽象的：数据结构可以是任何类型，只要其实现了某种接口，目标或源数据要能够被编码就必须实现 io.Writer 或 io.Reader 接口。由于 Go 语言中到处都实现了 Reader 和 Writer，因此 Encoder 和 Decoder 可被应用的场景非常广泛，例如读取或写入 HTTP 连接、websockets 或文件。

10. XML 数据格式

    - JSON 例子等价的 XML 版本：

      ```
         <Person>
            <FirstName>Laura</FirstName>
            <LastName>Lynn</LastName>
         </Person>
      ```

    - 如同 json 包一样，也有 `xml.Marshal()` 和 `xml.Unmarshal()` 从 XML 中编码和解码数据；但这个更通用，可以从文件中读取和写入（或者任何实现了 `io.Reader` 和 `io.Writer` 接口的类型）。
    - 和 JSON 的方式一样，XML 数据可以序列化为结构，或者从结构反序列化为 XML 数据。
    - `encoding/xml` 包实现了一个简单的 XML 解析器（SAX），用来解析 XML 数据内容。
    - 包中定义了若干 XML 标签类型：StartElement，Chardata（这是从开始标签到结束标签之间的实际文本），EndElement，Comment，Directive 或 ProcInst。
    - 包中同样定义了一个结构解析器：`NewParser()` 方法持有一个 `io.Reader`（这里具体类型是 `strings.NewReader`）并生成一个解析器类型的对象。还有一个 `Token()` 方法返回输入流里的下一个 XML token。在输入流的结尾处，会返回 (nil,io.EOF)
    - XML 文本被循环处理直到 `Token()` 返回一个错误，因为已经到达文件尾部，再没有内容可供处理了。通过一个 type-switch 可以根据一些 XML 标签进一步处理。Chardata 中的内容只是一个 `[]byte`，通过字符串转换让其变得可读性强一些。

11. 用 Gob 传输数据

    - Gob 是 Go 自己的以二进制形式序列化和反序列化程序数据的格式；可以在 encoding 包中找到。这种格式的数据简称为 Gob （即 Go binary 的缩写）。类似于 Python 的 "pickle" 和 Java 的 "Serialization"。
    - Gob 通常用于远程方法调用（RPCs）参数和结果的传输，以及应用程序和机器之间的数据传输。
    - Gob 特定地用于纯 Go 的环境中，例如，两个用 Go 写的服务之间的通信。这样的话服务可以被实现得更加高效和优化。
    - Gob 不是可外部定义，语言无关的编码方式。因此它的首选格式是二进制，而不是像 JSON 和 XML 那样的文本格式。
    - Gob 并不是一种不同于 Go 的语言，而是在编码和解码过程中用到了 Go 的反射。
    - Gob 文件或流是完全自描述的：里面包含的所有类型都有一个对应的描述，并且总是可以用 Go 解码，而不需要了解文件的内容。
    - 只有可导出的字段会被编码，零值会被忽略。在解码结构体的时候，只有同时匹配名称和可兼容类型的字段才会被解码。
    - 当源数据类型增加新字段后，Gob 解码客户端仍然可以以这种方式正常工作：解码客户端会继续识别以前存在的字段。并且还提供了很大的灵活性，比如在发送者看来，整数被编码成没有固定长度的可变长度，而忽略具体的 Go 类型。
    - 和 JSON 的使用方式一样，Gob 使用通用的 `io.Writer` 接口，通过 `NewEncoder()` 函数创建 Encoder 对象并调用 `Encode()`；相反的过程使用通用的 `io.Reader` 接口，通过 `NewDecoder()` 函数创建 Decoder 对象并调用 `Decode()`。

12. Go 中的密码学
    - 通过网络传输的数据必须加密，以防止被 hacker（黑客）读取或篡改，并且保证发出的数据和收到的数据检验和一致。 鉴于 Go 母公司的业务，我们毫不惊讶地看到 Go 的标准库为该领域提供了超过 30 个的包：
      - hash 包：实现了 adler32、crc32、crc64 和 fnv 校验；
      - crypto 包：实现了其它的 hash 算法，比如 md4、md5、sha1 等。以及完整地实现了 aes、blowfish、rc4、rsa、xtea 等加密算法。
