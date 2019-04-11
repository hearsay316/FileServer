package main

import (
	"fmt"
	"handler"
	"net/http"
)

func main() {
	fmt.Println("za")
	// 路由接口
	http.HandleFunc("/file/upload", handler.Uploaded)
	http.HandleFunc("/file/upload/suc", handler.UploadSuc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe err", err)
		return
	}
}
