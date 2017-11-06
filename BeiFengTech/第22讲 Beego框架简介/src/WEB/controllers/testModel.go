package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type UserInfo struct {
	Id       int64
	Username string
	Password string
}

type TestModelController struct {
	beego.Controller
}

func (this *TestModelController) Get() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))

	o := orm.NewOrm()

	//插入数据操作
	//user := UserInfo{Username: "zhangsan", Password: "112233445566"}
	//id, err := o.Insert(&user)

	//修改数据操作
	//user := UserInfo{}
	//user.Id = 1
	//user.Username="Lisi"
	//id,err := o.Update(&user)

	//查询数据的操作
	//user := UserInfo{}
	//user.Id = 1
	//o.Read(&user)
	//this.Ctx.WriteString(fmt.Sprintf("user的内容:%v", user))

	//删除数据操作
	user := UserInfo{}
	user.Id = 3
	num, err := o.Delete(&user)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("num:%v", num))
	}
}
