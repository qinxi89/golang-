package main
//变量赋值

import "fmt"

func main()  {
	//1、直接赋值
	var v1 int
	v1 = 1
	//2、：=
	v2 := 2
	//3、多重赋值
	var v3, v4 int
	v3, v4 = 3, 4
	fmt.Println(v1, v2, v3, v4)
}