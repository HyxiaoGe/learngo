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
test
Go 语言不是面向对象语言在于它的结构体不支持多态，它不能算是一个严格的面向对象语言。多态是指父类定义的方法可以调用子类实现的方法，不同的子类有不同的实现，从而给父类的方法带来了多样的不同行为。

## 接口
在go语言中接口是一种类型，是一种抽象类型

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

### 空接口

```
type Course struct {
name string
price int
url string
}
type Printer interface {
printInfo() string
}

func (c Course) printInfo() string{
return "课程信息"
}

//func print(x interface{}){
//	if v, ok := x.(int); ok{
//		fmt.Printf("%d(整数)\n", v)
//	}
//	if s, ok := x.(string); ok {
//		fmt.Printf("%s(字符串)\n", s)
//	}
//	//牵扯到go的另一个默认的问题
//	//fmt.Printf("%v\n", i)
//}

type AliOss struct {

}

type LocalFile struct {

}

func store(x interface{}){
switch v := x.(type) {
case AliOss:
//此处要做一些特殊的处理，我设置阿里云的权限问题
fmt.Println(v)
case LocalFile:
//检查路径的权限
fmt.Println(v)
}
}

func print(x interface{}){
switch v := x.(type) {
case string:
fmt.Printf("%s(字符串)\n", v)
case int:
fmt.Printf("%d(整数)\n", v)
}
}


func main()  {
//空接口
var i interface {} //空接口
//空接口可以类似于我们java和python中的object
//i = Course{}
//fmt.Println(i)
i = 10
print(i)
i = "bobby"
print(i)
i = []string{"django", "scrapy"}
print(i)
//空接口的第一个用途 可以把任何类型都赋值给空接口变量
//2. 参数传递
//3. 空接口可以作为map的值
var teacherInfo = make(map[string]interface{})
teacherInfo["name"] = "bobby"
teacherInfo["age"] = 18
teacherInfo["weight"] = 75.2
teacherInfo["courses"] = []string{"django", "scrapy", "sanic"}
fmt.Printf("%v", teacherInfo)
//类型断言
//接口的一个坑, 接口引入了
// 接口有一个默认的规范  接口的名称一般以 er结尾
//c := &Course{}
//var c Printer = Course{}
//c.printInfo()

}

```













