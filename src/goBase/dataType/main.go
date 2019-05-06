package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	//无符号整型
	var a uint8 =1
	var b int32 =1
	var c int =1
	fmt.Print(unsafe.Sizeof(a))
	fmt.Print(unsafe.Sizeof(b))
	fmt.Print(unsafe.Sizeof(c))

	var x float32=1.2
	fmt.Print(unsafe.Sizeof(x))

	var t bool=true
	fmt.Print(t)

	var i byte=1
	fmt.Print(unsafe.Sizeof(i))

	var j rune=1
	fmt.Print(unsafe.Sizeof(j))
}
