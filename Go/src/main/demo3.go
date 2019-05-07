package main
import (
	"fmt"
	"net/url"
	"time"
)
//定义结构体
type D4 struct { x,y float64 }
//定义内置，把一个结构体放入另一个结构体
type D5 struct {
	D4
	z float64
}

const Liby  = iota

//定义结构体
func main() {
	gurl, er := url.Parse("http://www.baidu.com")
    fmt.Println(gurl.User,er)
	fmt.Println(gurl.Path,er)
	fmt.Println(gurl.Host,er)
	//定义数组
	type Point [6]float64;
	type Line [4]Point;
	a := Point{1, 2}
	b := Point{3,4}
	l := Line{a,b}
	fmt.Println(a[0],l[1][1],b,l)
	c := a
	d := Line{a,b}
	fmt.Println(c,d)

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

	var test = append(bb,9)
	fmt.Println(test)

	str := "beijing,北京"
	for i := 0; i < len(str); i++{
		fmt.Println(i, " ", str[i])
	}

	//append 追加
	var s0 = []int{0,0}
	s1 := append(s0,2)
	fmt.Println(s1)
	var s2 = append(s1,5)
	fmt.Println(s2)

	var s3 []byte
	s3 = append(s3,"abcd"...)
	fmt.Println(s3)
   //切片复制
   demo := []int{10,11,12,13}
   demo1 := make([]int,3)
   copy(demo1,demo[2:])
   fmt.Println(demo1)

   //删除[i,j] 相当于删除i 到 j-1
    source := []string{"aaa","bbb","ccc","ddd","eee"}
    source1 := append(source[:1],source[3:]...)
    fmt.Println(source1)

    //删除第i个元素
    dest := []string{"a","b","c","d"}
    dest1 := append(dest[:1],dest[2:]...)
    fmt.Println(dest1)

    //扩展j个空元素
   vivo := []string{"vivo","vivo","vivo","vivo"}
   vivo1 := append(vivo,make([]string,3)...)
   fmt.Println(len(vivo1))

   //插入j个空元素
   ciso := []string{"ciso","ciso","ciso","ciso","ciso"}
   ciso1 := append(ciso[:3],append(make([]string,9),ciso[3:]...)...)
   fmt.Println(ciso1)

   //插入元素x
   huawei := []string{"huawei","huawei","huawei","huawei","huawei"}
   huawei1 := append(huawei[:3],append([]string{"ciso"},huawei[3:]...)...)
   fmt.Println(huawei1)

   //插入切片b
   zhongxing := []string{"zhongxing","zhongxing","zhongxing","zhongxing","zhong"}
   zhongxing1 := zhongxing[:3]
   zhongxing2 := append(zhongxing[:3],append(zhongxing1,zhongxing[3:]...)...)
   fmt.Println(zhongxing2)

   //弹出最后一个元素
   ali := zhongxing[len(zhongxing)-1]
   fmt.Println(ali)

   var d2 = D4{7,8}
   d3 := D5{d2,9}
   fmt.Println(d2)
   fmt.Println(d3)
   fmt.Println(d2.x,d2.y,d3.z)
   fmt.Println(d3.D4 == d2)

   //定义指针
   var jj int
   var zk *int
   var elk **int
   jj = 0
   zk = &jj
   elk = &zk
   fmt.Println(jj,zk,elk)

   langchao := 70
   langchao1 := &langchao
   fmt.Printf("langchao的值为 %v, langchao的指针是 %v ，langchao1指向的变量的值为 %v\n",langchao,langchao1,*langchao1)
	aaa := 3
	bbb := 4
	p1 := &aaa //获取变量a的内存地址，并将其赋值给变量p1
	p2 := &bbb //获取变量b的内存地址，并将其赋值给变量p2
	fmt.Printf("a的值为 %v, a的指针是 %v ，p1指向的变量的值为 %v\n",aaa,p1,*p1)
	fmt.Printf("b的值为 %v, b的指针是 %v ，p2指向的变量的值为 %v\n",bbb,p2,*p2)
	fmt.Println(demo3(p1,p2))
	fmt.Printf("a的值为 %v, a的指针是 %v ，p1指向的变量的值为 %v\n",aaa,p1,*p1)
	fmt.Printf("b的值为 %v, b的指针是 %v ，p2指向的变量的值为 %v\n",bbb,p2,*p2)

	 //定义空指针
    var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr  )
	if(ptr != nil){
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}

	//多重间接引用
	z1 := &aaa //获取变量a的内存地址，并将其赋值给变量p1
	z2 := &z1 //获取变量b的内存地址，并将其赋值给变量p2
	//aaa = 4
	fmt.Printf("aaa的值为 %v, z1的指针是 %v ，指向的变量的值为 %v\n",aaa,z1,*z1)
	fmt.Printf("aaa的值为 %v, z2的指针是 %v ，指向的变量的值为 %v\n",aaa,z2,**z2)

	//for 循环
	str = "str,message,ok"
	for i :=0;i<len(str) ;i++  {
		fmt.Println(str[i])
	}
	fmt.Println(say1("k8s"))

	tz := time.Date(2019,time.April,10,23,0,0,0,time.UTC)
	fmt.Println(tz)

    add := func(base int) func(int) int{
		return func(n int) int{
			return base + n
		}
	}
    add5 := add(5)
    fmt.Println(add5(10))
    fmt.Println(foo2(60))
}

func demo3(a,b *int)int  {
	*a = 5
	*b = 6
	return  *a * *b //这里出现连续的两个*，Go编译器会根据上下文自动识别乘法与两个引用
}
func say1(message string)string {
	return message
}

func foo1(i int) int {
	i = 100
	i = 200
	fmt.Printf("i的值为 %v",i)
	return i
}
//定义defer
func foo2(i int) int {
	i = 100
	defer foo1(4)
	i = 200
	return i
}


