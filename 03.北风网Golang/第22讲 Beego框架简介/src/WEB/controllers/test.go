package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (controller *TestController) Get() {
	//相当于php里面的echo 函数
	controller.Ctx.WriteString("<font style=\"color:red\">welcome test controller!</font>")
}
