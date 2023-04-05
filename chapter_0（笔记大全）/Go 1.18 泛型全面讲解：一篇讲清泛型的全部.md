# Go 1.18 泛型全面讲解：一篇讲清泛型的全部



## 1. 一切从函数的形参和实参说起

假设我们有个计算两数之和的函数
```go
func Add(a int, b int) int {
    return a + b
}
```

这个函数很简单，但是它有个问题——无法计算int类型之外的和。如果我们想计算浮点或者字符串的和该怎么办？解决办法之一就是像下面这样为不同类型定义不同的函数
```go
func AddFloat32(a float32, b float32) float32 {
    return a + b
}

func AddString(a string, b string) string {
    return a + b
}
```
可是除此之外还有没有更好的方法？答案是有的，我们可以来回顾下函数的 形参(parameter) 和 实参(argument) 这一基本概念：
```go
func Add(a int, b int) int {  
    // 变量a,b是函数的形参   "a int, b int" 这一串被称为形参列表
    return a + b
}

Add(100,200) // 调用函数时，传入的100和200是实参
```
我们知道，函数的 形参(parameter) 只是类似占位符的东西并没有具体的值，只有我们调用函数传入实参(argument) 之后才有具体的值。

那么，如果我们将 形参 实参 这个概念推广一下，给变量的类型也引入和类似形参实参的概念的话，问题就迎刃而解：在这里我们将其称之为 类型形参(type parameter) 和 类型实参(type argument)，如下：
```go
// 假设 T 是类型形参，在定义函数时它的类型是不确定的，类似占位符
func Add(a T, b T) T {  
    return a + b
}
```
在上面这段伪代码中， T 被称为 类型形参(type parameter)， 它不是具体的类型，在定义函数时类型并不确定。因为 T 的类型并不确定，所以我们需要像函数的形参那样，在调用函数的时候再传入具体的类型。这样我们不就能一个函数同时支持多个不同的类型了吗？在这里被传入的具体类型被称为 类型实参(type argument):

下面一段伪代码展示了调用函数时传入类型实参的方式：
```go
// [T=int]中的 int 是类型实参，代表着函数Add()定义中的类型形参 T 全都被 int 替换
Add[T=int](100, 200)  
// 传入类型实参int后，Add()函数的定义可近似看成下面这样：
func Add( a int, b int) int {
    return a + b
}

// 另一个例子：当我们想要计算两个字符串之和的时候，就传入string类型实参
Add[T=string]("Hello", "World") 
// 类型实参string传入后，Add()函数的定义可近似视为如下
func Add( a string, b string) string {
    return a + b
}
```
通过引入 类型形参 和 类型实参 这两个概念，我们让一个函数获得了处理多种不同类型数据的能力，这种编程方式被称为 泛型编程。

可能你会已奇怪，我通过Go的 接口+反射 不也能实现这样的动态数据处理吗？是的，泛型能实现的功能通过接口+反射也基本能实现。但是使用过反射的人都知道反射机制有很多问题：
1. 用起来麻烦
2. 失去了编译时的类型检查，不仔细写容易出错
3. 性能不太理想

而在泛型适用的时候，它能解决上面这些问题。但这也不意味着泛型是万金油，泛型有着自己的适用场景，当你疑惑是不是该用泛型的话，请记住下面这条经验：

> 如果你经常要分别为不同的类型写完全相同逻辑的代码，那么使用泛型将是最合适的选择



## 2. Go的泛型

通过上面的伪代码，我们实际上已经对Go的泛型编程有了最初步也是最重要的认识—— 类型形参 和 类型实参。而Go1.18也是通过这种方式实现的泛型，但是单纯的形参实参是远远不能实现泛型编程的，所以Go还引入了非常多全新的概念：

- 类型形参 (Type parameter)
- 类型实参(Type argument)
- 类型形参列表( Type parameter list)
- 类型约束(Type constraint)
- 实例化(Instantiations)
- 泛型类型(Generic type)
- 泛型接收器(Generic receiver)
- 泛型函数(Generic function)

等等等等。

啊，实在概念太多了头晕？没事请跟着我慢慢来，首先从 **泛型类型(generic type)** 讲起



## 3. 类型形参、类型实参、类型约束和泛型类型

```go
type IntSlice []int

var a IntSlice = []int{1, 2, 3} // 正确
var b IntSlice = []float32{1.0, 2.0, 3.0} // ✗ 错误，因为IntSlice的底层类型是[]int，浮点类型的切片无法赋值
```
这里定义了一个新的类型 IntSlice ，它的底层类型是 []int ，理所当然只有int类型的切片能赋值给 IntSlice 类型的变量。

接下来如果我们想要定义一个可以容纳 float32 或 string 等其他类型的切片的话该怎么办？很简单，给每种类型都定义个新类型：
```go
type StringSlice []string
type Float32Slie []float32
type Float64Slice []float64
```
但是这样做的问题显而易见，它们结构都是一样的只是成员类型不同就需要重新定义这么多新类型。那么有没有一个办法能只定义一个类型就能代表上面这所有的类型呢？答案是可以的，这时候就需要用到泛型了：
```go
type Slice[T int|float32|float64 ] []T
```
不同于一般的类型定义，这里类型名称 Slice 后带了中括号，对各个部分做一个解说就是：

- `T` 就是上面介绍过的**类型形参(Type parameter)**，在定义Slice类型的时候 T 代表的具体类型并不确定，类似一个占位符
- `int|float32|float64` 这部分被称为**类型约束(Type constraint)**，中间的 `|` 的意思是告诉编译器，类型形参 T 只可以接收 int 或 float32 或 float64 这三种类型的实参
- 中括号里的 `T int|float32|float64` 这一整串因为定义了所有的类型形参(在这个例子里只有一个类型形参T），所以我们称其为 **类型形参列表(type parameter list)**
- 这里新定义的类型名称叫 `Slice[T]`

这种类型定义的方式中带了类型形参，很明显和普通的类型定义非常不一样，所以我们将这种

> 类型定义中带 **类型形参** **的类型，称之为** 泛型类型(Generic type)**

泛型类型不能直接拿来使用，必须传入**类型实参(Type argument)** 将其确定为具体的类型之后才可使用。而传入类型实参确定具体类型的操作被称为 **实例化(Instantiations)** ：

```go
// 这里传入了类型实参int，泛型类型Slice[T]被实例化为具体的类型 Slice[int]
var a Slice[int] = []int{1, 2, 3}  
fmt.Printf("Type Name: %T",a)  //输出：Type Name: Slice[int]

// 传入类型实参float32, 将泛型类型Slice[T]实例化为具体的类型 Slice[string]
var b Slice[float32] = []float32{1.0, 2.0, 3.0} 
fmt.Printf("Type Name: %T",b)  //输出：Type Name: Slice[float32]

// ✗ 错误。因为变量a的类型为Slice[int]，b的类型为Slice[float32]，两者类型不同
a = b  

// ✗ 错误。string不在类型约束 int|float32|float64 中，不能用来实例化泛型类型
var c Slice[string] = []string{"Hello", "World"} 

// ✗ 错误。Slice[T]是泛型类型，不可直接使用必须实例化为具体的类型
var x Slice[T] = []int{1, 2, 3} 
```
对于上面的例子，我们先给泛型类型 Slice[T] 传入了类型实参 int ，这样泛型类型就被实例化为了具体类型 Slice[int] ，被实例化之后的类型定义可近似视为如下：
```go
type Slice[int] []int     // 定义了一个普通的类型 Slice[int] ，它的底层类型是 []int
```
我们用实例化后的类型 Slice[int] 定义了一个新的变量 a ，这个变量可以存储int类型的切片。之后我们还用同样的方法实例化出了另一个类型 Slice[float32] ，并创建了变量 b 。

因为变量 a 和 b 就是具体的不同类型了(一个 Slice[int] ，一个 Slice[float32]），所以 a = b 这样不同类型之间的变量赋值是不允许的。

同时，因为 Slice[T] 的类型约束限定了只能使用 int 或 float32 或 float64 来实例化自己，所以 Slice[string] 这样使用 string 类型来实例化是错误的。

上面只是个最简单的例子，实际上类型形参的数量可以远远不止一个，如下：
```go
// MyMap类型定义了两个类型形参 KEY 和 VALUE。分别为两个形参指定了不同的类型约束
// 这个泛型类型的名字叫： MyMap[KEY, VALUE]
type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE  

// 用类型实参 string 和 flaot64 替换了类型形参 KEY 、 VALUE，泛型类型被实例化为具体的类型：MyMap[string, float64]
var a MyMap[string, float64] = map[string]float64{
    "jack_score": 9.6,
    "bob_score":  8.4,
}
```

用上面的例子重新复习下各种概念的话：

- KEY和VALUE是**类型形参**
- `int|string` 是KEY的**类型约束**， `float32|float64` 是VALUE的**类型约束**
- `KEY int|string, VALUE float32|float64` 整个一串文本因为定义了所有形参所以被称为**类型形参列表**
- Map[KEY, VALUE] 是**泛型类型**，类型的名字就叫 Map[KEY, VALUE]
- `var a MyMap[string, float64] = xx` 中的string和float64是**类型实参**，用于分别替换KEY和VALUE，**实例化**出了具体的类型 `MyMap[string, float64]`

还有点头晕？没事，的确一下子有太多概念了，这里用一张图就能简单说清楚：

![img](https://segmentfault.com/img/bVcYRiO)

### 3.1 其他的泛型类型

所有类型定义都可使用类型形参，所以下面这种结构体以及接口的定义也可以使用类型形参：

```go
// 一个泛型类型的结构体。可用 int 或 sring 类型实例化
type MyStruct[T int | string] struct {  
    Name string
    Data T
}

// 一个泛型接口(关于泛型接口在后半部分会详细讲解）
type IPrintData[T int | float32 | string] interface {
    Print(data T)
}

// 一个泛型通道，可用类型实参 int 或 string 实例化
type MyChan[T int | string] chan T
```

### 3.2 类型形参的互相套用

类型形参是可以互相套用的，如下

```go
type WowStruct[T int | float32, S []T] struct {
    Data     S
    MaxValue T
    MinValue T
}
```

这个例子看起来有点复杂且难以理解，但实际上只要记住一点：任何泛型类型都必须传入类型实参实例化才可以使用。所以我们这就尝试传入类型实参看看：

```go
var ws WowStruct[int, []int]
// 泛型类型 WowStuct[T, S] 被实例化后的类型名称就叫 WowStruct[int, []int]
```

上面的代码中，我们为T传入了实参 `int`，然后因为 S 的定义是 `[]T` ，所以 S 的实参自然是 `[]int` 。经过实例化之后 WowStruct[T,S] 的定义类似如下：

```go
// 一个存储int类型切片，以及切片中最大、最小值的结构体
type WowStruct[int, []int] struct {
    Data     []int
    MaxValue int
    MinValue int
}
```

因为 S 的定义是 []T ，所以 T 一定决定了的话 S 的实参就不能随便乱传了，下面这样的代码是错误的：

```go
// 错误。S的定义是[]T，这里T传入了实参int, 所以S的实参应当为 []int 而不能是 []float32
ws := WowStruct[int, []float32]{
        Data:     []float32{1.0, 2.0, 3.0},
        MaxValue: 3,
        MinValue: 1,
    }
```

### 3.3 几种语法错误

1. 定义泛型类型的时候，**基础类型不能只有类型形参**，如下：

   ```go
   // 错误，类型形参不能单独使用
   type CommonType[T int|string|float32] T
   ```
   
2. 当类型约束的一些写法会被编译器误认为是表达式时会报错。如下：

   ```go
   //✗ 错误。T *int会被编译器误认为是表达式 T乘以int，而不是int指针
   type NewType[T *int] []T
   // 上面代码再编译器眼中：它认为你要定义一个存放切片的数组，数组长度由 T 乘以 int 计算得到
   type NewType [T * int][]T 
   
   //✗ 错误。和上面一样，这里不光*被会认为是乘号，| 还会被认为是按位或操作
   type NewType2[T *int|*float64] []T 
   
   //✗ 错误
   type NewType2 [T (int)] []T 
   ```

   为了避免这种误解，解决办法就是给类型约束包上 `interface{}` 或加上逗号消除歧义（关于接口具体的用法会在后半篇提及）

   ```go
   type NewType[T interface{*int}] []T
   type NewType2[T interface{*int|*float64}] []T 
   
   // 如果类型约束中只有一个类型，可以添加个逗号消除歧义
   type NewType3[T *int,] []T
   
   //✗ 错误。如果类型约束不止一个类型，加逗号是不行的
   type NewType4[T *int|*float32,] []T 
   ```

   因为上面逗号的用法限制比较大，这里推荐统一用 interface{} 解决问题

### 3.4 特殊的泛型类型

这里讨论种比较特殊的泛型类型，如下：

```go
type Wow[T int | string] int

var a Wow[int] = 123     // 编译正确
var b Wow[string] = 123  // 编译正确
var c Wow[string] = "hello" // 编译错误，因为"hello"不能赋值给底层类型int
```

这里虽然使用了类型形参，但因为类型定义是 `type Wow[T int|string] int` ，所以无论传入什么类型实参，实例化后的新类型的底层类型都是 int 。所以int类型的数字123可以赋值给变量a和b，但string类型的字符串 “hello” 不能赋值给c

这个例子没有什么具体意义，但是可以让我们理解泛型类型的实例化的机制

### 3.5 泛型类型的套娃

泛型和普通的类型一样，可以互相嵌套定义出更加复杂的新类型，如下：

```go
// 先定义个泛型类型 Slice[T]
type Slice[T int|string|float32|float64] []T

// ✗ 错误。泛型类型Slice[T]的类型约束中不包含uint, uint8
type UintSlice[T uint|uint8] Slice[T]  

// ✓ 正确。基于泛型类型Slice[T]定义了新的泛型类型 FloatSlice[T] 。FloatSlice[T]只接受float32和float64两种类型
type FloatSlice[T float32|float64] Slice[T] 

// ✓ 正确。基于泛型类型Slice[T]定义的新泛型类型 IntAndStringSlice[T]
type IntAndStringSlice[T int|string] Slice[T]  
// ✓ 正确 基于IntAndStringSlice[T]套娃定义出的新泛型类型
type IntSlice[T int] IntAndStringSlice[T] 

// 在map中套一个泛型类型Slice[T]
type WowMap[T int|string] map[string]Slice[T]
// 在map中套Slice[T]的另一种写法
type WowMap2[T Slice[int] | Slice[string]] map[string]T
```

### 3.6 类型约束的两种选择

观察下面两种类型约束的写法

```go
type WowStruct[T int|string] struct {
    Name string
    Data []T
}

type WowStruct2[T []int|[]string] struct {
    Name string
    Data T
}
```

仅限于这个例子，这两种写法和实现的功能其实是差不多的，实例化之后结构体相同。但是像下面这种情况的时候，我们使用前一种写法会更好：

```go
type WowStruct3[T int | string] struct {
    Data     []T
    MaxValue T
    MinValue T
}
```

### 3.7 匿名结构体不支持泛型	

我们有时候会经常用到匿名的结构体，并在定义好匿名结构体之后直接初始化：

```go
testCase := struct {
        caseName string
        got      int
        want     int
    }{
        caseName: "test OK",
        got:      100,
        want:     100,
    }
```

那么匿名结构体能不能使用泛型呢？答案是不能，下面的用法是错误的：

```go
testCase := struct[T int|string] {
        caseName string
        got      T
        want     T
    }[int]{
        caseName: "test OK",
        got:      100,
        want:     100,
    }
```

所以在使用泛型的时候我们只能放弃使用匿名结构体，对于很多场景来说这会造成麻烦（最主要麻烦集中在单元测试的时候，为泛型做单元测试会非常麻烦，这点我之后的文章将会详细阐述）



## 4. 泛型receiver

看了上的例子，你一定会说，介绍了这么多复杂的概念，但好像泛型类型根本没什么用处啊？

是的，单纯的泛型类型实际上对开发来说用处并不大。但是如果将泛型类型和接下来要介绍的泛型receiver相结合的话，泛型就有了非常大的实用性了

我们知道，定义了新的普通类型之后可以给类型添加方法。那么可以给泛型类型添加方法吗？答案自然是可以的，如下：

```go
type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
    var sum T
    for _, value := range s {
        sum += value
    }
    return sum
}
```

这个例子为泛型类型 `MySlice[T]` 添加了一个计算成员总和的方法 `Sum()` 。注意观察这个方法的定义：

- 首先看receiver `(s MySlice[T])` ，所以我们直接把类型名称 `MySlice[T]` 写入了receiver中
- 然后方法的返回参数我们使用了类型形参 T ****(实际上如果有需要的话，方法的接收参数也可以实用类型形参)
- 在方法的定义中，我们也可以使用类型形参 T （在这个例子里，我们通过 `var sum T` 定义了一个新的变量 `sum` )

对于这个泛型类型 `MySlice[T]` 我们该如何使用？还记不记得之前强调过很多次的，泛型类型无论如何都需要先用类型实参实例化，所以用法如下：

```go
var s MySlice[int] = []int{1, 2, 3, 4}
fmt.Println(s.Sum()) // 输出：10

var s2 MySlice[float32] = []float32{1.0, 2.0, 3.0, 4.0}
fmt.Println(s2.Sum()) // 输出：10.0
```

该如何理解上面的实例化？首先我们用类型实参 int 实例化了泛型类型 `MySlice[T]`，所以泛型类型定义中的所有 T 都被替换为 int，最终我们可以把代码看作下面这样：

```go
type MySlice[int] []int // 实例化后的类型名叫 MyIntSlice[int]

// 方法中所有类型形参 T 都被替换为类型实参 int
func (s MySlice[int]) Sum() int {
    var sum int 
    for _, value := range s {
        sum += value
    }
    return sum
}
```

用 float32 实例化和用 int 实例化同理，此处不再赘述。

通过泛型receiver，泛型的实用性一下子得到了巨大的扩展。在没有泛型之前如果想实现通用的数据结构，诸如：堆、栈、队列、链表之类的话，我们的选择只有两个：

- 为每种类型写一个实现
- 使用 接口+反射

而有了泛型之后，我们就能非常简单地创建通用数据结构了。接下来用一个更加实用的例子 —— 队列来讲解

### 4.1 基于泛型的队列

队列是一种先入先出的数据结构，它和现实中排队一样，数据只能从队尾放入、从队首取出，先放入的数据优先被取出来

```go
// 这里类型约束使用了空接口，代表的意思是所有类型都可以用来实例化泛型类型 Queue[T] (关于接口在后半部分会详细介绍）
type Queue[T interface{}] struct {
    elements []T
}

// 将数据放入队列尾部
func (q *Queue[T]) Put(value T) {
    q.elements = append(q.elements, value)
}

// 从队列头部取出并从头部删除对应数据
func (q *Queue[T]) Pop() (T, bool) {
    var value T
    if len(q.elements) == 0 {
        return value, true
    }

    value = q.elements[0]
    q.elements = q.elements[1:]
    return value, len(q.elements) == 0
}

// 队列大小
func (q Queue[T]) Size() int {
    return len(q.elements)
}
```

为了方便说明，上面是队列非常简单的一种实现方法，没有考虑线程安全等很多问题

```go
var q1 Queue[int]  // 可存放int类型数据的队列
q1.Put(1)
q1.Put(2)
q1.Put(3)
q1.Pop() // 1
q1.Pop() // 2
q1.Pop() // 3

var q2 Queue[string]  // 可存放string类型数据的队列
q2.Put("A")
q2.Put("B")
q2.Put("C")
q2.Pop() // "A"
q2.Pop() // "B"
q2.Pop() // "C"

var q3 Queue[struct{Name string}] 
var q4 Queue[[]int] // 可存放[]int切片的队列
var q5 Queue[chan int] // 可存放int通道的队列
var q6 Queue[io.Reader] // 可存放接口的队列
// ......
```

### 4.2 动态判断变量的类型

使用接口的时候经常会用到类型断言或 type swith 来确定接口具体的类型，然后对不同类型做出不同的处理，如：

```go
var i interface{} = 123
i.(int) // 类型断言

// type switch
switch i.(type) {
    case int:
        // do something
    case string:
        // do something
    default:
        // do something
    }
}
```

那么你一定会想到，对于 `valut T` 这样通过类型形参定义的变量，我们能不能判断具体类型然后对不同类型做出不同处理呢？答案是不允许的，如下：

```go
func (q *Queue[T]) Put(value T) {
    value.(int) // 错误。泛型类型定义的变量不能使用类型断言

    // 错误。不允许使用type switch 来判断 value 的具体类型
    switch value.(type) {
    case int:
        // do something
    case string:
        // do something
    default:
        // do something
    }
    
    // ...
}
```

虽然type switch和类型断言不能用，但我们可通过反射机制达到目的：

```go
func (receiver Queue[T]) Put(value T) {
    // Printf() 可输出变量value的类型(底层就是通过反射实现的)
    fmt.Printf("%T", value) 

    // 通过反射可以动态获得变量value的类型从而分情况处理
    v := reflect.ValueOf(value)

    switch v.Kind() {
    case reflect.Int:
        // do something
    case reflect.String:
        // do something
    }

    // ...
}
```

这看起来达到了我们的目的，可是当你写出上面这样的代码时候就出现了一个问题：

> 你为了避免使用反射而选择了泛型，结果到头来又为了一些功能在在泛型中使用反射

当出现这种情况的时候你可能需要重新思考一下，自己的需求是不是真的需要用泛型（毕竟泛型机制本身就很复杂了，再加上反射的复杂度，增加的复杂度并不一定值得）

当然，这一切选择权都在你自己的手里，根据具体情况斟酌

## 5. 泛型函数

在介绍完泛型类型和泛型receiver之后，我们来介绍最后一个可以使用泛型的地方——泛型函数。有了上面的知识，写泛型函数也十分简单。假设我们想要写一个计算两个数之和的函数：

```go
func Add(a int, b int) int {
    return a + b
}
```

这个函数理所当然只能计算int的和，而浮点的计算是不支持的。这时候我们可以像下面这样定义一个泛型函数：

```go
func Add[T int | float32 | float64](a T, b T) T {
    return a + b
}
```

上面就是泛型函数的定义。

> 这种带类型形参的函数被称为**泛型函数**

它和普通函数的点不同在于函数名之后带了类型形参。这里的类型形参的意义、写法和用法因为与泛型类型是一模一样的，就不再赘述了。

和泛型类型一样，泛型函数也是不能直接调用的，要使用泛型函数的话必须传入类型实参之后才能调用。

```go
Add[int](1,2) // 传入类型实参int，计算结果为 3
Add[float32](1.0, 2.0) // 传入类型实参float32, 计算结果为 3.0

Add[string]("hello", "world") // 错误。因为泛型函数Add的类型约束中并不包含string
```

或许你会觉得这样每次都要手动指定类型实参太不方便了。所以Go还支持类型实参的自动推导：

```go
Add(1, 2)  // 1，2是int类型，编译请自动推导出类型实参T是int
Add(1.0, 2.0) // 1.0, 2.0 是浮点，编译请自动推导出类型实参T是float32
```

自动推导的写法就好像免去了传入实参的步骤一样，但请记住这仅仅只是编译器帮我们推导出了类型实参，实际上传入实参步骤还是发生了的。

### 5.1 匿名函数不支持泛型

在Go中我们经常会使用匿名函数，如：

```go
fn := func(a, b int) int {
    return a + b 
}  // 定义了一个匿名函数并赋值给 fn 

fmt.Println(fn(1, 2)) // 输出: 3
```

那么Go支不支持匿名泛型函数呢？答案是不能——**匿名函数不能自己定义类型形参：**

```go
// 错误，匿名函数不能自己定义类型实参
fnGeneric := func[T int | float32](a, b T) T {
        return a + b
} 

fmt.Println(fnGeneric(1, 2))
```

但是匿名函数可以使用别处定义好的类型实参，如：

```go
func MyFunc[T int | float32 | float64](a, b T) {

    // 匿名函数可使用已经定义好的类型形参
    fn2 := func(i T, j T) T {
        return i*2 - j*2
    }

    fn2(a, b)
}
```

### 5.2 既然支持泛型函数，那么泛型方法呢？

既然函数都支持泛型了，那你应该自然会想到，方法支不支持泛型？很不幸，目前Go的方法并不支持泛型，如下：

```go
type A struct {
}

// 不支持泛型方法
func (receiver A) Add[T int | float32 | float64](a T, b T) T {
    return a + b
}
```

但是因为receiver支持泛型， 所以如果想在方法中使用泛型的话，目前唯一的办法就是曲线救国，迂回地通过receiver使用类型形参：

```go
type A[T int | float32 | float64] struct {
}

// 方法可以使用类型定义中的形参 T 
func (receiver A[T]) Add(a T, b T) T {
    return a + b
}

// 用法：
var a A[int]
a.Add(1, 2)

var aa A[float32]
aa.Add(1.0, 2.0)
```



## 前半小结

讲完了泛型类型、泛型receiver、泛型函数后，Go的泛型算是介绍完一半多了。在这里我们做一个概念的小结：

1. Go的泛型(或者或类型形参)目前可使用在3个地方
   1. 泛型类型 - 类型定义中带类型形参的类型
   2. 泛型receiver - 泛型类型的receiver
   3. 泛型函数 - 带类型形参的函数
2. 为了实现泛型，Go引入了一些新的概念：
   1. 类型形参
   2. 类型形参列表
   3. 类型实参
   4. 类型约束
   5. 实例化 - 泛型类型不能直接使用，要使用的话必须传入类型实参进行实例化

什么，这文章已经很长很复杂了，才讲了一半？是的，Go这次1.18引入泛型为语言增加了较大的复杂度，目前还只是新概念的介绍，下面后半段将介绍Go引入泛型后对接口做出的重大调整。那么做好心理准备，我们出发吧。
