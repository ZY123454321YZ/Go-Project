package main
import "fmt"
func main() {
	//变量申明
	var v1 int
	var v2 string
	var v3 [10]int  //数组
	var v4 []int    //切片
    var v5 struct{
    	f int
	}
	var v6 *int
	var v7 map[string]int
	var v8 func()
	var (
		v9 int
		v10 string
	)
	fmt.Print(v1,v2,v3,v4,v5,v6,v7,v8,v9,v10)
   //变量初始化 三种写法
	var a1 int = 10
	var a2 = 10
	a3 := 10
	fmt.Println(a1,a2,a3)
	//多重赋值
    var i =10
    var j = 9
    i,j = j,i
    fmt.Println(i,j)
    var m int
    var n int
    m,_ =vals()
    fmt.Println(m,n)
    //单返回值
    single := getName("cc")
    fmt.Println(single)
    //多返回值
     lisi,lisi1 := getParmter("abc","bc")
    fmt.Println("lisi 值为 :" + lisi, "lisi1 值 :" +lisi1)
     //命名返回值
	area, perimeter := rectProps(80,100)
	fmt.Println("area 的值为 ",area,"perimeter 的值为 ",perimeter)
}
//返回值一个
//func functionname(parametername type) returntype {
//// 函数体（具体实现的功能）
//}
func getName(a string) string{
	return a
}
//无返回值
//func functionname() {
//	// 译注: 表示这个函数不需要输入参数，且没有返回值
//}
func getParmter(a,b string)(x,z string){
	return a + b,b
}
//多重返回值
func vals()(int,int){
	return 3,7
}
//命名返回值
func rectProps(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}