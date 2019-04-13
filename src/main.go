package main

import (
	"fmt"
	"handler"
	"net/http"
)

func main() {
	fmt.Println("服务器启动了")
	// 路由接口
	http.HandleFunc("/file/upload", handler.Uploaded)
	http.HandleFunc("/file/upload/suc", handler.UploadSuc)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileUpdateMetaHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe err", err)
		return
	}
}
