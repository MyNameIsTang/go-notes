## 时间和日期

- time 包为我们提供了一个数据类型 time.Time（作为值使用）以及显示和测量时间和日期的功能函数。
- 当前时间可以使用 time.Now() 获取，或者使用 t.Day()、t.Minute() 等等来获取时间的一部分；你甚至可以自定义时间格式化字符串，例如： fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) 将会输出 21.07.2011。
- Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64。Location 类型映射某个时区的时间，UTC 表示通用协调世界时间。
- 包中的一个预定义函数 func (t Time) Format(layout string) string 可以根据一个格式化字符串来将一个时间 t 转换为相应格式的字符串，你可以使用一些预定义的格式，如：time.ANSIC 或 time.RFC822。
- 一般的格式化设计是通过对于一个标准时间的格式化描述来展现的，例如：`fmt.Println(t.Format("02 Jan 2006 15:04")) `。
- 如果你需要在应用程序在经过一定时间或周期执行某项任务（事件处理的特例），则可以使用 `time.After()` 或者 `time.Ticker`。`time.Sleep(d Duration)` 可以实现对某个进程（实质上是 goroutine）时长为 d 的暂停。
