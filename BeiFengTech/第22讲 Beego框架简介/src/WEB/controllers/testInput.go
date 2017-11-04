package controllers

import (
	"github.com/astaxie/beego"
	"API/models"
	"encoding/json"
)

type TestInputController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string
}

func (this *TestInputController) Get() {
	id := this.GetString("id")       //方法一
	name := this.Input().Get("name") //方法二
	this.Ctx.WriteString("<html>" + id + "<br/>")
	this.Ctx.WriteString(name + "</html>")
}

//func (this *TestInputController) Post() {
//	u := new(User)
//	if err := this.ParseForm(u); err != nil {
//		this.Ctx.WriteString("异常错误")
//	}
//	this.Ctx.WriteString("<html>" + u.Username + "<br/>")
//	this.Ctx.WriteString(u.Password + "</html>")
//}

//在配置文件里设置 copyrequestbody = true
func (this *TestInputController) Post() {
	var ob models.Object
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	this.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
	this.ServeJSON()
}
