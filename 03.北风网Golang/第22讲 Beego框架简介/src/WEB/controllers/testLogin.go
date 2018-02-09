package controllers

import (
	"github.com/astaxie/beego"
)

type TestLoginController struct {
	beego.Controller
}

type UserInfoV2 struct{
	Username string
	Password string
}

func (c *TestLoginController) Login(){
	name := c.Ctx.GetCookie("name")
	password := c.Ctx.GetCookie("password")

	//do verify work
	if name != ""{
		c.Ctx.WriteString("Username:" + name + " Password:" + password)
	}else{
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8082/testLogin" method="post">
							<input type="text" name="Username"/>
							<input type="password" name="Password"/>
							<input type="submit" value="提交"/>
					   </form></html>`)
	}
}


func (c *TestLoginController) Post(){
	u := UserInfoV2{}
	if err:=c.ParseForm(&u) ; err != nil{
		//process error
	}

	c.Ctx.SetCookie("name", u.Username, 100, "/")
	c.Ctx.SetCookie("password", u.Password, 100, "/")
	c.SetSession("name", u.Username)
	c.SetSession("password", u.Password)
	c.Ctx.WriteString("Username:" + u.Username + " Password:" + u.Password)
}
