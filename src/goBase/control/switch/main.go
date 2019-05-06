package main

import "fmt"

func main()  {
	switch 3 {
	case 1:
		fmt.Print("判断为1\n")
	case 2:
		fmt.Print("判断为2\n")
	default:
		fmt.Print("以上都不满足\n")
	}


	var a interface{}
	a="history"
	switch a.(type) {
	case int:
		fmt.Print("类型为整型")
	case string:
		fmt.Print("类型为字符串类型")
	default:
		fmt.Print("以上都不是")
	}
}
