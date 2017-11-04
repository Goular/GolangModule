package controllers

import "github.com/astaxie/beego"

type TestInputController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string
}

func (this *TestInputController) Get() {
	id:=this.GetString("id")//方法一
	name := this.Input().Get("name")//方法二
	this.Ctx.WriteString("<html>" + id + "<br/>")
	this.Ctx.WriteString(name + "</html>")
}
