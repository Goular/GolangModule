package main

import (
	"os"
	"bufio"
)

func main() {
	// fmt.Println(GetCurrentDirectory())

	// 数据内容
	data := "dsfdsfwr6ppop"
	// 文件夹名称
	folderName := "/public/"

	// 获取当前文件夹
	currentDir, _ := os.Getwd() //当前的目录
	// 判断文件夹是否存在
	var flag bool
	var err error
	if flag, err = FolderExists(folderName); err != nil {
		// todo：直接日志记录
		panic(err)
	}
	if !flag {
		//在当前目录下生成md目录
		if err := os.Mkdir(currentDir+folderName, os.ModePerm); err != nil {
			// todo：直接日志记录
			panic(err)
		}
	}
	// 文件名
	fileName := currentDir + folderName + "/test.txt"
	// 保存文件
	WriteWithBufio(fileName, []byte(data))
}

// 将字节数组写入文件 folderName："文件夹名称"
func storeDataIntoFile(folderName, fileName string, contents []byte) {

}

// 文件夹是否存在
func FolderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//使用bufio包中Writer对象的相关方法进行数据的写入
func WriteWithBufio(fileName string, contents []byte) {
	// 读写文件，创建文件和覆盖文件
	if fileObj, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644); err == nil {
		defer fileObj.Close()
		writeObj := bufio.NewWriter(fileObj)
		//使用Write方法,需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
		buf := contents
		if _, err := writeObj.Write(buf); err == nil {
			if err := writeObj.Flush(); err != nil {
				// todo：直接日志记录
				panic(err)
			}
		}
	}
}
