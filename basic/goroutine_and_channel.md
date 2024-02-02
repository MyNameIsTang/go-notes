## 协程 (goroutine) 与通道 (channel)

- 作为一门 21 世纪的语言，Go 原生支持应用之间的通信（网络，客户端和服务端，分布式计算和程序的并发。程序可以在不同的处理器和计算机上同时执行不同的代码段。Go 语言为构建并发程序的基本代码块是协程 (goroutine) 与通道 (channel)。他们需要语言，编译器，和 runtime 的支持。Go 语言提供的垃圾回收器对并发编程至关重要。
- 不要通过共享内存来通信，而通过通信来共享内存。
- 通信强制协作。

1. 并发、并行和协程

   1. 什么是协程
      - 一个应用程序是运行在机器上的一个进程；
      - 进程是一个运行在自己内存地址空间里的独立执行体。
      - 一个进程由一个或多个操作系统线程组成，这些线程其实是共享同一个内存地址空间的一起工作的执行体。
      - 几乎所有'正式'的程序都是多线程的，以便让用户或计算机不必等待，或者能够同时服务多个请求（如 Web 服务器），或增加性能和吞吐量（例如，通过对不同的数据集并行执行代码）。
      - 一个并发程序可以在一个处理器或者内核上使用多个线程来执行任务，但是只有同一个程序在某个时间点同时运行在多核或者多处理器上才是真正的并行。
      - 并行是一种通过使用多处理器以提高速度的能力。所以并发程序可以是并行的，也可以不是。
      - 公认的，使用多线程的应用难以做到准确，最主要的问题是内存中的数据共享，它们会被多线程以无法预知的方式进行操作，导致一些无法重现或者随机的结果（称作竞态）。
      - **不要使用全局变量或者共享内存，它们会给代码在并发运算的时候带来危险**。
      - 解决之道在于同步不同的线程，对数据加锁，这样同时就只有一个线程可以变更数据。在 Go 的标准库 sync 中有一些工具用来在低级别的代码中实现加锁；不过过去的软件开发经验告诉我们这会带来更高的复杂度，更容易使代码出错以及更低的性能，所以这个经典的方法明显不再适合现代多核/多处理器编程：`thread-per-connection` 模型不够有效。
      - Go 更倾向于其他的方式，在诸多比较合适的范式中，有个被称作 `Communicating Sequential Processes`（顺序通信处理）（CSP, C. Hoare 发明的）还有一个叫做 `message passing-model`（消息传递）（已经运用在了其他语言中，比如 Erlang）。
      - 在 Go 中，应用程序并发处理的部分被称作 goroutines（协程），它可以进行更有效的并发运算。在协程和操作系统线程之间并无一对一的关系：协程是根据一个或多个线程的可用性，映射（多路复用，执行于）在他们之上的；协程调度器在 Go 运行时很好的完成了这个工作。
      - 协程工作在相同的地址空间中，所以共享内存的方式一定是同步的；这个可以使用 sync 包来实现，不过我们很不鼓励这样做：Go 使用 channels 来同步协程。
      - 当系统调用（比如等待 I/O）阻塞协程时，其他协程会继续在其他线程上工作。协程的设计隐藏了许多线程创建和管理方面的复杂工作。
      - 协程是轻量的，比线程更轻。它们痕迹非常不明显（使用少量的内存和资源）：使用 4K 的栈内存就可以在堆中创建它们。因为创建非常廉价，必要的时候可以轻松创建并运行大量的协程（在同一个地址空间中 100,000 个连续的协程）。并且它们对栈进行了分割，从而动态的增加（或缩减）内存的使用；栈的管理是自动的，但不是由垃圾回收器管理的，而是在协程退出后自动释放。
      - 协程可以运行在多个操作系统线程之间，也可以运行在线程之内，可以很小的内存占用就可以处理大量的任务。由于操作系统线程上的协程时间片，可以使用少量的操作系统线程就能拥有任意多个提供服务的协程，而且 Go 运行时可以聪明的意识到哪些协程被阻塞了，暂时搁置它们并处理其他协程。
      - 存在两种并发方式：确定性的（明确定义排序）和非确定性的（加锁/互斥从而未定义排序）。Go 的协程和通道理所当然的支持确定性的并发方式（例如通道具有一个 sender 和一个 receiver）。
      - 协程是通过使用关键字 go 调用（执行）一个函数或者方法来实现的（也可以是匿名或者 lambda 函数）。这样会在当前的计算过程中开始一个同时进行的函数，在相同的地址空间中并且分配了独立的栈，比如：`go sum(bigArray)`，在后台计算总和。
      - 协程的栈会根据需要进行伸缩，不出现栈溢出；开发者不需要关心栈的大小。当协程结束的时候，它会静默退出：用来启动这个协程的函数不会得到任何的返回值。
      - 协程的栈会根据需要进行伸缩，不出现栈溢出；开发者不需要关心栈的大小。当协程结束的时候，它会静默退出：用来启动这个协程的函数不会得到任何的返回值。
      - 任何 Go 程序都必须有的 `main()` 函数也可以看做是一个协程，尽管它并没有通过 go 来启动。协程可以在程序初始化的过程中运行（在 `init()` 函数中）。
      - 在一个协程中，比如需要进行非常密集的运算，可以在运算循环中周期的使用 `runtime.Gosched()`：这会让出处理器，允许运行其他协程；并不会使当前协程挂起，所以会自动恢复执行。使用 `Gosched()` 可以使计算均匀分布，使通信不至于迟迟得不到响应。
   2. 并发和并行的差异
      - Go 的并发原语提供了良好的并发设计基础：表达程序结构以便表示独立地执行的动作；所以 Go 的重点不在于并行的首要位置：并发程序可能是并行的，也可能不是。并行是一种通过使用多处理器以提高速度的能力。但往往是，一个设计良好的并发程序在并行方面的表现也非常出色。
      - 必须使用 GOMAXPROCS 变量。这会告诉运行时有多少个协程同时执行。
      - 并且只有 gc 编译器真正实现了协程，适当的把协程映射到操作系统线程。使用 gccgo 编译器，会为每一个协程创建操作系统线程。
   3. 使用 GOMAXPROCS
      - 在 gc 编译器下（6g 或者 8g）必须设置 GOMAXPROCS 为一个大于默认值 1 的数值来允许运行时支持使用多于 1 个的操作系统线程，所有的协程都会共享同一个线程除非将 GOMAXPROCS 设置为一个大于 1 的数。
      - 当 GOMAXPROCS 大于 1 时，会有一个线程池管理许多的线程。
      - 通过 gccgo 编译器 GOMAXPROCS 有效的与运行中的协程数量相等。
      - 假设 n 是机器上处理器或者核心的数量。如果设置环境变量 `GOMAXPROCS>=n`，或者执行 `runtime.GOMAXPROCS(n)`，接下来协程会被分割（分散）到 n 个处理器上。更多的处理器并不意味着性能的线性提升。
      - 有这样一个经验法则，对于 n 个核心的情况设置 GOMAXPROCS 为 n-1 以获得最佳性能，也同样需要遵守这条规则：协程的数量 > 1 + GOMAXPROCS > 1。
      - 所以如果在某一时间只有一个协程在执行，不要设置 GOMAXPROCS！
      - 还有一些通过实验观察到的现象：在一台 1 颗 CPU 的笔记本电脑上，增加 GOMAXPROCS 到 9 会带来性能提升。在一台 32 核的机器上，设置 `GOMAXPROCS=8` 会达到最好的性能，在测试环境中，更高的数值无法提升性能。如果设置一个很大的 GOMAXPROCS 只会带来轻微的性能下降；设置 `GOMAXPROCS=100`，使用 top 命令和 H 选项查看到只有 7 个活动的线程。
      - 增加 GOMAXPROCS 的数值对程序进行并发计算是有好处的；
      - GOMAXPROCS 等同于（并发的）线程数量，在一台核心数多于 1 个的机器上，会尽可能有等同于核心数的线程在并行运行。
   4. 如何用命令行指定使用的核心数量
      - 使用 flags 包，如下：`var numCores = flag.Int("n", 2, "number of CPU cores to use")`。
      - 在 `main()` 中：
        ```
           flag.Parse()
           runtime.GOMAXPROCS(*numCores)
        ```
      - 协程可以通过调用 `runtime.Goexit()` 来停止，尽管这样做几乎没有必要。
      - 当 `main()` 函数返回的时候，程序退出：它不会等待任何其他非 main 协程的结束。这就是为什么在服务器程序中，每一个请求都会启动一个协程来处理，`server()` 函数必须保持运行状态。通常使用一个无限循环来达到这样的目的。
      - 另外，协程是独立的处理单元，一旦陆续启动一些协程，无法确定他们是什么时候真正开始执行的。的代码逻辑必须独立于协程调用的顺序。
      - 协程更有用的一个例子应该是在一个非常长的数组中查找一个元素。将数组分割为若干个不重复的切片，然后给每一个切片启动一个协程进行查找计算。这样许多并行的协程可以用来进行查找任务，整体的查找时间会缩短（除以协程的数量）。
   5. Go 协程 (goroutines) 和协程 (coroutines)
      - 在其他语言中，比如 C#，Lua 或者 Python 都有协程的概念。这个名字表明它和 Go 协程有些相似，不过有两点不同：
        - Go 协程意味着并行（或者可以以并行的方式部署），协程一般来说不是这样的。
        - Go 协程通过通道来通信；协程通过让出和恢复操作来通信。
      - Go 协程比协程更强大，也很容易从协程的逻辑复用到 Go 协程。

2. 协程间的信道

   1. 概念
      - 协程可以使用共享变量来通信，但是很不提倡这样做，因为这种方式给所有的共享内存的多线程都带来了困难。
      - Go 有一种特殊的类型，通道（channel），就像一个可以用于发送类型化数据的管道，由其负责协程之间的通信，从而避开所有由共享内存导致的陷阱；这种通过通道进行通信的方式保证了同步性。
      - 数据在通道中进行传递：**在任何给定时间，一个数据被设计为只有一个协程可以对其访问，所以不会发生数据竞争**。 数据的所有权（可以读写数据的能力）也因此被传递。
      - 通道服务于通信的两个目的：值的交换，同步的，保证了两个计算（协程）任何时候都是可知状态。
      - 通常使用这样的格式来声明通道：`var identifier chan datatype`。
      - 未初始化的通道的值是 nil。
      - 所以通道只能传输一种类型的数据，比如 `chan int` 或者 `chan string`，所有的类型都可以用于通道，空接口 `interface{}` 也可以，甚至可以（有时非常有用）创建通道的通道。
      - 通道实际上是类型化消息的队列：使数据得以传输。它是先进先出(FIFO) 的结构所以可以保证发送给他们的元素的顺序（通道可以比作 Unix shells 中的双向管道 (two-way pipe)）。通道也是引用类型，所以我们使用 `make()` 函数来给它分配内存：`ch1 := make(chan string)`。
      - 构建一个 int 通道的通道： `chanOfChans := make(chan chan int)`。
      - 函数通道：`funcChan := make(chan func())`。
      - 所以通道是第一类对象：可以存储在变量中，作为函数的参数传递，从函数返回以及通过通道发送它们自身。另外它们是类型化的，允许类型检查，比如尝试使用整数通道发送一个指针。
   2. 通信操作符 <-
      - 这个操作符直观的标示了数据的传输：信息按照箭头的方向流动。
      - 流向通道（发送），`ch <- int1` 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）
      - 从通道流出（接收），三种方式：
        - `int2 = <- ch` 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）
        - `<- ch` 可以单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证，所以以下代码是合法的：
          ```
            if <- ch != 1000{
                ...
             }
          ```
        - 同一个操作符 <- 既用于发送也用于接收，但 Go 会根据操作对象弄明白该干什么 。
      - 虽非强制要求，但为了可读性通道的命名通常以 ch 开头或者包含 chan 。通道的发送和接收都是原子操作：它们总是互不干扰地完成。
      - 运行时 (runtime) 会检查所有的协程是否在等待着什么东西（可从某个通道读取或者写入某个通道），这意味着程序将无法继续执行。这是死锁 (deadlock) 的一种形式，而运行时 (runtime) 可以为我们检测到这种情况。
      - 不要使用打印状态来表明通道的发送和接收顺序：由于打印状态和通道实际发生读写的时间延迟会导致和真实发生的顺序不同。
   3. 通道阻塞
      - 默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束。
      - 可以想象一个无缓冲的通道在没有空间来保存数据的时候：必须要一个接收者准备好接收通道的数据然后发送者可以直接把数据发送给接收者。所以通道的发送/接收操作在对方准备好之前是阻塞的：
        - 对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果 ch 中的数据无人接收，就无法再给通道传入其他数据：新的输入无法在通道非空的情况下传入。所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）。
        - 对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。
   4. 通过一个（或多个）通道交换数据进行协程同步。
      - 通信是一种同步形式：通过通道，两个协程在通信（协程会合）中某刻同步交换数据。无缓冲通道成为了多个协程同步的完美工具。
      - 甚至可以在通道两端互相阻塞对方，形成了叫做死锁的状态。Go 运行时会检查并 `panic()`，停止程序。死锁几乎完全是由糟糕的设计导致的。
      - 无缓冲通道会被阻塞。设计无阻塞的程序可以避免这种情况，或者使用带缓冲的通道。
   5. 同步通道-使用带缓冲的通道
      - 一个无缓冲通道只能包含 1 个元素，有时显得很局限。给通道提供了一个缓存，可以在扩展的 make 命令中设置它的容量，如下：`ch1 := make(chan string, 100)`，第二个参数是通道可以同时容纳的元素（这里是 string）个数。
      - 在缓冲满载（缓冲被全部使用）之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了。
      - 缓冲容量和类型无关，所以可以（尽管可能导致危险）给一些通道设置不同的容量，只要他们拥有同样的元素类型。内置的 `cap()` 函数可以返回缓冲区的容量。
      - 如果容量大于 0，通道就是异步的了：缓冲满载（发送）或变空（接收）之前通信不会阻塞，元素会按照发送的顺序被接收。如果容量是 0 或者未设置，通信仅在收发双方准备好的情况下才可以成功。
      - 同步：ch :=make(chan type, value)
        - `value == 0 -> synchronous`, unbuffered （阻塞）
        - `value > 0 -> asynchronous`, buffered（非阻塞）取决于 value 元素
      - 若使用通道的缓冲，程序会在“请求”激增的时候表现更好：更具弹性，专业术语叫：更具有伸缩性(scalable)。在设计算法时首先考虑使用无缓冲通道，只在不确定的情况下使用缓冲。
   6. 协程中用通道输出结果
      - 为了知道计算何时完成，可以通过信道回报。例：
        ```
           ch := make(chan int)
           go sum(bigArray, ch)
           sum := <-ch
        ```
      - 也可以使用通道来达到同步的目的，这个很有效的用法在传统计算机中称为信号量 (semaphore)。或者换个方式：通过通道发送信号告知处理已经完成（在协程中）。
      - 在其他协程运行时让 main 程序无限阻塞的通常做法是在 `main()` 函数的最后放置一个 `select {}`。
      - 也可以使用通道让 main 程序等待协程完成，就是所谓的信号量模式。
   7. 信号量模式

      - 协程通过在通道 ch 中放置一个值来处理结束的信号。`main()` 协程等待 `<-ch` 直到从中获取到值。
      - 期望从这个通道中获取返回的结果，像这样：

        ```
           func compute(ch chan int){
              ch <- someComputation() // when it completes, signal on the channel.
           }

           func main(){
              ch := make(chan int) 	// allocate a channel.
              go compute(ch)		// start something in a goroutines
              doSomethingElseForAWhile()
              result := <- ch
           }
        ```

      - 这个信号也可以是其他的，不返回结果，比如下面这个协程中的匿名函数 (lambda) 协程：
        ```
           ch := make(chan int)
           go func(){
              // doSomething
              ch <- 1 // Send a signal; value does not matter
           }()
           doSomethingElseForAWhile()
           <- ch	// Wait for goroutine to finish; discard sent value.
        ```
      - 或者等待两个协程完成，每一个都会对切片 s 的一部分进行排序，片段如下：
        ```
            done := make(chan bool)
            // doSort is a lambda function, so a closure which knows the channel done:
            doSort := func(s []int){
               sort(s)
               done <- true
            }
            i := pivot(s)
            go doSort(s[:i])
            go doSort(s[i:])
            <-done
            <-done
        ```
      - 用完整的信号量模式对长度为 N 的 float64 切片进行了 N 个 `doSomething()` 计算并同时完成，通道 sem 分配了相同的长度（且包含空接口类型的元素），待所有的计算都完成后，发送信号（通过放入值）。在循环中从通道 sem 不停的接收数据来等待所有的协程完成。
        ```
           type Empty interface {}
           var empty Empty
           ...
           data := make([]float64, N)
           res := make([]float64, N)
           sem := make(chan Empty, N)
           ...
           for i, xi := range data {
              go func (i int, xi float64) {
                 res[i] = doSomething(i, xi)
                 sem <- empty
              } (i, xi)
           }
           // wait for goroutines to finish
           for i := 0; i < N; i++ { <-sem }
        ```

   8. 实现并行的 for 循环
      - for 循环的每一个迭代是并行完成的：
        ```
           for i, v := range data {
              go func (i int, v float64) {
                 doSomething(i, v)
                 ...
              } (i, v)
           }
        ```
      - 在 for 循环中并行计算迭代可能带来很好的性能提升。不过所有的迭代都必须是独立完成的。
   9. 用带缓冲通道实现一个信号量

      - 信号量是实现互斥锁（排外锁）常见的同步机制，限制对资源的访问，解决读写问题，比如没有实现信号量的 sync 的 Go 包，使用带缓冲的通道可以轻松实现：
        - 带缓冲通道的容量和要同步的资源容量相同
        - 通道的长度（当前存放的元素个数）与当前资源被使用的数量相同
        - 容量减去通道的长度就是未处理的资源个数（标准信号量的整数值）
      - 不用管通道中存放的是什么，只关注长度；因此我们创建了一个长度可变但容量为 0（字节）的通道：
        ```
           type Empty interface {}
           type semaphore chan Empty
        ```
      - 将可用资源的数量 N 来初始化信号量 `semaphore：sem = make(semaphore, N)`
      - 然后直接对信号量进行操作：

        ```
           // acquire n resources
           func (s semaphore) P(n int) {
              e := new(Empty)
              for i := 0; i < n; i++ {
                 s <- e
              }
           }

           // release n resources
           func (s semaphore) V(n int) {
              for i:= 0; i < n; i++{
                 <- s
              }
           }
        ```

      - 可以用来实现一个互斥的例子：

        ```
           /* mutexes */
           func (s semaphore) Lock() {
              s.P(1)
           }

           func (s semaphore) Unlock(){
              s.V(1)
           }

           /* signal-wait */
           func (s semaphore) Wait(n int) {
              s.P(n)
           }

           func (s semaphore) Signal() {
              s.V(1)
           }
        ```

      - **习惯用法：通道工厂模式**
        - 编程中常见的另外一种模式如下：不将通道作为参数传递给协程，而用函数来生成一个通道并返回（工厂角色）；函数内有个匿名函数被协程调用。
        ```
            func pump() chan int {
               ch := make(chan int)
               go func() {
                  for i := 0; ; i++ {
                     ch <- i
                  }
               }()
               return ch
            }
        ```

   10. 给通道使用 for 循环

       - for 循环的 range 语句可以用在通道 ch 上，便可以从通道中获取值，像这样：

         ```
            for v := range ch {
               fmt.Printf("The value is %v\n", v)
            }
         ```

       - 它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。很明显，另外一个协程必须写入 ch（不然代码就阻塞在 for 循环了），而且必须在写入完成后才关闭。

         ```
            func suck(ch chan int) {
               go func() {
                  for v := range ch {
                     fmt.Println(v)
                  }
               }()
            }
         ```

       - **习惯用法：通道迭代器模式**
         - 通常，需要从包含了地址索引字段 items 的容器给通道填入元素。为容器的类型定义一个方法 `Iter()`，返回一个只读的通道 items：
           ```
              func (c *container) Iter () <- chan item {
                 ch := make(chan item)
                 go func () {
                    for i:= 0; i < c.Len(); i++{	// or use a for-range loop
                       ch <- c.items[i]
                    }
                 } ()
                 return ch
              }
           ```
         - 在协程里，一个 for 循环迭代容器 c 中的元素（对于树或图的算法，这种简单的 for 循环可以替换为深度优先搜索）。
         - 调用这个方法的代码可以这样迭代容器：`for x := range container.Iter() { ... }`。
         - 其运行在自己启动的协程中，所以上边的迭代用到了一个通道和两个协程（可能运行在不同的线程上）。 这样我们就有了一个典型的生产者-消费者模式。
         - 如果在程序结束之前，向通道写值的协程未完成工作，则这个协程不会被垃圾回收；这是设计使然。这种看起来并不符合预期的行为正是由通道这种线程安全的通信方式所导致的。如此一来，一个协程为了写入一个永远无人读取的通道而被挂起就成了一个 bug ，而并非预想中的那样被悄悄回收掉 (garbage-collected) 了。
       - **习惯用法：生产者消费者模式**
         - 假设有 Produce() 函数来产生 Consume() 函数需要的值。它们都可以运行在独立的协程中，生产者在通道中放入给消费者读取的值。整个处理过程可以替换为无限循环：
           ```
              for {
                 Consume(Produce())
              }
           ```

   11. 通道的方向

       - 通道类型可以用注解来表示它只发送或者只接收：
         - 只能发送：`var send_only chan<- int`
         - 只能接收：`var recv_only <-chan int`
       - 只接收的通道 (`<-chan T`) 无法关闭，因为关闭通道是发送者用来表示不再给通道发送值了，所以对只接收通道是没有意义的。
       - 通道创建的时候都是双向的，但也可以分配给有方向的通道变量：

         ```
            var c = make(chan int) // bidirectional
            go source(c)
            go sink(c)

            func source(ch chan<- int){
               for { ch <- 1 }
            }

            func sink(ch <-chan int) {
               for { <-ch }
            }
         ```

       - **习惯用法：管道和选择器模式**

         - 协程处理它从通道接收的数据并发送给输出通道：

           ```
              sendChan := make(chan int)
              receiveChan := make(chan string)
              go processChannel(sendChan, receiveChan)

              func processChannel(in <-chan int, out chan<- string) {
                 for inValue := range in {
                    result := ... /// processing inValue
                    out <- result
                 }
              }
           ```

         - 通过使用方向注解来限制协程对通道的操作。

3. 协程的同步：关闭通道-测试阻塞的通道

   - 通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。只有发送者需要关闭通道，接收者永远不会需要。
   - 我们如何在通道的 `sendData()` 完成的时候发送一个信号，`getData()` 又如何检测到通道是否关闭或阻塞？
     - 第一个可以通过函数 `close(ch)` 来完成：这个将通道标记为无法通过发送操作 `<-` 接受更多的值；给已经关闭的通道发送或者再次关闭都会导致运行时的 `panic()`。在创建一个通道后使用 defer 语句是个不错的办法：
     ```
        ch := make(chan float64)
        defer close(ch)
     ```
     - 第二个问题可以使用逗号 ok 模式用来检测通道是否被关闭：
       - 通常和 if 语句一起使用：
         ```
            if v, ok := <-ch; ok {
               process(v)
            }
         ```
       - 或者在 for 循环中接收的时候，当关闭的时候使用 break：
         ```
            v, ok := <-ch
            if !ok {
               break
            }
            process(v)
         ```
       - 而检测通道当前是否阻塞，需要使用 select
         ```
            select {
            case v, ok := <-ch:
            if ok {
               process(v)
            } else {
               fmt.Println("The channel is closed")
            }
            default:
            fmt.Println("The channel is blocked")
            }
         ```
       - 使用 for-range 语句来读取通道是更好的办法，因为这会自动检测通道是否关闭：
         ```
            for input := range ch {
               process(input)
            }
         ```
   - **阻塞和生产者-消费者模式**：
     - 两个协程经常是一个阻塞另外一个。如果程序工作在多核心的机器上，大部分时间只用到了一个处理器。可以通过使用带缓冲（缓冲空间大于 0）的通道来改善。比如，缓冲大小为 100，迭代器在阻塞之前，至少可以从容器获得 100 个元素。如果消费者协程在独立的内核运行，就有可能让协程不会出现阻塞。
     - 由于容器中元素的数量通常是已知的，需要让通道有足够的容量放置所有的元素。这样，迭代器就不会阻塞（尽管消费者协程仍然可能阻塞）。然而，这实际上加倍了迭代容器所需要的内存使用量，所以通道的容量需要限制一下最大值。记录运行时间和性能测试可以帮助找到最小的缓存容量带来最好的性能。

4. 使用 select 切换协程

   - 从不同的并发执行的协程中获取值可以通过关键字 select 来完成，它和 switch 控制语句非常相似也被称作通信开关；它的行为像是“你准备好了吗”的轮询机制；select 监听进入通道的数据，也可以是用通道发送值的时候。
     ```
        select {
           case u:= <- ch1:
               ...
           case v:= <- ch2:
               ...
               ...
           default: // no value ready to be received
               ...
        }
     ```
   - default 语句是可选的；fallthrough 行为，和普通的 switch 相似，是不允许的。在任何一个 case 中执行 break 或者 return，select 就结束了。
   - select 做的就是：选择处理列出的多个通信情况中的一个。
     - 如果都阻塞了，会等待直到其中一个可以处理
     - 如果多个可以处理，随机选择一个
     - 如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）。
   - 在 select 中使用发送操作并且有 default 可以确保发送不被阻塞！如果没有 default，select 就会一直阻塞。
   - select 语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过 break 语句使循环退出。
   - **习惯用法：后台服务模式**
     - 服务通常是是用后台协程中的无限循环实现的，在循环中使用 select 获取并处理通道中的数据：
       ```
          // Backend goroutine.
          func backend() {
             for {
                select {
                case cmd := <-ch1:
                   // Handle ...
                case cmd := <-ch2:
                   ...
                case cmd := <-chStop:
                   // stop server
                }
             }
          }
       ```
       - 在程序的其他地方给通道 ch1，ch2 发送数据，比如：通道 stop 用来清理结束服务程序。
     - 另一种方式（但是不太灵活）就是（客户端）在 chRequest 上提交请求，后台协程循环这个通道，使用 switch 根据请求的行为来分别处理：
       ```
          func backend() {
             for req := range chRequest {
                switch req.Subjext() {
                   case A1:  // Handle case ...
                   case A2:  // Handle case ...
                   default:
                   // Handle illegal request ..
                   // ...
                }
             }
          }
       ```

5. 通道、超时和计时器（Ticker）

   - time 包中有一些有趣的功能可以和通道组合使用。
   - 其中就包含了 `time.Ticker` 结构体，这个对象以指定的时间间隔重复的向通道 C 发送时间值：
     ```
        type Ticker struct {
           C <-chan Time // the channel on which the ticks are delivered.
           // contains filtered or unexported fields
           ...
        }
     ```
   - 时间间隔的单位是 ns（纳秒，int64），在工厂函数 `time.NewTicker` 中以 Duration 类型的参数传入：`func NewTicker(dur) *Ticker`。
   - 在协程周期性的执行一些事情（打印状态日志，输出，计算等等）的时候非常有用。
   - 调用 `Stop()` 使计时器停止，在 defer 语句中使用。这些都很好地适应 select 语句:
     ```
        ticker := time.NewTicker(updateInterval)
        defer ticker.Stop()
        ...
        select {
        case u:= <-ch1:
           ...
        case v:= <-ch2:
           ...
        case <-ticker.C:
           logState(status) // call some logging function logState
        default: // no value ready to be received
           ...
        }
     ```
   - `time.Tick()` 函数声明为 `Tick(d Duration) <-chan Time`，当想返回一个通道而不必关闭它的时候这个函数非常有用：它以 d 为周期给返回的通道发送时间，d 是纳秒数。
   - 像下边的代码一样，可以限制处理频率（函数 `client.Call()` 是一个 RPC 调用：

     ```
        import "time"

        rate_per_sec := 10
        var dur Duration = 1e9 / rate_per_sec
        chRate := time.Tick(dur) // a tick every 1/10th of a second
        for req := range requests {
           <- chRate // rate limit our Service.Method RPC calls
           go client.Call("Service.Method", req, ...)
        }
     ```

     - 这样只会按照指定频率处理请求：chRate 阻塞了更高的频率。每秒处理的频率可以根据机器负载（和/或）资源的情况而增加或减少。
     - 定时器 (Timer) 结构体看上去和计时器 (Ticker) 结构体的确很像（构造为 `NewTimer(d Duration)`），但是它只发送一次时间，在 `Dration d` 之后。
       - 还有 `time.After(d)` 函数，声明如下：`func After(d Duration) <-chan Time`。
       - 在 `Duration d` 之后，当前时间被发到返回的通道；所以它和 `NewTimer(d).C` 是等价的；它类似 `Tick()`，但是 `After()` 只发送一次时间。

   - **习惯用法：简单超时模式**
     - 要从通道 ch 中接收数据，但是最多等待 1 秒。先创建一个信号通道，然后启动一个 lambda 协程，协程在给通道发送数据之前是休眠的：
       ```
          timeout := make(chan bool, 1)
          go func() {
                time.Sleep(1e9) // one second
                timeout <- true
          }()
       ```
     - 然后使用 select 语句接收 ch 或者 timeout 的数据：如果 ch 在 1 秒内没有收到数据，就选择到了 time 分支并放弃了 ch 的读取。
       ```
          select {
             case <-ch:
                // a read from ch has occured
             case <-timeout:
                // the read from ch has timed out
                break
          }
       ```
     - 第二种形式：取消耗时很长的同步调用
       - 可以使用 `time.After()` 函数替换 timeout-channel
       - 在 select 中通过 `time.After()` 发送的超时信号来停止协程的执行。
       - 在 timeoutNs 纳秒后执行 select 的 timeout 分支后，执行 `client.Call`的协程也随之结束，不会给通道 ch 返回值：
         ```
            ch := make(chan error, 1)
            go func() { ch <- client.Call("Service.Method", args, &reply) } ()
            select {
            case resp := <-ch
               // use resp and reply
            case <-time.After(timeoutNs):
               // call timed out
               break
            }
         ```
       - 注意缓冲大小设置为 1 是必要的，可以避免协程死锁以及确保超时的通道可以被垃圾回收。
       - 此外，需要注意在有多个 case 符合条件时， select 对 case 的选择是伪随机的，如果上面的代码稍作修改如下，则 select 语句可能不会在定时器超时信号到来时立刻选中 time.After(timeoutNs) 对应的 case，因此协程可能不会严格按照定时器设置的时间结束。
         ```
            ch := make(chan int, 1)
            go func() { for { ch <- 1 } } ()
            L:
            for {
               select {
               case <-ch:
                  // do something
               case <-time.After(timeoutNs):
                  // call timed out
                  break L
               }
            }
         ```
     - 第三种形式：假设程序从多个复制的数据库同时读取。只需要一个答案，需要接收首先到达的答案，Query 函数获取数据库的连接切片并请求。并行请求每一个数据库并返回收到的第一个响应：
       ```
          func Query(conns []Conn, query string) Result {
             ch := make(chan Result, 1)
             for _, conn := range conns {
                go func(c Conn) {
                      select {
                      case ch <- c.DoQuery(query):
                      default:
                      }
                }(conn)
             }
             return <- ch
          }
       ```
       - 结果通道 ch 必须是带缓冲的：以保证第一个发送进来的数据有地方可以存放，确保放入的首个数据总会成功，所以第一个到达的值会被获取而与执行的顺序无关。正在执行的协程可以总是可以使用 `runtime.Goexit()` 来停止。
     - 在应用中缓存数据：
       - 应用程序中用到了来自数据库（或者常见的数据存储）的数据时，经常会把数据缓存到内存中，因为从数据库中获取数据的操作代价很高；
       - 如果数据库中的值不发生变化就没有问题。但是如果值有变化，我们需要一个机制来周期性的从数据库重新读取这些值：缓存的值就不可用（过期）了，而且我们也不希望用户看到陈旧的数据。

6. 协程和回复（recover）

   - 一个用到 `recover()` 的程序

     ```
        func server(workChan <-chan *Work) {
           for work := range workChan {
              go safelyDo(work)   // start the goroutine for that work
           }
        }

        func safelyDo(work *Work) {
           defer func() {
              if err := recover(); err != nil {
                    log.Printf("Work failed with %s in %v", err, work)
              }
           }()
           do(work)
        }
     ```

     - 如果 `do(work)` 发生 `panic()`，错误会被记录且协程会退出并释放，而其他协程不受影响。

   - 因为 `recover()` 总是返回 nil，除非直接在 defer 修饰的函数中调用，defer 修饰的代码可以调用那些自身可以使用 `panic()` 和 `recover()` 避免失败的库例程（库函数）。

7. 新旧模型对比：任务和 worker

   - 假设我们需要处理很多任务；一个 worker 处理一项任务。
     ```
        type Task struct {
           // some state
        }
     ```
   - 旧模式：使用共享内存进行同步

     - 由各个任务组成的任务池共享内存；为了同步各个 worker 以及避免资源竞争，我们需要对任务池进行加锁保护：
       ```
         type Pool struct {
            Mu      sync.Mutex
            Tasks   []*Task
         }
       ```
       - `sync.Mutex` 是互斥锁：它用来在代码中保护临界区资源：同一时间只有一个 go 协程 (goroutine) 可以进入该临界区。如果出现了同一时间多个 go 协程都进入了该临界区，则会产生竞争：Pool 结构就不能保证被正确更新。
     - 在传统的模式中（经典的面向对象的语言中应用得比较多，比如 C++，JAVA，C#），worker 代码可能这样写：
       ```
          func Worker(pool *Pool) {
             for {
                pool.Mu.Lock()
                // begin critical section:
                task := pool.Tasks[0]        // take the first task
                pool.Tasks = pool.Tasks[1:]  // update the pool of tasks
                // end critical section
                pool.Mu.Unlock()
                process(task)
             }
          }
       ```
     - 这些 worker 有许多都可以并发执行；他们可以在 go 协程中启动。
     - 加锁保证了同一时间只有一个 go 协程可以进入到 pool 中：一项任务有且只能被赋予一个 worker 。
     - 如果不加锁，则工作协程可能会在 `task:=pool.Tasks[0]` 发生切换，导致 `pool.Tasks=pool.Tasks[1:]` 结果异常：一些 worker 获取不到任务，而一些任务可能被多个 worker 得到。
     - 加锁实现同步的方式在工作协程比较少时可以工作得很好，但是当工作协程数量很大，任务量也很多时，处理效率将会因为频繁的加锁/解锁开销而降低。
     - 当工作协程数增加到一个阈值时，程序效率会急剧下降，这就成为了瓶颈。

   - 新模式：使用通道

     - 使用通道进行同步：使用一个通道接受需要处理的任务，一个通道接受处理完成的任务（及其结果）。worker 在协程中启动，其数量 N 应该根据任务数量进行调整。
     - 主线程扮演着 Master 节点角色，可能写成如下形式：
       ```
         func main() {
            pending, done := make(chan *Task), make(chan *Task)
            go sendWork(pending)       // put tasks with work on the channel
            for i := 0; i < N; i++ {   // start N goroutines to do work
               go Worker(pending, done)
            }
            consumeWork(done)          // continue with the processed tasks
         }
       ```
     - worker 的逻辑比较简单：从 pending 通道拿任务，处理后将其放到 done 通道中：
       ```
         func Worker(in, out chan *Task) {
            for {
               t := <-in
               process(t)
               out <- t
            }
         }
       ```
       - 这里并不使用锁：从通道得到新任务的过程没有任何竞争。
       - 随着任务数量增加，worker 数量也应该相应增加，同时性能并不会像第一种方式那样下降明显。
       - 在 pending 通道中存在一份任务的拷贝，第一个 worker 从 pending 通道中获得第一个任务并进行处理，这里并不存在竞争（对一个通道读数据和写数据的整个过程是原子性的）。
       - 某一个任务会在哪一个 worker 中被执行是不可知的，反过来也是。worker 数量的增多也会增加通信的开销，这会对性能有轻微的影响。
     - 第二种模式对比第一种模式而言，不仅性能是一个主要优势，而且还有个更大的优势：代码显得更清晰、更优雅。
     - 对于任何可以建模为 Master-Worker 范例的问题，一个类似于 worker 使用通道进行通信和交互、Master 进行整体协调的方案都能完美解决。如果系统部署在多台机器上，各个机器上执行 Worker 协程，Master 和 Worker 之间使用 netchan 或者 RPC 进行通信。

   - 怎么选择是该使用锁还是通道？普遍的经验法则：
     - 使用锁的情景：
       - 访问共享数据结构中的缓存信息
       - 保存应用程序上下文和状态信息数据
     - 使用通道的情景：
       - 与异步操作的结果进行交互
       - 分发任务
       - 传递数据所有权
     - 当发现锁使用规则变得很复杂时，可以反省使用通道会不会使问题变得简单些。

8. 惰性生成器的实现

   - 生成器是指当被调用时返回一个序列中下一个值的函数，例如：
     ```
        generateInteger() => 0
        generateInteger() => 1
        generateInteger() => 2
     ```
   - 生成器每次返回的是序列中下一个值而非整个序列；这种特性也称之为惰性求值：只在你需要时进行求值，同时保留相关变量资源（内存和 CPU）：这是一项在需要时对表达式进行求值的技术。
   - 有一个细微的区别是从通道读取的值可能会是稍早前产生的，并不是在程序被调用时生成的。如果确实需要这样的行为，就得实现一个请求响应机制。
   - 当生成器生成数据的过程是计算密集型且各个结果的顺序并不重要时，那么就可以将生成器放入到 go 协程实现并行化。但是得小心，使用大量的 go 协程的开销可能会超过带来的性能增益。

9. 实现 Futures 模式

   - 所谓 Futures 就是指：有时候在使用某一个值之前需要先对其进行计算。这种情况下，可以在另一个处理器上进行该值的计算，到使用时，该值就已经计算完毕了。
   - Futures 模式通过闭包和通道可以很容易实现，类似于生成器，不同地方在于 Futures 需要返回一个值。
   - 假设我们有一个矩阵类型，我们需要计算两个矩阵 A 和 B 乘积的逆，首先我们通过函数 `Inverse(M)` 分别对其进行求逆运算，再将结果相乘。
     ```
        func InverseProduct(a Matrix, b Matrix) {
           a_inv := Inverse(a)
           b_inv := Inverse(b)
           return Product(a_inv, b_inv)
        }
     ```
     - 调用 `Product()` 函数只需要等到 a_inv 和 b_inv 的计算完成。
       ```
         func InverseProduct(a Matrix, b Matrix) {
            a_inv_future := InverseFuture(a)   // start as a goroutine
            b_inv_future := InverseFuture(b)   // start as a goroutine
            a_inv := <-a_inv_future
            b_inv := <-b_inv_future
            return Product(a_inv, b_inv)
         }
       ```
     - `InverseFuture()` 函数以 goroutine 的形式起了一个闭包，该闭包会将矩阵求逆结果放入到 future 通道中：
       ```
         func InverseFuture(a Matrix) chan Matrix {
            future := make(chan Matrix)
            go func() {
               future <- Inverse(a)
            }()
            return future
         }
       ```
   - 当开发一个计算密集型库时，使用 Futures 模式设计 API 接口是很有意义的。在你的包使用 Futures 模式，且能保持友好的 API 接口。此外，Futures 可以通过一个异步的 API 暴露出来。这样你可以以最小的成本将包中的并行计算移到用户代码中。

10. 复用
    1. 典型的客户端/服务器（C/S）模式
       - 客户端-服务器应用正是 goroutines 和 channels 的亮点所在。
       - 客户端 (Client) 可以是运行在任意设备上的任意程序，它会按需发送请求 (request) 至服务器。服务器 (Server) 接收到这个请求后开始相应的工作，然后再将响应 (response) 返回给客户端。
       - 典型情况下一般是多个客户端（即多个请求）对应一个（或少量）服务器。例如我们日常使用的浏览器客户端，其功能就是向服务器请求网页。而 Web 服务器则会向浏览器响应网页数据。
       - 使用 Go 的服务器通常会在协程中执行向客户端的响应，故而会对每一个客户端请求启动一个协程。一个常用的操作方法是客户端请求自身中包含一个通道，而服务器则向这个通道发送响应。
    2. 卸载 (Teardown)：通过信号通道关闭服务器
       - 
