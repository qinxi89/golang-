package _const

import "fmt"

//1、标准声明格式
func main()  {

	var name string
	var age int
	var isOK bool
	fmt.Print(name,age,isOK)
}

//2、批量声明变量
var (
	a string
	b int
	c bool
	d float32
)
fmt.Println(a,b,c,d)

//3、声明变量同时指定初始值
var name1 string = "芹溪"
var age1 int = 18
fmt.Println(name1,age1)
var name2, age2 = "翠花", 20
fmt.Println(name2,age2)

//4、类型推导，让编译器根据变量的值推导出其类型
var name3 = "若初"
var age3 = 3
fmt.Println(name3, age3)









