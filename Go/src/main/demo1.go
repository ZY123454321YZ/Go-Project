package main

import "fmt"

func main() {
	//声明数组类型
	type Point [6]float32
	//定义数组
	a := Point{1,2,2,3,5,8}
	var i = 1000
	fmt.Println("demo")
	fmt.Println(i)
    fmt.Println(a[5])
	fmt.Println(len(a))
	//定义切片
	s := [4]int{1,2,3,4}
	t := s[1:3]
	fmt.Println(s[0],t,s[1],s[2])
	fmt.Println(len(s),cap(s))
	fmt.Println(len(t),cap(t))


	var aa = [...]int{0,1,2,3,5,6,7,8,9,10,11,12,13}
	var ss = make([]int,6)
	var bb = make([]byte,5)
	n1 := copy(ss,aa[0:])
	n2 := copy(ss,ss[2:])
	n3 := copy(bb,"demotest")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	//append(t,o) 添加新的切片
     //追加
	s0 := []int{0,0}
	s1 := append(s0,1)
	fmt.Println(s1)
	s2 := append(s1,3,5)
	fmt.Println(s2)
    s3 := append(s1,s0...)
    fmt.Println(s3)
    var b[]byte
    b = append(b,"ab"...)//b = []byte{'a','b'}
    fmt.Println(b)
    //添加切片
    c := []int{0,1,3}
    cc := append(c,2)
    fmt.Println(cc)
    //复制
    var ccc = make([]int ,3)
    copy(ccc,cc[1:2])
    fmt.Println(ccc)
    //删除[i,j]
    cccc := append(cc[:1],cc[3:]...)
    fmt.Println(cccc)
    //删除第2个元素
    ac := append(cc[:2],cc[3:]...)
    fmt.Println(ac)


}
