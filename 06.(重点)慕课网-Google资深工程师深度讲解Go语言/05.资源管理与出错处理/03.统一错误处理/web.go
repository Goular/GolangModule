package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"strings"

	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 输入的函数作为输出的函数来使用，这是函数式编程的一个点
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("panic:%v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request :%s", err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			// 文件不存在
			case os.IsNotExist(err):
				code = http.StatusNotFound
				//文件无权限访问
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

const prefix = "/lists/"

type userErr string

func (ue userErr) Error() string {
	return ue.Message()
}

func (ue userErr) Message() string {
	return string(ue)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userErr("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		//return
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		// panic(err)
		return err
	}
	writer.Write(all)
	return nil
}
