package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type TestInputController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string
}

//func (this *TestInputController) Get() {
//	//id := this.GetString("id")       //方法一
//	//name := this.Input().Get("name") //方法二
//	//this.Ctx.WriteString("<html>" + id + "<br/>")
//	//this.Ctx.WriteString(name + "</html>")
//
//	//读取session的内容
//	name := this.GetSession("name")
//	password := this.GetSession("password")
//
//	if nameString, ok := name.(string); ok && nameString != "" {
//		this.Ctx.WriteString("Name:" + name.(string) + " password:" + password.(string))
//	}else{
//		this.Ctx.WriteString(`<html><form action="http://127.0.0.1:8082/testInput" method="post">
//							<input type="text" name="Username"/>
//							<input type="password" name="Password"/>
//							<input type="submit" value="提交"/>
//					   </form></html>`)
//	}
//
//}

//方法三
//func (this *TestInputController) Post() {
//	u := new(User)
//	if err := this.ParseForm(u); err != nil {
//		this.Ctx.WriteString("异常错误")
//	}
//	this.Ctx.WriteString("<html>" + u.Username + "<br/>")
//	this.Ctx.WriteString(u.Password + "</html>")
//}

//方法四
//在配置文件里设置 copyrequestbody = true
//func (this *TestInputController) Post() {
//	var ob models.Object
//	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
//	objectid := models.AddOne(ob)
//	this.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
//	this.ServeJSON()
//}

func (c *TestInputController) Get() {
	//测试session的使用
	c.SetSession("name", "Goular44558")
	c.SetSession("password", "zhoiajo@dslkdjlk")
}

func (c *TestInputController) Get2() {
	name := c.GetSession("name").(string)
	password := c.GetSession("password").(string)

	fmt.Println(name)

	c.Ctx.WriteString("Username:" + name + " Password:" + password)
}

func (c *TestInputController) Post() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		//process error
	}

	c.Ctx.WriteString("Username:" + u.Username + " Password:" + u.Password)
}
