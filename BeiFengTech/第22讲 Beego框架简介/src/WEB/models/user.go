package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

var (
	db orm.Ormer
)

//由于model这个名字叫 UserInfo 那么操作的表其实 user_info
type UserInfo struct {
	Id       int64
	Username string
	Password string
}

func init() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))
	db = orm.NewOrm()
}

func AddUser(user_info *UserInfo) (int64, error) {
	id, err := db.Insert(user_info)
	return id, err
}

func ReadUserInfo(users *[]UserInfo) {
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("*").From("user_info")

	sql := qb.String()
	db.Raw(sql).QueryRows(users)
}
