## 网络、模板与网页应用

Go 在编写 web 应用方面非常得力。因为目前它还没有 GUI（Graphic User Interface 即图形化用户界面）的框架，通过文本或者模板展现的 html 页面是目前 Go 编写界面应用程序的唯一方式。

1. tcp 服务器

   - 一个 (web) 服务器应用需要响应众多客户端的并发请求：Go 会为每一个客户端产生一个协程用来处理请求。需要使用 net 包中网络通信的功能。它包含了处理 TCP/IP 以及 UDP 协议、域名解析等方法。
   - `net.Listen()`，实现了服务器的基本功能：用来监听和接收来自客户端的请求。
     - 返回一个 error 类型的错误变量
   - `listener.Accept()`，等待客户端的请求
     - 客户端的请求将产生一个 `net.Conn` 类型的连接变量。
   - `net.Dial()`，客户端创建了一个和服务器之间的连接。
     - 在网络编程中 `net.Dial()` 函数是非常重要的，一旦连接到远程系统，函数就会返回一个 `Conn` 类型的接口，可以用它发送和接收数据。
     - `Dial()` 函数简洁地抽象了网络层和传输层。所以不管是 IPv4 还是 IPv6，TCP 或者 UDP 都可以使用这个公用接口。

2. 简单的 web 服务器

   - http 是比 tcp 更高层的协议，它描述了网页服务器如何与客户端浏览器进行通信。Go 提供了 net/http 包。
   - 使用 `http.ListenAndServe("localhost:8080", nil)` 函数，如果成功会返回空，否则会返回一个错误（地址 localhost 部分可以省略，8080 是指定的端口号）。
   - `http.URL` 用于表示网页地址，其中字符串属性 Path 用于保存 url 的路径；`http.Request` 描述了客户端请求，内含一个 URL 字段。
   - 如果 req 是来自 html 表单的 POST 类型请求，“var1” 是该表单中一个输入域的名称，那么用户输入的值就可以通过 Go 代码 `req.FormValue("var1")`
     - 还有一种方法是先执行 `request.ParseForm()`，然后再获取 `request.Form["var1"]` 的第一个返回参数，就像这样：`var1, found := request.Form["var1"]`。
     - 第二个参数 found 为 true。如果 var1 并未出现在表单中，found 就是 false。
   - 表单属性实际上是 `map[string][]string` 类型
   - 网页服务器发送一个 `http.Response` 响应，它是通过 `http.ResponseWriter` 对象输出的，后者组装了 HTTP 服务器响应，通过对其写入内容，我们就将数据发送给了 HTTP 客户端。
   - 服务器必须做的事，即如何处理请求，通过 `http.HandleFunc()` 函数完成。
     - 可以为每一个特定的 url 定义一个单独的处理函数。这个函数需要两个参数：第一个是 ReponseWriter 类型的 w；第二个是请求 req。
     - 第一个参数是请求的路径，第二个参数是当路径被请求时，需要调用的处理函数的引用。
   - 注：
     - `http.ListenAndServe`可简写为：`http.ListenAndServe(":8080", http.HandlerFunc(HelloServer))`。
     - `fmt.Fprint()` 和 `fmt.Fprintf()`都是可以用来写入`http.ResponseWriter`的函数（他们实现了`io.Writer`），例如：`fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, body)`。
     - 如果需要使用安全的 https 连接，使用 `http.ListenAndServeTLS()` 代替 `http.ListenAndServe()`。
     - 除了 `http.HandleFunc("/", Hfunc)`，其中的 HFunc 是一个处理函数，签名为：
       ```
         func HFunc(w http.ResponseWriter, req *http.Request) {
           ...
         }
       ```
       - 也可以使用这种方式：`http.Handle("/", http.HandlerFunc(HFunc))`
         - HandlerFunc 只是定义了上述 HFunc 签名的别名：`type HandlerFunc func(ResponseWriter, *Request)`
         - 它是一个可以把普通的函数当做 HTTP 处理器 (Handler) 的适配器。如果函数 f 声明得合适，`HandlerFunc(f)` 就是一个执行 f 函数的 Handler 对象。
       - `http.Handle()` 的第二个参数也可以是 T 类型的对象 obj：`http.Handle("/", obj)`。
         - 如果 T 有 `ServeHTTP()` 方法，那就实现了 http 的 Handler 接口：
           ```
              func (obj *Typ) ServeHTTP(w http.ResponseWriter, req *http.Request) {
                ...
              }
           ```

3. 访问并读取页面数据

   - `http.Head()` 请求查看返回值；它的声明如下：`func Head(url string) (r *Response, err error)`。
   - 使用 `http.Get()` 获取并显示网页内容；`Get()` 返回值中的 Body 属性包含了网页内容，用 `ioutil.ReadAll()` 把它读出来。
   - `http.Redirect(w ResponseWriter, r *Request, url string, code int)`：这个函数会让浏览器重定向到 url（可以是基于请求 url 的相对路径），同时指定状态码。
   - `http.NotFound(w ResponseWriter, r *Request)`：这个函数将返回网页没有找到，HTTP 404 错误。
   - `http.Error(w ResponseWriter, error string, code int)`：这个函数返回特定的错误信息和 HTTP 代码。
   - 另一个 `http.Request` 对象 req 的重要属性：`req.Method`，这是一个包含 GET 或 POST 字符串，用来描述网页是以何种方式被请求的。
   - HTTP 状态码常量：
     ```
       http.StatusContinue		= 100
       http.StatusOK			= 200
       http.StatusFound		= 302
       http.StatusBadRequest		= 400
       http.StatusUnauthorized		= 401
       http.StatusForbidden		= 403
       http.StatusNotFound		= 404
       http.StatusInternalServerError	= 500
     ```
   - 使用 `w.header().Set("Content-Type", "../..")`（`w ResponseWriter`） 设置头信息。比如在网页应用发送 html 字符串的时候，在输出之前执行 `w.Header().Set("Content-Type", "text/html")`。

4. 写一个简单的网页应用 demo

   - simple_webserver.go
   - 当使用字符串常量表示 html 文本的时候，包含 `<html><body>...</body></html>` 对于让浏览器将它识别为 html 文档非常重要。
     - 更安全的做法是在处理函数中，在写入返回内容之前将头部的 content-type 设置为 `text/html`：`w.Header().Set("Content-Type", "text/html")`。
     - "Content-Type" 会让浏览器认为它可以使用函数` http.DetectContentType([]byte(form))` 来处理收到的数据。

5. 确保网页应用健壮

   - 当网页应用的处理函数发生 panic，服务器会简单地终止运行。
   - 首先能想到的是在每个处理函数中使用 defer/recover()，不过这样会产生太多的重复代码。
   - 使用闭包的错误处理模式是更优雅的方案。
   - 为增强代码可读性，为页面处理函数创建一个类型：`type HandleFnc func(http.ResponseWriter, *http.Request)`。
   - 错误处理函数应用闭包的模式：
     ```
       func logPanics(function HandleFnc) HandleFnc {
         return func(writer http.ResponseWriter, request *http.Request) {
           defer func() {
             if x := recover(); x != nil {
               log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
             }
           }()
           function(writer, request)
         }
       }
     ```
   - 用 logPanics() 来包装对处理函数的调用：
     ```
       http.HandleFunc("/test1", logPanics(SimpleServer))
       http.HandleFunc("/test2", logPanics(FormServer))
     ```

6. 用模板编写网页应用

   - https://golang.org/doc/articles/wiki/
   - wiki.go

7. 探索 template 包

   - https://golang.org/pkg/text/template/
   - 模板是一项更为通用的技术方案：数据驱动的模板被创建出来，以生成文本输出。HTML 仅是其中的一种特定使用案例。
   - 模板通过与数据结构的整合来生成，通常为结构体或其切片。
   - 当数据项传递给 `tmpl.Execute()` ，它用其中的元素进行替换， 动态地重写某一小段文本。只有被导出的数据项才可以被整合进模板中。可以在 `{{` 和 `}}` 中加入数据求值或控制结构。数据项可以是值或指针，接口隐藏了他们的差异。

   1. 字段替换：`{{.FieldName}}`
      - 要在模板中包含某个字段的内容，使用双花括号括起以点 (.) 开头的字段名。
      - 假设 Name 是某个结构体的字段，其值要在被模板整合时替换，则在模板中使用文本 `{{.Name}}`。
      - 当 Name 是 map 的键时这么做也是可行的。
      - 要创建一个新的 Template 对象，调用 `template.New()`，其字符串参数可以指定模板的名称。
      - `Parse()` 方法通过解析模板定义字符串，生成模板的内部表示。
      - 当使用包含模板定义字符串的文件时，将文件路径传递给 `ParseFiles()` 来解析。
      - 最后通过 `Execute()` 方法，数据结构中的内容与模板整合，并将结果写入方法的第一个参数中，其类型为 `io.Writer`。
      - 如果只是想简单地把 `Execute()` 方法的第二个参数用于替换，使用 `{{.}}`。
      - 当在浏览器环境中进行这些步骤，应首先使用 html 过滤器来过滤内容，例如 `{{html .}}`， 或者对 FieldName 过滤：`{{ .FieldName |html }}`。
      - `|html` 这部分代码，是请求模板引擎在输出 FieldName 的结果前把值传递给 html 格式化器，它会执行 HTML 字符转义（例如把 > 替换为 `&gt;`）。这可以避免用户输入数据破坏 HTML 文档结构。
   2. 验证模板格式
      - 为了确保模板定义语法是正确的，使用 `Must()` 函数处理 Parse 的返回结果。
      - 代码中常见到这 3 个基本函数被串联使用：`var strTempl = template.Must(template.New("TName").Parse(strTemplateHTML))`。
   3. If-else
      - 运行 `Execute()` 产生的结果来自模板的输出，它包含静态文本，以及被 `{{}}` 包裹的称之为管道的文本。
      - 对管道数据的输出结果用 if-else-end 设置条件约束：如果管道是空的，类似于：` {{if ``}} Will not print. {{end}} `。
      - 那么 if 条件的求值结果为 false，不会有输出内容。但如果是这样：`{{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}`，会输出 Print IF part.。
   4. 点号和 with-end
      - 点号 (.) 可以在 Go 模板中使用：其值 `{{.}}` 被设置为当前管道的值。
      - with 语句将点号设为管道的值。如果管道是空的，那么不管 with-end 块之间有什么，都会被忽略。在被嵌套时，点号根据最近的作用域取得值。
   5. 模板变量 $
      - 可以在模板内为管道设置本地变量，变量名以 $ 符号作为前缀。变量名只能包含字母、数字和下划线。
   6. range-end
      - range-end 结构格式为：`{{range pipeline}} T1 {{else}} T0 {{end}}`。
      - range 被用于在集合上迭代：管道的值必须是数组、切片或 map。如果管道的值长度为零，点号的值不受影响，且执行 T0；否则，点号被设置为数组、切片或 map 内元素的值，并执行 T1。
      - 例如：
        ```
          {{range .}}
            {{with .Author}}
              <p><b>{{html .}}</b> wrote:</p>
            {{else}}
              <p>An anonymous person wrote:</p>
            {{end}}
            <pre>{{html .Content}}</pre>
            <pre>{{html .Date}}</pre>
          {{end}}
        ```
   7. 模板预定义函数
      - 可以在模板代码中使用的预定义函数，例如 `printf()` 函数工作方式类似于 `fmt.Sprintf()`。

8. 用 rpc 实现远程过程调用

   - Go 程序之间可以使用 net/rpc 包实现相互通信，这是另一种客户端-服务器应用场景。
   - 它提供了一种方便的途径，通过网络连接调用远程函数。当然，仅当程序运行在不同机器上时，这项技术才实用。
   - rpc 包建立在 gob 包之上，实现了自动编码/解码传输的跨网络方法调用。
   - 服务器端需要注册一个对象实例，与其类型名一起，使之成为一项可见的服务：它允许远程客户端跨越网络或其他 I/O 连接访问此对象已导出的方法。总之就是在网络上暴露类型的方法。
   - rpc 包使用了 http 和 tcp 协议，以及用于数据传输的 gob 包。服务器端可以注册多个不同类型的对象（服务），但同一类型的多个对象会产生错误。

9. 基于网络的通道 netchan

   - 一项和 rpc 密切相关的技术是基于网络的通道。
   - netchan 包实现了类型安全的网络化通道：它允许一个通道两端出现由网络连接的不同计算机。其实现原理是，在其中一台机器上将传输数据发送到通道中，那么就可以被另一台计算机上同类型的通道接收。
   - 一个导出器 (exporter) 会按名称发布（一组）通道。导入器 (importer) 连接到导出的机器，并按名称导入这些通道。之后，两台机器就可按通常的方式来使用通道。网络通道不是同步的，它们类似于带缓存的通道。
   - 发送端示例代码：
     ```
       exp, err := netchan.NewExporter("tcp", "netchanserver.mydomain.com:1234")
       if err != nil {
         log.Fatalf("Error making Exporter: %v", err)
       }
       ch := make(chan myType)
       err := exp.Export("sendmyType", ch, netchan.Send)
       if err != nil {
         log.Fatalf("Send Error: %v", err)
       }
     ```
   - 接受端示例代码：
     ```
       imp, err := netchan.NewImporter("tcp", "netchanserver.mydomain.com:1234")
       if err != nil {
         log.Fatalf("Error making Importer: %v", err)
       }
       ch := make(chan myType)
       err = imp.Import("sendmyType", ch, netchan.Receive)
       if err != nil {
         log.Fatalf("Receive Error: %v", err)
       }
     ```

10. 与 websocket 通信

    - 与 http 协议相反，websocket 是通过客户端与服务器之间的对话，建立的基于单个持久连接的协议。然而在其他方面，其功能几乎与 http 相同。

11. 用 smtp 发送邮件
    - smtp 包实现了用于发送邮件的“简单邮件传输协议”（Simple Mail Transfer Protocol）。它有一个 Client 类型，代表一个连接到 SMTP
      - `Dial()` 方法返回一个已连接到 SMTP 服务器的客户端 Client
      - 设置 Mail（from，即发件人）和 Rcpt（to，即收件人）
      - `Data()` 方法返回一个用于写入数据的 Writer，这里利用 `buf.WriteTo(wc)` 写入
