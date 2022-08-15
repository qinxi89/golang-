package main

import (
	"fmt"
	"reflect"
)
//整型
func main()  {
	var v1 int32
	v1 = 123
	v2 := 64
	fmt.Println(reflect.TypeOf(v1))
	fmt.Println(reflect.TypeOf(v2))
}
//浮点型
func main()  {
	var f1 float64
	f1 = 12
	fmt.Println(f1)
	f2 := 12.5
	fmt.Println(reflect.TypeOf(f2))
}



