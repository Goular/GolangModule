package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello JiaGongWu Blog！"))
	})
	http.ListenAndServe("127.0.0.1:8080", nil);
}
