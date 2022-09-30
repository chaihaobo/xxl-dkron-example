// Package dkron
// @author： Boice
// @createTime：2022/9/30 09:59
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", func(writer http.ResponseWriter, request *http.Request) {
		println("hello world")
	})
	err := http.ListenAndServe(":2531", nil)
	if err != nil {
		panic(err)
	}

}
