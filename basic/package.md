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
       - 通过以下方式，一次性安装，并导入到代码中：`import goex "codesite.ext/author/goExample/goex"`。因此该包的 URL 将用作导入路径。
     - 在 http://golang.org/cmd/goinstall/ 的 `go install` 文档中列出了一些广泛被使用的托管在网络代码仓库的包的导入路径。
   - 包的初始化:
     - 程序的执行开始于导入包，初始化 main 包然后调用 `main()` 函数。
     - 一个没有导入的包将通过分配初始值给所有的包级变量和调用源码中定义的包级 `init()` 函数来初始化。一个包可能有多个 `init()` 函数甚至在一个源码文件中。它们的执行是无序的。这是最好的例子来测定包的值是否只依赖于相同包下的其他值或者函数。
     - `init()` 函数是不能被调用的。
     - 导入的包在包自身初始化前被初始化，而一个包在程序执行中只能初始化一次。

6. 为自定义包使用 godoc

   - godoc 工具在显示自定义包中的注释也有很好的效果：注释必须以 // 开始并无空行放在声明（包，类型，函数）前。godoc 会为每个文件生成一系列的网页。
   - 例如：
     - 在 doc_examples 目录下我们有用来排序的 go 文件，文件中有一些注释（文件需要未编译）。
     - 命令行下进入目录下并输入命令：`godoc -http=:6060 -goroot="."`。
     - . 是指当前目录，-goroot 参数可以是 /path/to/my/package1 这样的形式指出 package1 在源码中的位置或接受用冒号形式分隔的路径，无根目录的路径为相对于当前目录的相对路径
     - 在浏览器打开地址：http://localhost:6060
   - 如果在一个团队中工作，并且源代码树被存储在网络硬盘上，就可以使用 godoc 给所有团队成员连续文档的支持。通过设置 sync_minutes=n，甚至可以让它每 n 分钟自动更新您的文档！

7. 使用 `go install` 安装自定义包

   - `go install` 是 Go 中自动包安装工具：如需要将包安装到本地它会从远端仓库下载包：检出、编译和安装一气呵成。
   - 在包安装前的先决条件是要自动处理包自身依赖关系的安装。被依赖的包也会安装到子目录下，但是没有文档和示例：可以到网上浏览。
   - `go install` 使用了 GOPATH 变量。
   - 远端包
     - 需要创建目录在 Go 安装目录下，所以需要使用 root 或者 su 的身份执行命令。
     - 确保 Go 环境变量已经设置在 root 用户下的 ./bashrc 文件中。
     - 使用命令安装：`go install` tideland-cgl.googlecode.com/hg。
     - 可执行文件 hg.a 将被放到 $GOROOT/pkg/linux_amd64/tideland-cgl.googlecode.com 目录下，源码文件被放置在 $GOROOT/src/tideland-cgl.googlecode.com/hg 目录下，同样有个 hg.a 放置在 \_obj 的子目录下。
     - 现在就可以在 go 代码中使用这个包中的功能了，例如使用包名 cgl 导入：`import cgl "tideland-cgl.googlecode.com/hg"`。
     - 从 Go1 起 `go install` 安装 Google Code 的导入路径形式是：`"code.google.com/p/tideland-cgl"`。
   - 升级到新的版本
     - 更新到新版本的 Go 之后本地安装包的二进制文件将全被删除。如果想更新，重编译、重安装所有的 go 安装包可以使用：`go install -a`。
     - go 的版本发布的很频繁，所以需要注意发布版本和包的兼容性。go1 之后都是自己编译自己了。
     - `go install` 同样可以使用 `go install` 编译链接并安装本地自己的包。

8. 自定义包的目录结构、go install 和 go test

   1. 自定义包的目录结构：uc 代表通用包名, 名字为粗体的代表目录，斜体代表可执行文件

      ```
        /home/user/goprograms
            ucmain.go	(uc 包主程序)
            Makefile (ucmain 的 makefile)
            ucmain
            src/uc	 (包含 uc 包的 go 源码)
              uc.go
              uc_test.go
              Makefile (包的 makefile)
              uc.a
              _obj
                uc.a
              _test
                uc.a
            bin		(包含最终的执行文件)
              ucmain
            pkg/linux_amd64
              uc.a	(包的目标文件)
      ```

      - 将项目放在 goprograms 目录下 (可以创建一个环境变量 GOPATH：在 .profile 和 .bashrc 文件中添加 `export GOPATH=/home/user/goprograms`)，而项目将作为 src 的子目录。uc 包中的功能在 uc.go 中实现。
      - `uc.go`：

        ```
          package uc
          import "strings"

          func UpperCase(str string) string {
            return strings.ToUpper(str)
          }
        ```

      - 包通常附带一个或多个测试文件，在这创建了一个 uc_test.go 文件：

        ```
          package uc
          import "testing"

          type ucTest struct {
            in, out string
          }

          var ucTests = []ucTest {
            ucTest{"abc", "ABC"},
            ucTest{"cvo-az", "CVO-AZ"},
            ucTest{"Antwerp", "ANTWERP"},
          }

          func TestUC(t *testing.T) {
            for _, ut := range ucTests {
              uc := UpperCase(ut.in)
              if uc != ut.out {
                t.Errorf("UpperCase(%s) = %s, must be %s", ut.in, uc,
                ut.out)
              }
            }
          }
        ```

      - 通过指令编译并安装包到本地：`go install uc`, 这会将 uc.a 复制到 pkg/linux_amd64 下面。
      - 使用 make ，通过以下内容创建一个包的 Makefile 在 src/uc 目录下:

        ```
          include $(GOROOT)/src/Make.inc

          TARG=uc
          GOFILES=\
              uc.go\

          include $(GOROOT)/src/Make.pkg
        ```

      - 在该目录下的命令行调用: gomake， 这将创建一个 \_obj 目录并将包编译生成的存档 uc.a 放在该目录下。
      - 这个包可以通过 go test 测试。创建一个 uc.a 的测试文件在目录下，输出为 PASS 时测试通过。
      - 有可能当前的用户不具有足够的资格使用 go install（没有权限）。这种情况下，选择 root 用户 su。确保 Go 环境变量和 Go 源码路径也设置给 su，同样也适用普通用户
      - 创建主程序 ucmain.go:

        ```
          package main
          import (
            "./src/uc"
            "fmt"
          )

          func main() {
            str1 := "USING package uc!"
            fmt.Println(uc.UpperCase(str1))
          }
        ```

      - 然后在这个目录下输入 go install。
      - 另外复制 uc.a 到 /home/user/goprograms 目录并创建一个 Makefile 并写入文本：

        ```
          include $(GOROOT)/src/Make.inc
          TARG=ucmain
          GOFILES=\
            ucmain.go\

          include $(GOROOT)/src/Make.cmd
        ```

      - 执行 gomake 编译 ucmain.go 生成可执行文件 ucmain，运行 ./ucmain 显示: USING PACKAGE UC!。

   2. 本地安装包

      - 本地包在用户目录下，使用给出的目录结构，以下命令用来从源码安装本地包：
        ```
          go install /home/user/goprograms/src/uc # 编译安装 uc
          cd /home/user/goprograms/uc
          go install ./uc 	# 编译安装 uc（和之前的指令一样）
          cd ..
          go install .	# 编译安装 ucmain
        ```
      - 安装到 $GOPATH 下：如果我们想安装的包在系统上的其他 Go 程序中被使用，它一定要安装到 $GOPATH 下。 这样做，在 .profile 和 .bashrc 中设置 `export GOPATH=/home/user/goprograms`。
      - 然后执行 `go install uc` 将会复制包存档到 `$GOPATH/pkg/LINUX_AMD64/uc`。
      - 现在，uc 包可以通过 `import "uc"` 在任何 Go 程序中被引用。

   3. 依赖系统的代码
      - 在不同的操作系统上运行的程序以不同的代码实现是非常少见的：绝大多数情况下语言和标准库解决了大部分的可移植性问题。
      - 去写平台特定的代码，例如汇编语言。这种情况下，按照下面的约定是合理的：
        ```
          prog1.go
          prog1_linux.go
          prog1_darwin.go
          prog1_windows.go
        ```
      - `prog1.go` 定义了不同操作系统通用的接口，并将系统特定的代码写到 `prog1_os.go` 中。 对于 Go 工具可以指定 `prog1_$GOOS.go` 或 `prog1_$GOARCH.go` 或在平台 Makefile 中：`prog1*$(GOOS).go\` 或 `prog1*$(GOARCH).go\`。

9. 通过 Git 打包和安装

   1. 安装到 GitHub
      - 在 Linux 和 OS X 的机器上 Git 是默认安装的，在 Windows 上必须先自行安装。
      - 进入到 uc 包目录下并创建一个 Git 仓库在里面: `git init`。信息提示: `Initialized empty git repository in $PWD/uc`。
      - 每一个 Git 项目都需要一个对包进行描述的 README.md 文件，所以需要打开文本编辑器（gedit、notepad 或 LiteIde）并添加一些说明进去。
        - 添加所有文件到仓库：`git add README.md uc.go uc_test.go Makefile`。
        - 标记为第一个版本：`git commit -m "initial rivision"`。
      - 在云端创建一个新的 uc 仓库;发布的指令为（NNNN 替代用户名）:
        ```
          git remote add origin git@github.com:NNNN/uc.git
          git push -u origin master
        ```
      - 操作完成后检查 GitHub 上的包页面: `http://github.com/NNNN/uc`。
   2. 从 GitHub 安装
      - 从远端项目到本地机器，打开终端并执行（NNNN 是在 GitHub 上的用户名）：`go get github.com/NNNN/uc`。
      - 这样现在这台机器上的其他 Go 应用程序也可以通过导入路径：`"github.com/NNNN/uc"` 代替 `"./uc/uc"` 来使用。
      - 也可以将其缩写为：`import uc "github.com/NNNN/uc"`。
      - 然后修改 Makefile: 将 `TARG=uc` 替换为 `TARG=github.com/NNNN/uc`。
      - Gomake（和 go install）将通过 $GOPATH 下的本地版本进行工作。
      - 网站和版本控制系统的其他的选择(括号中为网站所使用的版本控制系统)：
        - BitBucket(hg/Git)
        - GitHub(Git)
        - Google Code(hg/Git/svn)
        - Launchpad(bzr)
      - 版本控制系统可以选择熟悉的或者本地使用的代码版本控制。Go 核心代码的仓库是使用 Mercurial(hg) 来控制的，所以它是一个最可能保证可以得到开发者项目中最好的软件。Git 也很出名，同样也适用。如果从未使用过版本控制，这些网站有一些很好的帮助并且可以通过在谷歌搜索 "{name} tutorial"（name 为想要使用的版本控制系统）得到许多很好的教程。
