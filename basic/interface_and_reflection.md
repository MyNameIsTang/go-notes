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
