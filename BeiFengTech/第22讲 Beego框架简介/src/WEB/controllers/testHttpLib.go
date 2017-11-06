package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type TestHttpLibController struct {
	beego.Controller
}

func (this *TestHttpLibController) Get() {
	req := httplib.Get("http://www.baidu.com")
	result, err := req.String()
	if err != nil {
		//
	}
	this.Ctx.WriteString(result)
}
