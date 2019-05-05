//单行注释 package main
package main

import "fmt"

//常量定义 常量定义建议大写
const NAME="history"

//main函数外定义的变量  是全局变量
var mainName="main name"
/*
	多行注释
	main func
*/
func main()  {
	fmt.Println("learn")
	fmt.Print(mainName)
	fmt.Print(NAME)
}