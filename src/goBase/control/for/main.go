package main

import (
	"fmt"
	"time"
)

func main()  {
	//无限循环
	for {
		fmt.Print("历史\n")
	}
	//有限循环
	for i:=1;i<10;i++ {
		fmt.Print("历史\n")
		//1秒钟停止一次
		time.Sleep(1*time.Second)
	}

	//配合range达到for each的效果
	a :=[]string{"夏商","西周","春秋","战国","秦汉","魏蜀吴","二晋","南北朝","隋唐","宋元明清"}
	for _,value:=range a{

		fmt.Print("value值为：")
		fmt.Print(value)
		fmt.Print("\n")
	}
}
