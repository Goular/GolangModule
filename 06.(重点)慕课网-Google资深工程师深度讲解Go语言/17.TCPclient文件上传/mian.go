package main

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	tonydon := &User{Name: "TangXiaodong", Age: 100}
	//object := reflect.ValueOf(tonydon)
	//myref := object.Elem()
	//typeOfType := myref.Type()
	//for i := 0; i < myref.NumField(); i++ {
	//	field := myref.Field(i)
	//	fmt.Printf("%d. %s %s = %v \n", i, typeOfType.Field(i).Name, field.Type(), field.Interface())
	//}
	//tonydon.SayHello()
	//v := object.MethodByName("SayHello")
	//v.Call([]reflect.Value{})
	getRequestBody(tonydon)
}

type User struct {
	Name string
	Age  int
	Id   string
}

func (u *User) SayHello() {
	fmt.Println("I'm " + u.Name + ", Id is " + u.Id + ". Nice to meet you! ")
}

// 获取需要插入或者更新的map
func getUpdateMap(v interface{}) bson.M {
	object := reflect.ValueOf(v)
	myref := object.Elem()
	typeOfType := myref.Type()
	m := bson.M{}
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		key := typeOfType.Field(i).Name
		value := field.Interface()
		if field.Interface() != reflect.Zero(field.Type()).Interface() {
			m[key] = value
		}
	}
	return m
}
