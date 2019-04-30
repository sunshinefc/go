package main

import "github.com/kataras/iris"

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

