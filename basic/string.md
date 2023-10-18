## 字符串

- 字符串是 UTF-8 字符的一个序列（当字符为 ASCII 码时则占用 1 个字节，其它字符根据需要占用 2-4 个字节）。
- UTF-8 是被广泛使用的编码格式，是文本文件的标准编码，其它包括 XML 和 JSON 在内，也都使用该编码。
  - 由于该编码对占用字节长度的不定性，Go 中的字符串里面的字符也可能根据需要占用 1 至 4 个字节，这与其它语言如 C++、Java 或者 Python 不同（Java 始终使用 2 个字节）。
  - Go 这样做的好处是不仅减少了内存和硬盘空间占用，同时也不用像其它语言那样需要对使用 UTF-8 字符集的文本进行编码和解码。
- 字符串是一种值类型，且值不可变，即创建某个文本后你无法再次修改这个文本的内容；更深入地讲，字符串是字节的定长数组。
- Go 支持以下 2 种形式的字面值：
  - 解释字符串，该类字符串使用双引号括起来，其中的相关的转义字符将被替换，这些转义字符包括：
    - \n：换行符
    - \r：回车符
    - \t：tab 键
    - \u 或 \U：Unicode 字符
    - \\：反斜杠自身
  - 非解释字符串，该类字符串使用反引号括起来，支持换行，例如 \`This is a raw string \n\`中的`\n\` 会被原样输出。
- 和 C/C++不一样，Go 中的字符串是根据长度限定，而非特殊字符 \0。
- string 类型的零值为长度为零的字符串，即空字符串 ""。
- 一般的比较运算符（==、!=、<、<=、>=、>）通过在内存中按字节比较来实现字符串的对比。
- 通过函数 len() 来获取字符串所占的字节长度，例如：`len(str)`。
- 字符串的内容（纯字节）可以通过标准索引法来获取，在中括号 [] 内写入索引，索引从 0 开始计数：
  - 字符串 str 的第 1 个字节：`str[0]`。
  - 第 i 个字节：`str[i - 1]`。
  - 最后 1 个字节：`str[len(str)-1]`。
  - 这种转换方案只对纯 ASCII 码的字符串有效
- 获取字符串中某个字节的地址的行为是非法的，例如：`&str[i]`。
- 字符串拼接符 +。
  - 通过 `s := s1 + s2` 拼接在一起。
  - 多行的字符串进行拼接： `str := "Beginning of the string " +
"second part of the string"`。
  - 拼接的简写形式 += ：`s += "world!"`。
- 在循环中使用加号 + 拼接字符串并不是最高效的做法，更好的办法是使用函数 strings.Join()，使用字节缓冲（bytes.Buffer）拼接更加给力。

  ```
    buffer := bytes.NewBufferString("")

    for i := 0; i < 5; i++ {
      buffer.WriteString(fmt.Sprint(i))
    }

    r := buffer.String()
    fmt.Print(r)
  ```

## strings 和 strconv 包

作为一种基本数据结构，每种语言都有一些对于字符串的预定义处理函数。Go 中使用 strings 包来完成对字符串的主要操作。

1. 前缀和后缀

   - HasPrefix() 判断字符串 s 是否以 prefix 开头：`strings.HasPrefix(s, prefix string) bool`。
   - HasSuffix() 判断字符串 s 是否以 suffix 结尾：`strings.HasSuffix(s, suffix string) bool`。

2. 字符串包含关系

   - Contains() 判断字符串 s 是否包含 substr：`strings.Contains(s, substr string) bool`。

3. 判断子字符串或字符在父字符串中出现的位置（索引）

   - Index() 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：`strings.Index(s, str string) int`。
   - LastIndex() 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：`strings.LastIndex(s, str string) int`。
   - 如果需要查询非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位：`strings.IndexRune(s string, r rune) int`。

4. 字符串替换

   - Replace() 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new：`strings.Replace(str, old, new string, n int) string`。

5. 统计字符串出现次数

   - Count() 用于计算字符串 str 在字符串 s 中出现的非重叠次数：`strings.Count(s, str string) int`。

6. 重复字符串

   - Repeat() 用于重复 count 次字符串 s 并返回一个新的字符串：`strings.Repeat(s, count int) string`。

7. 修改字符串大小写

   - ToLower() 将字符串中的 Unicode 字符全部转换为相应的小写字符：`strings.ToLower(s) string`。
   - ToUpper() 将字符串中的 Unicode 字符全部转换为相应的大写字符：`strings.ToUpper(s) string`。

8. 修剪字符串

   - 使用 strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号。
   - 如果你想要剔除指定字符，则可以使用 strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉。
   - 只想剔除开头或者结尾的字符串，则可以使用 TrimLeft() 或者 TrimRight() 来实现。

9. 分割字符串

   - strings.Fields(s) 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。
   - strings.Split(s, sep) 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。

10. 拼接 slice 到字符串

    - strings.Join() 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串：`strings.Join(sl []string, sep string) string`。

11. 从字符串中读取内容

    - strings.NewReader(str) 用于生成一个 Reader 并读取字符串中的内容，然后返回指向该 Reader 的指针，从其它类型读取内容的函数还有：
      - Read() 从 []byte 中读取内容。
      - ReadByte() 和 ReadRune() 从字符串中读取下一个 byte 或者 rune。

12. 字符串与其他类型的转换

    - 与字符串相关的类型转换都是通过 strconv 包实现的。
    - 该包包含了一些变量用于获取程序运行的操作系统平台下 int 类型所占的位数，如：strconv.IntSize。
    - 任何类型 T 转换为字符串总是成功的。
    - 从数字类型转换到字符串：
      - strconv.Itoa(i int) string 返回数字 i 所表示的字符串类型的十进制数。
      - strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串，其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64。
    - 将字符串转换为其它类型 tp 并不总是可能的，可能会在运行时抛出错误 parsing "…": invalid argument。
    - 从字符串类型转换为数字类型：
      - strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型。
      - strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型。
