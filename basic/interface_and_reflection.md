## 接口 (interface)与反射 (reflection)

1. 接口的定义

   - Go 语言不是一种 “传统” 的面向对象编程语言：它里面没有类和继承的概念。
   - 但是 Go 语言里有非常灵活的 **接口** 概念，通过它可以实现很多面向对象的特性。接口提供了一种方式来 **说明** 对象的行为：如果谁能搞定这件事，它就可以用在这儿。
   - 接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
   - 通过如下格式定义接口：
     ```
       type Namer interface {
         Method1(param_list) return_type
         Method2(param_list) return_type
         ...
       }
     ```
   - （按照约定，只包含一个方法的）接口的名字由方法名加 er 后缀组成，例如 Printer、Reader、Writer、Logger、Converter 等等。还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，此时接口名以 able 结尾，或者以 I 开头（像 .NET 或 Java 中那样）。
   - Go 语言中的接口都很简短，通常它们会包含 0 个、最多 3 个方法。
   - 不像大多数面向对象编程语言，在 Go 语言中接口可以有值，一个接口类型的变量或一个 接口值 ：`var ai Namer`，ai 是一个多字（multiword）数据结构，它的值是 nil。它本质上是一个指针，虽然不完全是一回事。指向接口值的指针是非法的，它们不仅一点用也没有，还会导致代码错误。
   - 此处的方法指针表是通过运行时反射能力构建的。
   - 类型（比如结构体）可以实现某个接口的方法集；这个实现可以描述为，该类型的变量上的每一个具体方法所组成的集合，包含了该接口的方法集。实现了 Namer 接口的类型的变量可以赋值给 ai（即 receiver 的值），方法表指针（method table ptr）就指向了当前的方法实现。当另一个实现了 Namer 接口的类型的变量被赋给 ai，receiver 的值和方法表指针也会相应改变。
   - 类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。
   - 实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
   - 一个类型可以实现多个接口。
   - 接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。
   - 即使接口在类型之后才定义，二者处于不同的包中，被单独编译：只要类型实现了接口中的方法，它就实现了此接口。
   - 所有这些特性使得接口具有很大的灵活性。
   - 接口变量里包含了接收者实例的值和指向对应方法表的指针。
   - 这是 多态 的 Go 版本，多态是面向对象编程中一个广为人知的概念：根据当前的类型选择正确的方法，或者说：同一种类型在不同的实例上似乎表现出不同的行为。
   - 一个标准库的例子
     - io 包里有一个接口类型 Reader:
       ```
         type Reader interface {
           Read(p []byte) (n int, err error)
         }
       ```
     - 定义变量 r： `var r io.Reader`
       ```
         	 var r io.Reader
           r = os.Stdin    // see 12.1
           r = bufio.NewReader(r)
           r = new(bytes.Buffer)
           f,_ := os.Open("test.txt")
           r = bufio.NewReader(f)
       ```
     - 上面 r 右边的类型都实现了 `Read()` 方法，并且有相同的方法签名，r 的静态类型是 `io.Reader`。
   - 备注：从某个类型的角度来看，它的接口指的是：它的所有导出方法，只不过没有显式地为这些导出方法额外定一个接口而已。

2. 接口嵌套接口

   - 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
   - 比如接口 File 包含了 ReadWrite 和 Lock 的所有方法，它还额外有一个 `Close()` 方法。

     ```
      type ReadWrite interface {
        Read(b Buffer) bool
        Write(b Buffer) bool
      }

      type Lock interface {
        Lock()
        Unlock()
      }

      type File interface {
        ReadWrite
        Lock
        Close()
      }
     ```
