package main
import (
	"fmt"
	"net/url"
)
//go映射相当于python 里的字典
const (
	China int = iota
	USA
	JPAN
	Total
)
//结构体定义
type URL struct {
	RAW string
	Schedl string
	Host string
	flax bool
	User string
	IP int64
}
type D2 struct {
	 x,y float32
}
type D3 struct {
	D2
	z float32
}
func main() {
	GDPof := map[string]float64{
		"USA": 14058,
		"JAPAN": 5.45,
	}
	fmt.Println(len(GDPof))
	gurl, er := url.Parse("http://www.baidu.com")
	fmt.Println(gurl.Host,er)
	fmt.Println(gurl.User,er)
	var d2 = D2{1,2}
	fmt.Println(d2)
	d3 := D3{d2,3}
	fmt.Println(d3)
	//也可以这样赋值
	d4 := D3{D2{1,4},5}
	fmt.Println(d4)
	//结构体不可以比较大小，只可以比较项是否相等，但必须是同一个类型，所有项都相等，才能说明结构体相等

	//指针
	var i int
	var p *int  //声明int指针类型
	var pp **int //声明int指针的指针类型
	i = 10
	p = &i //取变量地址使用&操作符
	pp = &p
	*p++
	fmt.Println(i,p,*p,pp,*pp)

	//流程控制
    var ii uint8
	for ii < 8  {
		fmt.Printf("%08b\n",1<<ii)
		ii++
	}

	for i:= 0; i < 7 ;i++  {
		fmt.Println(i)
	}
	say("test")
	fmt.Println(getString("fanhui"))
}
//定义函数
func  say(s string )  {
	 fmt.Println(s)
}
//返回值
func g(oo int) int {
	return 1;
}
func getString(oo string) string {
   return oo;
}




