package main
//本文讲述匿名变量
import (
	"fmt"
)

func pp() (int, int)  {
	return 1, 2
}
func main()  {
	v1, _ := pp()
	fmt.Println(v1)
}
