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

3. 类型断言：如何检测和转换接口变量的类型

   - 一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的 动态 类型，即运行时在变量中存储的值的实际类型。
   - 在执行过程中动态类型可能会有所不同，但是它总是可以分配给接口变量本身的类型。通常我们可以使用 **类型断言** 来测试在某个时刻 varI 是否包含类型 T 的值：`v := varI.(T) `。
   - varI 必须是一个接口变量，否则编译器会报错。
   - 类型断言可能是无效的，虽然编译器会尽力检查转换是否有效，但是它不可能预见所有的可能性。如果转换在程序运行时失败会导致错误发生。更安全的方式是使用以下形式来进行类型断言：
     ```
       if v, ok := varI.(T); ok {  // checked type assertion
         Process(v)
         return
       }
     ```
     - 如果转换合法，v 是 varI 转换到类型 T 的值，ok 会是 true；否则 v 是类型 T 的零值，ok 是 false，也没有运行时错误发生。
     - 应该总是使用上面的方式来进行类型断言。

4. 类型判断：type-switch

   - 接口变量的类型也可以使用一种特殊形式的 switch 来检测：type-switch
   - 变量 t 得到了 接口变量 的值和类型，所有 case 语句中列举的类型（nil 除外）都必须实现对应的接口，如果被检测类型没有在 case 语句列举的类型中，就会执行 default 语句。
   - 可以用 type-switch 进行运行时类型分析，但是在 type-switch 不允许有 fallthrough 。
   - 如果仅仅是测试变量的类型，不用它的值，那么就可以不需要赋值语句。
     ```
       switch areaIntf.(type) {
       case *Square:
         // TODO
       case *Circle:
         // TODO
       ...
       default:
         // TODO
       }
     ```
   - 下面的代码片段展示了一个类型分类函数，它有一个可变长度参数，可以是任意类型的数组，它会根据数组元素的实际类型执行不同的动作：
     ```
       func classifier(items ...interface{}) {
         for i, x := range items {
           switch x.(type) {
           case bool:
             fmt.Printf("Param #%d is a bool\n", i)
           case float64:
             fmt.Printf("Param #%d is a float64\n", i)
           case int, int64:
             fmt.Printf("Param #%d is a int\n", i)
           case nil:
             fmt.Printf("Param #%d is a nil\n", i)
           case string:
             fmt.Printf("Param #%d is a string\n", i)
           default:
             fmt.Printf("Param #%d is unknown\n", i)
           }
         }
       }
     ```
   - 在处理来自于外部的、类型未知的数据时，比如解析诸如 JSON 或 XML 编码的数据，类型测试和转换会非常有用。

5. 测试一个值是否实现了某个接口

   - 假定 v 是一个值，然后我们想测试它是否实现了 Stringer 接口，可以这样做：

     ```
       type Stringer interface {
          String() string
       }

       if sv, ok := v.(Stringer); ok {
          fmt.Printf("v implements String(): %s\n", sv.String())
       }
     ```

   - `Print()` 函数就是如此检测类型是否可以打印自身的。
   - 接口是一种契约，实现类型必须满足它，它描述了类型的行为，规定类型可以做什么。接口彻底将类型能做什么，以及如何做分离开来，使得相同接口的变量在不同的时刻表现出不同的行为，这就是多态的本质。
   - 编写参数是接口变量的函数，这使得它们更具有一般性。
   - 使用接口使代码更具有普适性。
   - 标准库里到处都使用了这个原则，如果对接口概念没有良好的把握，是不可能理解它是如何构建的。

6. 使用方法集与接口

   - 作用于变量上的方法实际上是不区分变量到底是指针还是值的。当碰到接口类型值时，这会变得有点复杂，原因是接口变量中存储的具体值是不可寻址的，幸运的是，如果使用不当编译器会给出错误。
   - 在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以根据具体类型 P 直接辨识的：
     - 指针方法可以通过指针调用
     - 值方法可以通过值调用
     - 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
     - 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
   - 将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。
   - Go 语言规范定义了接口方法集的调用规则：
     - 类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
     - 类型 T 的可调用方法集包含接受者为 T 的所有方法
     - 类型 T 的可调用方法集不包含接受者为 \*T 的方法

7. 例子：读和写

   - 读和写是软件中很普遍的行为，提起它们会立即想到读写文件、缓存（比如字节或字符串切片）、标准输入输出、标准错误以及网络连接、管道等等，或者读写我们的自定义类型。
   - io 包提供了用于读和写的接口 `io.Reader` 和 `io.Writer`：

     ```
      type Reader interface {
          Read(p []byte) (n int, err error)
      }

      type Writer interface {
          Write(p []byte) (n int, err error)
      }
     ```

   - 只要类型实现了读写接口，提供 Read 和 Write 方法，就可以从它读取数据，或向它写入数据。
   - 一个对象要是可读的，它必须实现 `io.Reader` 接口，这个接口只有一个签名是 `Read(p []byte) (n int, err error)` 的方法，它从调用它的对象上读取数据，并把读到的数据放入参数中的字节切片中，然后返回读取的字节数和一个 error 对象，如果没有错误发生返回 nil，如果已经到达输入的尾端，会返回 `io.EOF("EOF")`，如果读取的过程中发生了错误，就会返回具体的错误信息。
   - 类似地，一个对象要是可写的，它必须实现 `io.Writer` 接口，这个接口也只有一个签名是 `Write(p []byte) (n int, err error)` 的方法，它将指定字节切片中的数据写入调用它的对象里，然后返回实际写入的字节数和一个 error 对象（如果没有错误发生就是 nil）。
   - io 包里的 Readers 和 Writers 都是不带缓冲的，bufio 包里提供了对应的带缓冲的操作，在读写 UTF-8 编码的文本文件时它们尤其有用。
   - 在实际编程中尽可能的使用这些接口，会使程序变得更通用，可以在任何实现了这些接口的类型上使用读写方法。

8. 空接口

   1. 概念
      - **空接口或者最小接口** 不包含任何方法，它对实现不做任何要求：`type Any interface {}`
      - 任何其他类型都实现了空接口（它不仅仅像 Java/C# 中 Object 引用类型），any 或 Any 是空接口一个很好的别名或缩写。
      - 空接口类似 Java/C# 中所有类的基类： `Object` 类，二者的目标也很相近。
      - 可以给一个空接口类型的变量 `var val interface {}` 赋任何类型的值。
      - 每个 `interface {}` 变量在内存中占据两个字长：一个用来存储它包含的类型，另一个用来存储它包含的数据或者指向数据的指针。
   2. 构建通用类型或包含不同类型变量的数组
      - 给空接口定一个别名类型 Element：`type Element interface{}`。
      - 然后定义一个容器类型的结构体 Vector，它包含一个 Element 类型元素的切片：
        ```
          type Vector struct {
            a []Element
          }
        ```
      - Vector 里能放任何类型的变量，因为任何类型都实现了空接口，实际上 Vector 里放的每个元素可以是不同类型的变量。我们为它定义一个 `At()` 方法用于返回第 i 个元素：
        ```
          func (p *Vector) At(i int) Element {
            return p.a[i]
          }
        ```
      - 再定一个 `Set()` 方法用于设置第 i 个元素的值：
        ```
          func (p *Vector) Set(i int, e Element) {
            p.a[i] = e
          }
        ```
      - Vector 中存储的所有元素都是 Element 类型，要得到它们的原始类型（unboxing：拆箱）需要用到类型断言。
   3. 复制数据切片至空接口切片
      - 假设有一个 myType 类型的数据切片，想将切片中的数据复制到一个空接口切片中，类似：
        ```
          var dataSlice []myType = FuncReturnSlice()
          var interfaceSlice []interface{} = dataSlice
        ```
      - 可惜不能这么做，编译时会出错：`cannot use dataSlice (type []myType) as type []interface { } in assignment`。
      - 原因是它们俩在内存中的布局是不一样的。
      - 必须使用 for-range 语句来一个一个显式地赋值：
        ```
          var dataSlice []myType = FuncReturnSlice()
          var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
          for i, d := range dataSlice {
              interfaceSlice[i] = d
          }
        ```
   4. 通用类型的节点数据结构

      - 现在可以使用空接口作为数据字段的类型，这样就能写出通用的代码。
      - 下面是实现一个二叉树的部分代码：通用定义、用于创建空节点的 NewNode 方法，及设置数据的 SetData 方法：

        ```
          type Node struct {
            le   *Node
            data interface{}
            ri   *Node
          }

          func NewNode(left, right *Node) *Node {
            return &Node{left, nil, right}
          }

          func (n *Node) SetData(data interface{}) {
            n.data = data
          }
        ```

   5. 接口到接口

      - 一个接口的值可以赋值给另一个接口变量，只要底层类型实现了必要的方法。
      - 这个转换是在运行时进行检查的，转换失败会导致一个运行时错误：这是 Go 语言动态的一面，可以拿它和 Ruby 和 Python 这些动态语言相比较。
      - 下面是函数调用的一个例子：

        ```
          type myPrintInterface interface {
            print()
          }

          func f3(x myInterface) {
            x.(myPrintInterface).print() // type assertion to myPrintInterface
          }
        ```

        - x 转换为 myPrintInterface 类型是完全动态的：只要 x 的底层类型（动态类型）定义了 print 方法这个调用就可以正常运行（译注：若 x 的底层类型未定义 print 方法，此处类型断言会导致 panic，最佳实践应该为 if mpi, ok := x.(myPrintInterface); ok { mpi.print() }）。

9. 反射包

   1. 方法和类型的反射
      - 反射是用程序检查其所拥有的结构，尤其是类型的一种能力，这是元编程的一种形式。
      - 反射可以在运行时检查类型和变量，例如：它的大小、它的方法以及它能“动态地”调用这些方法。
      - 这对于没有源代码的包尤其有用。这是一个强大的工具，除非真得有必要，否则应当避免使用或小心使用。
      - 变量的最基本信息就是类型和值：反射包的 Type 用来表示一个 Go 类型，反射包的 Value 为 Go 值提供了反射接口。
      - 两个简单的函数，`reflect.TypeOf` 和 `reflect.ValueOf`，返回被检查对象的类型和值。例如，x 被定义为：`var x float64 = 3.4`，那么 `reflect.TypeOf(x)` 返回 ` float64`，`reflect.ValueOf(x) ` 返回 `<float64 Value>`
      - 实际上，反射是通过检查一个接口的值，变量首先被转换成空接口。这从下面两个函数签名能够很明显的看出来：
        ```
          func TypeOf(i interface{}) Type
          func ValueOf(i interface{}) Value
        ```
      - 接口的值包含一个 type 和 value。
      - 反射可以从接口值反射到对象，也可以从对象反射回接口值。
        - `reflect.Type` 和 `reflect.Value` 都有许多方法用于检查和操作它们。
        - Value 有一个 `Type()` 方法返回 `reflect.Value` 的 Type 类型。
        - Type 和 Value 都有 `Kind()` 方法返回一个常量来表示类型：Uint、Float64、Slice 等等。
        - 同样 Value 有叫做 `Int()` 和 `Float()` 的方法可以获取存储在内部的值（跟 int64 和 float64 一样）
          ```
            const (
              Invalid Kind = iota
              Bool
              Int
              Int8
              Int16
              Int32
              Int64
              Uint
              Uint8
              Uint16
              Uint32
              Uint64
              Uintptr
              Float32
              Float64
              Complex64
              Complex128
              Array
              Chan
              Func
              Interface
              Map
              Ptr
              Slice
              String
              Struct
              UnsafePointer
            )
          ```
        - 对于 float64 类型的变量 x，如果 `v:=reflect.ValueOf(x)`，那么 `v.Kind()` 返回 `reflect.Float64` ，所以下面的表达式是 `true：v.Kind() == reflect.Float64`。
   2. 通过反射修改（设置）值
      - 假设要把 x 的值改为 3.1415。Value 有一些方法可以完成这个任务，但是必须小心使用：`v.SetFloat(3.1415)`。这将产生一个错误：reflect.Value.SetFloat using unaddressable value。
      - 问题的原因是 v 不是可设置的（这里并不是说值不可寻址）。是否可设置是 Value 的一个属性，并且不是所有的反射值都有这个属性：可以使用 `CanSet()` 方法测试是否可设置。
      - 当 `v := reflect.ValueOf(x)` 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。要想 v 的更改能作用到 x，那就必须传递 x 的地址 `v = reflect.ValueOf(&x)`。
      - 通过 `Type()` 我们看到 v 现在的类型是 `*float64` 并且仍然是不可设置的。
      - 要想让其可设置我们需要使用 `Elem()` 函数，这间接地使用指针：`v = v.Elem()`。
      - 现在 `v.CanSet()` 返回 true 并且 `v.SetFloat(3.1415)` 设置成功了！
   3. 反射结构
      - 有些时候需要反射一个结构类型。`NumField()` 方法返回结构内的字段数量；通过一个 for 循环用索引取得每个字段的值 `Field(i)`。
      - 同样能够调用签名在结构上的方法，例如，使用索引 n 来调用：`Method(n).Call(nil)`。

10. Printf() 和反射

    - fmt 包中的 `Printf()`（以及其他格式化输出函数）都会使用反射来分析它的 ... 参数。
    - `Printf()` 的函数声明为：`func Printf(format string, args ... interface{}) (n int, err error)`。
    - `Printf()` 中的 ... 参数为空接口类型。`Printf()` 使用反射包来解析这个参数列表。所以，`Printf()` 能够知道它每个参数的类型。因此格式化字符串中只有 %d 而没有 %u 和 %ld，因为它知道这个参数是 unsigned 还是 long。

11. 接口与动态类型

    1. Go 的动态类型
       - 在经典的面向对象语言（像 C++，Java 和 C#）中数据和方法被封装为类的概念：类包含它们两者，并且不能剥离。
       - Go 没有类：数据（结构体或更一般的类型）和方法是一种松耦合的正交关系。
       - Go 中的接口跟 Java/C# 类似：都是必须提供一个指定方法集的实现。但是更加灵活通用：任何提供了接口方法实现代码的类型都隐式地实现了该接口，而不用显式地声明。
       - 和其它语言相比，Go 是唯一结合了接口值，静态类型检查（是否该类型实现了某个接口），运行时动态转换的语言，并且不需要显式地声明类型是否满足某个接口。该特性允许我们在不改变已有的代码的情况下定义和使用新接口。
       - 接收一个（或多个）接口类型作为参数的函数，其实参可以是任何实现了该接口的类型的变量。 实现了某个接口的类型可以被传给任何以此接口为参数的函数。
       - 类似于 Python 和 Ruby 这类动态语言中的动态类型 (duck typing)；这意味着对象可以根据提供的方法被处理（例如，作为参数传递给函数），而忽略它们的实际类型：它们能做什么比它们是什么更重要。
    2. 动态方法调用

       - 像 Python，Ruby 这类语言，动态类型是延迟绑定的（在运行时进行）：方法只是用参数和变量简单地调用，然后在运行时才解析（它们很可能有像 responds_to 这样的方法来检查对象是否可以响应某个方法，但是这也意味着更大的编码量和更多的测试工作）。
       - Go 的实现与此相反，通常需要编译器静态检查的支持：当变量被赋值给一个接口类型的变量时，编译器会检查其是否实现了该接口的所有函数。如果方法调用作用于像 `interface{}` 这样的“泛型”上，你可以通过类型断言来检查变量是否实现了相应接口。
       - 例如，用不同的类型表示 XML 输出流中的不同实体。然后我们为 XML 定义一个如下的“写”接口：

         ```
           type xmlWriter interface {
             WriteXML(w io.Writer) error
           }
            // Exported XML streaming function.
            func StreamXML(v interface{}, w io.Writer) error {
              if xw, ok := v.(xmlWriter); ok {
                // It’s an  xmlWriter, use method of asserted type.
                return xw.WriteXML(w)
              }
              // No implementation, so we have to use our own function (with perhaps reflection):
              return encodeToXML(v, w)
            }

            // Internal XML encoding function.
            func encodeToXML(v interface{}, w io.Writer) error {
              // ...
            }
         ```

       - 因此 Go 提供了动态语言的优点，却没有其他动态语言在运行时可能发生错误的缺点。
       - 对于动态语言非常重要的单元测试来说，这样即可以减少单元测试的部分需求，又可以发挥相当大的作用。
       - Go 的接口提高了代码的分离度，改善了代码的复用性，使得代码开发过程中的设计模式更容易实现。用 Go 接口还能实现“依赖注入模式”。

    3. 接口的提取
       - *提取接口*是非常有用的设计模式，可以减少需要的类型和方法数量，而且不需要像传统的基于类的面向对象语言那样维护整个的类层次结构。
       - Go 接口可以让开发者找出自己写的程序中的类型。假设有一些拥有共同行为的对象，并且开发者想要抽象出这些行为，这时就可以创建一个接口来使用。
       - 所以不用提前设计出所有的接口；整个设计可以持续演进，而不用废弃之前的决定。类型要实现某个接口，它本身不用改变，只需要在这个类型上实现新的方法。
    4. 显式地指明类型实现了某个接口
       - 如果希望满足某个接口的类型显式地声明它们实现了这个接口，可以向接口的方法集中添加一个具有描述性名字的方法。例如：
         ```
           type Fooer interface {
             Foo()
             ImplementsFooer()
           }
         ```
       - 类型 Bar 必须实现 ImplementsFooer 方法来满足 Fooer 接口，以清楚地记录这个事实：
         ```
           type Bar struct{}
           func (b Bar) ImplementsFooer() {}
           func (b Bar) Foo() {}
         ```
       - 大部分代码并不使用这样的约束，因为它限制了接口的实用性。但是有些时候，这样的约束在大量相似的接口中被用来解决歧义。
    5. 空接口和函数重载
       - 在 Go 语言中函数重可以用可变参数 `...T` 作为函数最后一个参数来实现。
       - 如果我们把 T 换为空接口，那么可以知道任何类型的变量都是满足 T (空接口）类型的，这样就允许我们传递任何数量任何类型的参数给函数，即重载的实际含义。
       - 函数 `fmt.Printf` 就是这样做的：`fmt.Printf(format string, a ...interface{}) (n int, errno error)`。
       - 这个函数通过枚举 slice 类型的实参动态确定所有参数的类型，并查看每个类型是否实现了 `String()` 方法，如果是就用于产生输出信息。
    6. 接口的继承
       - 当一个类型包含（内嵌）另一个类型（实现了一个或多个接口）的指针时，这个类型就可以使用（另一个类型）所有的接口方法。
         ```
           type Task struct {
             Command string
             *log.Logger
           }
         ```
         - 这个类型的工厂方法像这样：
           ```
             func NewTask(command string, logger *log.Logger) *Task {
               return &Task{command, logger}
             }
           ```
         - 当 `log.Logger` 实现了 `Log()` 方法后，Task 的实例 task 就可以调用该方法：`task.Log()`。
       - 类型可以通过继承多个接口来提供像多重继承一样的特性：
         ```
           type ReaderWriter struct {
             *io.Reader
             *io.Writer
           }
         ```
       - 上面概述的原理被应用于整个 Go 包，多态用得越多，代码就相对越少。这被认为是 Go 编程中的重要的最佳实践。
       - 有用的接口可以在开发的过程中被归纳出来。添加新接口非常容易，因为已有的类型不用变动（仅仅需要实现新接口的方法）。
       - 已有的函数可以扩展为使用接口类型的约束性参数：通常只有函数签名需要改变。对比基于类的 OO 类型的语言在这种情况下则需要适应整个类层次结构的变化。

12. Go 中的面向对象

    - Go 没有类，而是松耦合的类型、方法对接口的实现。
    - OO 语言最重要的三个方面分别是：封装、继承和多态，在 Go 中它们是怎样表现的呢？
      - 封装（数据隐藏）：和别的 OO 语言有 4 个或更多的访问层次相比，Go 把它简化为了 2 层：
        - 包范围内的：通过标识符首字母小写，对象只在它所在的包内可见。
        - 可导出的：通过标识符首字母大写，对象对所在包以外也可见。
      - 类型只拥有自己所在包中定义的方法。
      - 继承：用组合实现：内嵌一个（或多个）包含想要的行为（字段和方法）的类型；多重继承可以通过内嵌多个类型实现。
      - 多态：用接口实现：某个类型的实例可以赋给它所实现的任意接口类型的变量。类型和接口是松耦合的，并且多重继承可以通过实现多个接口实现。Go 接口不是 Java 和 C# 接口的变体，而且接口间是不相关的，并且是大规模编程和可适应的演进型设计的关键。

13. 结构体、集合和高阶函数

    - 在应用中定义了一个结构体，那么也可能需要这个结构体的（指针）对象集合，比如：

      ```
        type Any interface{}
        type Car struct {
          Model        string
          Manufacturer string
          BuildYear    int
          // ...
        }

        type Cars []*Car
      ```

    - 使用高阶函数，实际上也就是把函数作为定义所需方法（其他函数）的参数，例如：

      - 定义一个通用的 `Process()` 函数，它接收一个作用于每一辆 car 的 f 函数作参数：
        ```
          func (cs Cars) Process(f func(car *Car)) {
            for _, c := range cs {
              f(c)
            }
          }
        ```
      - 在上面的基础上，实现一个查找函数来获取子集合，并在 `Process()` 中传入一个闭包执行（这样就可以访问局部切片 cars）：

        ```
          func (cs Cars) FindAll(f func(car *Car) bool) Cars {

            cars := make([]*Car, 0)
            cs.Process(func(c *Car) {
              if f(c) {
                cars = append(cars, c)
              }
            })
            return cars
          }
        ```

      - 实现对应作用的功效 (Map-functionality)，从每个 car 对象当中产出某些东西：
        ```
          func (cs Cars) Map(f func(car *Car) Any) []Any {
            result := make([]Any, 0)
            ix := 0
            cs.Process(func(c *Car) {
              result[ix] = f(c)
              ix++
            })
            return result
          }
        ```
      - 现在我们可以定义下面这样的具体查询：
        ```
          allNewBMWs := allCars.FindAll(func(car *Car) bool {
            return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
          })
        ```
      - 也可以根据参数返回不同的函数。也许我们想根据不同的厂商添加汽车到不同的集合，但是这（这种映射关系）可能会是会改变的。所以我们可以定义一个函数来产生特定的添加函数和 map 集：

        ```
          func MakeSortedAppender(manufacturers []string)(func(car *Car),map[string]Cars) {
            // Prepare maps of sorted cars.
            sortedCars := make(map[string]Cars)
            for _, m := range manufacturers {
              sortedCars[m] = make([]*Car, 0)
            }
            sortedCars["Default"] = make([]*Car, 0)
            // Prepare appender function:
            appender := func(c *Car) {
              if _, ok := sortedCars[c.Manufacturer]; ok {
                sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
              } else {
                sortedCars["Default"] = append(sortedCars["Default"], c)
              }

            }
            return appender, sortedCars
          }
        ```
      - 
