/*
curl http://localhost:8080/
curl --data "users=yifan,yifan2" http://localhost:8080/import
curl http://localhost:8080/lucky

*/
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"math/rand"
	"strings"
	"time"
)

var userList []string

//iris的上下文环境
type lotteryController struct{
	Ctx iris.Context
}

//iris的应用
func newApp() *iris.Application  {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main(){
	//调用
	app :=newApp()
	//声明变量
	//userList=make([]string,0)
	userList=[]string{}

	//运行
	app.Run(iris.Addr(":8080"))
}
func (c *lotteryController) Get() string{
	count :=len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数: %d\n",count)
}

//导入用户名单
func (c*lotteryController) PostImport() string {
	strUsers :=c.Ctx.FormValue("users")
	users :=strings.Split(strUsers,"")
	count1 :=len(userList)
	for _, u :=range users{
		u=strings.TrimSpace(u)
	}
	count2 :=len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数： %d，成功导入的用户数： %d/n",
		count2,(count2-count1))
}
//GET http://localhost:8080/lucky
func (c *lotteryController) GetLucky() string  {
	count :=len(userList)
	if count>1{
		seed :=time.Now().UnixNano()
		index :=rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user :=userList[index]
		userList=append(userList[0:index],userList[index+1:]...)
		return fmt.Sprintf("当前中奖用户:%s,剩余用户数:%d\n",user,count-1)
	}else if count==1{
		user :=userList[0]
		return fmt.Sprintf("当前中奖用户:%s,剩余用户数:%d\n",user,count-1)
	}else{
		return fmt.Sprintf("已经没有参与用户，请先通过 /import 导入用户\n",)
	}

}

