package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"fmt"

	"github.com/go-errors/errors"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserErr string

func (ue testingUserErr) Error() string {
	return ue.Message()
}

func (ue testingUserErr) Message() string {
	return string(ue)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserErr("User Error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("Unkonwn error")
}

func noErr(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "User Error"},
	{errNotFound, 400, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
}

// 这个仅仅是模拟request对方法进行测试
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)
		// 这里err就不管了，因为我们本来就是测试，死循环的关注err没意义
		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("Expect(%d,%s); Get(%d %s)", tt.code, tt.message, response.Code, body)
		}
	}
}

//这个是直接启动一个server来测试 使用的是httptest的包中的方法的newServer
// 这个是测试整个服务器跑起来的处理，上面那个仅仅是方法
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.StatusCode != tt.code || body != tt.message {
			t.Errorf("Expect(%d,%s); Get(%d %s)", tt.code, tt.message, response.StatusCode, body)
		}
	}
}
