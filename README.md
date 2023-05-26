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
oldSlice := make([]int, 0)
fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))    // len=0, cap=0
oldSlice = append(oldSlice, 1)
fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))    // len=1, cap=1
oldSlice = append(oldSlice, 2)
fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))    // len=2, cap=2
oldSlice = append(oldSlice, 3)
fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))    // len=3, cap=4
oldSlice = append(oldSlice, 4)
oldSlice = append(oldSlice, 5)
oldSlice = append(oldSlice, 6)
oldSlice = append(oldSlice, 7)
fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))    // len=7, cap=8
oldSlice = append(oldSlice, 8)
oldSlice = append(oldSlice, 9)
fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))    // len=7, cap=8
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
























