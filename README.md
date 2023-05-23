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




































































