package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

type TestConfigController struct {
	beego.Controller
}

func (this *TestConfigController)Get()  {
	iniconf, err := config.NewConfig("ini", "testini.conf")
	if err != nil {
		//t.Fatal(err)
	}
	str := iniconf.String("appname")
	this.Ctx.WriteString(str)
}


