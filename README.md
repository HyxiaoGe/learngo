# learngo

## 变量
### 如何定义变量
#### 单声明变量

var名称类型是声明单个变量的语法。
1. 指定变量类型，声明后若不赋值，使用默认值
   var name string
   name = "张三"
2. 根据值自行判定变量类型（类型推断 Type inference）
   如果一个变量有一个初始值，Go将自动通过初始值来推断该变量的类型。因此，如果变量具有初始值，则是可以省略变量声明种的类型。
   var name = "张三"
3. 省略 var, 注意 := 左侧的变量不应该是已经声明过的（多个变量同时声明时，至少保证一个是新变量），否则会导致编译错误（简短声明）
   var a int = 10
   var b = 10
   c := 10
   需要注意的是：:= 这种方式只能用于函数体内，不能用在全局变量的声明与赋值
   
#### 多声明变量
1. 以逗号分隔，声明与赋值分开，若不赋值，存在默认值
```
var name1, name2, name3 string
name1, name2, name3 = "张三", "李四", "王五"
```
2. 直接赋值，变量类型可以是不同的类型
```
var name, age, score = "张三", 12, 100
```
3. 集合类型
var (
  name string
  age int
  score float32
)
***
注意：
* 变量必须先定义才能使用
* go 是静态语言，要求变量的类型和赋值的类型必须一致
* 变量名不能冲突（同一个作用域内不能冲突）
* := 不能用于全局变量
* 变量的零值也叫默认值
* 只要定义了变量，就要使用，否则无法通过编译

#### 匿名变量
在使用多重赋值时候，如果不需要在左值中接收变量，可以使用匿名变量 _ 下划线代替
```
a := [3]int{5, 10, 15}
for _, x := range a {
  fmt.Println(x)
}
// output:
// 5  
// 10
// 15
```
## 常量
常量是一个简单值的标识符，在程序运行时，不会被修改的量。用 const 修饰符来修饰
```
// 显式类型定义： 
const b string = "abc"
// 隐式类型定义： 
const b = "abc"
```
也可以作为枚举，常量组
```
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```
***
注意：
* 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
* 不曾使用的常量，在编译的时候，是不会报错的

## iota
iota，特殊常量，默认是int类型，可以认为是一个可以被编译器修改的常量，类似于计数器的原理
可以被用作枚举值：
```
const (
    Unknown = iota
    Female = iota
    Male = iota
)
```
第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1，因此Unknown = 0，Female = 1，Male = 2；也可以简写为如下形式：
```
const (
    Unknown = iota
    Female
    Male
)
```
***
```
const (
    Unknown = iota
    Female
    Male
)
const (
    a = iota
    b
    c
)
```
注意，这里的 a 是从0开始了，因为每次 const 出现时，都会让 iota 初始化为0.

## 基本数据类型
### bool 类型
布尔型的值只可以是常量 true 或者 false
```
var b bool = true
```
### 数值型
* int8 有符号 8 位整型 (-128 到 127) 长度：8bit
* int16 有符号 16 位整型 (-32768 到 32767)
* int32 有符号 32 位整型 (-2147483648 到 2147483647)
* int64 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
* uint8 无符号 8 位整型 (0 到 255) 8位都用于表示数值：
* uint16 无符号 16 位整型 (0 到 65535)
* uint32 无符号 32 位整型 (0 到 4294967295)
* uint64 无符号 64 位整型 (0 到 18446744073709551615)
### 浮点型
* float32  32位浮点型数
* float64  64位浮点型数
### 其他
* byte 等于 uint8
* rune 等于 int32
* uint  32 或 64 位
### 字符
Goland 中没有专门的字符类型；如果要存储单个字符（字母），一般使用byte来保存。字符串一般是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。也就是说对于传统的字符串是由字符组成的，而Go的字符串不同，它是由字节组成的。
```
	var a byte
	a = 'a'
	fmt.Println(a) // 输出的是ascii对应码值，即 97
```

## 数据类型的转换
### 简单的数值转换
valueOfTypeB = typeB(valueOfTypeA)

```
// 浮点数
a := 5.0

// 转换为int类型
b := int(a)
fmt.Println(a, b)

var c int32 = 1
var d int64 = 3
d = int64(c)
fmt.Println(c, d)

```
需要注意的是
* 不是所有数据类型都能转换的，例如字母格式的string类型"abcd"转换为int肯定会失败
* 低精度转换为高精度时是安全的，高精度的值转换为低精度时会丢失精度。例如int32转换为int16，float32转换为int
* 这种简单的转换方式不能对int(float)和string进行互转，要跨大类型转换，可以使用strconv包提供的函数
### strconv
#### Itoa 和 Atoi
Itoa
```
fmt.Println("I'm " + strconv.Itoa(18) + " year old") // I'm 18 year old
```
Atoi
```
i, _ := strconv.Atoi("5")
fmt.Println(5 + i)  // 10
```
Atoi(int, error) 这个函数有两个返回值：第一个返回值是转换成int的值，第二个返回值判断是否转换成功

### if语句
go语言的常用控制流程有if和for，没有while， 而switch和goto是为了简化代码，降低重复代码，属于扩展的流程控制。
```
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
}

if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
} else {
  /* 在布尔表达式为 false 时执行 */
}

if 布尔表达式1 {
   /* 在布尔表达式1为 true 时执行 */
} else if 布尔表达式2{
   /* 在布尔表达式1为 false ,布尔表达式2为true时执行 */
} else{
   /* 在上面两个布尔表达式都为false时，执行*/
}
```
除此之外，还有这一种也常用到
```
// 这种写法可 以将返回值与判断放在一行进行处理，而且返回值的作用范围被限制在if、 else 语句组合中。
if err := doSomethoing(); err != nil {
    
}
```
提示：在编程中，变量在其实现了变量的功能后，作用范围越小，所造成的问题可能性越小，每一个变量代表一个状态，有状态的地方，状态就会被修改，函数的局部变量只会影响一个函数的执行，但全局变量可能会影响所有代码的执行状态，因此限制变量的作用范围对代码的稳定性有很大的帮助。

### for循环
Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。
```
//和 C 语言的 for 一样：
for init; condition; post { }
//和 C 的 while 一样：
for condition { }
//和 C 的 for(;;) 一样：
for { }
```
* init： 一般为赋值表达式，给控制变量赋初值；
* condition： 关系表达式或逻辑表达式，循环控制条件；
* post： 一般为赋值表达式，给控制变量增量或减量。

for 循环语句执行过程
1. 执行 init 代码，先赋初值；
2. 判别 condition，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post；
3. 进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

除此之外，for 循环还有另一种写法
for ... range 格式可以对 map、切片、数组、字符串等进行迭代循环。格式如下：
```
for key, value := range map {
    map[key] = value
}
//
numbers := [6]int{1, 2, 3, 4, 5, 6}
for index, num:= range numbers {
   	fmt.Printf("第 %d 位 x 的值 = %d\n", index, num)
} 
```
### switch语句
switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止
```
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```
### 数组
Go 语言提供了数组类型的数据结构。 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。

数组元素可以通过索引（位置）来读取（或者修改），索引从0开始，第一个元素索引为 0，第二个索引为 1，以此类推。数组的下标取值范围是从0开始，到长度减1。

数组一旦定义后，大小不能更改。

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。

数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。

#### 声明和初始化数组
需要指明数组的大小和存储的数据类型。
```
var array1 [10] float32
var array2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
var array3 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```
初始化数组中 {} 中的元素个数不能大于 [] 中的数字。如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

您甚至可以忽略声明中数组的长度并将其替换为...让编译器为你找到长度。这是在下面的程序中完成的。

数组的其他创建方式：
```
var a [4] float32 // 等价于：var arr2 = [4]float32{}
fmt.Println(a) // [0 0 0 0]
var b = [5] string{"ruby", "王二狗", "rose"}
fmt.Println(b) // [ruby 王二狗 rose  ]
var c = [5] int{'A', 'B', 'C', 'D', 'E'} // byte
fmt.Println(c) // [65 66 67 68 69] ascii码
d := [...] int{1,2,3,4,5}// 根据元素的个数，设置数组的大小
fmt.Println(d)//[1 2 3 4 5]
e := [5] int{4: 100} // [0 0 0 0 100]
fmt.Println(e)
f := [...] int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
fmt.Println(f)
```
#### 数组的长度
通过将数组作为参数传递给 len 函数，可以获得数组的长度。
```
var array2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
fmt.Println("length of a is",len(array2))
```
#### 数组是值类型
数组是值类型 Go中的数组是值类型，而不是引用类型。这意味着当它们被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映。

数组的大小是类型的一部分。因此[5]int和[25]int是不同的类型。因此，数组不能被调整大小。

## 切片
### 什么是切片
Go 数组的长度不可变，在特定场景中这样的集合就不太适用，因此Go提供了一种灵活、功能强悍的内置类型切片（”动态数组“），与数组相比，切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大

数组的传递是值传递 切片是引用传递
### 定义切片
```
//  第一种
var identifier []type
//  第二种
var slice1 []type = make([]type, len)
//  也可简写成
slice1 := make([]type, len)

//  第三种，通过对数组操作返回
course := [5]string{"Java", "Python", "C++", "PHP", "Go"}
subCourse := course[1:2]    //  [1:2] 左闭右开区间，即 取第二个值
fmt.Println(subCourse) // [Python]

// 第四种，new 
subCourse2 := new([]int)
fmt.Println(subCourse2)
```
### 切片是引用传递
```
oldArr := [3]int{1, 2, 3}
newArr := oldArr
newArr[0] = 5
fmt.Println(newArr, oldArr) // [5 2 3] [1 2 3]
// ----------------------------------------------------------------
oldSlice := []int{1, 2, 3}
newSlice := oldSlice
newSlice[0] = 5
fmt.Println(oldSlice, newSlice) // [5 2 3] [5 2 3]
```
### 切片的 append 方法
```
newCourse := []string{"Java", "Python", "C++", "Go", "PHP"}
appendCourse := []string{"Vue", "React", "angular"}
newCourse = append(newCourse, appendCourse...)  //  函数的传递规则
fmt.Println(newCourse) // [Java Python C++ Go PHP Vue React angular]
```
### 切片如何实现新切片的赋值
```
//  拷贝的时候 目标对象长度需要设置好
oldCourse := []string{"Java", "Python", "C++", "Go", "PHP"}
newCourse := make([]string, len(oldCourse))
copy(newCourse, oldCourse)
```
### 切片如何删除元素
```
deleteCourses := [6]string{"Java", "Python", "Vue", "C++", "Go", "PHP"}
courseSlice := deleteCourses[:]
courseSlice = append(courseSlice[:2], courseSlice[3:]...)
fmt.Println(courseSlice) // [Java Python C++ Go PHP]
```
### 切片的扩容机制
```
a := []int{1, 2, 3}
b := a[:]
b = append(b, 4)
b[0] = 5
fmt.Println(a)  // [1 2 3]
fmt.Println(b)  // [5 2 3 4]
```
由上述代码可以看出，append函数没有影响到原来的数组；这是因为切片 b 产生扩容机制，扩容机制一旦产生，这个时候切片 b 就会指向新的内存地址。
### 切片的容量 cap
cap指的是容量，跟长度 len 不是一个概念，
```
data := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
slice := data[2:4]
newSlice := data[3:6]
fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))  //  len=2, cap=8
fmt.Printf("len=%d, cap=%d\n", len(newSlice), cap(newSlice))    //  len=3, cap=7
```
//  TODO 解释
```
oldSlice := make([]int, 6, 12)
newSlice := oldSlice[:]
newSlice = append(newSlice, 1)

newSlice[0] = 1
fmt.Printf("oldSlice: %v\n", oldSlice)	// oldSlice: [1 0 0 0 0 0]
fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice)) // len=6, cap=12 
fmt.Printf("newSlice: %v\n", newSlice) // newSlice: [1 0 0 0 0 0 1]
fmt.Printf("len=%d, cap=%d\n", len(newSlice), cap(newSlice)) // len=7, cap=12 
fmt.Printf("----------------------------------------------------------------\n")
newSlice = append(newSlice, 2)
newSlice = append(newSlice, 3)
newSlice = append(newSlice, 4)
newSlice = append(newSlice, 5)
newSlice = append(newSlice, 6)
newSlice = append(newSlice, 7)
// 此时触发扩容机制，newSlice 重新分配内存，这就会导致 newSlice 和 oldSlice 不再共享内存
// 所以对 newSlice 做任何修改，对 oldSlice 不会再起作用
newSlice[0] = 2
fmt.Printf("oldSlice: %v\n", oldSlice) // oldSlice: [1 0 0 0 0 0]
fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice)) // len=6, cap=12
fmt.Printf("newSlice: %v\n", newSlice) // newSlice: [2 0 0 0 0 0 1 2 3 4 5 6 7]
fmt.Printf("len=%d, cap=%d\n", len(newSlice), cap(newSlice)) // len=13, cap=24
```
Go 中切片扩容的策略是这样的：
如果旧切片的长度小于1024，则最终容量(newCap)就是旧容量(oldCap)的两倍，即（newCap = oldCap * 2）

如果旧切片长度大于等于1024，则最终容量（newCap）从旧容量（oldCap）开始循环增加原来的 1/4，即 (newCap = oldCap + 1/4*oldCap) = ( oldCap * 1.25)

## map
要求所有的key的数据类型相同，所有value数据类型相同(注：key与value可以有不同的数据类型)
### map创建
```
//1. 直接声明
m1 := map[string]string{
	"name": "zhangsan",
}
fmt.Printf("%v\n", m1)  //map[name:zhangsan]

//2. make函数
m2 := make(map[string]int) //创建时，里面不需要添加元素
m2["age"] = 18
fmt.Printf("%v\n", m2)  //  map[age:18]
```
### map中key的类型
map中的每个key在keys的集合中是唯一的，而且需要支持 == or != 操作

key的常用类型：int, rune, string, 结构体(每个元素需要支持 == or != 操作), 指针, 基于这些类型自定义的类型

### map的增删查改
```
m := map[string]string{
	"a": "va",
	"b": "vb",
	"d":"",
}

//  1. 进行增加，修改
m["c"] = "vc"
m["b"] = "vb1"
fmt.Printf("%v\n", m)

//  查询，你返回空的字符串到底是没有获取到还是值本身就是这样空字符串呢
v, ok := m["d"]
if ok {
	fmt.Println("找到了", v)
}else{
	fmt.Println("没找到")
}
fmt.Println(v, ok)  //  true

//  删除
delete(m, "a")

//  遍历
for k, v := range m {
	fmt.Println(k, v)
}
```

## 指针
### 什么是指针
指针是一种特殊的变量类型，这个变量只能保存地址值
指针的默认值是 nil，所以可通过判断指针地址是否为 nil 来确认有没有赋值
指针变量中涉及到两个符号 &(取地址) 和 * (取值)
```
if ip != nil {
   ...
}
```

```
a := 10
var num *int 	// num 变量里面只能存放地址类型的值
num = &a	// 将变量a的地址传递给num
... 
// 修改指针所执行的变量的值 
*num = 20
```

下面通过完整的例子来讲解
```
func swap(a *int, b *int){
     //用于交换a和b
     c := *a
     *a = *b
     *b = c
}

func main() {
     a := 10
     b := 20
     swap(&a, &b)
     fmt.Println(a, b) // 20 10
}
```
将a，b两个变量的地址作为参数来调用 swap 函数，以此来达到交换 a 和 b 的值，由于在swap 函数里面，是对 a 和 b 两个变量地址的值做修改，同时也会影响到在 main 函数中声明的 a 和 b，所以在执行 swap(&a, &b) 完成之后，a 和 b 的值也正常完成了交换。

对于指针来说，或者说其他的默认值是0的情况来说，可以通过new函数来声明并分配内存。
当 go 的编译器看到new之后，就明白需要预先申请一个内存空间，并且将内存中的值全部设置为 0，new函数返回的是这个值的地址
如果只是单单声明变量，没有声明new函数来申请内存空间的话，那么该变量还是会占用内存空间，只不过占用的是 nil 
```
var p *int = new(int)
*p = 10
```

## 函数
### 定义函数

函数的几个要素： 1. 函数名 2. 参数 3. 返回值
下面来看看常用的函数定义方法
```
// 第一种
func add(a, b int) int {
     var sum int
     sum = a + b
     return sum
}

//第二种
func add2(a, b int) (sum int) {
     sum = a + b
     return sum
}

//第三种
func add3(a, b int) (sum int) {
	sum = a + b
	return
}

//第4种
func div2(a, b int) (result int,err error) {
     if b == 0 {
     	err = errors.New("被除数不能为0")
     }else{
	result = a / b
     }

     return
}
```

### 匿名函数

```
fmt.Println(func (a, b int) int{
	return a+b 
}(1,2))
```

### 函数的参数类型

可以使用 省略号 来起到可变参数的作用
```
func add(params ...int) (sum int){
     for _, v := range params {
	sum += v
     }
     params[0] = 9
     return
}
```

函数也可以当做参数传递给一个函数
```
var mySub sub = func (a, b int) int {
	return a - b
}
fmt.Println(mySub(1,2))
```

```
score := []int{10, 50, 80, 90, 85}
// 写一个函数过滤掉不合格的成绩
fmt.Println(filter(score, func (a int) bool{
	if a>=90{
		return true
	}else {
		return false
	}
}))

func filter(score []int, f func(int) bool) []int{
	reSlice := make([]int, 0)
	for _, v := range score {
		if f(v) {
			reSlice = append(reSlice, v)
		}
	}
	return reSlice
}
```

### defer函数

defer 语句是go语言用的一种用于注册延迟调用的机制，可以让当前函数执行完成之后再执行
```
func main() {
	fmt.Println("test1")
	//defer之后只能是函数调用 不能是表达式 比如 a++
	defer fmt.Println("defer test1")
	defer fmt.Println("defer test2")
	defer fmt.Println("defer test3")
	//此处有大量的逻辑需要读取文件
	fmt.Println("test2")
}
```
// 输出
test1
test2
defer test3
defer test2
defer test1

```
func main()  {
     //defer语句执行时的拷贝机制
     test := func () {
     	fmt.Println("test1")
     }
     defer test()
     test = func () {
     	fmt.Println("test2")
     }
     fmt.Println("test3")
}
```
// 输出
test3
test1

defer本质上是注册了一个延迟函数，defer函数的执行顺序已经确定，类似于栈的数据结构，先进后出

## 结构体
Go 语言不是面向对象的语言，它里面不存在类的概念，结构体正是类的替代品。
- 结构体内部变量的大小写，首字母大写是公开变量，首字母小写是内部变量，分别相当于Java的类成员变量的 Public 和 Private 类别。内部变量只有属于同一个 package（简单理解就是同一个目录）的代码才能直接访问。
- 结构体的方法只能和结构体在同一个包中
- 结构体是值传递类型
- 结构体的接收者有两种形式 值接收者和指针接收者
```
type Course struct {
	Name  string
	Price int
	Url   string
}

func (c Course) printCourseInfo() {
	fmt.Printf("课程名:%s, 课程价格: %d, 课程的地址:%s", c.Name, c.Price, c.Url)
}

func (c *Course) setPrice(price int) {
	c.Price = price
}

func main() {

	// 1. 实例化- kv形式
	var c Course = Course{
		Name:  "Go从入门到放弃",
		Price: 9.9,
		Url:   "https://go.dev/",
	}

	//2. 第二种实例化方式 - 顺序形式
	c2 := Course{"Go从入门到放弃", 9.9, "https://go.dev/"}
	fmt.Println(c2.Name, c2.Price, c2.Url)

	//3. 如果一个指向结构体的指针, 通过结构体指针获取对象的值
	c3 := &Course{"Go从入门到放弃", 9.9, "https://go.dev/"}
	//fmt.Printf("%T\n", c3)
	fmt.Println((*c3).Name, (*c3).Price, (*c3).Url) // 等同于下面 
	fmt.Println(c3.Name, c3.Price, c3.Url)  //这里其实是go语言的一个语法糖 go语言内部会将c3.Name转换成 (*c3).Name

	//4. 零值 如果不给结构体赋值， go语言会默认给每个字段采用默认值
	c4 := Course{}
	fmt.Println(c4.Price)

	//5. 多种方式零值初始结构体
	var c5 Course = Course{}
	var c6 Course
	var c7 *Course = new(Course)
	var c8 *Course = &Course{}

	fmt.Println(c5.Price, c6.Price, c7.Price, c8.Price)

	//6. 结构体是值传递类型
	c9 := Course{"scrapy", 110, "https://go.dev/"}
	c10 := c9
	fmt.Println(c9)
	fmt.Println(c10)
	c8.Price = 200
	fmt.Println(c9)
	fmt.Println(c10)
```

### 结构体的大小

结构体可以使用sizeof来查看对象占用内存的长度

```
	// int 占用8个长度
	fmt.Println(unsafe.Sizeof(1))   // 8

	// go语言string的本质 其实string是个结构体
	//type string struct {
	//	Data uintptr //指针占8个长度
	//	Len  int     //长度64位系统8个长度
	//}
	fmt.Println(unsafe.Sizeof("scrapy"))    // 16

	//8.slice的大小
	type slice struct {
		array unsafe.Pointer // 底层数组的地址 8个长度
		len   int            // 长度 8个长度
		cap   int            // 容量 8个长度
	}

	s1 := []string{"scrapy", "django", "tornado", "flask", "beego", "gin"}
	fmt.Println(unsafe.Sizeof(s1)) // 24

    // map 固定占用8个长度，跟kv无关
	m1 := map[string]string{
		"bobby1": "django",
		"bobby2": "tornado",
		"bobby3": "scrapy",
		"bobby4": "celery",
	}
	fmt.Println(unsafe.Sizeof(m1))  // 8

```
#### 零值结构体和 nil 结构体

nil 结构体是指结构体指针变量没有指向一个实际存在的内存。这样的指针变量只会占用 1 个指针的存储空间，也就是一个机器字的内存大小。（如果只是单单声明变量，没有声明new函数来申请内存空间的话，那么该变量还是会占用内存空间，只不过占用的是 nil ）
而零值结构体是会实实在在占用内存空间的，只不过每个字段都是零值。如果结构体里面字段非常多，那么这个内存空间占用肯定也会很大。

## 继承

Go 里面没有继承的概念，但是也可以达到伪继承的效果。结构体作为一种变量它可以放进另外一个结构体作为一个字段来使用，这种内嵌结构体的形式在 Go 语言里称之为「组合」。
下面我们来看看内嵌结构体的基本使用方法

```
type Teacher struct {
	Name  string
	Age   int
	Title string
}

func (t *Teacher) PrintInfo() {
	println(t.Name, t.Age, t.Title)
}

type Course struct {
	Teacher *Teacher
	Name    string
	Price   int
	Url     string
}

func (c Course) courseInfo() {
	fmt.Printf("课程名:%s, 价格:%d, 链接:%s, 讲师信息：%s %d %s", c.Name, c.Price, c.Url, c.Teacher.Name, c.Teacher.Age, c.Teacher.Title)
}

func main() {

	t := Teacher{"zhangsan", 18, "程序员"}
	c := Course{Teacher: &t, Name: "django", Price: 100, Url: "https://go.dev/"}
	c.courseInfo()
}
```

## 多态

Go 语言不是面向对象语言在于它的结构体不支持多态，它不能算是一个严格的面向对象语言。多态是指父类定义的方法可以调用子类实现的方法，不同的子类有不同的实现，从而给父类的方法带来了多样的不同行为。

## 接口
在go语言中接口是一种类型，是一种抽象类型
- 类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。
- 实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
- 一个类型可以实现多个接口。

代码实现
```
type Programmer interface {
	Coding() string //方法只是申明
	Debug() string
}
type Designer interface {
	Design() string
}
type Manger interface {
	Programmer
	Designer
	Manage() string
}

```

### 总结

在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以根据具体类型 P 直接辨识的：

- 指针方法可以通过指针调用
- 值方法可以通过值调用
- 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
- 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
- 将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。

### 译注

Go 语言规范定义了接口方法集的调用规则：

- 类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
- 类型 T 的可调用方法集包含接受者为 T 的所有方法
- 类型 T 的可调用方法集不包含接受者为 *T 的方法

### 空接口

空接口或最小接口 不包含任何方法，它对实现不做任何要求：

```
type Any interface {}
```

空接口类似于 Java 里所有类的基类 Object 类，可以给一个空接口类型的变量 var val interface {} 赋任何类型的值

```
var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func main() {

	var val Any
	val = i
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)

	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)

	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}

}
```

## 反射

### 方法和类型的反射

反射是用程序检查其所拥有的结构，尤其是类型的一种能力；这是元编程的一种形式。反射可以在运行时检查类型和变量，例如：他的大小、它的方法以及它能”动态地“调用这些方法。这对于没有源代码的包尤其有用。

变量的最基本信息就是类型和值：反射包的 Type 用来表示一个 Go 类型，反射包的 Value 为 Go 值提供了反射接口。

两个简单的函数，reflect.TypeOf 和 reflect.ValueOf，返回被检查对象的类型和值。例如，x被定义为：var x float64 = 3.4，那么reflect.TypeOf(x) 返回 float64，reflect.ValueOf(x) 返回 <float64 Value>

实际上，反射是通过检查一个接口的值，变量首先被转换成空接口。

```
func TypeOf(i interface{}) Type
func Value(i interface{}) Value
```

接口的值包含一个 type 和 value

反射可以从接口值反射到对象，也可以从对象反射回接口值

### 通过反射修改（设置）值

```
func main() {

	var x float64 = 3.14
	v := reflect.ValueOf(x)
	//v.SetFloat(1.0) //panic: reflect: reflect.Value.SetFloat using unaddressable value

	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	v = v.Elem()
	fmt.Println("settability of v:", v.CanSet())
	fmt.Println("The Elem of v is:", v)
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
}
```

以上代码当修改 v 的值时，会报 reflect.Value.SetFloat using unaddressable value ；
类似于Java里修改 private 修饰的变量，要想通过反射来修改它，只能特定的方法来开启安全检查的开关，Java 是 field.setAccessible(true);
go 里面是用 Elem() 来开启，就可以正常地赋值了

### 反射结构

当需要反射一个结构类型。NumField() 方法返回结构内的字段数量；通过一个 for 循环用索引取得每个字段的值 Field(i)。

```
type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)

	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
	results := value.Method(0).Call(nil)
	fmt.Println(results)

}
```

## 总结：Go中的面向对象

Go 没有类，而是松耦合的类型、方法对接口的实现
- 封装（数据隐藏）：包范围内的：通过标识符首字母小写，对象只在它所在的包内可见；包范围外的，即可以直接import的，通过标识符首字母大写，对象所在包以外也可见
- 继承：用组合实现：内嵌一个（或多个）包含想要的行为（字段和方法）的类型；多重继承可以通过内嵌多个类型实现
- 多态：用接口实现：某个类型的实例可以赋给它所实现的任意接口类型的变量。类型和接口是松耦合的，并且多重继承可以通过实现多个接口实现。Go 接口不是 Java 和 C# 接口的变体，而且接口间是不相关的，并且是大规模编程和可适应的演进型设计的关键。

## goroutine

### 协程
在 Go 中，应用程序并发处理的部分被称作 goroutines(协程)，它可以进行更有效的并发运算。在协程和操作系统线程之间并无一对一的关系：协程是根据一个或多个线程的可用性，映射（多路复用，执行于）在他们之上的；协程调度器在 Go 运行时很好的完成了这个工作。

Go 使用 channels 来同步协程

当系统调用（比如等待I/O）堵塞协程时，其他协程会继续在其他线程上工作。协程的设计隐藏了许多线程创建和管理方面的复杂工作。

协程可以运行在多个操作系统线程之间，也可以运行在线程之内，让你可以用很小的内存占用就可以处理大量的任务。

协程是通过使用关键字 go 调用（执行）一个函数或者方法来实现（也可以匿名或者lambda函数）。这样会在当前的计算过程中开始一个同时进行的函数，在相同的地址空间中并且分配了独立的栈，比如：go sum(num) 在后台计算总和

### 通道 channel

Go 有一种特殊的类型，通道（channel）就像一个可以用于发送类型化数据的管道，由其负责协程之间的通信；这种通过信道进行通信的方式保证了同步性。数据在通道中进行传递：在任何给定时间，一个数据被设计为只有一个协程可以对其访问，所以不会发生数据竞争。

声明通道：var identifier chan datatype，未初始化的通道的是 nil。

通道只能传输一种类型的数据，比如 chan int 或者 chan string，所有的类型都可以用于通道，空接口 interface{} 也可以。

通道实际上是类型化消息的队列：使数据得以传输。它是先进先出（FIFO）的结构所以可以保证发送给他们的元素的顺序。通道也是引用类型，所以我们使用 make() 函数来给他分配内存。

这里先声明了一个字符串通道 ch1，然后使用 make 创建了它（实例化）:

```
var ch1 chan string
ch1 = make(chan string)
或者更简短点的
ch1 := make(chan string)
再或者是函数通道
funcChan := make(chan func())
```
#### 通信指示符 <-

这个操作符直观的表示了数据的传输：信息按照箭头的方向流动

流向通道（发送）

ch <- int1 表示：用通道 ch 发送变量 int1 （双目运算符）

int2 := <- ch 表示：定义一个变量 int2，并从通道 ch 接受数据（获取新值）

<- ch 可以单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证，所以以下代码是合法的：
```
if <- ch != 1000{
...
}
```
- 同一个操作符 <- 即用于发送也用于接收
- 通道的发送和接收都是原子操作，互不干扰

### 通道堵塞

默认情况下，通信是同步且无缓冲的：在有接收者接收数据之前，发送不会结束。

一个无缓冲的通道在没有空间来保存数据的时候，必须要一个接收者准备好接收通道的数据，然后发送者可以直接把数据发送给接收者。所以通道的发送/接收操作在对方准备好之前都是堵塞的：

1. 对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是堵塞的：如果 channel 中的数据无人接收，就无法再给通道传入其他数据；新的输入无法在通道非空的情况下传入。
2. 对于同一个通道，接收操作是堵塞的（协程或者函数中的），直到发送者可用；如果通道中没有数据，接收者就堵塞了。

### 通过一个（或多个）通道交换数据进行协程同步。

通道是一种同步形式，通过通道，两个协程在通信（协程会合）中某刻同步交换数据。无缓冲通道成为了多个协程同步的完美工具。

甚至可以在通道两端互相堵塞对方，形成了 deadlock（死锁）的状态。

无缓冲通道会被堵塞，可以通过带缓冲的通道来避免死锁。

### 同步通道-使用带缓冲的通道

可以在扩展的 make 命令中设置它的容量，如下
```
buf := 100
ch1 := make(chan string, buf)
```
buf 是通道可以同时容纳的元素个数

在缓冲满载（缓冲被全部使用）之前，给一个带缓冲的通道发送数据是不会堵塞的，而从通道读取数据也不会堵塞，直到缓冲空了。 

缓冲容量和类型无关，所以可以给一些通道设置不同的容量，只要他们拥有同样的元素类型。内置的 cap() 函数可以返回缓冲区的容量。

如果容量大于 0，通道就是异步的了：缓冲满载（发送）或变空（接收）之前通信不会阻塞，元素会按照发送的顺序被接收。如果容量是 0 或者未设置，通信仅在收发双方准备好的情况下才可以成功。

下面看下通道不为0的时候，延迟发送数据
```

import (
"fmt"
"time"
)

func getData(ch chan int) {
    input := <-ch
    fmt.Println(input)
}

func main() {
    ch1 := make(chan int, 1)
    ch1 <- 1
    go getData(ch1)
    time.Sleep(time.Second * 5)
}
```

#### 协程的同步-关闭通道

通道可以被显式的关闭；尽管它们和文件不同，不必每次都关闭。只有在需要告诉接收者不会再提供新的值的时候，才需要关闭通道。只有发送者需要关闭通道，接收者不需要。

可以通过函数 close(ch) 来完成，将通道标记为无法通过发送操作 <- 接收更多的值，给已经关闭的通道发送或者再次关闭都会导致运行时的 panic()。
```
ch := make(chan float64)
defer close(ch)
```

可以通过逗号 ok模式用来检测通道是否被关闭
通常和 if 语句一起使用
```
if v, ok := <- ch; ok{
   process(v)
}
```
或者在for循环中接收的时候，使用break来控制
```
v, ok := <- ch
if !ok {
   break
}
process(v)
```
检测通道当前是否堵塞，需要使用select
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
### 使用 select 切换协程

从不同的并发执行的协程中获取值可以通过关键字 select 来完成，它和 switch 控制语句非常相似也称作通信开关；它的行为类似于轮询机制；select 监听进入通道的数据，也可以是用通道发送值的时候。
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
在任何一个 case 中执行 break 或者 return，select 就结束了。

select 做的是：选择处理列出的多个通信情况中的一个。

- 如果都堵塞了，会等待直到其中一个可以处理
- 如果多个可以选择，随机选择一个
- 如果没有通道操作可以处理并且写了 default 语句，可以保证发送不被堵塞！如果没有 default，select 就会一直堵塞。

select 语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过 break 语句使循环退出

以下代码中，有 2 个通道 ch1 和 ch2，三个协程 pump1()、pump()2 和 suck1()。这是一个典型的生产者消费者模式。在无限循环中，ch1 和 ch2 通过 pump1() 和 pump()2 填充整数；suck1 也是在无限循环中轮询输入的，通过 select 语句获取 ch1 和 ch2 的整数并输出。
选择哪一个 case 取决于哪一个通道收到了信息。
```
package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck1(ch1, ch2)

	time.Sleep(10 * time.Second)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck1(ch1, ch2 chan int) {
	for true {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}

```

## TCP Socket

一个 Web 服务器应用需要响应众多客户端的并发请求，Go 会为每一个客户端产生一个协程来处理请求。

需要用到 net 包中网络通信的功能，包含了处理 TCP/IP 以及 UDP 协议、域名解析等办法。

server 代码
```
import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server...")
	// 创建 listener
	listener, err := net.Listen("tcp", `localhost:7890`)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		return // 终止程序
	}
	// 监听并接收来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error listening: ", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		content, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error listening: ", err.Error())
			return // 终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:content]))
	}
}
```
在 main() 中创建了一个 net.Listen 类型的变量 listener，它实现了服务器的基本功能：用来监听和接收来自客户端的请求（基于 TCP 协议下，位于 IP 地址为 127.0.0.1，端口为 50000 的 localhost）。

Listen() 函数可以返回一个error类型的错误变量。用一个无限 for 循环的listener.Accept() 来等待客户端的请求。客户端的请求将产生一个 net.Conn 类型的连接变量。然后一个独立的协程使用这个连接执行 doServerStuff(),
开始使用一个 512 字节的缓冲 data 来读取客户端发送来的数据，并且把它们打印到服务器的终端，len() 获取客户端发送的数据字节数；当客户端发送的所有数据都被读取完时，协程就结束了。

这段程序会为每一个客户端连接创建一个独立的协程。必须先运行服务器代码，再运行客户端代码。

----------------------------------------------------------------

client 代码

```
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	// 打开连接
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	fmt.Printf("CLIENTNAME %s", clientName)
	trimmedClient := strings.Trim(clientName, "\r\n")
	// 给服务器发送信息直到程序退出：
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		fmt.Printf("input:--%s--", input)
		fmt.Printf("trimmedInput:--%s--", trimmedInput)
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says：" + trimmedInput))
	}

}

```

客户端通过 net.Dial() 创建了一个和服务器之间的连接。

通过无限循环从 os.Stin 接收来自键盘的输入，直到输入了"Q"。裁剪后的输入被 connection 的 Write() 方法发送到服务器。

当然，服务器必须先启动好，如果服务器并未开始监听，客户端是无法成功连接的。

在网络编程中 net.Dial() 函数是非常重要的，一旦你连接到远程系统，函数就会返回一个 Conn 类型的接口，我们可以用它发送和接收数据。Dial() 函数简洁地抽象了网络层和传输层。

所以不管是 IPv4 还是 IPv6，TCP 或者 UDP 都可以使用这个公用接口。

## 访问并读取页面数据

发送一个简单的 http.Head() 请求查看返回值

```
import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.baidu.com",
	"https://weibo.com/",
	"https://www.taobao.com/",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error: ", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
```

使用 http.Get() 获取并显示网页内容；Get() 返回值中的 Body 属性包含了网页内容，然后我们用 ioutil.ReadAll() 把它读出来：

```
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://www.baidu.com")
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got: %q", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
```

## 写一个简单的网页应用

下边的程序在端口 8088 上启动了一个网页服务器；SimpleServer() 会处理 url/test1 使它在浏览器输出 hello world。FormServer 会处理
url/test2；如果 url 最初由浏览器请求，那么它是一个 GET 请求，返回一个 form 常量，包含了简单的 input 表单，这个表单里有一个文本框
和一个提交按钮。当在文本框输入一些东西并点击提交按钮的时候，会发起一个 POST 请求。FormServer() 中的代码用到了 switch 来区分两种情况。
请求为 POST 类型时，name 属性为 inp 的文本框的内容可以这样获取: request.FormValue("inp")。然后将其写回浏览器页面中。在控制台启动
程序，然后到浏览器中打开 url http://localhost:8088/test2 来测试这个程序：

```
import (
	"io"
	"net/http"
)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case http.MethodGet:
		io.WriteString(w, form)
	case http.MethodPost:
		io.WriteString(w, request.FormValue("in"))
	}
}

func main() {

	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)

	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}

}
```

## Gin 

### 介绍

Gin 是一个用Go(Golang)编写的Web框架。它具有类似的 martini 的API，性能要好得多，多亏了 httprouter ，速度提高了 40 倍。

### 特性

- 快速，基于 Radix 树（压缩前缀树）的路由，小内存占用。没有反射。可预测的API性能
- 支持中间件，传入的 HTTP 请求可以由一系列中间件和最终操作来处理。例如：Logger，Authorization，GZIP，最终操作DB。
- Crash 处理，Gin 可以 catch 一个发生在 HTTP 请求中的 panic 并 recover 它。保证服务器一直可用。
- JSON 验证，Gin 可用解析并验证请求的 JSON，例如检查所需值的存在。
- 路由组，更好地组织路由。是否需要授权，不同的 API 版本...此外，这些组可以无限制地嵌套而不会降低性能。
- 错误处理，Gin 提供了一种方便的方法来收集HTTP请求期间发生的所有错误。最终，中间件可以将他们写入日志文件，数据库并通过网络发送。
- 内置渲染，Gin 为JSON，XML和HTML渲染提供了易于使用的API。

### 开始

```
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// "github.com/gin-gonic/gin"
func main() {

    // Default 使用 Logger 和 Recovery 中间件
	router := gin.Default()
    // gin.H 是 map[string]interface{} 的一种快捷方式
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 默认在 0.0.0.0:8080 上启动服务并监听
	router.Run()

}
```

### JSONP

使用 JSONP 向不同域的服务器请求数据。如果查询参数存在回调，则将回调添加到响应体中。

```
func main() {

	router := gin.Default()
	router.GET("/JSONP", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		context.JSON(http.StatusOK, data)
	})

	router.Run()
}
```
### Multipart/Urlencoded 绑定

```
func main() {

	router := gin.Default()
	router.POST("/login", func(context *gin.Context) {
		// 可以使用显式绑定声明绑定 multipart form
		// c.shouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定
		var form LoginForm
		// 在这种情况下，将自动选择合适地绑定
		if context.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.Run()
}
```
### Multipart/Urlencoded 表单

```
func main() {

	router := gin.Default()

	router.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")

		context.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run()
}
```
### XML/JSON/YAML 渲染

```
func main() {

	router := gin.Default()

	router.GET("/someJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hey", "status": http.StatusOK,
		})
	})

	router.GET("/moreJSON", func(context *gin.Context) {

		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		context.JSON(http.StatusOK, msg)
	})

	router.GET("/someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.Run()

}
```
### 上传文件

#### 单文件

```
func main() {

	router := gin.Default()
	// 为 multipart form 设置较低的内存限制（默认是 32MiB）
	router.MaxMultipartMemory = 8 << 20 // 8MiB
	router.POST("/upload", func(context *gin.Context) {
		// 单文件
		file, _ := context.FormFile("file")
		log.Println(file.Filename)

		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		context.SaveUploadedFile(file, dst)

		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})

	router.Run()

}
```
#### 多文件

```
func main() {

	router := gin.Default()
	// 为 multipart form 设置较低的内存限制（默认是 32MiB）
	router.MaxMultipartMemory = 8 << 20 // 8MiB
	router.POST("/upload", func(context *gin.Context) {
		// multipart form
		form, _ := context.MultipartForm()
		files := form.File["files[]"]

		for _, file := range files {
			log.Println(file.Filename)

			dst := "./" + file.Filename
			// 上传文件至指定目录
			context.SaveUploadedFile(file, dst)
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))
	})

	router.Run()

}
```
### 使用 BasicAuth 中间件
```
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {

	router := gin.Default()
	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是map[string]string 的一种快捷方式
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 ”localhost:8080/admin/secrets“
	authorized.GET("/secrets", func(context *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := context.MustGet(gin.AuthUserKey).(string)
		if secrets, ok := secrets[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": secrets})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRETS :("})
		}
	})

	router.Run(":8083")
}
```

### 使用 HTTP 方法
```
func main() {

	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run()
	// router.Run(":8081") hardcode 端口号
}
```
### 使用中间件
```
func main() {

	// 新建一个没有任何默认中间件的路由
	router := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置成 release
	router.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic，如果有 panic 的话，会写入 500
	router.Use(gin.Recovery())

	// 你甚至为每个路由添加任意数量的中间件
	//router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 认证路由组
	// authorized := router.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样
	//authorized := router.Group("/")
	// 路由组中间件
	// AuthRequired() 中间件
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)
	//
	//	// 嵌套路由组
	//	testing := authorized.Group("testing")
	//	testing.GET("/analytics", analyticsEndpoint)
	//}

	router.Run()
}
```
### 只绑定 url 查询字符串

ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据
```
type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()
	router.Any("/testing", startPage)
	router.Run(":8083")
}

func startPage(context *gin.Context) {
	var person Person
	if context.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	context.String(http.StatusOK, "Success")
}
```

请求示例：http://127.0.0.1:8083/testing?name=xieguangkun&address=xiangyashan

### 记录日志
```
func main() {
	// 记录到文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	router.Run(":8083")
}
```

### 映射查询字符串或表单参数
```
func main() {

	router := gin.Default()

	router.POST("/post", func(context *gin.Context) {

		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	router.Run()
}
```
### 映射查询字符串或表单参数
```
func main() {

	router := gin.Default()

	// 使用现有的基础请求对象解析查询字符串参数
	// 示例 URL: /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "guest")
		lastname := context.Query("lastname") // context.URL.Query("lastname") 的一种快捷方式

		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run()
}
```
### 模型绑定和验证

要将请求体绑定到结构体中，使用模型绑定。Gin 目前支持 JSON、XML、YAML 和标准表单值的绑定（foo=bar&boo=baz）

使用时，需要在要绑定的所有字段上，设置响应的tag。例如，使用 JSON 绑定时，设置字段标签为 json:"fieldname"。

Gin 提供了两类绑定方法：

- Type - Must bind
    - Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML
    - Behavior - 这些方法属于 MustBindWith 的具体调用。如果发生绑定错误，则请求终止，并触发 c.AbortWithError(400, err).SetType(ErrorTypeBind)。响应状态码被设置为 400 并且 Content-Type 被设置为 text/plain; charset=utf-8。如果您在此之后尝试设置响应状态码，Gin 会输出日志 [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422。如果你希望更好地控制绑定，考虑使用 ShouldBind 等方法。
- Type - Should bind
    - Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAMl
    - Behavior - 这些方法属于 ShouldBindWith 的具体调用。如果发生绑定错误，Gin 会返回错误并由开发者处理错误和请求。

使用 Bind 方法时，Gin 会尝试根据 Content-Type 推断如何绑定。如果你明确知道要绑定什么，可以使用 MustBindWith 或 ShouldBindWith。

你也可以指定必须绑定的字段。如果一个字段的 tag 加上了 binding:"required"，但绑定时是空值，Gin 会报错。
```
func main() {

	router := gin.Default()

	// 绑定 JSON({"user": "admin", "password": "123456"})
	router.POST("/loginJSON", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if json.User != "admin" || json.Password != "123456" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "you are logger in"})
	})

	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginForm", func(context *gin.Context) {
		var form Login
		// 根据 Content-Type Header 推断使用哪个绑定器
		if err := context.ShouldBind(&form); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "admin" || form.Password != "123456" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "you are logger in"})
	})

	router.Run()
}
```
### 绑定 Uri

```
func main() {

	router := gin.Default()
	router.GET("/:name/:id", func(context *gin.Context) {
		var person Person
		if err := context.ShouldBindUri(&person); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})
	router.Run()

}
```
### 绑定查询字符串或表单数据
```
type Person struct {
	Name     string `form:"name"`
	Address  string `form:"address"`
	Birthday string `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {

	router := gin.Default()
	router.GET("/testing", startPage)
	router.Run()

}

func startPage(context *gin.Context) {
	var person Person
	// 如果是 GET 请求，只能使用 Form 绑定引擎
	// 如果是 POST 请求，首先检查 content-type 是否为 JSON 或者 XML，然后再使用 Form（form-data）
	if context.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	context.JSON(http.StatusOK, "Success")
}
```
### 绑定表单数据至自定义结构体
```
type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		Field string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(context *gin.Context) {
	var b StructB
	context.Bind(&b)
	context.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(context *gin.Context) {
	var c StructC
	context.Bind(&c)
	context.JSON(http.StatusOK, gin.H{
		"a": c.NestedStructPointer,
		"c": c.FieldC,
	})
}

func GetDataD(context *gin.Context) {
	var d StructD
	context.Bind(&d)
	context.JSON(http.StatusOK, gin.H{
		"a": d.NestedAnonyStruct,
		"d": d.FieldD,
	})
}

func main() {

	router := gin.Default()
	router.GET("/getb", GetDataB)
	router.GET("/getc", GetDataC)
	router.GET("/getd", GetDataD)

	router.Run()
}
```

使用 curl 命令结果：

$ curl "http://localhost:8080/getb?field_a=hello&field_b=world"
{"a":{"FieldA":"hello"},"b":"world"}

$ curl "http://localhost:8080/getc?field_a=hello&field_c=world"
{"a":{"FieldA":"hello"},"c":"world"}

$ curl "http://localhost:8080/getd?field_x=hello&field_d=world"
{"d":"world","x":{"FieldX":"hello"}}

### 自定义 HTTP 配置

直接使用 http.ListenAndServe()
```
func main() {
    router := gin.Default()
    http.ListenAndServe(":8080", router)
}
```
或
```
func main() {
    router := gin.Default()
    
    s := &http.Server{
        Addr: "http://localhost:8080",
        Handler: router,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
```
### 自定义中间件
```
func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		// 设置 example 变量
		context.Set("example", "12345")
		// 请求前
		context.Next()
		// 请求后
		latency := time.Since(t)
		log.Print("latency: ", latency)

		// 获取发送的 status
		status := context.Writer.Status()
		log.Print("status: ", status)
	}
}

func main() {

	router := gin.New()
	router.Use(Logger())

	router.GET("/test", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)

		// 打印：”12345“
		log.Println("example", example)
	})

	router.Run()
}
```

### 设置和获取 Cookie
```
func main() {

	router := gin.Default()
	router.GET("/cookie", func(context *gin.Context) {

		cookie, err := context.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			context.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}
```
### 路由组

```
func main() {

	router := gin.Default()

	// 简单的路由组：v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndPoint)
		v1.POST("/submit", submitEndPoint)
		v1.POST("/read", readEndPoint)
	}

	// 简单的路由组：v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndPoint)
		v2.POST("/submit", submitEndPoint)
		v2.POST("/read", readEndPoint)
	}

	router.Run()
}
```

### 运行多个服务

```
var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "Welcome server 01",
		})
	})
	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "Welcome server 02",
		})
	})
	return e
}

func main() {

	server01 := &http.Server{
		Addr:              ":8080",
		Handler:           router01(),
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	server02 := &http.Server{
		Addr:              ":8081",
		Handler:           router02(),
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
```
### 重定向

HTTP 重定向很容易。内部、外部重定向均支持。
```
router.GET("/redirect", func(context *gin.Context) {
	context.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
})
```

通过 POST 方法进行 HTTP 重定向。
```
router.POST("/test", func(c *gin.Context) {
	c.Redirect(http.StatusFound, "/foo")
})
```

路由内部重定向，使用 HandleContext:
```
router.GET("/test", func(context *gin.Context) {
	context.Request.URL.Path = "/test2"
	router.HandleContext(context)
})
router.GET("/test2", func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"hello": "world"})
})
```

### 运行多个服务

```
func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
```

## 性能提示

- 如果需要把数字转换成字符串，使用 strconv.Itoa() 比 fmt.Sprintf() 要快一倍左右。
- 尽可能避免把 String 转换成 []Byte，这个转换会导致性能下降。
- 如果在 for-loop 里对某个 Slice 使用 append()，请先把 Slice 的容量扩充到位，这样可以避免内存重新分配以及系统自动按 2 的N次方幂进行扩展但又用不到的情况，从而避免浪费内存。
- 使用 StringBuffer 或是 StringBuild 来拼接字符串，性能会比使用 + 或 += 高三到四个数量级。
- 尽可能使用并发的 goroutine，然后使用 sync.WaitGroup 来同步分片操作。
- 避免在热代码中进行内存分配，这样会导致 gc 很忙。尽可能使用 sync.Pool 来重用对象。
- 使用 lock-free 的操作，避免使用 mutex，尽可能使用 sync/Atomic 包
- 使用 I/O 缓冲，I/O 是个非常慢的操作，使用 bufio.NewWrite() 和 bufio.NewReader() 可以带来更高的性能
- 对于在 for-loop 里的固定的正则表达式，一定要使用 regexp.Compile() 编译正则表达式。性能会提升两个数量级
- 如果需要更高性能的协议，就要考虑使用 protobuf 或 msgp 而不是 JSON，因为JSON的序列化和反序列化里使用了反射。
- 在使用Map的时候，使用整型的key会比字符串的要快，因为整型比较 会比 字符串比较要快

## 错误处理

Go 语言的函数支持多返回值，所以，可以在返回接口把业务语义（业务返回值）和控制语义（出错返回值）区分开。Go 语言的很多函数都会返回 result、err 两个值，于是就有以下几点：

- 参数上基本上就是入参，而返回接口就是把结果和错误分离，这样使得函数的接口语义清晰；
- 而且，Go 语言中的错误参数如果要忽略，需要显式地忽略，用 _ 这样的变量来忽略；
- 另外，因为返回的是 error 是个接口（其中只有一个方法 Error()，返回一个string），所以你可以扩展自定义的错误处理。

另外，如果一个函数返回了多个不同类型的 error，你也可以使用下面这样的方式：

```
if err != nil {
  switch err.(type) {
    case *json.SyntaxError:
      ...
    case *ZeroDivisionError:
      ...
    case *NullPointerError:
      ...
    default:
      ...
  }
}
```

























