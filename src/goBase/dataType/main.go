package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type 历史 int32
func main()  {
	var i 历史
	var j int32

	fmt.Print("i 的默认值")
	fmt.Print(i)
	fmt.Print("\n")
	fmt.Print("i变量的数据类型为：")
	fmt.Print(reflect.TypeOf(i))
	fmt.Print("\n")
	fmt.Print("历史别名占用空间大小为：")
	fmt.Print(unsafe.Sizeof(i))

	fmt.Print("\n")
	fmt.Print("----------------------------------")
	fmt.Print("\n")

	fmt.Print("j 的默认值")
	fmt.Print(j)
	fmt.Print("\n")
	fmt.Print("j变量的数据类型为：")
	fmt.Print(reflect.TypeOf(j))
	fmt.Print("\n")
	fmt.Print("j占用空间大小为：")
	fmt.Print(unsafe.Sizeof(j))


}
