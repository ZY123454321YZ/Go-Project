package main
import "fmt"
func main() {
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
}
