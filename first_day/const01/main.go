package main
//本文记录变量声明
import "fmt"

func main() {

	//1、声明变量
	var v1 int
	var v2 int

	//2、一次定义多个变量
	var v3, v4 int

	//3、集中定义多个变量
	var (
		v5 int
		v6 int
	)
	fmt.Println(v1, v2, v3, v4, v5, v6)
}