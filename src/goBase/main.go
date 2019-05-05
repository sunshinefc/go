package main

import (
	"fmt"
	"goBase/show2"
)

func init()  {
	fmt.Print("在main()函数执行之前\n")
}

func main()  {

	show2.Show2()
	fmt.Print("hello imooc")
}