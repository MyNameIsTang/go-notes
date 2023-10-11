# 编辑器、集成开发环境与其它工具

## 调试器
1. 在合适的位置使用打印语句输出相关变量的值：fmt.Print、fmt.Println、fmt.Printf
2. 在fmt.Printf中使用下面的说明符打印有关变量的相关信息
  * %+v 打印包括字段在内的实例的完整信息
  * %#v 打印包括字段和限定类型名称在内的实例的完整信息
  * %T 打印某个类型的完整说明
3. 使用panic()语句来获取栈跟踪信息（直到painc()时所有被调用函数的列表）
4. 使用关键字defer来跟踪代码执行过程

## 生成代码文档（http://golang.org/cmd/godoc）
1. godoc 工具会从Go程序和包文件中提取顶级声明的首行注释以及每个对象的相关注释，并生成相关文档
2. 可以作为一个提供在线文档浏览的web服务器：godoc -http=:6060
3. 用法：
  * go doc package 获取包的文档注释
  * go doc package/subpackage 获取子包的文档注释
  * go doc package function 获取某个函数在某个包中的文档注释
4. godoc 只能获取在Go安装目录../go/src中的注释内容
5. godoc 也可以用于生成非标准库的Go源码文件的文档注释

## 其他工具
1. go install 安装Go包的工具，类似Ruby中的rubygems。主要用于安装非标准库的包文件，将源代码编译为对象文件
2. go fix 用于将Go代码从旧的发行版迁移到最新的发行版，主要负责简单的、重复的、枯燥的修改工作
3. go test 是一个轻量级的单元测试框架