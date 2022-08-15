package main
//const     声明常量  iota 常量生成器和枚举类型

import (
	"fmt"

)
//定义在函数外
const filename = "abcd.txt"

//常量定义
func consts()  {
	const (
		filename = "abcd.txt"
		a, b = 3, 4
	)
	fmt.Println(filename)
}
//枚举
func enums1()  {
	const (
		python = 0
		java   = 1
		golang = 2

		)
	fmt.Println(python, java, golang)
}
//可以省略不赋值，跟前面一样
func enums2() {
	const (
		python = 0
		java
		golang = 1
		php
		)

}
//定义的自增枚举类型
//iota 默认值为0，往下一次自增
func enums3()  {
	const (
		python = iota
		java
		golang
		)
	fmt.Println(python, java ,golang )
}
func enums4()  {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tbsdfsf
		pb
		)
	fmt.Println(b, kb, mb, gb, tbsdfsf, pb)
}

func main()  {
	enums1()
	enums2()
	enums3()
	enums4()
}