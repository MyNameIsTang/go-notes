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
