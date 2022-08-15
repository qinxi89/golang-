package main

import "fmt"

func main()  {
	//iota接着走，不重新来
	const (
		a = iota
		b
		c
		d = 8
		f = iota
		g
	)
	fmt.Println(a, b, c, d, f, g)

	// iota 不在同一个const里， 不是同一个
	const (
		x = iota
		y
	)
	fmt.Println(x, y)
}
