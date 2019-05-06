package main

import "fmt"

const (
	a,b  =iota,iota+3
	c,d
	f=iota
)
func main()  {
	fmt.Print("a的常量值为：")
	fmt.Print(a)
	fmt.Print("\n")

	fmt.Print("b的常量值为：")
	fmt.Print(b)
	fmt.Print("\n")

	fmt.Print("c的常量值为：")
	fmt.Print(c)
	fmt.Print("\n")

	fmt.Print("d的常量值为：")
	fmt.Print(d)
	fmt.Print("\n")

	fmt.Print("f的常量值为：")
	fmt.Print(f)
	fmt.Print("\n")
}
