## Map

- map 是一种特殊的数据结构：一种元素对 (pair) 的无序集合，pair 的一个元素是 key，对应的另一个元素是 value，所以这个结构也称为关联数组或字典。
- 这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位。
- map 这种数据结构在其他编程语言中也称为字典 (Python) 、hash 和 HashTable 等。

1. 概念

   - map 是引用类型，可以使用如下声明：`var map1 map[keytype]valuetype`，`[keytype]` 和 `valuetype` 之间允许有空格，但是 gofmt 移除了空格。
   - 在声明的时候不需要知道 map 的长度，map 是可以动态增长的。未初始化的 map 的值是 nil。
   - key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float32(64)。所以数组、切片和结构体不能作为 key （含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的），但是指针和接口类型可以。如果要用结构体作为 key 可以提供 `Key()` 和 `Hash()` 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key。
   - value 可以是任意类型的；通过使用空接口类型，可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言。
   - map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以如果很在乎性能的话还是建议用切片来解决问题。
   - map 也可以用函数作为自己的值，这样就可以用来做分支结构：key 用来选择要执行的函数。
   - 如果 key1 是 map1 的 key，那么 `map1[key1]` 就是对应 key1 的值，就如同数组索引符号一样（数组可以视为一种简单形式的 map，key 是从 0 开始的整数）。key1 对应的值可以通过赋值符号来设置为 val1：`map1[key1] = val1`。
   - 令 `v := map1[key1]` 可以将 key1 对应的值赋值给 v；如果 map 中没有 key1 存在，那么 v 将被赋值为 map1 的值类型的空值。
   - 常用的 `len(map1)` 方法可以获得 map 中的 pair 数目，这个数目是可以伸缩的，因为 map-pairs 在运行时可以动态添加和删除。
   - map 可以用 `{key1: val1, key2: val2}` 的描述方法来初始化，就像数组和结构体一样。
   - map 是 引用类型 的： 内存用 `make()` 方法来分配。map 的初始化：`var map1 = make(map[keytype]valuetype)`。或者简写为：`map1 := make(map[keytype]valuetype)`。
   - 不要使用 `new()`，永远用 `make()` 来构造 map。
   - 如果错误地使用 `new()` 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：`mapCreated := new(map[string]float32)`。

   1. map 容量

      - 和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。但是也可以选择标明 map 的初始容量 capacity，就像这样：`make(map[keytype]valuetype, cap)`。例如：`map2 := make(map[string]float32, 100)`。
      - 当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。

   2. 用切片作为 map 的值

      - 既然一个 key 只能对应一个 value，而 value 又是一个原始类型，那么如果一个 key 要对应多个值怎么办？例如，当要处理 Unix 机器上的所有进程，以父进程（pid 为整型）作为 key，所有的子进程（以所有子进程的 pid 组成的切片）作为 value。通过将 value 定义为 []int 类型或者其他类型的切片，就可以优雅地解决这个问题。
      - 定义这种 map 的例子：
        ```
          mp1 := make(map[int][]int)
          mp2 := make(map[int]*[]int)
        ```

2. 测试键值对是否存在及删除元素

   - 区分到底是 key1 不存在还是它对应的 value 就是空值：`val1, isPresent = map1[key1]`。isPresent 返回一个 bool 值：如果 key1 存在于 map1，val1 就是 key1 对应的 value 值，并且 isPresent 为 true；如果 key1 不存在，val1 就是一个空值，并且 isPresent 会返回 false。
   - 只是想判断某个 key 是否存在而不关心它对应的值到底是多少：`_, ok := map1[key1] `。或者和 if 混合使用：`if _, ok := map1[key1]; ok {}`
   - 从 map1 中删除 key1：`delete(map1, key1)`，如果 key1 不存在，该操作不会产生错误。

3. for-range 的配套用法

   - 可以使用 for 循环读取 map：`for key, value := range map1 { }`。
   - map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。
   - map 的本质是散列表，而 map 的增长扩容会导致重新进行散列，这就可能使 map 的遍历结果在扩容前后变得不可靠，Go 设计者为了让大家不依赖遍历的顺序，每次遍历的起点--即起始 bucket 的位置不一样，即不让遍历都从某个固定的 bucket0 开始，所以即使未扩容时我们遍历出来的 map 也总是无序的。

4. map 类型的切片

   - 想获取一个 map 类型的切片，必须使用两次 make() 函数，第一次分配切片，第二次分配切片中每个 map 元素：
     ```
       items := make([]map[int]int, 5)
       for i := range items {
         items[i] = make(map[int]int, 1)
         items[i][1] = 2
       }
     ```
   - 需要注意的是，应通过索引使用切片的 map 元素。

5. map 的排序

   - map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序。
   - 如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序，然后可以使用切片的 for-range 方法打印出所有的 key 和 value。
     ```
       var barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
         "delta": 87, "echo": 56, "foxtrot": 12,
         "golf": 34, "hotel": 16, "indio": 87,
         "juliet": 65, "kili": 43, "lima": 98}
       keys := make([]string, len(barVal))
       i := 0
       for k := range barVal {
         keys[i] = k
         i++
       }
       sort.Strings(keys)
       fmt.Println(keys)
       for _, k := range keys {
         fmt.Printf("key: %v, Value is %v\n", k, barVal[k])
       }
     ```
   - 但是如果想要一个排序的列表，那么最好使用结构体切片，这样会更有效：
     ```
       type name struct {
         key string
         value int
       }
     ```

6. 将 map 的键值对调

   - 对调是指调换 key 和 value。如果 map 的值类型可以作为 key 且所有的 value 是唯一的，那么通过下面的方法可以简单的做到键值对调。

     ```
       var barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
         "delta": 87, "echo": 56, "foxtrot": 12,
         "golf": 34, "hotel": 16, "indio": 87,
         "juliet": 65, "kili": 43, "lima": 98}

       invMap := make(map[int]string, len(barVal))

       for key, val := range barVal {
         invMap[val] = key
       }
     ```

   - 如果原始 value 值不唯一那这么做肯定会出问题；这种情况下不会报错，但是当遇到不唯一的 key 时应当直接停止对调，且此时对调后的 map 很可能没有包含原 map 的所有键值对！一种解决方法就是仔细检查唯一性并且使用多值 map，比如使用 `map[int][]string` 类型。
