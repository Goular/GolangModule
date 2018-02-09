package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte("123456"))
	Result := Md5Inst.Sum([]byte("")) //添加加密因子
	fmt.Printf("%x\n\n",Result)
}
