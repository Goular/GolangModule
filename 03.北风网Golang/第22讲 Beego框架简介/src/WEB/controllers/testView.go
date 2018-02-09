package controllers

import (
	"github.com/astaxie/beego"
	"WEB/models"
)

type TestViewController struct {
	beego.Controller
}

func (this *TestViewController) Get() {
	var users []models.UserInfo
	models.ReadUserInfo(&users)
	this.Data["Users"] = users
	this.Data["title"] = "欢迎来到加工屋"
	this.Data["len"] = len(users)
	this.Data["Content1"] = "沉默是金111"
	this.Data["Content2"] = "沉默是金222"
	this.TplName = "test_view.tpl"
}
