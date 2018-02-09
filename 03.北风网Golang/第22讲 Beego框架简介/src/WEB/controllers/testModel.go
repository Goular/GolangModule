package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"WEB/models"
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
	orm.Debug = true

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
	//user := UserInfo{}
	//user.Id = 3
	//num, err := o.Delete(&user)
	//if err != nil {
	//	this.Ctx.WriteString(fmt.Sprintf("num:%v", num))
	//}

	//原生SQL查询
	//var maps []orm.Params
	//o.Raw("select * from user_info").Values(&maps)
	//
	//for _, v := range maps {
	//	this.Ctx.WriteString(fmt.Sprintf("User info %v",v))
	//}

	//原生SQL查询 二
	//var users []UserInfo
	//o.Raw("select * from user_info").QueryRows(&users)
	//for _, v := range users {
	//	this.Ctx.WriteString(fmt.Sprintf("User info：%v", v))
	//}

	//采用queryBuilder方式进行读取
	var users []UserInfo
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("password").From("user_info").Where("username=?").Limit(1)
	sql := qb.String()
	o.Raw(sql, "lisi").QueryRows(&users)
	this.Ctx.WriteString(fmt.Sprintf("User info：%v", users))
}

func (this *TestModelController) Get2() {
	user := models.UserInfo{Username: "liusi", Password: "8755522k"}
	models.AddUser(&user)
	this.Ctx.WriteString("Call Model Success")
}
