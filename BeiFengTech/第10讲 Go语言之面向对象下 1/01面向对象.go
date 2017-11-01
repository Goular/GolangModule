package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (person Person) getNameAndAge() (string, int) {
	return person.name, person.age
}

type Student struct {
	Person
	speciality string
}

func (student Student) getSpeciality() string {
	return student.speciality
}


func main() {
	student := new(Student)
	student.name = "ZhangSan"
	student.age = 26
	student.speciality = "Hello Golang！"
	name,age := student.getNameAndAge()
	speciality := student.getSpeciality()
	fmt.Println(name,age)
	fmt.Println(speciality)
}
