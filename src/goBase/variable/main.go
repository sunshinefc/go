package main

import (
	"fmt"
	"reflect"
)


func main()  {

	 var b float32=3.01
	 c:=int32(b)

	fmt.Print(c)
	fmt.Print("\n")
	fmt.Print(reflect.TypeOf(c))


}
