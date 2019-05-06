package main

import (
	"fmt"
	"reflect"
)


func main()  {
	var a,b,c  = 11,21.2,31
	fmt.Print("a变量的类型为：")
	fmt.Print(reflect.TypeOf(a))
	fmt.Print("\n")

	fmt.Print("b变量的类型为：")
	fmt.Print(reflect.TypeOf(b))
	fmt.Print("\n")

	fmt.Print(a)
	fmt.Print("\n")

	fmt.Print(b)
	fmt.Print("\n")

	fmt.Print(c)


}
