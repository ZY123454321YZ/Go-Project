package main
import "fmt"

//变量申明
var a = 1234
var b string = "test"
var c bool
//
var x,y int  //类型相同的多个变量
//类型不同的多个变量
var (
	d string = "string"
	f bool
)
var g,i int =1,2
var gg,ii = 1,"string" //类型推断
func main() {
	uu,hh := 123,"hello"
	fmt.Println(a,b,c)
	fmt.Println(f,uu,hh)
	sa := []int{1,2,3}
	i := 0
	i,sa[i] =1,2

	sb := []int{1, 5, 3}
	j := 0
	sb[j], j = 2, 1  // sets sb[0] = 2, j = 1

	//map 引用类型
	//var m map[string]int
	//m["a"]=1
    var a = 10
    if a > 10 {
		fmt.Println("True")
	} else {
		fmt.Println("false")
	}

}
