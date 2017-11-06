package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type TestContextController struct {
	beego.Controller
}

func (this *TestContextController) Get() {
	this.Ctx.WriteString(this.Ctx.Input.IP() + ":" + strconv.Itoa(this.Ctx.Input.Port()))
	this.Ctx.WriteString(this.Ctx.Input.Query("name")) //等价于php中的 $_REQUEST["name"]
	m := make(map[string]float64)
	m["zhangsna"] = 98.7
	this.Ctx.Output.JSON(m, false, false)
}
