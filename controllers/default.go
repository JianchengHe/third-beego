package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
	//获取get类型的请求参数
	name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
	//以admin和18为正确数据验证
	if name != "admin" || age != "18"{
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据验证成功"))
	//c.Data["Website"] = "www.baidu.com"
	//c.Data["Email"] = "1019015673@qq.com"
	//c.TplName = "index.tpl"
}


//该post方法是处理post类型的请求的时候
func (c *MainController) Post(){
	fmt.Println("post类型的请求")
	user := c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为：",user)
	pwd := c.Ctx.Request.FormValue("pwd")
	fmt.Println("密码为：",pwd)

	//与固定值进行比较
	if user != "admin" || pwd != "123456"{
		//失败页面
		c.Ctx.ResponseWriter.Write([]byte("对不起数据不正确"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据正确"))
	//request 请求，response 响应

}