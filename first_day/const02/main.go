package main

import "fmt"

//本文讲述变量初始化
func main()  {
	//1、初始化方式1
	var v1 int =1

	//2、自动推导
	var v2 = 2

	//3、：= 声明并初始化
	v3 := 3
	fmt.Print(v1, v2, v3)
}




