package main

func main() {
	x:="Hello World!"
	for _,str := range x{
		println(string(str))
	}
	y:=[10]float32{12.2,15.3,2,2.6,8.77}
	for _,value := range y{
		println(value)
	}
	z:=make(map[string]string)
	z["GuangZhou"]="花城"
	z["BeiJing"]="京城"
	z["ShangHai"]="申城"
	for key,value := range z{
		println(key,value)
	}


	//如何让两个数进行交换
	a:=12
	b:=19
	a,b = b,a
	println(a,b)

}
