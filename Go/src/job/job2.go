package main

import (
	"errors"
	"fmt"
	"os"
)

// PersonInfo  是一个包含个人详细信息的类型
type PersonInfo struct {
	ID string
	Name string
	Address string
}

//var v1 int
//var v2 string
//var v3 [10]int  //数组
//var v4 []int    //数组切片
////结构体
//var v5 struct{
//	f int
//	b string
//}
//var v6 *int //指针
//var v7 map[string]int //map,key为string类型，value为int类型
//var v8 func(a int) int   //函数
//var(
//	v10 int
//	v11 string
//)

func main() {
	//2.12  变量初始化
	var v1 int = 10 // 正确方式式1
	var v2 = 10  // 正确方式2
	v3 := 10  // 正确方式3
	fmt.Println(v1,v2,v3)

	//2.1.3 变量赋值
	//下面这个交换 i 和 j 变量
	//i, j = j, i

	//2.1.4  匿名变量
	_, _, nickName := GetName()
	fmt.Println(nickName)

	//2.2  常量
	//2.2.1 字面常量

//	2.2.2 常量定义
//	通过 const 关键字，你可以给字面常量指定一个友好的名字：常量只能在编译器获得，不可从运行期获得
	const Pi float64 = 3.14159265358979323846
	const zero = 0.0 // 无类型浮点常量
	const (
		size int64 = 1024
		eof = -1 // 无类型整型常量
	)
	const u, v float32 = 0, 3 // u = 0.0, v = 3.0,常量的多重赋值
	const a, b, c = 3, 4, "foo"
	// a = 3, b = 4, c = "foo", 无类型整型和字符串常量
	fmt.Println(Pi,zero,size,eof,a,b,c)
	//const  string :=os.Getenv("HOME") //错误
	string :=os.Getenv("HOME")  //正确
	fmt.Println(string)



	//2.2.3 预定义常量
	/*Go语言?定义了这些常量： true 、 false 和 iota 。
	iota 比较特殊，可以被认为是一个可被编译器修改的常量，在每一个 const 关键字出现时被
	重置为0，然后在下一个 const 出现之前，每出现一次 iota ，其所代表的数字会自动增1。
	从以下的例子可以基本理解 iota 的用法：*/
	const ( // iota被重设为0
		c0 = iota  // c0 == 0
		c1 = iota  // c1 == 1
		c2 = iota  // c2 == 2
	)
	const (
		uu = iota * 42 // u == 0
		vv float64 = iota * 42 // v == 42.0
	    ww = iota * 42 // w == 84
	)
    fmt.Println(c0,c1,c2,uu,vv,ww)
	const ( // iota被重置为0      简写方式
		c00 = iota // c0 == 0
		c10 // c1 == 1
		c20 // c2 == 2
	)


//	2.2.4 枚举
//	下面是一个常规的枚举表示法，其中定义了一系列整型常量：
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays // 这个常量没有导出
	)
	fmt.Println(Sunday,Monday,Thursday,numberOfDays)

	//2.3 类型
	/*Go语言内置以下这些基础类型：
	 布尔类型： bool 。
	整型： int8 、 byte 、 int16 、 int 、 uint 、 uintptr 等。
	浮点类型： float32 、 float64 。

	复数类型： complex64 、 complex128 。
	字符串： string 。
	字符类型： rune 。
	错误类型： error 。
	此外，Go语言也支持以下这些复合类型：
	指针（pointer）
	数组（array）
	切片（slice）
	字典（map）
	通道（chan）
	结构体（struct）
	接口（interface）*/

	//布尔类型
	var v1v1 bool
	//v1v1 = true
	v2v2 := (1==2)  //自动推导boolean类型
	fmt.Println(v1v1,v2v2)

    //比较运算
    ii,jj := 1,1
    if ii == jj {
    	fmt.Println(true)
	}

	//两个不同类型的整型数不能直接比较,比如 int8 类型的数和 int 类型的数不能直接比较，但
	//各种类型的整型变量都可以直接与字面常量（literal）进行比较，比如：
	var xz int32
    var zx int64
    xz,zx = 10,20
    //if xz == zx {
	//	fmt.Println(xz,zx)
	//}    //编译错误
	if xz == 10 || zx == 20{
		fmt.Println(xz,zx)
	}    //编译通过


	//2.3.3 浮点类型
	/*Go语言定义了两个类型 float32 和 float64 ，其中 float32 等价于C语言的 float 类型，
	float64 等价于C语言的 double 类型。*/
	var fvalue1 float32
	var ff float64
	//fvalue1 = 12
	fvalue2 := 12.000// 如果不加小数点，fvalue2会自动推导为整型，而不是浮点型
	fmt.Println(fvalue1,fvalue2)
	//使用强类型转换
    fvalue1 = float32(fvalue2)
    ff = float64(fvalue1)
	fmt.Println(fvalue1,fvalue2,ff)


    //2.3.4 复合类型
	/*复数实际上由两个实数（在计算机中用?点数表示）构成，一个表示实部（real），一个表示
	虚部（imag）。如果了解了数学上的复数是怎么回事，那么Go语言的复数就非常容易理解了。*/
	var value1 complex64 // 由2个float32构成的复数类型
	value1 = 3.2 + 12i
	fmt.Println(value1)
	value2 := 3.2 + 12i // value2是complex128
	value3 := complex(3.2, 12) // value3结果同value2
	fmt.Println(value2,value3)



	//2.3.5  字符串
	//var str string //申明字符串变量
	str := "hello go"
	ch := str[0]   //取字符串第一个字符
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

	/*strmessage := "Hello world" //
	strmessage[0] = 'X'  //编译错误，不能修改*/

	//字符串遍历
	n:= len(str)
	for ix := 0; ix < n; ix++{
		ch := str[ix]
		fmt.Printf("The first character of \"%s\" is %c.\n",ix, ch)
	}


	////2.3.7 数组
	//[32]byte   // 长度为32的数组,每个元素为一个字节
	//[10] struct { x, y int32 }  // 复杂类型数组
	//[1000]*float64  // 指针数组
	//[3][5]int  // 二维数组
	//[2][2][2]float64  // 等同于[2]([2]([2]float64))

	//定义切片
	s := [4]int{0,1,2,3}
	t := s[1:3]
	fmt.Println(s[0],t,s[:3],t[1:])


	//make copy append
	var aa = [...]int{1,2,3,4,5,6,7,8}
	var bb = make([]int,6)
	var cc = make([]byte,5)
	var n1 = copy(bb,aa[0:])
	var n2 = copy(bb,bb[2:])
	var n3 = copy(cc,"src")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	//2.3.9  map
	var personDB map[int] PersonInfo
	personDB = make(map[int] PersonInfo)
    fmt.Println(personDB)
    personDB[0] = PersonInfo{"id","name","address"}
    person,ok := personDB[0]
    if ok {  //找到了
    	fmt.Println(person.Address)
	}
    fmt.Println(person,ok)



    /*
    关于条件语句，需要注意以下几点：
     条件语句不需要使用括号将条件包含起来 () ；
     无论语句体内有几条语句，花括号 {} 都是必须存在的；
     花括号 { 必须与 if 或者 else 处于同一行；
     在 if 之后，条件语句之前，可以添加变量初始化语句，使用 ; 间隔；   看下面示例
     在有返回值的函数中，不允许将“最终的” return 语句包含在 if...else... 结构中，
     否则会编译??：
     */
   //流程控制
   if aa := 3;aa< 10{

   }


	//2.4.2 选择语句
	/*在使用 switch 结构时，我们需要?意以下几点：
	  左花括号 { 必须与 switch 处于同一行；
	  条件表达式不限制为常量或者整数；
	  单个 case 中，可以出现多个结果选?；
	  与C语言等规则相反，Go语言不需要用 break 来明确退出一个 case ；
	  只有在 case 中明确?加 fallthrough 关键字，才会继续执行??的下一个 case ；
	  可以不设定 switch 之后的?件表达式，在此种情况下，整个 switch 结构与多个
	   if...else... 的逻辑作用等同。*/
	i := 1
	switch i {
	case 0:
		fmt.Printf("0")
	case 1:
		fmt.Printf("1")
	case 2:
		fallthrough
	case 3:
		fmt.Printf("3")
	case 4, 5, 6:
		fmt.Printf("4, 5, 6")
	default:
		fmt.Printf("Default")
	}

    //2.4.3 循环语句
	/*与多数语言不同的是，Go语言中的??语句只支持 for 关键字，而不支持 while 和 do-while
	结构。关键字 for 的基本使用方法与C和C++中非常接近：*/
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
    fmt.Println(sum)

	/*
	可以看到比?大的一个不同在于 for 后面的?件表达式不需要用?括号 () 包含起来。Go语言
还进一步考虑到无限循环的场景，让开发者不用写无聊的 for (;;) {}  和 do {} while(1); ，
而直接简化为如下的写法：
	 */
	summ := 0
	for {
		summ++
		if summ > 1000 {
			fmt.Println(summ)
			break
		}
	}


	//2.4.4    goto 跳转语句
    myfunc()

    //2.5   函数
	// 2.5.1 函数定义
    bx,xb := Add(1,2)
    fmt.Println(bx,xb)










}

//2.5.1 函数定义 函数定义
func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 { // 假设这个函数只支持2个非负数数字相加
		err= errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil  // 支持多重返回
}

func myfunc() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

































































































































































































































































































































































































































































































































































//2.1.4  匿名变量
func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

//函数
func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}