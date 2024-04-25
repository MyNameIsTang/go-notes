## 模式

1. 逗号 ok 模式

   - 在一个需要赋值的 if 条件语句中，使用这种模式去检测第二个参数值会让代码显得优雅简洁。这种模式在 Go 语言编码规范中非常重要。
   - 在函数返回时检测错误：
     ```
       value, err := pack1.Func1(param1)
       if err != nil {
         fmt.Printf("Error %s: ", err.Error())
         return
       }
     ```
   - 检测映射中是否存在一个键值：
     ```
       if value, ok := map1[key1]; ok {
         Process(value)
       }
     ```
   - 检测一个接口类型变量 varI 是否包含了类型 T：类型断言
     ```
       if value, ok := varI.(T); ok {
         Process(value)
       }
     ```
   - 检测一个通道 ch 是否关闭

     ```
       for input := range ch {
         Process(input)
       }

       for {
        if input, open := <-ch; !open {
          break
        }
        Process(input)
       }
     ```

2. defer 模式

   - 使用 defer 可以确保资源不再需要时，都会被恰当地关闭或归还到“池子”中。更重要的一点是，它可以恢复 panic。
   - 关闭一个文件流：`defer f.Close()`
   - 解锁一个被锁定的资源（mutex）：`mu.Lock();defer mu.Unlock()`
   - 关闭一个通道：`defer close(ch)`
   - 从 panic 恢复：
     ```
       defer func(){
         if err := recover(); err != nil {
           log.Printf('panic: %v', err)
         }
       }()
     ```
   - 停止一个计时器
     ```
       tick1 := time.NewTicker(updateInterval)
       defer tick1.Stop()
     ```
   - 释放一个进程 p：
     ```
       p, err:= os.StartProcess(...)
       defer p.Release()
     ```
   - 停止 CPU 性能分析并立即写入：
     ```
       pprof.StartCPUProfile()
       defer pprof.StopCPUProfile()
     ```
   - defer 也可以在打印报表时避免忘记输出页脚。

3. 可见性模式

   - 使用可见性规则控制对类型成员的访问，可以是 Go 变量或函数。
   - 在单独的包中定义类型时，强制使用工厂函数。

4. 运算符模式和接口
   - 运算符是一元或二元函数，它返回一个新对象而不修改其参数，类似 C++ 中的 + 和 \*，特殊的中缀运算符（+，-，\_ 等）可以被重载以支持类似数学运算的语法。
   - 但除了一些特殊情况，Go 语言并不支持运算符重载：为了克服该限制，运算符必须由函数来模拟。
   1. 函数作为运算符
      - 运算符由包级别的函数实现，以操作一个或两个参数，并返回一个新对象。
      - 函数针对要操作的对象，在专门的包中实现。
      - 如果想在这些运算中区分不同类型的矩阵（稀疏或稠密），由于没有函数重载，不得不给函数起不同的名称，例如：
        ```
          func addSparseToDense (a *sparseMatrix, b *denseMatrix) *denseMatrix
          func addDenseToDense (a *denseMatrix, b *denseMatrix) *denseMatrix
          func addSparseToSparse (a *sparseMatrix, b *sparseMatrix) *sparseMatrix
        ```
      - 最佳方案是将它们隐藏起来，作为包的私有函数，并暴露单一的 `Add()` 函数作为公共 API。可以在嵌套的 switch 断言中测试类型，以便在任何支持的参数组合上执行操作：
        ```
          func Add(a Matrix, b Matrix) Matrix {
            switch a.(type) {
            case sparseMatrix:
              switch b.(type) {
              case sparseMatrix:
                return addSparseToSparse(a.(sparseMatrix), b.(sparseMatrix))
              case denseMatrix:
                return addSparseToDense(a.(sparseMatrix), b.(denseMatrix))
              …
              }
            default:
              // 不支持的参数
              …
            }
          }
        ```
      - 更优雅和优选的方案是将运算符作为方法实现，标准库中到处都运用了这种做法。
   2. 方法作为运算符
      - 根据接收者类型不同，可以区分不同的方法。因此可以为每种类型简单地定义 Add 方法，来代替使用多个函数名称：
        ```
          func (a *sparseMatrix) Add(b Matrix) Matrix
          func (a *denseMatrix) Add(b Matrix) Matrix
        ```
      - 每个方法都返回一个新对象，成为下一个方法调用的接收者，因此可以使用链式调用表达式：`m := m1.Mult(m2).Add(m3)`
      - 正确的实现同样可以基于类型，通过 switch 类型断言在运行时确定：
        ```
          func (a *sparseMatrix) Add(b Matrix) Matrix {
            switch b.(type) {
            case sparseMatrix:
              return addSparseToSparse(a.(sparseMatrix), b.(sparseMatrix))
            case denseMatrix:
              return addSparseToDense(a.(sparseMatrix), b.(denseMatrix))
            …
            default:
              // 不支持的参数
              …
            }
          }
        ```
   3. 使用接口
      - 当在不同类型上执行相同的方法时，创建一个通用化的接口以实现多态的想法，就会自然产生。
        ```
          type Algebraic interface {
            Add(b Algebraic) Algebraic
            Min(b Algebraic) Algebraic
            Mult(b Algebraic) Algebraic
            …
            Elements()
          }
        ```
      - 为 matrix 类型定义 `Add()`，`Min()`，`Mult()`，……等方法。
      - 每种实现上述 Algebraic 接口类型的方法都可以链式调用。每个方法实现都应基于参数类型，使用 switch 类型断言来提供优化过的实现。另外，应该为仅依赖于接口的方法，指定一个默认处理分支：
        ```
          func (a *denseMatrix) Add(b Algebraic) Algebraic {
            switch b.(type) {
            case sparseMatrix:
              return addDenseToSparse(a, b.(sparseMatrix))
            …
            default:
              for x in range b.Elements() …
            }
          }
        ```
      - 如果一个通用的功能无法仅使用接口方法来实现，可能正在处理两个不怎么相似的类型，此时应该放弃这种运算符模式。遇到这种情况，把包拆分成两个，提供单独的接口。
